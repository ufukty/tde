package models

import (
	"go/ast"
	"maps"
	"tde/internal/astw/clone/clean"
	"tde/internal/utilities"

	"github.com/google/uuid"
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
	Parent       Sid
	File         []byte // product of AST
	AST          TargetAst
	Fitness      Fitness
	ExecTimeInMs int
}

func (s Subject) Clone() *Subject {
	return &Subject{
		Sid:    Sid(uuid.New().String()),
		Parent: s.Sid,
		AST: TargetAst{
			Package:  s.AST.Package,
			File:     s.AST.File,
			FuncDecl: clean.FuncDecl(s.AST.FuncDecl),
		},
	}
}

func (c Subject) IsValidIn(layer Layer) bool {
	return c.Fitness.Layer() >= layer
}

type Subjects map[Sid]*Subject // To make Subjects accessible by CIDs

func (s Subjects) Add(subj *Subject) {
	(s)[subj.Sid] = subj
}

func (s Subjects) Join(s2 Subjects) {
	maps.Copy(s, s2)
}

func (s Subjects) Diff(subtract Subjects) Subjects {
	diff := utilities.MapDiff(map[Sid]*Subject(s), map[Sid]*Subject(subtract))
	return (Subjects)(diff)
}
