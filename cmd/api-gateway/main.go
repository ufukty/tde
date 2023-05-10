package main

import (
	config_reader "tde/internal/microservices/config-reader"
	"tde/internal/router"

	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

func RedirectClosure(targetURL string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		yaml.NewEncoder(os.Stdout).Encode(r)
	}
}

func RegisterRedirect(router *mux.Router, src, target string) {
	router.PathPrefix(src).HandlerFunc(RedirectClosure(target))
}

func main() {
	var (
		config = config_reader.GetConfig()
	)

	router.StartRouter(config.APIGateway.RouterPublic, func(r *mux.Router) {
		sub := r.PathPrefix("/api").Subrouter()

		RegisterRedirect(sub, "/api/customs", config.Customs.RouterPublic)

		sub.PathPrefix("/evolve").HandlerFunc(RedirectClosure("newURI"))
	})

	router.Wait(config.APIGateway.GracePeriod)
}
