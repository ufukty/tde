package models

import (
	"tde/internal/astw/clone"
	astw_utl "tde/internal/astw/utilities"

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
	Package  *ast.Package  // NOT cloned from original, read access only
	File     *ast.File     // cloned from original, safe to manipulate
	FuncDecl *ast.FuncDecl // cloned from original, safe to manipulate
	// AllowedPackages []string
}

type CandidateID string

type BreedID string

type Candidate struct {
	UUID         CandidateID
	BreedID      BreedID
	File         []byte // product of AST
	AST          CandidateASTRepresentation
	Fitness      Fitness
	ExecTimeInMs int
}

func NewCandidate(pkg *ast.Package, file *ast.File, funcDecl *ast.FuncDecl) (*Candidate, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not create an UUID for new Individual"))
	}

	cloneFile := clone.File(file)
	clondeFuncDecl, err := astw_utl.FindFuncDecl(cloneFile, funcDecl.Name.Name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find 'funcDecl.Name.Name' in the 'file' after cloned it")
	}

	return &Candidate{
		UUID: CandidateID(newUUID.String()),
		AST: CandidateASTRepresentation{
			Package:  pkg,
			File:     cloneFile,
			FuncDecl: clondeFuncDecl,
		},
	}, nil
}
