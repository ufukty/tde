package models

import (
	"go/ast"
)

type TargetAst struct {
	Package  *ast.Package  // NOT cloned from original, read access only
	File     *ast.File     // cloned from original, safe to manipulate
	FuncDecl *ast.FuncDecl // cloned from original, safe to manipulate
	// AllowedPackages []string
}

type CandidateID string

type Candidate struct {
	UUID         CandidateID
	File         []byte // product of AST
	AST          TargetAst
	Fitness      Fitness
	ExecTimeInMs int
}
