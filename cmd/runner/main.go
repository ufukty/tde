package main

import (
	"tde/cmd/runner/controllers/batch"
	"tde/internal/microservices/config_reader"
	"tde/internal/router"

	"github.com/gorilla/mux"
)

func main() {
	var (
		config = config_reader.GetConfig()
	)

	router.StartRouter(config.Runner.RouterPublic, func(r *mux.Router) {
		r.PathPrefix("/batch").Methods("POST").HandlerFunc(batch.Handler)
		r.HandleFunc("/", router.NotFound)
	})

	router.Wait(config.Runner.GracePeriod)
}
