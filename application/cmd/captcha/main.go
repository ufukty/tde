package main

import (
	config_reader "tde/internal/microservices/config-reader"
	"tde/internal/router"

	"github.com/gorilla/mux"
)

func main() {
	var (
		config = config_reader.GetConfig()
	)

	config_reader.Print(config.Captcha)
	router.StartRouter(config.Captcha.RouterPublic, &config.Captcha.RouterParameters, func(r *mux.Router) {
		r.PathPrefix("/").Methods("GET").HandlerFunc(router.NotFound)
		r.PathPrefix("/").Methods("POST").HandlerFunc(router.NotFound)
	})

	router.Wait(&config.Captcha.RouterParameters)
}
