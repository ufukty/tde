package main

import (
	"tde/cmd/customs/endpoints"
	"tde/cmd/customs/endpoints/volmng"
	"tde/config"
	"tde/internal/microservices/cfgreader"
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
		cfg = cfgreader.GetConfig()
		vm  = volmng.NewVolumeManager(cfg.Customs.MountPath)
		em  = endpoints.NewManager(vm)
	)

	// dbo.Connect()
	// defer dbo.Close()

	cfgreader.Print(cfg.Customs)

	var handlers = map[paths.Endpoint]http.HandlerFunc{
		config.CustomsModuleUpload:      em.UploadHandler(),
		config.CustomsModuleDownload:    em.DownloadHandler(),
		config.CustomsModuleList:        em.ListHandler(),
		config.CustomsModuleAstFuncDecl: http.NotFound,
		config.CustomsModuleAstFile:     http.NotFound,
		config.CustomsModuleAstPackage:  em.AstPackageHandler(),
		config.CustomsModuleContext:     em.ContextHandler(),
	}

	router.StartRouter(":"+cfg.Customs.RouterPrivate, &cfg.Customs.RouterParameters, paths.RouteRegisterer(handlers))
	router.Wait(&cfg.Customs.RouterParameters)
}
