package astwutl

import (
	"fmt"
	"go/ast"
	"go/token"
	"testing"
)

var (
	TEST_TREE_OBJ_FUNC = ast.NewObj(ast.Fun, "Addition")
	TEST_TREE_OBJ_A    = ast.NewObj(ast.Var, "a")
	TEST_TREE_OBJ_B    = ast.NewObj(ast.Var, "b")
	TEST_TREE          = &ast.File{
		Name: &ast.Ident{
			Name: "main",
		},
		Decls: []ast.Decl{
			0: &ast.FuncDecl{
				Name: &ast.Ident{
					Name: "Addition",
					Obj:  TEST_TREE_OBJ_FUNC,
				},
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							0: {
								Names: []*ast.Ident{
									0: {
										Name: "a",
										Obj:  TEST_TREE_OBJ_A,
									},
									1: {
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
							0: {
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
						0: &ast.ReturnStmt{
							Results: []ast.Expr{
								0: &ast.BinaryExpr{
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
)

func Test_String(t *testing.T) {
	_, err := String(TEST_TREE)
	if err != nil {
		t.Error(fmt.Errorf("failed String: %w", err))
	}
}
