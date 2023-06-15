package customs

import (
	balanced_forwarder "tde/internal/microservices/balanced-forwarder"
	config_reader "tde/internal/microservices/config-reader"
	"tde/internal/microservices/logger"
	service_discovery "tde/internal/microservices/service-discovery"
	"tde/internal/microservices/service-discovery/models/services"

	"net/http"

	"github.com/go-chi/chi/middleware"
)

var (
	log    = logger.NewLogger("api-gateway/forwarders/customs")
	config *config_reader.Config
	sd     *service_discovery.ServiceDiscovery
	bf     *balanced_forwarder.BalancedForwarder
)

func Register(config_ *config_reader.Config, sd_ *service_discovery.ServiceDiscovery) {
	config = config_
	sd = sd_
	var hosts = sd.ListPrivateIPs(services.Customs)
	if len(hosts) == 0 {
		log.Fatalf("Not enough servers found for '%s' services\n", "customs")
	}
	bf = balanced_forwarder.New(sd, "customs", config.Customs.RouterPrivate, "/api/v1.0.0/customs", "")
}

func Forwarder(w http.ResponseWriter, r *http.Request) {
	log.Printf("Redirecting request '%s' to '%s%s'", middleware.GetReqID(r.Context()), r.Host, r.URL.Path)
	bf.Forward(w, r)

}
