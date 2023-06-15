package session_post

import (
	case_manager "tde/cmd/evolver/internal/case-manager"
	customs_proxy "tde/cmd/evolver/internal/customs-proxy"
	runner_communicator "tde/cmd/evolver/internal/runner-communicator"
)

var (
	runnerCommunicator *runner_communicator.RunnerCommunicator
	caseManager        *case_manager.CaseManager
	customsProxyCache  *customs_proxy.Cache
)

func Register(rc *runner_communicator.RunnerCommunicator, cm *case_manager.CaseManager, cpc *customs_proxy.Cache) {
	runnerCommunicator = rc
	caseManager = cm
	customsProxyCache = cpc
}
