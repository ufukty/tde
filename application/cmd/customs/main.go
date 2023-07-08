package main

import (
	"tde/cmd/customs/endpoints"
	"tde/cmd/customs/endpoints/volmng"
	"tde/config"
	config_reader "tde/internal/microservices/config-reader"
	"tde/internal/microservices/paths"
	"tde/internal/microservices/router"

	"net/http"
)

// TODO: accepts uploaded files
// TODO: , puts in sandboxed directory
// TODO: , checks against maliciousness
// TODO: , transforms into AST
// TODO: uploads AST to evolver server when requested

func main() {
	var (
		cfg = config_reader.GetConfig()
		vm  = volmng.NewVolumeManager(cfg.Customs.MountPath)
		em  = endpoints.NewManager(vm)
	)

	// dbo.Connect()
	// defer dbo.Close()

	config_reader.Print(cfg.Customs)

	var handlers = map[paths.Endpoint]http.HandlerFunc{
		config.CustomsModuleUpload:      em.HandleUpload,
		config.CustomsModuleDownload:    em.HandleDownload,
		config.CustomsModuleList:        em.HandleList,
		config.CustomsModuleAstFuncDecl: http.NotFound,
		config.CustomsModuleAstFile:     http.NotFound,
		config.CustomsModuleAstPackage:  em.HandleAstPackage,
		config.CustomsModuleContext:     em.HandleContext,
	}

	router.StartRouter(cfg.Customs.RouterPrivate, &cfg.Customs.RouterParameters, paths.RouteRegisterer(handlers))
	router.Wait(&cfg.Customs.RouterParameters)
}
