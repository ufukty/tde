package endpoints

import (
	"net/http"
	volume_manager "tde/cmd/customs/endpoints/volmng"
	"tde/internal/microservices/logger"
)

type Handlers struct {
	vm  *volume_manager.VolumeManager
	log *logger.Logger
}

func New(vm *volume_manager.VolumeManager) *Handlers {
	return &Handlers{
		vm:  vm,
		log: logger.NewLogger("customs/endpoints/module/ast/get"),
	}
}

func (h Handlers) HandleAstFile(w http.ResponseWriter, r *http.Request) {

}

func (h Handlers) HandleAstFuncDecl(w http.ResponseWriter, r *http.Request) {

}
