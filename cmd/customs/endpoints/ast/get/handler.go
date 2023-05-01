package ast

import (
	"net/http"
	"tde/cmd/customs/internal/volume_manager"
)

var volumeManager *volume_manager.VolumeManager

func RegisterVolumeManager(vm *volume_manager.VolumeManager) {
	volumeManager = vm
}

func Handler(w http.ResponseWriter, r *http.Request) {

}
