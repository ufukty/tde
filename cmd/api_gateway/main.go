package main

import (
	handler_evolve "tde/cmd/api_gateway/handlers/evolve"
	"tde/internal/router"

	"github.com/gorilla/mux"
)

func main() {
	router.StartRouter(":80", func(r *mux.Router) {
		sub := r.PathPrefix("/api").Subrouter()

		sub.PathPrefix("/evolve").HandlerFunc(handler_evolve.Handler)
	})

	router.Wait()
}
