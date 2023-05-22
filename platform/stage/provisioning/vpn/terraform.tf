terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.27.1"
    }
  }
}

# MARK: Variables

variable "project_prefix" { type = string }

variable "ssh_fingerprints" { type = list(string) }

variable "digitalocean" {
  type = object({
    activated_regions = object({
      vpc = set(string)
      vpn = set(string)
    })
    config = object({
      vpc = object({
        sfo2 = object({ range = string })
        sfo3 = object({ range = string })
        tor1 = object({ range = string })
        nyc1 = object({ range = string })
        nyc3 = object({ range = string })
        lon1 = object({ range = string })
        ams3 = object({ range = string })
        fra1 = object({ range = string })
        blr1 = object({ range = string })
        sgp1 = object({ range = string })
      })
      vpn = object({
        sfo2 = object({ subnet_address = string })
        sfo3 = object({ subnet_address = string })
        tor1 = object({ subnet_address = string })
        nyc1 = object({ subnet_address = string })
        nyc3 = object({ subnet_address = string })
        lon1 = object({ subnet_address = string })
        ams3 = object({ subnet_address = string })
        fra1 = object({ subnet_address = string })
        blr1 = object({ subnet_address = string })
        sgp1 = object({ subnet_address = string })
      })
    })
  })
}

# MARK: Data gathering

data "digitalocean_droplet_snapshot" "vpn_image" {
  for_each = var.digitalocean.activated_regions.vpn

  name_regex  = "^packer-vpn-.*"
  region      = each.value
  most_recent = true
}

data "digitalocean_vpc" "vpc" {
  for_each = var.digitalocean.activated_regions.vpn

  name = "${var.project_prefix}-${each.value}"
}

# MARK: Resource creation

resource "digitalocean_droplet" "vpn-server" {
  for_each = var.digitalocean.activated_regions.vpn

  ipv6        = true
  name        = "${each.value}-vpn"
  size        = "s-1vcpu-1gb"
  image       = data.digitalocean_droplet_snapshot.vpn_image[each.value].id
  region      = each.value
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = var.ssh_fingerprints
  vpc_uuid    = data.digitalocean_vpc.vpc[each.value].id
  tags        = ["vpn", "thesis"]
}

resource "local_file" "ssh-config" {
  content = templatefile(
    "${path.module}/templates/ssh.conf.tftpl",
    {
      providers = {
        digitalocean = digitalocean_droplet.vpn-server
      }
    }
  )
  filename = abspath("${path.module}/../../artifacts/ssh.conf.d/0.vpn.conf")
}

resource "terraform_data" "ssh_config_aggregate" {
  provisioner "local-exec" {
    command     = "cat ssh.conf.d/* > ssh.conf"
    working_dir = "${path.module}/../../artifacts"
  }
}


