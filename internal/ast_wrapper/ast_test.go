package ast_wrapper

import (
	"fmt"
	"go/ast"
	"go/token"
	"testing"
)

func Test_PrintAST2(t *testing.T) {
	testFunction := ast.FuncDecl{
		Doc:  nil,
		Recv: nil,
		Name: ast.NewIdent("Addition"),
		Type: &ast.FuncType{
			Func:       token.NoPos,
			TypeParams: nil, // related with generics?
			Params: &ast.FieldList{
				Opening: token.NoPos,
				List: []*ast.Field{{
					Doc:     nil,
					Names:   []*ast.Ident{ast.NewIdent("a"), ast.NewIdent("b")},
					Type:    ast.NewIdent("int"),
					Tag:     nil,
					Comment: nil,
				}},
				Closing: token.NoPos,
			},
			Results: &ast.FieldList{
				Opening: token.NoPos,
				List: []*ast.Field{{
					Doc:     nil,
					Names:   nil,
					Type:    ast.NewIdent("int"),
					Tag:     nil,
					Comment: nil,
				}},
				Closing: token.NoPos,
			},
		},
		Body: &ast.BlockStmt{
			Lbrace: token.NoPos,
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Return: token.NoPos,
					Results: []ast.Expr{
						&ast.BinaryExpr{
							X:     ast.NewIdent("a"),
							OpPos: token.NoPos,
							Op:    token.ADD,
							Y:     ast.NewIdent("b"),
						},
					},
				},
			},
			Rbrace: token.NoPos,
		},
	}

	fmt.Println(StringAST(&testFunction))
}
