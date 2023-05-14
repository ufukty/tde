packer {
  required_plugins {
    digitalocean = {
      source  = "github.com/digitalocean/digitalocean"
      version = ">=1.1.1"
    }
  }
}

locals {
  dir_name      = basename(abspath(path.root))
  snapshot_name = replace("packer-${local.dir_name}-${formatdate("YY-MM-DD-'T'-hh-mm-ss-ZZZ", timestamp())}", "_", "_")
}

source "digitalocean" "droplet" {
  image            = "ubuntu-20-04-x64"
  region           = "fra1"
  size             = "s-1vcpu-1gb"
  snapshot_name    = local.snapshot_name
  snapshot_regions = ["fra1"]
  ssh_username     = "root"
  tags             = [local.dir_name]
}

build {
  sources = ["source.digitalocean.droplet"]

  provisioner "ansible" {
    playbook_file    = "ansible/playbook.yml"
    ansible_env_vars = ["ANSIBLE_CONFIG=ansible/ansible.cfg"]
  }

  provisioner "file" {
    destination = "/etc/ssl/certs/localhost.crt"
    source      = "../../../../certificates/localhost.crt"
  }

  provisioner "file" {
    destination = "/etc/ssl/private/localhost.key"
    source      = "../../../../certificates/localhost.key"
  }
}
