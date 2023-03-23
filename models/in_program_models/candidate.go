package in_program_models

import (
	"go/ast"
	"log"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Fitness struct {
	AST      float64 // rel. printing errors (from ast)
	Code     float64 // rel. syntax errors (compile)
	Program  float64 // rel. runtime errors
	Solution float64 // rel. passed tests (user-provided)
}

func (f Fitness) Flat() float64 {
	if f.AST != 0.0 {
		return 3.0 + f.AST
	} else if f.Code != 0.0 {
		return 2.0 + f.Code
	} else if f.Program != 0.0 {
		return 1.0 + f.Program
	} else {
		return f.Program
	}
}

type CandidateASTRepresentation struct {
	Package  *ast.Package
	File     *ast.File
	FuncDecl *ast.FuncDecl
	// AllowedPackages []string
}

type CandidateID string

type Candidate struct {
	UUID         CandidateID
	File         []byte // product of AST
	AST          CandidateASTRepresentation
	Fitness      Fitness
	ExecTimeInMs int
}

func NewCandidate() *Candidate {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not create an UUID for new Individual"))
	}
	return &Candidate{
		UUID: CandidateID(newUUID.String()),
	}
}

// func (c *Candidate) Measure() {
// 	if !c.CheckSyntax() {
// 		c.Fitness = 1.1 // fitness for invalid-syntax programs exceeds the "1.0" treshold
// 		return
// 	}

// 	// t := &Testing{}
// 	// var timeStart = time.Now()
// 	// (*(i.TestFunction))(t)

// 	// i.Fitness = float64(t.TotalErrors) / float64(t.TotalCalls)
// 	// i.ExecTimeInMs = int(time.Since(timeStart))
// }
