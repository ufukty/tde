package main

import (
	"tde/cmd/api-gateway/forwarders/customs"
	"tde/cmd/api-gateway/forwarders/evolver"
	config_reader "tde/internal/microservices/config-reader"
	"tde/internal/microservices/router"
	service_discovery "tde/internal/microservices/service-discovery"

	"github.com/gorilla/mux"
)

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func main() {
	var (
		config = config_reader.GetConfig()
		sd     = service_discovery.NewServiceDiscovery(config.APIGateway.ServiceDiscoveryConfig, config.APIGateway.ServiceDiscoveryUpdatePeriod)
	)

	config_reader.Print(config.APIGateway)
	customs.Register(config, sd)

	router.StartRouter(config.APIGateway.RouterPublic, &config.APIGateway.RouterParameters, func(r *mux.Router) {
		sub := r.PathPrefix("/api/v1.0.0").Subrouter()
		sub.PathPrefix("/customs").HandlerFunc(customs.Forwarder)
		sub.PathPrefix("/evolve").HandlerFunc(evolver.Forwarder)
		sub.PathPrefix("/ping").HandlerFunc(pong)
		sub.PathPrefix("/").HandlerFunc(router.NotFound)
	})

	router.Wait(&config.APIGateway.RouterParameters)
}
