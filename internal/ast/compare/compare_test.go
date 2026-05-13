package compare

import (
	"fmt"
	"go/ast"
	"go/token"
	"testing"

	"tde/internal/ast/clone"
)

var (
	OBJ_FUNC  = ast.NewObj(ast.Fun, "Addition")
	OBJ_A     = ast.NewObj(ast.Var, "a")
	OBJ_B     = ast.NewObj(ast.Var, "b")
	FUNC_NAME = &ast.Ident{Name: "Addition", Obj: OBJ_FUNC}
	FUNC_TYPE = &ast.FuncType{
		Params: &ast.FieldList{List: []*ast.Field{{
			Names: []*ast.Ident{{Name: "a", Obj: OBJ_A}, {Name: "b", Obj: OBJ_B}},
			Type:  &ast.Ident{Name: "int"},
		}}},
		Results: &ast.FieldList{
			Opening: token.NoPos,
			List:    []*ast.Field{{Type: &ast.Ident{Name: "int"}}},
			Closing: token.NoPos,
		},
	}
	FUNC_BODY = &ast.BlockStmt{
		List: []ast.Stmt{&ast.ReturnStmt{Results: []ast.Expr{&ast.BinaryExpr{
			X:  &ast.Ident{Name: "a", Obj: OBJ_A},
			Op: token.ADD,
			Y:  &ast.Ident{Name: "b", Obj: OBJ_B},
		}}}},
	}
	FUNC_DECL = &ast.FuncDecl{Name: FUNC_NAME, Type: FUNC_TYPE, Body: FUNC_BODY}
	FILE      = &ast.File{
		Name:       &ast.Ident{Name: "main"},
		Decls:      []ast.Decl{FUNC_DECL},
		Scope:      &ast.Scope{Objects: map[string]*ast.Object{"Addition": OBJ_FUNC}},
		Unresolved: []*ast.Ident{},
	}
)

func compareSlices(a, b []ast.Node) error {
	if len(a) != len(b) {
		return fmt.Errorf("lengths mismatch a=%d b=%d", len(a), len(b))
	}
	for i := range len(a) {
		if a[i] != b[i] {
			return fmt.Errorf("items mismatch: %d", i)
		}
	}
	return nil
}

func TestChildren(t *testing.T) {
	type tc struct {
		input  ast.Node
		output []ast.Node
	}
	testCases := map[string]tc{
		"children of file": {
			input:  FILE,
			output: []ast.Node{FILE.Name, FUNC_DECL},
		},
		"children of func decl": {
			input:  FILE.Decls[0],
			output: []ast.Node{FUNC_NAME, FUNC_TYPE, FUNC_BODY},
		},
	}
	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			got := children(tc.input)
			if err := compareSlices(got, tc.output); err != nil {
				t.Errorf("assert: %v", err)
			}
		})
	}
}

func TestRecursively(t *testing.T) {
	t.Run("same AST", func(t *testing.T) {
		if !Recursively(FILE, FILE) {
			t.Error("assertion, expected true")
		}
	})
	t.Run("cloned AST", func(t *testing.T) {
		if !Recursively(FILE, clone.File(FILE)) {
			t.Error("assertion, expected true")
		}
	})
}

func TestRecursivelyWithAddresses(t *testing.T) {
	t.Run("same AST", func(t *testing.T) {
		if !RecursivelyWithAddresses(FILE, FILE) {
			t.Error("assertion, expected true")
		}
	})
	t.Run("cloned AST", func(t *testing.T) {
		if RecursivelyWithAddresses(FILE, clone.File(FILE)) {
			t.Error("assertion, expected false")
		}
	})
	t.Run("cloned and edited AST", func(t *testing.T) {
		cloned := clone.File(FILE)
		cloned.Decls = append(cloned.Decls, &ast.GenDecl{})
		if RecursivelyWithAddresses(FILE, cloned) {
			t.Error("assertion, expected false")
		}
	})
}
