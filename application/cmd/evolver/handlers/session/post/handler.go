package session_post

import (
	"errors"
	"fmt"
	"net/http"
	customs_proxy "tde/cmd/evolver/internal/customs-proxy"
	runner_communicator "tde/cmd/evolver/internal/runner-communicator"
	sessions "tde/cmd/evolver/internal/sessions"
	"tde/internal/microservices/logger"
)

var (
	ErrTargetInaccessible = errors.New("Target service is unaccassable or can not accept the request at the moment")
)

var (
	runnerCommunicator *runner_communicator.RunnerCommunicator
	sessionStore       *sessions.Store
	customsProxyCache  *customs_proxy.Cache
	log                = logger.NewLogger("endpoint/session/post")
)

func Register(rc *runner_communicator.RunnerCommunicator, ss *sessions.Store, cpc *customs_proxy.Cache) {
	runnerCommunicator = rc
	sessionStore = ss
	customsProxyCache = cpc
}

func Handler(w http.ResponseWriter, r *http.Request) {
	reqDTO := &Request{}
	err := reqDTO.ParseRequest(r)
	if err != nil {
		fmt.Fprintln(w, "Error")
		return
	}

	resDTO, bucket := Controller(reqDTO)
	if bucket.IsAny() {
		log.Println(bucket.Log())
		http.Error(w, bucket.Error(), http.StatusInternalServerError)
	}
	err = resDTO.SerializeIntoResponseWriter(w)
	if err != nil {
		fmt.Fprintln(w, "Error")
		// if errors.Is(err, ErrTargetInaccessible) {
		// 	queueRequest(reqDTO)
		// }
		return
	}

}
