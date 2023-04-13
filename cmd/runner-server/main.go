package main

import (
	"net/http"
	"tde/internal/router"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ready and loaded"))
}

func main() {
	router.StartRouter(":8081", func(r *mux.Router) {
		// r.HandleFunc("/build", router.NotFound) // respond directory listing with 404
		// r.PathPrefix("/build").Handler(http.StripPrefix("/build", http.FileServer(http.Dir("build"))))

		r.PathPrefix("/").HandlerFunc(handler)
	})
}
