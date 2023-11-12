package models

import "go/ast"

type GeneticOperationContext struct {
	Package         *ast.Package
	File            *ast.File
	FuncDecl        *ast.FuncDecl
	AllowedPackages []string
}

type GeneticOperation func(*GeneticOperationContext) bool
