
variables {
  token_digitalocean = "${env("DO_API_KEY")}"
  base_image_name    = "${env("DR_ENV_VAR_PACKER_IMAGE_DO")}"
  vpc_uuid           = "${env("DR_ENV_VAR_VPC_UUID")}"

  snapshot_name = "${env("DR_SNAPSHOT_NAME")}"
  project_name  = "${env("DR_PROJECT_NAME")}"
  image_name    = "${env("DR_IMAGE_NAME")}"
}

locals {
  sudo_user = "a4v95e281o7hvmc"
  timestamp = regex_replace(timestamp(), "[- TZ:]", "")
}

source "digitalocean" "droplet" {
  api_token = var.token_digitalocean

  image  = var.base_image_name
  region = "fra1"
  size   = "s-1vcpu-1gb"

  snapshot_name    = replace(var.snapshot_name, "_", "-")
  snapshot_regions = ["fra1", "nyc3"]

  private_networking      = true
  vpc_uuid                = var.vpc_uuid
  connect_with_private_ip = true

  ssh_agent_auth = true
  ssh_username   = local.sudo_user

  tags = [var.image_name, var.project_name]
}

build {
  sources = ["source.digitalocean.droplet"]

  provisioner "file" {
    source      = "./provisioner_files"
    destination = "~"
  }

  provisioner "file" {
    source      = "${path.root}/../../../companions/shell_commons.sh"
    destination = "~/provisioner_files/shell_commons.sh"
  }

  provisioner "file" {
    source      = "${path.root}/../../../certificates/sync/ssh-key-for-database-access/sync-application-server.private"
    destination = "~/provisioner_files/sync-application-server.private"
  }

  provisioner "file" {
    source      = "${path.root}/../../../certificates/sync/tls-cert-for-api-server/" // NOTE: don't remove trailing slash
    destination = "~/provisioner_files"
  }

  provisioner "shell" {
    environment_vars = [
      "EXECUTION_MODE=GOLDEN_IMAGE",
      "POSTGRES_USER=nil",
      "POSTGRES_SERVER_PRIVATE_IP=nil",
      "IPTABLES_PRIVATE_ETHERNET_INTERFACE=eth1"
    ]
    inline = [
      "cd ~/provisioner_files && sudo -u root --preserve-env bash shell.sh"
    ]
  }

  provisioner "breakpoint" {
    disable = true
    note    = "Last stop before start to taking the snapshot."
  }
}
