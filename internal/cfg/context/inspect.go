package context

import (
	"go/ast"
	"go/token"
	"tde/internal/ast_wrapper"
)

// TODO: Detect FuncLit's in code and add to Context
func InspectWithContext(startNode ast.Node, callback func(ctx Context, node ast.Node)) {
	ctx := NewContext()
	ast_wrapper.InspectTwiceWithTrace(startNode,
		// before visit children
		func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
			switch currentNode := currentNode.(type) {
			case
				*ast.BlockStmt,
				*ast.FuncDecl,
				*ast.ForStmt,
				*ast.IfStmt,
				*ast.SwitchStmt,
				*ast.TypeSwitchStmt:
				ctx.ScopeIn()

			case
				*ast.Package:
				for _, file := range currentNode.Files {
					for _, decl := range file.Decls {
						if funcDecl, ok := decl.(*ast.FuncDecl); ok {
							ctx.AddVariable(*funcDecl.Name)
						}
					}
				}

			case
				*ast.File:
				for _, decl := range currentNode.Decls {
					if funcDecl, ok := decl.(*ast.FuncDecl); ok {
						ctx.AddVariable(*funcDecl.Name)
					}
				}
			}

			callback(ctx, currentNode)
			return true
		},
		// after visit children
		func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) {
			switch node := currentNode.(type) {
			case
				*ast.BlockStmt,
				*ast.FuncDecl,
				*ast.ForStmt,
				*ast.IfStmt,
				*ast.SwitchStmt,
				*ast.TypeSwitchStmt:
				ctx.ScopeOut()

			case
				*ast.AssignStmt:
				if node.Tok == token.DEFINE {
					for _, expr := range node.Lhs {
						if ident, ok := expr.(*ast.Ident); ok {
							ctx.AddVariable(*ident)
						}
					}
				}

			case
				*ast.FuncType:
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
