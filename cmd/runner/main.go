package main

import (
	"net/http"
	"tde/cmd/runner/controllers/batch"
	"tde/internal/router"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router.StartRouter(":8082", func(r *mux.Router) {
		r.PathPrefix("/batch").Methods("POST").HandlerFunc(batch.Handler)
	})
}
