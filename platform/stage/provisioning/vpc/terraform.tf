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

variable "vpc_details" {
  type = object({
    do = object({
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
  })
}

resource "digitalocean_vpc" "vpc" {
  for_each = toset(keys(var.vpc_details.do))

  name     = "${var.project_prefix}-${each.value}"
  region   = each.value
  ip_range = var.vpc_details.do[each.value].range
}
