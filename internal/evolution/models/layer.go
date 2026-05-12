package models

//go:generate stringer -type Layer
type Layer int

const ( // don't change ordering
	AST       = Layer(0) // Subject stuck printing
	Code      = Layer(1) // Subject stuck compilation
	Program   = Layer(2) // Subject stuck execution
	Candidate = Layer(3) // Subject stuck on assertions
	Solution  = Layer(4)
)
