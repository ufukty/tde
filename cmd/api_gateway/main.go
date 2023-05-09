package main

import (
	handler_evolve "tde/cmd/api_gateway/handlers/evolve"
	"tde/internal/microservices/config_reader"
	"tde/internal/router"

	"github.com/gorilla/mux"
)

func main() {
	var (
		config = config_reader.GetConfig()
	)

	router.StartRouter(config.APIGateway.RouterPublic, func(r *mux.Router) {
		sub := r.PathPrefix("/api").Subrouter()

		sub.PathPrefix("/evolve").HandlerFunc(handler_evolve.Handler)
	})

	router.Wait(config.APIGateway.GracePeriod)
}
