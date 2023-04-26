package evolution

import (
	"tde/cmd/evolver-server/internal/case_manager"
	"tde/cmd/evolver-server/internal/runner_communicator"
)

var (
	runnerCommunicator *runner_communicator.RunnerCommunicator
	caseManager        *case_manager.CaseManager
)

func RegisterRunnerCommunicator(rc *runner_communicator.RunnerCommunicator) {
	runnerCommunicator = rc
}

func RegisterCaseManager(cm *case_manager.CaseManager) {
	caseManager = cm
}
