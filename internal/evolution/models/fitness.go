package models

type Fitness struct {
	AST       float64 // rel. printing errors (from ast)
	Code      float64 // rel. syntax errors (compile)
	Program   float64 // rel. runtime errors
	Candidate float64 // rel. passed tests (user-provided)
}

func (f Fitness) Flat() float64 {
	if f.AST != 0.0 {
		return 3.0 + f.AST
	} else if f.Code != 0.0 {
		return 2.0 + f.Code
	} else if f.Program != 0.0 {
		return 1.0 + f.Program
	} else {
		return f.Candidate
	}
}

func (f Fitness) Layer() Layer {
	if f.AST != 0.0 {
		return AST
	} else if f.Code != 0.0 {
		return Code
	} else if f.Program != 0.0 {
		return Program
	} else if f.Candidate != 0.0 {
		return Candidate
	} else {
		return Solution
	}
}

func (f Fitness) InLayer(layer Layer) float64 {
	switch layer {
	case AST:
		return f.AST
	case Code:
		return f.Code
	case Program:
		return f.Program
	case Candidate:
		return f.Candidate
	}
	panic("unhandled case")
}
