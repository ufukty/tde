package state_machine

import "github.com/google/uuid"

type CaseID string

func NewCaseID() CaseID {
	return CaseID(uuid.NewString())
}

//go:generate stringer -type state_machine.State .

type (
	State  int
	Action int

	Transitions        map[Action]State      // Action -> Next state
	StateMachineConfig map[State]Transitions // Current State -> Transitions

	StateMachine struct {
		cases  map[CaseID]State
		config StateMachineConfig
	}
)

func NewStateMachine(config StateMachineConfig) *StateMachine {
	return &StateMachine{
		cases:  map[CaseID]State{},
		config: config,
	}
}

func (machine *StateMachine) NewCase(initState State) CaseID {
	caseID := NewCaseID()
	machine.cases[caseID] = initState
	return caseID
}

func (machine *StateMachine) Translate(caseID CaseID, action Action) (success bool) {
	var (
		current     State
		next        State
		ok          bool
		transitions Transitions
	)

	if current, ok = machine.cases[caseID]; !ok {
		return false
	}

	if transitions, ok = machine.config[current]; !ok {
		return false
	}

	if next, ok = transitions[action]; !ok {
		return false
	}

	machine.cases[caseID] = next

	return true
}
