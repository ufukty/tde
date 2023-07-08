package endpoints

import (
	volume_manager "tde/cmd/customs/endpoints/volmng"
	"tde/internal/microservices/logger"

	"net/http"
)

type EndpointsManager struct {
	vm  *volume_manager.VolumeManager
	log *logger.Logger
}

func NewManager(vm *volume_manager.VolumeManager) *EndpointsManager {
	return &EndpointsManager{
		vm:  vm,
		log: logger.NewLogger("customs/endpoints/module/ast/get"),
	}
}

func (em EndpointsManager) HandleAstFile(w http.ResponseWriter, r *http.Request) {

}

func (em EndpointsManager) HandleAstFuncDecl(w http.ResponseWriter, r *http.Request) {

}
