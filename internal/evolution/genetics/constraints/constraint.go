package traverse

import (
	"slices"
)

//go:generate stringer -type constraint

type constraint int

const (
	File = constraint(iota) // only the Decls are allowed
	GenDecl
	FuncDecl
	FuncTypeDef     // io parameter list + receivers
	StructTypeDef   //
	ControlFlowCond // while, for, if, else if, switch
	ControlFlowInit // AssignStmt
	ImportSpec
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
