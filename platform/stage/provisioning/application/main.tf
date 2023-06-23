terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.27.1"
    }
  }
}

# ------------------------------------------------------------- #
# Variables
# ------------------------------------------------------------- #

variable "project_prefix" { type = string }

# ------------------------------------------------------------- #
# Locals
# ------------------------------------------------------------- #

locals {
  region = "fra1"
  slug   = "s-1vcpu-1gb"
  instances = {
    runner      = 3,
    evolver     = 1,
    api-gateway = 1,
  }
  ssh_fingerprints = ["42:75:b8:ad:c1:76:4b:58:07:ec:e9:85:66:27:9b:e6"]
}

# ------------------------------------------------------------- #
# Main
# ------------------------------------------------------------- #

data "digitalocean_droplet_snapshot" "internal" {
  name_regex  = "^packer-internal-.*"
  region      = local.region
  most_recent = true
}

data "digitalocean_droplet_snapshot" "gateway" {
  name_regex  = "^packer-gateway-.*"
  region      = local.region
  most_recent = true
}

# data "digitalocean_droplet_snapshot" "application" {
#   name_regex  = "^packer-application-.*"
#   region      = local.region
#   most_recent = true
# }

data "digitalocean_droplet_snapshot" "combined" {
  name_regex  = "^packer-combined-.*"
  region      = local.region
  most_recent = true
}

# data "digitalocean_droplet_snapshot" "compiled" {
#   name_regex  = "^packer-compiled-.*"
#   region      = local.region
#   most_recent = true
# }

data "digitalocean_vpc" "vpc" {
  name = "${var.project_prefix}-${local.region}"
}

resource "digitalocean_droplet" "runner" {
  count = local.instances.runner

  image  = data.digitalocean_droplet_snapshot.internal.id
  name   = "${local.region}-runner-${count.index}"
  region = local.region
  size   = local.slug
  tags   = ["thesis", "runner"]

  ipv6        = true
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = local.ssh_fingerprints
  vpc_uuid    = data.digitalocean_vpc.vpc.id
}

resource "digitalocean_droplet" "evolver" {
  count = local.instances.evolver

  image  = data.digitalocean_droplet_snapshot.combined.id
  name   = "${local.region}-evolver-${count.index}"
  region = local.region
  size   = local.slug
  tags   = ["thesis", "evolver"]

  ipv6        = true
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = local.ssh_fingerprints
  vpc_uuid    = data.digitalocean_vpc.vpc.id
}

data "digitalocean_volume" "customs_storage_volume" {
  name   = "volume-fra1-01"
  region = "fra1"
}

resource "digitalocean_droplet" "customs" {
  count = 1

  image      = data.digitalocean_droplet_snapshot.combined.id
  name       = "${local.region}-customs-${count.index}"
  region     = local.region
  size       = local.slug
  volume_ids = [data.digitalocean_volume.customs_storage_volume.id]
  tags       = ["thesis", "customs"]

  ipv6        = true
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = local.ssh_fingerprints
  vpc_uuid    = data.digitalocean_vpc.vpc.id
}

resource "digitalocean_droplet" "api-gateway" {
  count = local.instances.api-gateway

  image  = data.digitalocean_droplet_snapshot.gateway.id
  name   = "${local.region}-api-gateway-${count.index}"
  region = local.region
  size   = local.slug
  tags   = ["thesis", "api-gateway"]

  ipv6        = true
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = local.ssh_fingerprints
  vpc_uuid    = data.digitalocean_vpc.vpc.id
}

resource "local_file" "inventory" {
  content = templatefile(
    "${path.module}/templates/inventory.cfg.tftpl",
    {
      providers = {
        digitalocean = {
          fra1 = {
            vpc = data.digitalocean_vpc.vpc
            services = {
              api-gateway = digitalocean_droplet.api-gateway
              customs     = digitalocean_droplet.customs
              evolver     = digitalocean_droplet.evolver
              runner      = digitalocean_droplet.runner
            }
          }
        }
      }
    }
  )
  filename = abspath("${path.module}/../../artifacts/deployment/inventory.cfg")
}

resource "local_file" "ssh-config" {
  content = templatefile(
    "${path.module}/templates/ssh.conf.tftpl",
    {
      providers = {
        digitalocean = {
          fra1 = {
            vpc = data.digitalocean_vpc.vpc
            services = {
              api-gateway = digitalocean_droplet.api-gateway
              customs     = digitalocean_droplet.customs
              evolver     = digitalocean_droplet.evolver
              runner      = digitalocean_droplet.runner
            }
          }
        }
      }
    }
  )
  filename = abspath("${path.module}/../../artifacts/ssh.conf.d/0.application.conf")
}

resource "local_file" "service_discovery" {
  content = templatefile(
    "${path.module}/templates/service_discovery.json.tftpl",
    {
      content = jsonencode({
        digitalocean = {
          fra1 = {
            vpc = data.digitalocean_vpc.vpc
            services = {
              api-gateway = digitalocean_droplet.api-gateway
              customs     = digitalocean_droplet.customs
              evolver     = digitalocean_droplet.evolver
              runner      = digitalocean_droplet.runner
            }
          }
        }
      })
    }
  )
  filename = abspath("${path.module}/../../artifacts/deployment/service_discovery.json")
}
