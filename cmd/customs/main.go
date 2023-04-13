package main

import (
	"tde/internal/router"

	"github.com/gorilla/mux"
)

// TODO: accepts uploaded files
// TODO: , puts in sandboxed directory
// TODO: , checks against maliciousness
// TODO: , transforms into AST
// TODO: uploads AST to evolver server when requested

func main() {
	router.StartRouter(":8080", func(r *mux.Router) {
		r.HandleFunc("/upload", uploadFile)
	})
}
