package loadbalancer

import (
	"tde/internal/microservices/logger"
	"tde/internal/microservices/serviced"
	"tde/internal/microservices/serviced/models/services"
	"tde/internal/utilities"

	"errors"
)

var (
	ErrNoHostAvailable = errors.New("no hosts are available right now")
)

type LoadBalancer struct {
	sd            *serviced.ServiceDiscovery
	index         int
	targetService services.ServiceName
}

var log = logger.NewLogger("LoadBalancer")

// service: ip address and host
// hosts: ip addresses of available hosts
// port: port which will be used as forwarded target
func New(sd *serviced.ServiceDiscovery, targetService services.ServiceName) *LoadBalancer {
	return &LoadBalancer{
		sd:            sd,
		index:         0,
		targetService: targetService,
	}
}

func (lb *LoadBalancer) Next() (string, error) {
	var hosts = lb.sd.ListPrivateIPs(lb.targetService)
	if len(hosts) == 0 {
		return "", ErrNoHostAvailable
	}
	if len(hosts) <= lb.index {
		lb.index = utilities.URandIntN(len(hosts))
	}
	var next = hosts[lb.index]
	lb.index = (lb.index + 1) % len(hosts)
	return next, nil
}
