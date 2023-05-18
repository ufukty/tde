package flow_control

import (
	"tde/internal/microservices/state-machine"

	"fmt"
)

const ( // FIXME: Move each into appropriate endpoint's package
	Init = state_machine.State(iota)
	SentToRunner
	InReplayQueue
)

const ( // FIXME:
	RunnerReturnsInProgress = state_machine.Action(iota)
	RunnerReturnsInternalError
)

func CreateStateMachine() {

	config := state_machine.StateMachineConfig{

		Init: {
			RunnerReturnsInProgress:    SentToRunner,
			RunnerReturnsInternalError: InReplayQueue,
		},

		SentToRunner: {
			RunnerReturnsInProgress:    Init,
			RunnerReturnsInternalError: SentToRunner,
		},

		InReplayQueue: {
			RunnerReturnsInProgress: InReplayQueue,
		},
	}

	var machine = state_machine.NewStateMachine(config)

	var req1 = machine.NewCase(Init)
	if !machine.Translate(req1, RunnerReturnsInProgress) {
		fmt.Println("could not perform action")
	}
	if !machine.Translate(req1, RunnerReturnsInProgress) {
		fmt.Println("could not perform action")
	}
}
