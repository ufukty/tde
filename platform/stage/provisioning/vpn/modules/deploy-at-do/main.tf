
# MARK: Variables

variable "region" { type = string }
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
variable "base_image_name" { type = string }
variable "droplet_tag_project_name" { type = string }
variable "ssh_config_path" { type = string }
variable "ssh_config_hostname" { type = string }

# MARK: Locals

locals {
  slug             = "s-1vcpu-1gb"
  date_time_string = formatdate("YY-MM-DD-hh-mm-ss", timestamp())
}

# MARK: Main

data "digitalocean_droplet_snapshot" "snapshot_name" {
  name_regex  = "^${var.base_image_name}-[0-9]*"
  region      = var.region
  most_recent = true
}

resource "digitalocean_droplet" "created_droplet" {
  ipv6        = true
  name        = var.droplet_name
  size        = local.slug
  image       = data.digitalocean_droplet_snapshot.snapshot_name.id
  region      = var.region
  backups     = false
  monitoring  = true
  resize_disk = false
  ssh_keys    = var.ssh_fingerprints
  vpc_uuid    = var.vpc_uuid
  tags        = ["vpn", "terraform", var.droplet_tag_project_name]

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
        sudo --preserve-env bash deployment.sh \
        client-a client-b client-c
      EOF
    ]
  }

  // Download configuration files from server
  provisioner "local-exec" {
    command = <<EOF
      mkdir -p ~/vpn
      scp ${self.ipv4_address}:~/otp-uri.txt ~/vpn/${var.droplet_tag_project_name}-do-${var.region}-${local.date_time_string}-otp-uri.txt
      scp ${self.ipv4_address}:~/client-a.ovpn ~/vpn/${var.droplet_tag_project_name}-do-${var.region}-${local.date_time_string}-client-a.ovpn
      scp ${self.ipv4_address}:~/client-b.ovpn ~/vpn/${var.droplet_tag_project_name}-do-${var.region}-${local.date_time_string}-client-b.ovpn
      scp ${self.ipv4_address}:~/client-c.ovpn ~/vpn/${var.droplet_tag_project_name}-do-${var.region}-${local.date_time_string}-client-c.ovpn
    EOF
  }

  // Clean up in server & last changes
  provisioner "remote-exec" {
    inline = [
      "rm -rf ~/otp-uri.txt ~/*.ovpn ~/provisioner_files",
      "sudo systemctl restart systemd-journald",
      "sudo sed -i \"s;${var.username}\\(.*\\)NOPASSWD:\\(.*\\);${var.username} \\1 \\2;\" /etc/sudoers",
    ]
  }

  // Update ssh.conf
  provisioner "local-exec" {
    interpreter = ["/bin/bash", "-c"]
    working_dir = path.root
    command     = <<HERE
      RECORD_TO_ADD="Host ${var.ssh_config_hostname}\n  HostName ${self.ipv4_address_private}\n"

      grep -q "Host ${var.ssh_config_hostname}" ${var.ssh_config_path} &&
      perl -i -p0e "s/Host ${var.ssh_config_hostname}\n  HostName .*/$RECORD_TO_ADD/g" ${var.ssh_config_path} ||
      (
        echo -e "$RECORD_TO_ADD" | 
        cat - ${var.ssh_config_path} >config.bak && 
        mv config.bak ${var.ssh_config_path}
      )

      ssh-keygen -R ${self.ipv4_address_private}
      ssh-keyscan ${self.ipv4_address_private} >> ~/.ssh/known_hosts

      OTP_FILES="$(\ls -1 ~/vpn/*otp-uri.txt | sort -r | rev | uniq -s 29 | rev)";
      source ~/venv/bin/activate;
      echo "$OTP_FILES" | while read -r OTP_FILE; do
          cat "$OTP_FILE" | qr > "$OTP_FILE.png"
          OTP_FILENAME="$(basename "$OTP_FILE")"
          echo "Opening: $\{OTP_FILENAME:0:7}"
          open -W "$OTP_FILE.png";
          rm "$OTP_FILE.png" "$OTP_FILE"
      done
    HERE
  }
}
