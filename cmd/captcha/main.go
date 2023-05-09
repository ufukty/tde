package main

import (
	"tde/internal/microservices/config_reader"
	"tde/internal/router"

	"github.com/gorilla/mux"
)

func main() {
	var (
		config = config_reader.GetConfig()
	)

	router.StartRouter(config.Captcha.RouterPublic, func(r *mux.Router) {
		r.PathPrefix("/").Methods("GET").HandlerFunc(router.NotFound)
		r.PathPrefix("/").Methods("POST").HandlerFunc(router.NotFound)
	})

	router.Wait(config.Captcha.GracePeriod)
}
