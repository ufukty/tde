package main

import (
	session_post "tde/cmd/evolver/handlers/session/post"
	test_results_post "tde/cmd/evolver/handlers/session/test-results/post"
	case_manager "tde/cmd/evolver/internal/case-manager"
	runner_communicator "tde/cmd/evolver/internal/runner-communicator"
	config_reader "tde/internal/microservices/config-reader"
	"tde/internal/microservices/router"
	service_discovery "tde/internal/microservices/service-discovery"

	"log"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func main() {
	var (
		config = config_reader.GetConfig()
		sd     = service_discovery.NewServiceDiscovery(config.Evolver.ServiceDiscoveryConfig, config.Evolver.ServiceDiscoveryUpdatePeriod)
		cm     = case_manager.NewCaseManager()
	)

	config_reader.Print(config.Evolver)
	rc, err := runner_communicator.NewRunnerCommunicator(sd)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed on launch"))
	}

	session_post.RegisterCaseManager(cm)
	session_post.RegisterRunnerCommunicator(rc)

	router.StartRouter(config.Evolver.RouterPublic, &config.Evolver.RouterParameters, func(r *mux.Router) {
		// r.PathPrefix("/session/hall-of-fame").Methods("GET").HandlerFunc(handler_results.Handler)
		// r.PathPrefix("/session/generation/{:individual-id}").Methods("GET").HandlerFunc(handler_results.Handler)
		// r.PathPrefix("/session/generation").Methods("GET").HandlerFunc(handler_results.Handler)
		r.PathPrefix("/session/test-results").Methods("POST").HandlerFunc(test_results_post.Handler)
		r.PathPrefix("/session").Methods("POST").HandlerFunc(session_post.Handler)
		r.PathPrefix("/").HandlerFunc(router.NotFound)
	})

	router.Wait(&config.Evolver.RouterParameters)
}
