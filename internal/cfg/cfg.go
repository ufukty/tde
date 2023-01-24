package cfg

import (
	"go/ast"
)

// Code Fragment Generator
//
// Planned implementations vary with 4 methods:
//   - Random characters
//   - Random tokens
//   - Random AST-nodes
//   - Context-Aware Random AST-nodes
type CFG interface {
	Develop(*ast.File, *ast.FuncDecl)
	Exchange(*ast.File, *ast.FuncDecl, *ast.File, *ast.FuncDecl)
}
