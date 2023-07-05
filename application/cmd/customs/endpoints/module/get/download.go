package module

import (
	"tde/cmd/customs/internal/utilities"
	volume_manager "tde/cmd/customs/internal/volume-manager"
	"tde/internal/microservices/logger"

	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var (
	vm  *volume_manager.VolumeManager
	log = logger.NewLogger("customs/endpoints/module/get/handler")
)

func RegisterVolumeManager(vm_ *volume_manager.VolumeManager) {
	vm = vm_
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		ok          bool
		err         error
		digest      string
		archiveId   string
		fileHandler *os.File
		vars        map[string]string
	)

	vars = mux.Vars(r)

	if archiveId, ok = vars["id"]; !ok {
		log.Println(errors.Wrap(err, "Invalid request body"))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var bundleExists, zipExists, extractExists = vm.CheckIfExists(archiveId)
	if !(bundleExists && zipExists && extractExists) {
		log.Printf("Got asked for non-existent archive '%s'\n", archiveId)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var _, zip, _ = vm.FindPath(archiveId)
	fileHandler, err = os.Open(zip)
	if err != nil {
		log.Println(errors.Wrap(err, "opening file to read"))
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	defer fileHandler.Close()

	w.WriteHeader(http.StatusOK)
	http.ServeContent(w, r, "file.zip", time.Now(), fileHandler)

	digest, err = utilities.MD5(fileHandler)
	if err != nil {
		log.Println(errors.Wrap(err, "checking md5sum"))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Digest", fmt.Sprintf("md5=%s", digest))
	w.Header().Set("Content-Disposition", "attachment; filename=file.zip")
}
