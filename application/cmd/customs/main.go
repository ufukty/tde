package main

import (
	"net/http"
	volume_manager "tde/cmd/customs/internal/volume-manager"
	config_reader "tde/internal/microservices/config-reader"
	"tde/internal/microservices/router"

	module_package_get "tde/cmd/customs/endpoints/module/ast/package/get"
	module_get "tde/cmd/customs/endpoints/module/get"
	module_package_list_get "tde/cmd/customs/endpoints/module/list/get"
	module_post "tde/cmd/customs/endpoints/module/post"

	"github.com/gorilla/mux"
)

// TODO: accepts uploaded files
// TODO: , puts in sandboxed directory
// TODO: , checks against maliciousness
// TODO: , transforms into AST
// TODO: uploads AST to evolver server when requested

type Endpoint interface {
	Register(*mux.Router)
	Handler(http.ResponseWriter, http.Request)
}

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
	module_package_get.RegisterVolumeManager(volumeManager)

	router.StartRouter(config.Customs.RouterPrivate, &config.Customs.RouterParameters, func(r *mux.Router) {
		r.PathPrefix("/module/{id}/ast/package/{package}/file/{file}/function/{function}").Methods("GET").HandlerFunc(module_package_get.Handler)
		r.PathPrefix("/module/{id}/ast/package/{package}/file/{file}").Methods("GET").HandlerFunc(module_package_get.Handler)
		r.PathPrefix("/module/{id}/ast/package/{package}").Methods("GET").HandlerFunc(module_package_get.Handler)
		r.PathPrefix("/module/{id}/context/package/{package}/file/{file}").Methods("GET").HandlerFunc(module_package_get.Handler)
		r.PathPrefix("/module/{id}/list").Methods("GET").HandlerFunc(module_package_list_get.Handler)
		r.PathPrefix("/module/{id}").Methods("GET").HandlerFunc(module_get.Handler)
		r.PathPrefix("/module").Methods("POST").HandlerFunc(module_post.Handler)
	})

	router.Wait(&config.Customs.RouterParameters)
}
