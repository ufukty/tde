package main

import (
	ast_get "tde/cmd/customs/endpoints/ast/get"
	module_get "tde/cmd/customs/endpoints/module/get"
	module_post "tde/cmd/customs/endpoints/module/post"
	"tde/cmd/customs/internal/volume-manager"
	config_reader "tde/internal/microservices/config-reader"
	"tde/internal/router"

	"github.com/gorilla/mux"
)

// TODO: accepts uploaded files
// TODO: , puts in sandboxed directory
// TODO: , checks against maliciousness
// TODO: , transforms into AST
// TODO: uploads AST to evolver server when requested

func main() {
	var (
		config        = config_reader.GetConfig()
		volumeManager = volume_manager.NewVolumeManager(config.Customs.MountPath)
	)

	module_post.RegisterVolumeManager(volumeManager)
	module_get.RegisterVolumeManager(volumeManager)
	ast_get.RegisterVolumeManager(volumeManager)

	router.StartRouter(config.Customs.RouterPrivate, func(r *mux.Router) {
		r.PathPrefix("/module").Methods("POST").HandlerFunc(module_post.Handler)
		r.PathPrefix("/module").Methods("GET").HandlerFunc(module_get.Handler)
		r.PathPrefix("/ast").Methods("GET").HandlerFunc(ast_get.Handler)
	})

	router.Wait(config.Customs.GracePeriod)
}
