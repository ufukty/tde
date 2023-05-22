terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.27.1"
    }
  }
}

data "digitalocean_droplet_snapshot" "golden_base" {
  name_regex  = "^packer-base-.*"
  region      = local.region
  most_recent = true
}

resource "digitalocean_vpc" "vpc" {
  name   = "vpc"
  region = local.region
}

resource "digitalocean_droplet" "runner" {
  count = local.instances.runner

  image  = data.digitalocean_droplet_snapshot.golden_base.id
  name   = "runner-${count.index}"
  region = local.region
  size   = local.slug
  tags   = ["thesis", "runner"]

  ipv6        = true
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = local.ssh_fingerprints
  vpc_uuid    = digitalocean_vpc.vpc.id
}

resource "digitalocean_droplet" "evolver" {
  count = local.instances.evolver

  image  = data.digitalocean_droplet_snapshot.golden_base.id
  name   = "evolver-${count.index}"
  region = local.region
  size   = local.slug
  tags   = ["thesis", "evolver"]

  ipv6        = true
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = local.ssh_fingerprints
  vpc_uuid    = digitalocean_vpc.vpc.id
}

data "digitalocean_volume" "customs_storage_volume" {
  name   = "volume-fra1-01"
  region = "fra1"
}

resource "digitalocean_droplet" "customs" {
  count = 1

  image      = data.digitalocean_droplet_snapshot.golden_base.id
  name       = "customs-${count.index}"
  region     = local.region
  size       = local.slug
  volume_ids = [data.digitalocean_volume.customs_storage_volume.id]
  tags       = ["thesis", "customs"]

  ipv6        = true
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = local.ssh_fingerprints
  vpc_uuid    = digitalocean_vpc.vpc.id
}

resource "digitalocean_droplet" "api-gateway" {
  count = local.instances.api-gateway

  image  = data.digitalocean_droplet_snapshot.golden_base.id
  name   = "api-gateway-${count.index}"
  region = local.region
  size   = local.slug
  tags   = ["thesis", "api-gateway"]

  ipv6        = true
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = local.ssh_fingerprints
  vpc_uuid    = digitalocean_vpc.vpc.id
}

resource "local_file" "inventory" {
  content = templatefile(
    "${path.module}/templates/inventory.cfg.tftpl",
    {
      providers = {
        digitalocean = {
          fra1 = {
            vpc = digitalocean_vpc.vpc
            services = {
              api-gateway = digitalocean_droplet.api-gateway
              customs     = digitalocean_droplet.customs
              evolver     = digitalocean_droplet.evolver
              runner      = digitalocean_droplet.runner
              #   vpn         = digitalocean_droplet.vpn
            }
          }
        }
      }
    }
  )
  filename = abspath("${path.module}/../../deployment/inventory.cfg")
}

resource "local_file" "ssh-config" {
  content = templatefile(
    "${path.module}/templates/ssh.conf.tftpl",
    {
      providers = {
        digitalocean = {
          fra1 = {
            vpc = digitalocean_vpc.vpc
            services = {
              api-gateway = digitalocean_droplet.api-gateway
              customs     = digitalocean_droplet.customs
              evolver     = digitalocean_droplet.evolver
              runner      = digitalocean_droplet.runner
              vpn         = digitalocean_droplet.vpn
            }
          }
        }
      }
    }
  )
  filename = abspath("${path.module}/../../artifacts/ssh.0.application.conf")
}

resource "terraform_data" "ssh_config_aggregate" {
  provisioner "local-exec" {
    command     = "cat ssh.*.conf > ssh.conf"
    working_dir = "${path.module}/../../artifacts"
  }
}

resource "local_file" "service_discovery" {
  content = templatefile(
    "${path.module}/templates/service_discovery.json.tftpl",
    {
      content = jsonencode({
        digitalocean = {
          fra1 = {
            vpc = digitalocean_vpc.vpc
            services = {
              api-gateway = digitalocean_droplet.api-gateway
              customs     = digitalocean_droplet.customs
              evolver     = digitalocean_droplet.evolver
              runner      = digitalocean_droplet.runner
              #   vpn         = digitalocean_droplet.vpn
            }
          }
        }
      })
    }
  )
  filename = abspath("${path.module}/../../deployment/service_discovery.json")
}
