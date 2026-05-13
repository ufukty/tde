package compare

import (
	"go/ast"
	"go/token"
	"testing"

	"tde/internal/astw/clone"
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

func TestChildren(t *testing.T) {
	compareSlices := func(a, b []ast.Node) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	testCases := []struct {
		input  ast.Node
		output []ast.Node
	}{{
		input: TEST_TREE,
		output: []ast.Node{
			TEST_TREE.Name,
			TEST_TREE.Decls[0].(*ast.FuncDecl),
		},
	}, {
		input: TEST_TREE.Decls[0],
		output: []ast.Node{
			TEST_TREE.Decls[0].(*ast.FuncDecl).Name,
			TEST_TREE.Decls[0].(*ast.FuncDecl).Type,
			TEST_TREE.Decls[0].(*ast.FuncDecl).Body,
		},
	}}
	for testIndex, testCase := range testCases {
		got := children(testCase.input)
		if !compareSlices(got, testCase.output) {
			t.Errorf("Failed on comparison want == got for test #%d\n", testIndex)
		}
	}
}

func TestRecursively(t *testing.T) {
	if Recursively(TEST_TREE, TEST_TREE) != true {
		t.Error("Failed for same inputs")
	}
	if RecursivelyWithAddresses(TEST_TREE, TEST_TREE) != true {
		t.Error("Failed for same inputs")
	}

	TEST_TREE_NEW := clone.File(TEST_TREE)
	if Recursively(TEST_TREE, TEST_TREE_NEW) != true {
		t.Error("Failed for same inputs")
	}
	if RecursivelyWithAddresses(TEST_TREE, TEST_TREE_NEW) != false {
		t.Error("Failed for same inputs")
	}

	TEST_TREE_NEW.Decls = append(TEST_TREE_NEW.Decls, &ast.GenDecl{})
	if RecursivelyWithAddresses(TEST_TREE, TEST_TREE_NEW) != false {
		t.Error("Failed for changed inputs")
	}
}
