terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.27.1"
    }
  }
}

resource "digitalocean_volume" "fra1-customs" {
  for_each = toset([
    "fra1-customs-0-a",
    "fra1-customs-0-b",
    "fra1-customs-1-a",
    "fra1-customs-1-b",
  ])

  region                  = "fra1"
  name                    = each.value
  size                    = 1
  initial_filesystem_type = "ext4"
  description             = "object storage for user uploads"
}
