package models

import (
	"fmt"
	"go/ast"
)

type MutationParameters struct {
	Package         *ast.Package
	File            *ast.File
	FuncDecl        *ast.FuncDecl
	AllowedPackages []string
}

type GeneticOperation func(*MutationParameters) error

var (
	ErrUnsupportedMutation = fmt.Errorf("operation has returned false")
	ErrNoChangeNeeded      = fmt.Errorf("no change is needed")
)
