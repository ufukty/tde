package traverse

import (
	"slices"
)

//go:generate stringer -type constraint

type constraint int

const (
	file                 = constraint(iota) // only the Decls are allowed
	controlFlowCondition                    // while, for, if, else if, switch
	controlFlowInit                         // AssignStmt
	typeDefinition                          //
	funcTypeDefinition                      // io parameter list + receivers
)

type constraints []constraint

func (tc *constraints) Clone(c ...constraint) constraints {
	n := slices.Clone(*tc)
	n = append(n, c...)
	return n
}

func (tc *constraints) Print() (s string) {
	for i, c := range *tc {
		if i != 0 {
			s += ", "
		}
		s += c.String()
	}
	return
}
