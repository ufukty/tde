package genetics

import (
	"go/ast"
	"go/token"
	"tde/internal/ast_wrapper"
	"tde/internal/utilities"
)

func MutateStringLiteral(s *string) {
	// partially regenerate
	// one char mutate
}

func MutateLiteralInteger(i *string) {
	// random assignment

	// 10% + -
}

func MutateBlockStatement(stm *ast.BlockStmt) {
	// switch lines

	// partial regenerate
}

func Mutate(fn *ast.FuncDecl) {

	node := **utilities.Pick(ast_wrapper.ListSubnodes(fn.Body))
	mutated := false

	for !mutated {

		switch node := node.(type) {
		case *ast.BasicLit:
			switch node.Kind {
			case token.STRING:
				MutateStringLiteral(&node.Value)
			case token.INT:
				MutateLiteralInteger(&node.Value)
			}
		case *ast.BlockStmt:
			MutateBlockStatement(node)
		}
	}

}
