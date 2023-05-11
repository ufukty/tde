package models

import "tde/internal/microservices/service-discovery/models/provider"

type Service struct {
	Digitalocean []provider.DigitaloceanDroplet `json:"digitalocean"`
	Local        string                         `json:"local"`
}

func (s Service) GetIPs() (ips []string) {
	for _, droplet := range s.Digitalocean {
		ips = append(ips, droplet.Ipv4AddressPrivate)
	}
	if s.Local != "" {
		ips = append(ips, s.Local)
	}
	return
}

type File struct {
	Runner  Service `json:"runner"`
	Evolver Service `json:"evolver"`
	Customs Service `json:"customs"`
}
