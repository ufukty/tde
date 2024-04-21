package main

import (
	"tde/config/reader"
	"tde/internal/microservices/router"
	"tde/internal/microservices/serviced"

	"github.com/gorilla/mux"
)

func main() {
	var (
		cfg = reader.GetConfig()
		sd  = serviced.NewServiceDiscovery(cfg.APIGateway.ServiceDiscoveryConfig, cfg.APIGateway.ServiceDiscoveryUpdatePeriod)
	)

	reader.Print(cfg.APIGateway)

	router.StartRouter(":"+cfg.APIGateway.RouterPublic, &cfg.APIGateway.RouterParameters, func(r *mux.Router) {
		r = r.UseEncodedPath()
		RegisterForwarders(sd, cfg, r.PathPrefix("/api/v1.0.0").Subrouter())
	})

	router.Wait(&cfg.APIGateway.RouterParameters)
}
