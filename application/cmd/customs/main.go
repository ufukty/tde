package main

import (
	ast_module_get "tde/cmd/customs/endpoints/ast/module/get"
	module_get "tde/cmd/customs/endpoints/module/get"
	module_post "tde/cmd/customs/endpoints/module/post"
	volume_manager "tde/cmd/customs/internal/volume-manager"
	config_reader "tde/internal/microservices/config-reader"
	"tde/internal/microservices/router"

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

	// dbo.Connect()
	// defer dbo.Close()

	config_reader.Print(config.Customs)
	module_post.RegisterVolumeManager(volumeManager)
	module_get.RegisterVolumeManager(volumeManager)
	ast_module_get.RegisterVolumeManager(volumeManager)

	router.StartRouter(config.Customs.RouterPrivate, &config.Customs.RouterParameters, func(r *mux.Router) {
		r.PathPrefix("/module").Methods("POST").HandlerFunc(module_post.Handler)
		r.PathPrefix("/module").Methods("GET").HandlerFunc(module_get.Handler)
		r.PathPrefix("/ast/module").Methods("GET").HandlerFunc(ast_module_get.Handler)
		r.PathPrefix("/ast/module/package").Methods("GET").HandlerFunc(ast_module_get.Handler)
	})

	router.Wait(&config.Customs.RouterParameters)
}
