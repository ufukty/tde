package customs

import (
	config_reader "tde/internal/microservices/config-reader"
	load_balancer "tde/internal/microservices/load-balancer"
	"tde/internal/microservices/logger"
	service_discovery "tde/internal/microservices/service-discovery"

	"fmt"
	"net/http"
)

var (
	log    = logger.NewLogger("api-gateway/forwarders/customs")
	config *config_reader.Config
	sd     *service_discovery.ServiceDiscovery
	lb     *load_balancer.LoadBalancer
)

func Register(config_ *config_reader.Config, sd_ *service_discovery.ServiceDiscovery) {
	config = config_
	sd = sd_
	var hosts = sd.Customs.ListPrivateIPs()
	if len(hosts) == 0 {
		log.Fatalf("Not enough servers found for '%s' services\n", "customs")
	}
	lb = load_balancer.New(sd, "customs", config.Customs.RouterPrivate, "/api/v1.0.0/customs", "")
}

func Forwarder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Redirecting request to", r.Host, r.URL.Path)
	lb.Forward(w, r)

}
