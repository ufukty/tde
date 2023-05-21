terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.27.1"
    }
  }
}

variable "project_prefix" {
  type = string
}

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

resource "digitalocean_vpc" "vpc" {
  for_each = var.digitalocean.activated_regions.vpc

  name     = "${var.project_prefix}-${each.value}"
  region   = each.value
  ip_range = var.digitalocean.config.vpc[each.value].range
}
