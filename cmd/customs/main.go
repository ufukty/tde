package main

import (
	"flag"
	"log"
	handler_upload "tde/cmd/customs/handlers/upload"
	"tde/cmd/customs/internal/volume_manager"
	"tde/internal/router"

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

	handler_upload.RegisterVolumeManager(volumeManager)

	router.StartRouter(":8083", func(r *mux.Router) {
		r.HandleFunc("/upload", handler_upload.Handler)
	})
}
