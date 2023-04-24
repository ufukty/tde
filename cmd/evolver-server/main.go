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
	w.Write([]byte("Not found."))
}

func main() {
	rc, err := runner_communicator.NewRunnerCommunicator()
	if err != nil {
		utilities.Terminate(errors.Wrap(err, "failed on launch"))
	}

	evolution.RegisterRunnerCommunicator(rc)

	router.StartRouter(SERVER_ADDRESS, func(r *mux.Router) {
		r.PathPrefix("/evolution").Methods("POST").HandlerFunc(evolution.Handler)
		r.PathPrefix("/results").Methods("POST").HandlerFunc(handler)
		r.PathPrefix("/").HandlerFunc(handler)
	})
}
