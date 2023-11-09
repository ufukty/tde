package models

import (
	"go/ast"
	"tde/internal/astw/clone/clean"

	"github.com/google/uuid"
)

type Sid string // SubjectId

type Subject struct {
	Sid          Sid
	Parent       Sid
	AST          *ast.FuncDecl // to manipulate
	Imports      []*ast.ImportSpec
	Code         []byte // printed AST
	Fitness      Fitness
	ExecTimeInMs int
}

func (s Subject) Clone() *Subject {
	return &Subject{
		Sid:    Sid(uuid.New().String()),
		Parent: s.Sid,
		AST:    clean.FuncDecl(s.AST),
	}
}

func (c Subject) IsValidIn(layer Layer) bool {
	return c.Fitness.Layer() >= layer
}
