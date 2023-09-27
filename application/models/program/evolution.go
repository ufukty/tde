package models

type Parameters struct {
	Population  int
	Generations int
	Size        int // max. code size in bytes
	Packages    []string

	Cc int // cap. for code search
	Cp int // cap. for program search
	Cs int // cap. for solution search

	Dc int // max depth for code search
	Dp int // max depth for program search
	Ds int // max depth for solution search

	Rc float64 // reproduction rate for code search
	Rp float64 // reproduction rate for program search
	Rs float64 // reproduction rate for solution search
}

type Layer int

const (
	AST = Layer(iota)
	Code
	Program
	Solution
)

type SearchId string
