package node_constructor_v2

import (
	"go/ast"
	"go/token"
)

// if `x.(type)` is intended assertedType should be nil
func TypeSwitchStmtWithAssignStmt(ident *ast.Ident, assertedType ast.Expr, caseClauses []ast.Stmt) ast.TypeSwitchStmt {
	return ast.TypeSwitchStmt{
		Switch: 0,
		Init:   nil,
		Assign: &ast.AssignStmt{
			Lhs:    []ast.Expr{ident},
			TokPos: 0,
			Tok:    token.DEFINE,
			Rhs: []ast.Expr{
				&ast.TypeAssertExpr{
					X:      ident,
					Lparen: 0,
					Type:   assertedType,
					Rparen: 0,
				},
			},
		},
		Body: &ast.BlockStmt{
			Lbrace: 0,
			List:   caseClauses,
			Rbrace: 0,
		},
	}
}

// if `x.(type)` is intended assertedType should be nil
func TypeSwitchStmtWithoutAssignStmt(ident *ast.Ident, assertedType ast.Expr, caseClauses []ast.Stmt) ast.TypeSwitchStmt {
	return ast.TypeSwitchStmt{
		Switch: 0,
		Init:   nil,
		Assign: &ast.ExprStmt{
			X: &ast.TypeAssertExpr{
				X:      ident,
				Lparen: 0,
				Type:   assertedType,
				Rparen: 0,
			},
		},
		Body: &ast.BlockStmt{
			Lbrace: 0,
			List:   caseClauses,
			Rbrace: 0,
		},
	}
}
