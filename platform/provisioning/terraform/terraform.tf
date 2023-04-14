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

variable "DIGITALOCEAN_THESIS_TOKEN" {
  // env var
}

# ------------------------------------------------------------- #
# Locals
# ------------------------------------------------------------- #

locals {
  region          = "fra1"
  slug            = "s-1vcpu-1gb"
  base_image_name = "thesis-base-focal-64"
  instances = {
    runner    = 1,
    evolution = 1,
  }
  ssh_fingerprints = ["42:75:b8:ad:c1:76:4b:58:07:ec:e9:85:66:27:9b:e6"]
}

# ------------------------------------------------------------- #
# Main
# ------------------------------------------------------------- #

provider "digitalocean" {
  # Configuration options
  token = var.DIGITALOCEAN_THESIS_TOKEN
}

data "digitalocean_droplet_snapshot" "last_snapshot" {
  name_regex  = "^packer-${local.base_image_name}-[0-9]*"
  region      = local.region
  most_recent = true
}

resource "digitalocean_vpc" "vpc" {
  name   = "thesis-vpc"
  region = local.region
}

resource "digitalocean_droplet" "runners" {
  count = local.instances.runner

  image  = data.digitalocean_droplet_snapshot.last_snapshot.id
  name   = "thesis-runner-${count.index}"
  region = local.region
  size   = local.slug
  tags   = ["thesis", "thesis-runner"]

  ipv6        = true
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = local.ssh_fingerprints
  vpc_uuid    = digitalocean_vpc.vpc.id
}

resource "digitalocean_droplet" "evolution" {
  count = local.instances.evolution

  image  = data.digitalocean_droplet_snapshot.last_snapshot.id
  name   = "thesis-evolution-${count.index}"
  region = local.region
  size   = local.slug
  tags   = ["thesis", "thesis-evolution"]

  ipv6        = true
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = local.ssh_fingerprints
  vpc_uuid    = digitalocean_vpc.vpc.id
}

resource "local_file" "inventory" {
  content = templatefile(
    "${path.module}/inventory.template.cfg",
    {
      runner_hosts    = digitalocean_droplet.runners
      evolution_hosts = digitalocean_droplet.evolution
    }
  )
  filename = abspath("${path.module}/../ansible/inventory.cfg")
}
