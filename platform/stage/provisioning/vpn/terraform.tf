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

locals {
  public_ethernet_interface  = "eth0"
  private_ethernet_interface = "eth1"
}

# lan_network_cidr    = var.vpc_details.do[each.value].range
# lan_network_address = cidrhost(var.vpc_details.do[each.value].range, 0)
# subnet_address      = var.vpn_details.do[each.value].subnet_address

# username                   = "a4v95e281o7hvmc"
# droplet_name               = "pic-do-${each.value}-vpn"
# easyrsa_server_name        = "pic_do_${each.value}_vpn"
# easyrsa_server_common_name = "pic-do-${each.value}-vpn-rcn"
# otp_uri_issuer_name        = "pic-do-${each.value}"
# base_image_name            = "packer-common-images-openvpn-focal-64"
# droplet_tag_project_name   = "picarus"
# ssh_config_hostname        = "pic.do.${each.value}.vpn"
# ssh_config_path            = "${path.root}/../ssh.conf"

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

  provisioner "remote-exec" {
    inline = [
      <<EOF
        cd provisioner_files && \
            USER_ACCOUNT_NAME=\
                "${var.sudo_user}" \
            SERVER_NAME=\
                "${var.project_prefix}-do-${each.value}-vpn" \
            PUBLIC_IP=\
                "${self.ipv4_address}" \
            PRIVATE_IP=\
                "${self.ipv4_address_private}" \
            VPC_CIDR=\
                "${data.digitalocean_vpc.vpc[each.value].ip_range}" \
            VPC_ADDRESS=\
                "${cidrhost(data.digitalocean_vpc.vpc[each.value].ip_range, 0)}" \
            SUBNET_ADDRESS=\
                "${var.digitalocean.config.vpn[each.value].subnet_address}" \
            PUBLIC_ETHERNET_INTERFACE=\
                "${local.public_ethernet_interface}" \
            PRIVATE_ETHERNET_INTERFACE=\
                "${local.private_ethernet_interface}" \
            OVPN_USERNAME=\
                "${var.OVPN_USER}" \
            OVPN_HASH=\
                "${var.OVPN_HASH}" \
            sudo --preserve-env bash deployment.sh 
      EOF
    ]
  }

  #   provisioner "local-exec" {
  #     command = <<EOF
  #       mkdir -p ~/vpn
  #       scp ${self.ipv4_address}:~/otp-uri.txt ~/vpn/${var.droplet_tag_project_name}-do-${var.region}-${local.date_time_string}-otp-uri.txt
  #       scp ${self.ipv4_address}:~/client-a.ovpn ~/vpn/${var.droplet_tag_project_name}-do-${var.region}-${local.date_time_string}-client-a.ovpn
  #       scp ${self.ipv4_address}:~/client-b.ovpn ~/vpn/${var.droplet_tag_project_name}-do-${var.region}-${local.date_time_string}-client-b.ovpn
  #       scp ${self.ipv4_address}:~/client-c.ovpn ~/vpn/${var.droplet_tag_project_name}-do-${var.region}-${local.date_time_string}-client-c.ovpn
  #     EOF
  #   }

  #   // Clean up in server & last changes
  #   provisioner "remote-exec" {
  #     inline = [
  #       "rm -rf ~/otp-uri.txt ~/*.ovpn ~/provisioner_files",
  #       "sudo systemctl restart systemd-journald",
  #       "sudo sed -i \"s;${var.username}\\(.*\\)NOPASSWD:\\(.*\\);${var.username} \\1 \\2;\" /etc/sudoers",
  #     ]
  #   }
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


