package main

import (
	"tde/cmd/runner/endpoints/batch/batch_post"
	config_reader "tde/internal/microservices/config-reader"
	"tde/internal/microservices/router"

	"github.com/gorilla/mux"
)

func main() {
	var (
		config = config_reader.GetConfig()
	)

	config_reader.Print(config.Runner)

	router.StartRouter(config.Runner.RouterPublic, &config.Runner.RouterParameters, func(r *mux.Router) {
		r.PathPrefix("/batch").Methods("POST").HandlerFunc(batch_post.Handler)
		r.HandleFunc("/", router.NotFound)
	})

	router.Wait(&config.Runner.RouterParameters)
}
