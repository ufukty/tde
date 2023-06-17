package main

import (
	session_post "tde/cmd/evolver/handlers/session/post"
	test_results_post "tde/cmd/evolver/handlers/session/test-results/post"
	customs_proxy "tde/cmd/evolver/internal/customs-proxy"
	runner_communicator "tde/cmd/evolver/internal/runner-communicator"
	"tde/cmd/evolver/internal/sessions"
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
		sm     = sessions.NewStore()
		cpc    = customs_proxy.New(config, sd)
	)

	config_reader.Print(config.Evolver)
	rc, err := runner_communicator.NewRunnerCommunicator(sd)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed on launch"))
	}

	session_post.Register(rc, sm, cpc)

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
