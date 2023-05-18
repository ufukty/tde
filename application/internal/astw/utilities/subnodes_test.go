package utilities

import (
	"go/ast"
	"testing"
)

func TestChildNodes(t *testing.T) {
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
		got := ChildNodes(testCase.input)
		if !compareSlices(got, testCase.output) {
			t.Errorf("Failed on comparison want == got for test #%d\n", testIndex)
		}
	}
}
