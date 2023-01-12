package context

import (
	"go/ast"
	"go/token"
)

var TEST_TREE_OBJ_FUNC = ast.NewObj(ast.Fun, "Addition")
var TEST_TREE_OBJ_A = ast.NewObj(ast.Var, "a")
var TEST_TREE_OBJ_B = ast.NewObj(ast.Var, "b")
var TEST_TREE_OBJ_C = ast.NewObj(ast.Var, "c")
var TEST_TREE = &ast.File{
	Name: &ast.Ident{
		Name: "main",
	},
	Decls: []ast.Decl{
		&ast.FuncDecl{
			Name: &ast.Ident{
				Name: "Addition",
				Obj:  TEST_TREE_OBJ_FUNC,
			},
			Type: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								{
									Name: "a",
									Obj:  TEST_TREE_OBJ_A,
								},
								{
									Name: "b",
									Obj:  TEST_TREE_OBJ_B,
								},
							},
							Type: &ast.Ident{
								Name: "int",
							},
						},
					},
				},
				Results: &ast.FieldList{
					Opening: token.NoPos,
					List: []*ast.Field{
						{
							Type: &ast.Ident{
								Name: "int",
							},
						},
					},
					Closing: token.NoPos,
				},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{
								Name: "c",
								Obj:  TEST_TREE_OBJ_C,
							},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.BinaryExpr{
								X: &ast.Ident{
									Name: "a",
									Obj:  TEST_TREE_OBJ_A,
								},
								Op: token.ADD,
								Y: &ast.Ident{
									Name: "b",
									Obj:  TEST_TREE_OBJ_B,
								},
							},
						},
					},
					&ast.ReturnStmt{
						Results: []ast.Expr{
							&ast.Ident{
								Name: "c",
								Obj:  TEST_TREE_OBJ_C,
							},
						},
					},
				},
			},
		},
	},
	Scope: &ast.Scope{
		Objects: map[string]*ast.Object{
			"Addition": TEST_TREE_OBJ_FUNC,
		},
	},
	Unresolved: []*ast.Ident{},
}
