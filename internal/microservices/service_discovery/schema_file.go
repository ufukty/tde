package service_discovery

type providers struct {
	Digitalocean []digitaloceanDroplet `json:"digitalocean"`
}

func (schPro providers) GetIPs() (ips []string) {
	for _, droplet := range schPro.Digitalocean {
		ips = append(ips, droplet.Ipv4AddressPrivate)
	}
	return
}

type serviceDiscoveryFile struct {
	Runner  providers `json:"runner"`
	Evolver providers `json:"evolver"`
	Customs providers `json:"customs"`
}
