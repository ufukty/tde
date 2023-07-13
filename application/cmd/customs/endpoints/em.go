package endpoints

import (
	"tde/cmd/customs/endpoints/volmng"
	"tde/internal/microservices/logger"

	"net/http"
)

const ResponseMalformedRequest = "Request is malformed"

type EndpointsManager struct {
	vm  *volmng.VolumeManager
	log *logger.Logger
}

func NewManager(vm *volmng.VolumeManager) *EndpointsManager {
	return &EndpointsManager{
		vm:  vm,
		log: logger.NewLogger("customs/endpoints/module/ast/get"),
	}
}

func (em EndpointsManager) HandleAstFile(w http.ResponseWriter, r *http.Request) {

}
