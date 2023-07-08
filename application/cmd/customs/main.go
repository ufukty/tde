package main

import (
	volume_manager "tde/cmd/customs/internal/volume-manager"
	"tde/config"
	config_reader "tde/internal/microservices/config-reader"
	"tde/internal/microservices/paths"
	"tde/internal/microservices/router"

	mastg "tde/cmd/customs/endpoints/module/ast/package/get"
	mg "tde/cmd/customs/endpoints/module/get"
	mp "tde/cmd/customs/endpoints/module/post"
)

// TODO: accepts uploaded files
// TODO: , puts in sandboxed directory
// TODO: , checks against maliciousness
// TODO: , transforms into AST
// TODO: uploads AST to evolver server when requested

func main() {
	var (
		cfg           = config_reader.GetConfig()
		volumeManager = volume_manager.NewVolumeManager(cfg.Customs.MountPath)
	)

	// dbo.Connect()
	// defer dbo.Close()

	config_reader.Print(cfg.Customs)
	mp.RegisterVolumeManager(volumeManager)
	mg.RegisterVolumeManager(volumeManager)
	mastg.RegisterVolumeManager(volumeManager)

	router.StartRouter(cfg.Customs.RouterPrivate, &cfg.Customs.RouterParameters,
		paths.RouteRegisterer(config.Handlers[config.Customs]),
	)
	router.Wait(&cfg.Customs.RouterParameters)
}
