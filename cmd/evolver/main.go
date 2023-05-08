package main

import (
	handler_evolution "tde/cmd/evolver/handlers/evolution"
	handler_results "tde/cmd/evolver/handlers/results"
	"tde/cmd/evolver/internal/case_manager"
	"tde/cmd/evolver/internal/runner_communicator"
	"tde/internal/microservices/config_reader"
	"tde/internal/microservices/service_discovery"
	"tde/internal/router"
	"tde/internal/utilities"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type Config struct {
	ServiceDiscoveryConfig string `yaml:"service-discovery-config"`
	RouterPublic           string `yaml:"router_public"`
	RouterPrivate          string `yaml:"router_private"`
}

func main() {
	var (
		config = config_reader.FillAndReturn(&Config{})
		sd     = service_discovery.NewServiceDiscovery(config.ServiceDiscoveryConfig)
		cm     = case_manager.NewCaseManager()
	)

	rc, err := runner_communicator.NewRunnerCommunicator(sd)
	if err != nil {
		utilities.Terminate(errors.Wrap(err, "failed on launch"))
	}

	handler_evolution.RegisterCaseManager(cm)
	handler_evolution.RegisterRunnerCommunicator(rc)

	router.StartRouter(":8081", func(r *mux.Router) {
		var public = r.Host(config.RouterPublic)

		public.PathPrefix("/session/hall-of-fame").Methods("GET").HandlerFunc(handler_results.Handler)
		public.PathPrefix("/session/generation/{:individual-id}").Methods("GET").HandlerFunc(handler_results.Handler)
		public.PathPrefix("/session/generation").Methods("GET").HandlerFunc(handler_results.Handler)
		public.PathPrefix("/session").Methods("POST").HandlerFunc(handler_evolution.Handler)
		public.PathPrefix("/").HandlerFunc(router.NotFound)

		var private = r.Host(config.RouterPrivate)
		private.PathPrefix("/session/runner-results").Methods("POST").HandlerFunc(handler_results.Handler)
		public.PathPrefix("/").HandlerFunc(router.NotFound)
	})
}
