package main

import (
	"tde/cmd/runner/controllers/batch"
	"tde/internal/router"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world.")
}

func main() {
	router.StartRouter(":8082", func(r *mux.Router) {
		r.PathPrefix("/batch").Methods("POST").HandlerFunc(batch.Handler)
		r.HandleFunc("/", handler)
	})
}
