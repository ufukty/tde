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

variable "DIGITALOCEAN_TOKEN" {
  // env var
}

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

provider "digitalocean" {}

data "digitalocean_droplet_snapshot" "last_snapshot" {
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

  image  = data.digitalocean_droplet_snapshot.last_snapshot.id
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

  image  = data.digitalocean_droplet_snapshot.last_snapshot.id
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

  image      = data.digitalocean_droplet_snapshot.last_snapshot.id
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

  image  = data.digitalocean_droplet_snapshot.last_snapshot.id
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
    "${path.module}/templates/inventory.template.cfg",
    {
      runner      = digitalocean_droplet.runner
      evolver     = digitalocean_droplet.evolver
      customs     = digitalocean_droplet.customs
      api-gateway = digitalocean_droplet.api-gateway
    }
  )
  filename = abspath("${path.module}/../2-deployment/inventory.cfg")
}

resource "local_file" "service_discovery" {
  content = templatefile(
    "${path.module}/templates/service_discovery.json.tftpl",
    {
      content = jsonencode({
        runner      = { digitalocean = digitalocean_droplet.runner }
        evolver     = { digitalocean = digitalocean_droplet.evolver }
        customs     = { digitalocean = digitalocean_droplet.customs }
        api-gateway = { digitalocean = digitalocean_droplet.api-gateway }
      })
    }
  )
  filename = abspath("${path.module}/../2-deployment/service_discovery.json")
}
