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

variable "DIGITALOCEAN_THESIS_TOKEN" { // env var
}

# ------------------------------------------------------------- #
# Locals
# ------------------------------------------------------------- #

locals {
  region           = "fra1"
  slug             = "s-1vcpu-1gb"
  user_name        = "a4v95e281o7hvmc"
  base_image_name  = "thesis-base-focal-64"
  date_time_string = formatdate("YY-MM-DD-hh-mm-ss", timestamp())
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
  count = 1

  image  = data.digitalocean_droplet_snapshot.last_snapshot.id
  name   = "thesis-runner-${count.index}"
  region = local.region
  size   = "s-1vcpu-1gb"
  tags   = ["thesis", "thesis-runner"]

  ipv6        = true
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = local.ssh_fingerprints
  vpc_uuid    = digitalocean_vpc.vpc.id
}

resource "digitalocean_droplet" "evolution" {
  image  = data.digitalocean_droplet_snapshot.last_snapshot.id
  name   = "thesis-evolution"
  region = local.region
  size   = "s-1vcpu-1gb"
  tags   = ["thesis", "thesis-evolution"]

  ipv6        = true
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = local.ssh_fingerprints
  vpc_uuid    = digitalocean_vpc.vpc.id
}

resource "local_file" "inventory" {
  content = templatefile("${path.module}/inventory.template.cfg",
    {
      runners   = digitalocean_droplet.runners.*.ipv4_address
      evolution = digitalocean_droplet.evolution.*.ipv4_address
    }
  )
  filename = "inventory.cfg"
}
