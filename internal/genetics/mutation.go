package genetics

import (
	"go/ast"
	ast_utl "tde/internal/astw/utilities"
	utl "tde/internal/utilities"
)

func MutateBlockStatement(stm *ast.BlockStmt) {
	// switch lines

	// partial regenerate
}

func Mutate(fn *ast.FuncDecl) {

	node := *utl.Pick(ast_utl.ListSubtree(fn.Body))
	mutated := false

	for !mutated {

		switch node := node.(type) {
		case *ast.BlockStmt:
			MutateBlockStatement(node)
		}
	}

}

// TODO: call cfg/node_constructor for choosen node
func RegenerateSubtree() {}

// TODO: remove an entry from a []ast.Stmt
func RemoveLine() {}

// TODO: swap subsequent nodes in a []ast.Stmt
func SwitchLines() {}
