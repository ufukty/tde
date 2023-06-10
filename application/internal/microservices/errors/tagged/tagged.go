package tagged

import (
	"tde/internal/microservices/errors/loggable"

	"fmt"
)

type TaggedError struct {
	tag        loggable.Loggable
	underlying loggable.Loggable
}

func New(tag, underlying loggable.Loggable) *TaggedError {
	return &TaggedError{
		tag:        tag,
		underlying: underlying,
	}
}

func (t *TaggedError) Tag(another loggable.Loggable) *TaggedError {
	return New(another, t)
}

func (t *TaggedError) Error() string {
	return fmt.Sprintf("%s: %s", t.tag.Error(), t.underlying.Error())
}

func (t *TaggedError) Log() string {
	return fmt.Sprintf("%s: %s", t.tag.Log(), t.underlying.Log())
}
