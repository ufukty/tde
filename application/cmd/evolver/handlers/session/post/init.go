package session_post

import (
	customs_proxy "tde/cmd/evolver/internal/customs-proxy"
	runner_communicator "tde/cmd/evolver/internal/runner-communicator"
	"tde/cmd/evolver/internal/sessions"
)

var (
	runnerCommunicator *runner_communicator.RunnerCommunicator
	sessionStore       *sessions.Store
	customsProxyCache  *customs_proxy.Cache
)

func Register(rc *runner_communicator.RunnerCommunicator, ss *sessions.Store, cpc *customs_proxy.Cache) {
	runnerCommunicator = rc
	sessionStore = ss
	customsProxyCache = cpc
}
