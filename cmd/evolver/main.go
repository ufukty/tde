package main

import (
	handler_evolution "tde/cmd/evolver/handlers/evolution"
	handler_results "tde/cmd/evolver/handlers/results"
	"tde/cmd/evolver/internal/case_manager"
	"tde/cmd/evolver/internal/runner_communicator"
	"tde/internal/microservices/service_discovery"
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
	var (
		sd = service_discovery.NewServiceDiscovery()
		cm = case_manager.NewCaseManager()
	)

	rc, err := runner_communicator.NewRunnerCommunicator(sd)
	if err != nil {
		utilities.Terminate(errors.Wrap(err, "failed on launch"))
	}

	handler_evolution.RegisterCaseManager(cm)
	handler_evolution.RegisterRunnerCommunicator(rc)

	router.StartRouter(SERVER_ADDRESS, func(r *mux.Router) {
		r.PathPrefix("/evolution").Methods("POST").HandlerFunc(handler_evolution.Handler)
		r.PathPrefix("/results").Methods("POST").HandlerFunc(handler_results.Handler)
		r.PathPrefix("/").HandlerFunc(handler)
	})
}
