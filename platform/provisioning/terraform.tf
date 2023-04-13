terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.27.1"
    }
  }
}

variable "DSDS" {}

variable "DIGITALOCEAN_THESIS_TOKEN" {} // env var

provider "digitalocean" {
  # Configuration options
  token = var.DIGITALOCEAN_THESIS_TOKEN
}

resource "digitalocean_droplet" "runners" {
  count = 1

  image  = "ubuntu-uptodate-focal-64-1681391363"
  name   = "thesis-runner-${count.index}"
  region = "fra1" // ams3
  size   = "s-1vcpu-1gb"
  tags   = ["thesis"]
}

resource "digitalocean_droplet" "evolution" {
  image  = "ubuntu-uptodate-focal-64-1681391363"
  name   = "thesis-evolution"
  region = "fra1" // ams3
  size   = "s-1vcpu-1gb"
  tags   = ["thesis"]
}
