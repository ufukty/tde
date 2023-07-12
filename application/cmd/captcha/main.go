package main

import (
	"tde/internal/microservices/cfgreader"
	"tde/internal/microservices/router"

	"github.com/gorilla/mux"
)

func main() {
	var (
		config = cfgreader.GetConfig()
	)

	cfgreader.Print(config.Captcha)
	router.StartRouter(config.Captcha.RouterPublic, &config.Captcha.RouterParameters, func(r *mux.Router) {
		r.PathPrefix("/").Methods("GET").HandlerFunc(router.NotFound)
		r.PathPrefix("/").Methods("POST").HandlerFunc(router.NotFound)
	})

	router.Wait(&config.Captcha.RouterParameters)
}
