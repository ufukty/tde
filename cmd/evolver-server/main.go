package main

import (
	"tde/cmd/evolver-server/handlers/evolution"
	"tde/cmd/evolver-server/internal/runner_communicator"
	"tde/internal/router"
	"tde/internal/utilities"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const SERVER_ADDRESS = ":8081"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Evolver is ready and loaded"))
}

func main() {
	rc, err := runner_communicator.NewRunnerCommunicator()
	if err != nil {
		utilities.Terminate(errors.Wrap(err, "failed on launch"))
	}

	router.StartRouter(SERVER_ADDRESS, func(r *mux.Router) {
		// r.HandleFunc("/build", router.NotFound) // respond directory listing with 404
		// r.PathPrefix("/build").Handler(http.StripPrefix("/build", http.FileServer(http.Dir("build"))))

		r.PathPrefix("/evolution").HandlerFunc(evolution.Handler)
		r.PathPrefix("/results").HandlerFunc(handler)
	})
}
