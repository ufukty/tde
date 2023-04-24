package runner_communicator

type RequestState int

const (
	SENT = RequestState(iota)
	RESPONSE_RECEIVED
)

type CaseID string

type Request struct {
	CaseID   string
	RunnerIP string
	State    RequestState
}