package main

import (
	"tde/cmd/runner/endpoints/batch/batch_post"
	"tde/config/reader"
	"tde/internal/microservices/router"

	"github.com/gorilla/mux"
)

func main() {
	var (
		config = reader.GetConfig()
	)

	reader.Print(config.Runner)

	router.StartRouter(config.Runner.RouterPublic, &config.Runner.RouterParameters, func(r *mux.Router) {
		r.PathPrefix("/batch").Methods("POST").HandlerFunc(batch_post.Handler)
		r.HandleFunc("/", router.NotFound)
	})

	router.Wait(&config.Runner.RouterParameters)
}
