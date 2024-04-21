package detailed

import (
	"fmt"
)

type DetailedError struct {
	Base    string // logs & response
	Details string // logs
}

func New(base, detail string) *DetailedError {
	return &DetailedError{
		Base:    base,
		Details: detail,
	}
}

func (e *DetailedError) Error() string {
	return e.Base
}

func (e *DetailedError) Log() string {
	return fmt.Sprintf("%s (%s)", e.Base, e.Details)
}

func AddBase(e error, base string) *DetailedError {
	return &DetailedError{
		Base:    base,
		Details: e.Error(),
	}
}

func AddDetail(e error, detail string) *DetailedError {
	return &DetailedError{
		Base:    e.Error(),
		Details: detail,
	}
}
