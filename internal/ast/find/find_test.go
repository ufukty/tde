package find

import (
	"go/ast"
	"go/token"
	"testing"

	"tde/internal/ast/parse"
)

const TEST_FILE = `package main
	
func Addition(a, b int) int {
	return a + b
}
`

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

func TestFuncDecl(t *testing.T) {
	_, n, err := parse.String(TEST_FILE)
	if err != nil {
		t.Errorf("prep, parse: %v", err)
	}
	f, ok := n.(*ast.File)
	if !ok {
		t.Errorf("prep, assert: expected %q got %T", "*ast.File", n)
	}
	fd, err := FunctionInFile(f, "Addition")
	if err != nil {
		t.Errorf("failed find.Function: %v", err)
	}
	if fd.Name.Name != "Addition" {
		t.Errorf("failed name check: %v", err)
	}
}
