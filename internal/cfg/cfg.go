package cfg

import (
	"go/ast"
	"tde/internal/code"
)

// Code Fragment Generator
//
// Planned implementations vary with 4 methods:
//   - Random characters
//   - Random tokens
//   - Random AST-nodes
//   - Context-Aware Random AST-nodes
type CFG interface {
	Develop(code.Code, ast.BlockStmt)
	PickCutPoint(code.Code) int
}
