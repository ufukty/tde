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
  vpc_uuid      = "${env("VPC_UUID")}"
}

locals {
  sudo_user     = "olwgtzjzhnvexhpr"
  dir_name      = basename(abspath(path.root))
  now           = formatdate("YY-MM-DD-'T'-hh-mm-ss-ZZZ", timestamp())
  snapshot_name = replace("packer-${local.dir_name}-${local.now}", "_", "_")
}

source "digitalocean" "droplet" {
  image                   = var.base_image_id
  region                  = "fra1"
  size                    = "s-1vcpu-1gb"
  snapshot_name           = local.snapshot_name
  snapshot_regions        = ["fra1", "nyc3"]
  private_networking      = true
  vpc_uuid                = var.vpc_uuid
  connect_with_private_ip = true
  ssh_agent_auth          = true
  ssh_username            = local.sudo_user
  tags                    = [local.dir_name]
}

build {
  sources = ["source.digitalocean.droplet"]

  provisioner "file" {
    source      = "${path.root}/provisioner-files"
    destination = "~"
  }

  provisioner "file" {
    source      = "${path.root}/../../secrets/image/ssh-app-db/app-db"
    destination = "~/provisioner-files/ssh-application-db"
  }

  provisioner "file" {
    source      = "${path.root}/../../../secrets/pki/issued/app-db.crt"
    destination = "~/provisioner-files/app-db.crt"
  }

  provisioner "file" {
    source      = "${path.root}/../../../secrets/pki/private/app-db.key"
    destination = "~/provisioner-files/app-db.key"
  }

  provisioner "shell" {
    environment_vars = [
      "POSTGRES_USER=nil",
      "POSTGRES_SERVER_PRIVATE_IP=nil",
      "IPTABLES_PRIVATE_ETHERNET_INTERFACE=eth1"
    ]
    inline = [
      "cd ~/provisioner-files && sudo -u root --preserve-env bash golden-image.sh"
    ]
  }

  provisioner "breakpoint" {
    disable = true
    note    = "Last stop before start to taking the snapshot."
  }
}
