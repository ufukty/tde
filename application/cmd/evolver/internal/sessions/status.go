package sessions

//go:generate stringer -type=Status

type Status int

const (
	Initialized = Status(iota)
	InIteration
)
