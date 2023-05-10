package main

import (
	handler_evolution "tde/cmd/evolver/handlers/evolution"
	handler_results "tde/cmd/evolver/handlers/results"
	case_manager "tde/cmd/evolver/internal/case-manager"
	runner_communicator "tde/cmd/evolver/internal/runner-communicator"
	config_reader "tde/internal/microservices/config-reader"
	service_discovery "tde/internal/microservices/service-discovery"
	"tde/internal/router"

	"log"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func main() {
	var (
		config = config_reader.GetConfig()
		sd     = service_discovery.NewServiceDiscovery(config.Evolver.ServiceDiscoveryConfig)
		cm     = case_manager.NewCaseManager()
	)

	rc, err := runner_communicator.NewRunnerCommunicator(sd)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed on launch"))
	}

	handler_evolution.RegisterCaseManager(cm)
	handler_evolution.RegisterRunnerCommunicator(rc)

	router.StartRouter(config.Evolver.RouterPublic, func(r *mux.Router) {
		r.PathPrefix("/session/hall-of-fame").Methods("GET").HandlerFunc(handler_results.Handler)
		r.PathPrefix("/session/generation/{:individual-id}").Methods("GET").HandlerFunc(handler_results.Handler)
		r.PathPrefix("/session/generation").Methods("GET").HandlerFunc(handler_results.Handler)
		r.PathPrefix("/session").Methods("POST").HandlerFunc(handler_evolution.Handler)
		r.PathPrefix("/").HandlerFunc(router.NotFound)
	})

	router.StartRouter(config.Evolver.RouterPrivate, func(r *mux.Router) {
		r.PathPrefix("/session/runner-results").Methods("POST").HandlerFunc(handler_results.Handler)
		r.PathPrefix("/").HandlerFunc(router.NotFound)
	})

	router.Wait(config.Evolver.GracePeriod)
}
