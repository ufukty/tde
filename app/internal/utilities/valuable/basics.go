// provides data types complies [flag.Value] interface
// to make flag package to work with multiple values
package valuable

import "strings"

type Strings []string

func (s *Strings) String() string {
	return strings.Join(*s, ", ")
}

func (s *Strings) Set(v string) error {
	*s = append(*s, v)
	return nil
}
