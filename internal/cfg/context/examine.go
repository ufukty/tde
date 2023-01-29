package context

import (
	"go/ast"
	"go/token"
)

func ExamineEnteringNode(ctx *Context, node ast.Node) {
	switch node := node.(type) {
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
		for _, file := range node.Files {
			for _, decl := range file.Decls {
				if funcDecl, ok := decl.(*ast.FuncDecl); ok {
					ctx.AddVariable(funcDecl.Name)
				}
			}
		}

	case
		*ast.File:
		for _, decl := range node.Decls {
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				ctx.AddVariable(funcDecl.Name)
			}
		}
	}
}

func ExamineLeavingNode(ctx *Context, node ast.Node) {
	switch node := node.(type) {
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
					ctx.AddVariable(ident)
				}
			}
		}

	case
		*ast.FuncType:
		if fieldList := node.Params.List; fieldList != nil {
			for _, field := range fieldList {
				for _, ident := range field.Names {
					ctx.AddVariable(ident)
				}
			}
		}
	}
}
