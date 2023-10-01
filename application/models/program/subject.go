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

type Sid string // SubjectId

type Subject struct {
	Sid          Sid
	File         []byte // product of AST
	AST          TargetAst
	Fitness      Fitness
	ExecTimeInMs int
}

type Subjects = map[Sid]*Subject // To make Subjects accessible by CIDs

func (c Subject) IsValidIn(layer Layer) bool {
	return c.Fitness.Layer() >= layer
}
