package main

import (
	"tde/config/reader"
	"tde/internal/microservices/router"

	"github.com/gorilla/mux"
)

func main() {
	var (
		config = reader.GetConfig()
	)

	reader.Print(config.Captcha)
	router.StartRouter(config.Captcha.RouterPublic, &config.Captcha.RouterParameters, func(r *mux.Router) {
		r.PathPrefix("/").Methods("GET").HandlerFunc(router.NotFound)
		r.PathPrefix("/").Methods("POST").HandlerFunc(router.NotFound)
	})

	router.Wait(&config.Captcha.RouterParameters)
}
