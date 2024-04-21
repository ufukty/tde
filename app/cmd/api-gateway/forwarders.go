package main

import (
	"tde/config"
	"tde/config/reader"
	"tde/internal/microservices/forwarder"
	"tde/internal/microservices/serviced"
	"tde/internal/microservices/serviced/models/services"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func RegisterForwarders(sd *serviced.ServiceDiscovery, cfg *reader.Config, sub *mux.Router) {
	var (
		err         error
		fwdsCustoms http.HandlerFunc
		fwdsEvolver http.HandlerFunc
	)

	fwdsCustoms, err = forwarder.NewLoadBalancedProxy(sd, services.Customs, cfg.Customs.RouterPrivate, config.Customs.PathByGateway())
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Creating forwarder for Customs"))
	}
	fwdsEvolver, err = forwarder.NewLoadBalancedProxy(sd, services.Evolver, cfg.Evolver.RouterPrivate, config.Evolver.PathByGateway())
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Creating forwarder for Evolver"))
	}

	sub.PathPrefix("/customs").HandlerFunc(fwdsCustoms)
	sub.PathPrefix("/evolve").HandlerFunc(fwdsEvolver)
}
