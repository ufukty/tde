package models

import (
	"tde/internal/microservices/serviced/models/provider"
	"tde/internal/microservices/serviced/models/provider/digitalocean"
	"tde/internal/microservices/serviced/models/provider/local"
	"tde/internal/microservices/serviced/models/services"
)

type ServiceDiscoveryFile struct {
	Digitalocean digitalocean.Digitalocean `json:"digitalocean"`
	Local        local.Local               `json:"local"`
}

func (f ServiceDiscoveryFile) ListPrivateIPs(service services.ServiceName) (ips []string) {
	var providers = []provider.Provider{&f.Digitalocean, &f.Local}
	for _, provider := range providers {
		ips = append(ips, provider.ListPrivateIPs(service)...)
	}
	return
}
