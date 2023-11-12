package models

import "go/ast"

type MutationParameters struct {
	Package         *ast.Package
	File            *ast.File
	FuncDecl        *ast.FuncDecl
	AllowedPackages []string
}

type GeneticOperation func(*MutationParameters) bool
