package main

import (
	"tde/internal/microservices/cfgreader"
	"tde/internal/microservices/router"
	"tde/internal/microservices/serviced"

	"github.com/gorilla/mux"
)

func main() {
	var (
		cfg = cfgreader.GetConfig()
		sd  = serviced.NewServiceDiscovery(cfg.APIGateway.ServiceDiscoveryConfig, cfg.APIGateway.ServiceDiscoveryUpdatePeriod)
	)

	cfgreader.Print(cfg.APIGateway)

	router.StartRouter(":"+cfg.APIGateway.RouterPublic, &cfg.APIGateway.RouterParameters, func(r *mux.Router) {
		r = r.UseEncodedPath()
		RegisterForwarders(sd, cfg, r.PathPrefix("/api/v1.0.0").Subrouter())
	})

	router.Wait(&cfg.APIGateway.RouterParameters)
}
