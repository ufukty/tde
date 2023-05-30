packer {
  required_plugins {
    digitalocean = {
      source  = "github.com/digitalocean/digitalocean"
      version = ">=1.1.1"
    }
  }
}

variables {
  base_image_id = "${env("BASE_IMAGE_ID")}"
}

locals {
  sudo_user     = "2iuFDs13YDedYc3N"
  dir_name      = basename(abspath(path.root))
  now           = formatdate("YY-MM-DD-'T'-hh-mm-ss-ZZZ", timestamp())
  snapshot_name = replace("packer-${local.dir_name}-${local.now}", "_", "_")
}

source "digitalocean" "droplet" {
  image              = var.base_image_id
  region             = "fra1"
  size               = "s-1vcpu-1gb"
  snapshot_name      = local.snapshot_name
  snapshot_regions   = ["nyc3", "fra1"]
  tags               = [local.dir_name]
  ssh_username       = local.sudo_user
  ssh_agent_auth     = true
  private_networking = true
}

build {
  sources = ["source.digitalocean.droplet"]

  provisioner "file" {
    source      = "${path.root}/provisioner-files"
    destination = "~"
  }

  provisioner "shell" {
    inline = ["cd ~/provisioner-files && sudo --preserve-env bash golden-image.sh"]
  }

  provisioner "breakpoint" {
    disable = true
    note    = "Last stop before start to taking the snapshot."
  }
}
