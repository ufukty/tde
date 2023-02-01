package genetics

import (
	ast_utl "tde/internal/astw/utilities"
	utl "tde/internal/utilities"

	"go/ast"
	"go/token"
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

	node := *utl.Pick(ast_utl.ListSubnodes(fn.Body))
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
