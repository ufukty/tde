terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.27.1"
    }
  }
}

# MARK: Variables

variable "vpc" {
  type = object({
    region   = string
    ip_range = string
    uuid     = string
  })
}

variable "regions" {
  type = object({
    do = set(string)
  })
}
variable "vpc_uuid" { type = string }
variable "ssh_fingerprints" { type = list(string) }
variable "lan_network_cidr" { type = string }
variable "lan_network_address" { type = string }
variable "subnet_address" { type = string }
variable "username" { type = string }
variable "droplet_name" { type = string }
variable "easyrsa_server_name" { type = string }
variable "easyrsa_server_common_name" { type = string }
variable "otp_uri_issuer_name" { type = string }

# MARK: Main

data "digitalocean_droplet_snapshot" "golden_vpn" {
  name_regex  = "^packer-vpn-.*"
  region      = var.region
  most_recent = true
}

resource "digitalocean_droplet" "vpn-server" {
  for_each = var.regions.do

  ipv6        = true
  name        = var.droplet_name
  size        = "s-1vcpu-1gb"
  image       = data.digitalocean_droplet_snapshot.golden_vpn.id
  region      = each.value
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = var.ssh_fingerprints
  vpc_uuid    = var.vpc_uuid
  tags        = ["vpn", "thesis"]

  connection {
    host    = self.ipv4_address
    user    = var.username
    type    = "ssh"
    agent   = true
    timeout = "2m"
  }

  // Run provisioner script
  provisioner "remote-exec" {
    inline = [
      <<EOF
        cd provisioner_files
        EXECUTION_MODE="DEPLOYMENT" \
        USER_ACCOUNT_NAME="${var.username}" \
        PUBLIC_IP="${self.ipv4_address}" \
        PRIVATE_IP="${self.ipv4_address_private}" \
        VPC_CIDR="${var.lan_network_cidr}" \
        VPC_ADDRESS="${var.lan_network_address}" \
        SUBNET_ADDRESS="${var.subnet_address}" \
        DNS_ADDRESS="${cidrhost(format("%s%s", var.subnet_address, "/16"), 1)}" \
        PUBLIC_ETHERNET_INTERFACE="eth0" \
        PRIVATE_ETHERNET_INTERFACE="eth1" \
        SERVER_NAME="${var.easyrsa_server_name}" \
        EASYRSA_REQ_CN="${var.easyrsa_server_common_name}" \
        OTP_URI_ISSUER_NAME="${var.otp_uri_issuer_name}" \
        sudo --preserve-env bash shell.sh \
        client-a client-b client-c
      EOF
    ]
  }

  // Download configuration files from server
  provisioner "local-exec" {
    command     = <<EOF
      mkdir -p ../artifacts/vpn
      scp ${self.ipv4_address}:~/otp-uri.txt ../artifacts/vpn/do-${var.region}-otp-uri.txt
      scp ${self.ipv4_address}:~/client-a.ovpn ../artifacts/vpn/do-${var.region}-client-a.ovpn
      scp ${self.ipv4_address}:~/client-b.ovpn ../artifacts/vpn/do-${var.region}-client-b.ovpn
      scp ${self.ipv4_address}:~/client-c.ovpn ../artifacts/vpn/do-${var.region}-client-c.ovpn
      
      ssh-keygen -R ${self.ipv4_address_private}
      ssh-keyscan ${self.ipv4_address_private} >> ~/.ssh/known_hosts
    EOF
    interpreter = ["/bin/bash", "-c"]
    working_dir = path.root
  }

  // Clean up in server & last changes
  provisioner "remote-exec" {
    inline = [
      "rm -rf ~/otp-uri.txt ~/*.ovpn ~/provisioner_files",
      "sudo systemctl restart systemd-journald",
      "sudo sed -i \"s;${var.username}\\(.*\\)NOPASSWD:\\(.*\\);${var.username} \\1 \\2;\" /etc/sudoers",
    ]
  }
}
