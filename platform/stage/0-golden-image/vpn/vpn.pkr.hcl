packer {
  required_plugins {
    digitalocean = {
      source  = "github.com/digitalocean/digitalocean"
      version = ">=1.1.1"
    }
  }
}

variables {
  base_image_name = "${env("DR_ENV_VAR_PACKER_IMAGE_DO")}"
  vpc_uuid        = "${env("DR_ENV_VAR_VPC_UUID")}"
}

locals {
  sudo_user     = "ufukty"
  dir_name      = basename(abspath(path.root))
  now           = formatdate("YY-MM-DD-'T'-hh-mm-ss-ZZZ", timestamp())
  snapshot_name = replace("packer-${local.dir_name}-${local.now}", "_", "_")
}

source "digitalocean" "droplet" {
  image              = var.base_image_name
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
    source      = "./provisioner_files"
    destination = "~"
  }

  provisioner "file" {
    source      = "${path.root}/../../../../shell/provisioner-utilities.sh"
    destination = "~/provisioner_files/provisioner_utilities.sh"
  }

  provisioner "shell" {
    environment_vars = [
      "EXECUTION_MODE=GOLDEN_IMAGE",
      "USER_ACCOUNT_NAME=ufukty",
      "PUBLIC_IP=nil-for-golden-image",
      "PRIVATE_IP=nil-for-golden-image",
      "VPC_CIDR=nil-for-golden-image",
      "VPC_ADDRESS=nil-for-golden-image",
      "SUBNET_ADDRESS=nil-for-golden-image",
      "DNS_ADDRESS=nil-for-golden-image",
      "PUBLIC_ETHERNET_INTERFACE=nil-for-golden-image",
      "PRIVATE_ETHERNET_INTERFACE=nil-for-golden-image",
      "SERVER_NAME=nil-for-golden-image",
      "EASYRSA_REQ_CN=nil-for-golden-image",
      "OTP_URI_ISSUER_NAME=nil-for-golden-image",
    ]
    inline = [
      "cd provisioner_files; sudo -u root --preserve-env bash shell.sh client-a client-b client-c"
    ]
  }

  provisioner "breakpoint" {
    disable = true
    note    = "Last stop before start to taking the snapshot."
  }
}
