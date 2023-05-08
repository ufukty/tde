package main

import (
	ast_get "tde/cmd/customs/endpoints/ast/get"
	module_get "tde/cmd/customs/endpoints/module/get"
	module_post "tde/cmd/customs/endpoints/module/post"
	"tde/cmd/customs/internal/volume_manager"
	"tde/internal/router"

	"flag"
	"log"

	"github.com/gorilla/mux"
)

// TODO: accepts uploaded files
// TODO: , puts in sandboxed directory
// TODO: , checks against maliciousness
// TODO: , transforms into AST
// TODO: uploads AST to evolver server when requested

func main() {

	var mountPath = flag.String("volume", "", "path to the root of destination, expected the attached volume under /mnt")
	flag.Parse()

	if *mountPath == "" {
		log.Fatalln("volume path not found in arg list. read help.")
	}

	var volumeManager = volume_manager.NewVolumeManager(*mountPath)

	module_post.RegisterVolumeManager(volumeManager)
	module_get.RegisterVolumeManager(volumeManager)
	ast_get.RegisterVolumeManager(volumeManager)

	router.StartRouter(":8083", func(r *mux.Router) {
		r.PathPrefix("/module").Methods("POST").HandlerFunc(module_post.Handler)
		r.PathPrefix("/module").Methods("GET").HandlerFunc(module_get.Handler)
		r.PathPrefix("/ast").Methods("GET").HandlerFunc(ast_get.Handler)
	})

	router.Wait()
}
