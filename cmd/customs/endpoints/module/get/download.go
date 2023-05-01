package module

import (
	"net/http"
	"tde/cmd/customs/internal/volume_manager"
)

const (
	HTTPErrBadRequest = "Bad request"
)

const (
	maxAllowedContentLength = 40 * 1024 * 1024
)

var volumeManager *volume_manager.VolumeManager

func RegisterVolumeManager(vm *volume_manager.VolumeManager) {
	volumeManager = vm
}

func Handler(w http.ResponseWriter, r *http.Request) {

}
