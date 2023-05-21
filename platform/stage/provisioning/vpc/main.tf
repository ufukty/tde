terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.27.1"
    }
  }
}

variable "regions" {
  type = object({
    do = set(string)
  })
}

resource "digitalocean_vpc" "vpc" {
  for_each = toset(var.regions.do)

  name   = "thesis-${each.value}"
  region = each.value
}

output "output" {
  value = [for vpc in digitalocean_vpc.vpc : {
    region   = vpc.region,
    ip_range = vpc.ip_range,
    uuid     = vpc.uuid,
  }]
}
