package loggable

type Loggable interface {
	error
	Log() string
}
