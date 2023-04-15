package main

import (
	"tde/internal/evolution"
	"tde/internal/router"

	"context"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO: Listen endpoints for:
//   TODO: Creating a generation (request from client)
//   TODO: Download test results from runner

func GetParametersForUser() {}

func StartEvolution(input string, test string, populationSize int) { // FIXME: actual input is a package, contains tests and target function headers along with irrelevant code
	e := evolution.Evolution{}
	e.InitPopulation(populationSize)
	e.IterateLoop(context.Background())
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ready and loaded"))
}

func main() {
	router.StartRouter(":8081", func(r *mux.Router) {
		// r.HandleFunc("/build", router.NotFound) // respond directory listing with 404
		// r.PathPrefix("/build").Handler(http.StripPrefix("/build", http.FileServer(http.Dir("build"))))

		r.PathPrefix("/").HandlerFunc(handler)
	})

	// server.NewServer(6000)

	// t := Testing{}

	// provisionEnvironmentImage [x]
	// run tests

	// fitness := t.Calculate()
	// fmt.Println(fitness)

}
