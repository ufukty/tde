package evolution

import "tde/cmd/evolver-server/internal/runner_communicator"

var rc *runner_communicator.RunnerCommunicator

func RegisterRunnerCommunicator(runner_communicator *runner_communicator.RunnerCommunicator) {
	rc = runner_communicator
}
