package command

import (
	"strings"
)

type MultiString []string

func (s *MultiString) Set(v string) error {
	*s = append(*s, v)
	return nil
}

func (s *MultiString) String() string {
	return strings.Join(*s, "\n")
}
