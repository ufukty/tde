package evolution

import (
	"tde/cmd/evolver-server/internal/runner_communicator"
)

var (
	rc *runner_communicator.RunnerCommunicator
	// replayQueue *queue.Replay[dto.EvolverService_Evolve_Request]
)

func RegisterRunnerCommunicator(runner_communicator *runner_communicator.RunnerCommunicator) {
	rc = runner_communicator
	// replayQueue = queue.NewReplayQueue(dequeueRequest)
}

// func queueRequest(req dto.EvolverService_Evolve_Request) {

// }

// func dequeueRequest(es dto.EvolverService_Evolve_Request) {
// 	return queue.DONE
// }
