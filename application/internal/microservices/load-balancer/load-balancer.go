package load_balancer

import (
	"errors"
	"fmt"
	"tde/internal/microservices/logger"
	service_discovery "tde/internal/microservices/service-discovery"
	"tde/internal/microservices/service-discovery/models/services"
	"tde/internal/utilities"
)

var (
	ErrNoHostAvailable = errors.New("no hosts are available right now")
)

type LoadBalancer struct {
	sd            *service_discovery.ServiceDiscovery
	index         int
	targetPort    string
	targetService services.ServiceName
	addToPath     string
}

var log = logger.NewLogger("BalancedForwarder")

// service: ip address and host
// hosts: ip addresses of available hosts
// port: port which will be used as forwarded target
func New(sd *service_discovery.ServiceDiscovery, targetService services.ServiceName, targetPort, addToPath string) *LoadBalancer {
	return &LoadBalancer{
		sd:    sd,
		index: 0,
	}
}

func (lb *LoadBalancer) appendSuffixes(host string) string {
	return fmt.Sprintf("%s%s/%s", host, lb.targetPort, lb.addToPath)
}

func (lb *LoadBalancer) Next() (string, error) {
	var hosts = lb.sd.ListPrivateIPs(lb.targetService)
	if len(hosts) == 0 {
		return "", ErrNoHostAvailable
	} else if len(hosts) <= lb.index {
		lb.index = utilities.URandIntN(len(hosts))
	}
	var next = hosts[lb.index]
	lb.index = (lb.index + 1) % len(hosts)
	return lb.appendSuffixes(next), nil
}
