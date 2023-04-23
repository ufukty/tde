package main

import (
	"fmt"
	"net/http"
	"tde/internal/router"
	models "tde/models/transfer"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	req := models.RunnerService_NewTest_Request{}
	req.ParseRequest(r)

	fmt.Println("****\nNew request")
	spew.Println(req)

	res := models.RunnerService_NewTest_Response{
		Registered: true,
	}

	res.SerializeIntoResponseWriter(w)
}

func main() {
	router.StartRouter(":8081", func(r *mux.Router) {
		// r.HandleFunc("/build", router.NotFound) // respond directory listing with 404
		// r.PathPrefix("/build").Handler(http.StripPrefix("/build", http.FileServer(http.Dir("build"))))

		r.PathPrefix("/").HandlerFunc(handler)
	})
}
