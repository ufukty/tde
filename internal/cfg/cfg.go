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
	PickCutPoint(*ast.File, *ast.FuncDecl) *ast.Node
}
