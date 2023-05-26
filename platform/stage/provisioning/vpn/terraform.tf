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
variable "sudo_user" { type = string }
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
variable "OVPN_USER" { type = string }
variable "OVPN_HASH" { type = string }
variable "openvpn_client_name" { type = string }

locals {
  public_ethernet_interface  = "eth0"
  private_ethernet_interface = "eth1"
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

  connection {
    host    = self.ipv4_address
    user    = var.sudo_user
    type    = "ssh"
    agent   = true
    timeout = "2m"
  }

  provisioner "file" {
    source      = "${path.module}/provisioner-files"
    destination = "/home/${var.sudo_user}"
  }

  provisioner "remote-exec" {
    inline = [
      <<EOF
        cd ~/provisioner-files && \
                     USER_ACCOUNT_NAME="${var.sudo_user}" \
                           SERVER_NAME="${var.project_prefix}-do-${each.value}-vpn" \
                             PUBLIC_IP="${self.ipv4_address}" \
                            PRIVATE_IP="${self.ipv4_address_private}" \
                OPENVPN_SUBNET_ADDRESS="${var.digitalocean.config.vpn[each.value].subnet_address}" \
                   OPENVPN_SUBNET_MASK="255.255.255.0" \
             PUBLIC_ETHERNET_INTERFACE="${local.public_ethernet_interface}" \
            PRIVATE_ETHERNET_INTERFACE="${local.private_ethernet_interface}" \
                         OVPN_USERNAME="${var.OVPN_USER}" \
                             OVPN_HASH="${var.OVPN_HASH}" \
            sudo --preserve-env bash deployment.sh 

        cd ~/provisioner-files && \
            USER_ACCOUNT_NAME="${var.sudo_user}" \
                    PUBLIC_IP="${self.ipv4_address}" \
                  CLIENT_NAME="${var.openvpn_client_name}" \
            sudo --preserve-env bash new_client.sh
      EOF   

    ]
  }

  provisioner "local-exec" {
    command = <<EOF
      ssh-keygen -R ${self.ipv4_address_private}
      ssh-keyscan ${self.ipv4_address_private} >> ~/.ssh/known_hosts

      mkdir -p ../../artifacts/vpn
      scp ${var.sudo_user}@${self.ipv4_address}:~/artifacts/totp-share.txt \
          ${path.module}/../../artifacts/vpn/${var.project_prefix}-do-${each.value}-totp-share.txt
      scp ${var.sudo_user}@${self.ipv4_address}:~/artifacts/${var.openvpn_client_name}.ovpn \
          ${path.module}/../../artifacts/vpn/${var.project_prefix}-do-${each.value}-${var.openvpn_client_name}.ovpn
    EOF
  }

  // Clean up in server & last changes
  provisioner "remote-exec" {
    inline = [
      "rm -rf ~/artifacts ~/provisioner_files",
      #   "sudo systemctl restart systemd-journald",
      #   "sudo sed -i \"s;${var.username}\\(.*\\)NOPASSWD:\\(.*\\);${var.username} \\1 \\2;\" /etc/sudoers",
    ]
  }
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
  depends_on = [local_file.ssh-config]

  provisioner "local-exec" {
    command     = "cat ssh.conf.d/* > ssh.conf"
    working_dir = "${path.module}/../../artifacts"
  }
}


