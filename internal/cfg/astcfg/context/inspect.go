package context

import (
	"go/ast"
	"tde/internal/ast_wrapper"
)

func InspectWithContext(startNode ast.Node, callback func(ctx Context, node ast.Node)) {
	ctx := NewContext()
	ast_wrapper.InspectTwiceWithTrace(startNode, func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {

		switch currentNode.(type) {
		case *ast.BlockStmt:
			ctx.ScopeIn()
		case *ast.DeclStmt:

		}

		return true
	}, func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		switch currentNode.(type) {
		case *ast.BlockStmt:
			ctx.ScopeOut()
		}

		return true
	})

}
