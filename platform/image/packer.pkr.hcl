
variables {
  token_digitalocean = "${env("DO_API_KEY")}"
  base_image_name    = "${env("DR_ENV_VAR_PACKER_IMAGE_DO")}"
  dir_name           = "${env("DIR_NAME")}"
}

locals {
  sudo_user     = "a4v95e281o7hvmc"
  timestamp     = regex_replace(timestamp(), "[- TZ:]", "")
  register_name = "python-worker-${local.timestamp}"
}

source "digitalocean" "droplet" {
  api_token = "${var.token_digitalocean}"

  image  = "${var.base_image_name}"
  region = "fra1"
  size   = "s-1vcpu-1gb"

  snapshot_name    = "${replace(local.register_name, "_", "-")}"
  snapshot_regions = ["fra1"]

  private_networking      = false
  connect_with_private_ip = false

  ssh_agent_auth = true
  ssh_username   = "${local.sudo_user}"

  tags = ["${var.dir_name}"]
}

build {
  sources = ["source.digitalocean.droplet"]

  provisioner "file" {
    source      = "files"
    destination = "~/provision-files"
  }

  provisioner "shell" {
    inline = [
      "cd ~/provision-files && sudo -u root --preserve-env -H bash shell.sh",
      "rm -rfv ~/provision-files"
    ]
  }
}
