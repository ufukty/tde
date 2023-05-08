package main

import (
	"tde/internal/router"

	"github.com/gorilla/mux"
)

func main() {
	router.StartRouter(":8085", func(r *mux.Router) {
		r.PathPrefix("/").Methods("GET").HandlerFunc(router.NotFound)
		r.PathPrefix("/").Methods("POST").HandlerFunc(router.NotFound)
	})
}
