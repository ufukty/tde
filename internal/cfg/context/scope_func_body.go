package context

import (
	"go/ast"
	"go/token"
)

func ExamineAssignStmt(ctx *Context, stmt *ast.AssignStmt) {
	if stmt.Tok == token.DEFINE {
		for i := 0; i < len(stmt.Lhs); i++ {
			lhs, rhs := stmt.Lhs[i], stmt.Rhs[i]
			if lhs, ok := lhs.(*ast.Ident); ok {
				switch rhs.(type) {
				case *ast.FuncLit:
				}
				ctx.AddVariable(lhs)
			}
		}
	}
}

func ExamineExprStmt(ctx *Context, exprStmt *ast.ExprStmt) {
	switch exprStmt.X.(type) {
	case *ast.CallExpr:
	}
}

func ExamineDeclStmt(ctx *Context, declStmt *ast.DeclStmt) {
	switch declStmt.Decl.(type) {
	case *ast.GenDecl:
	case *ast.FuncDecl:
	}
}
