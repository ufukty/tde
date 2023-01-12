package context

import (
	"go/ast"
	"go/token"
	"tde/internal/ast_wrapper"
)

// TODO: Checking for if entering to or leaving BlockStmt might not be enough for detecting scopes; since parent node of block could be the scope defining one; eg. FuncDecl
func InspectWithContext(startNode ast.Node, callback func(ctx Context, node ast.Node)) {
	ctx := NewContext()
	ast_wrapper.InspectTwiceWithTrace(startNode,
		// before visit children
		func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
			switch currentNode.(type) {
			case *ast.BlockStmt:
				ctx.ScopeIn()
			}
			callback(ctx, currentNode)
			return true
		},
		// after visit children
		func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) {
			switch node := currentNode.(type) {
			case *ast.BlockStmt:
				ctx.ScopeOut()
			case *ast.AssignStmt:
				if node.Tok == token.DEFINE {
					for _, expr := range node.Lhs {
						if ident, ok := expr.(*ast.Ident); ok {
							ctx.AddVariable(*ident)
						}
					}
				}
			case *ast.FuncType:
				if fieldList := node.Params.List; fieldList != nil {
					for _, field := range fieldList {
						for _, ident := range field.Names {
							ctx.AddVariable(*ident)
						}
					}
				}
			}
		})

}
