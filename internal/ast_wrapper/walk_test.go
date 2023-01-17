package ast_wrapper

import (
	"go/ast"
	"reflect"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_Walk(t *testing.T) {

	testCasesChildIndexTrace := map[ast.Node][]int{
		TEST_TREE:          {},
		TEST_TREE.Name:     {0},
		TEST_TREE.Decls[0]: {1},
		// TEST_TREE.Decls[0].(*ast.FuncDecl).Recv: {1, 0}, // will FAIL because can't compare nil values
		TEST_TREE.Decls[0].(*ast.FuncDecl).Name: {1, 1},
		TEST_TREE.Decls[0].(*ast.FuncDecl).Type: {1, 2},
		TEST_TREE.Decls[0].(*ast.FuncDecl).Body: {1, 3},
	}

	isFaulty := func(node ast.Node, got []int, t *testing.T) {
		for forNode, want := range testCasesChildIndexTrace {
			if node == forNode {
				if slices.Compare(want, got) != 0 {
					t.Errorf("failure for %s: want %v got: %v", reflect.TypeOf(node).String(), want, got)
				}
			}
		}
	}

	Walk(TEST_TREE, false, func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		// fmt.Printf("%-60s %-20s %v %#v\n", fmt.Sprintf("%+v", parentTrace), fmt.Sprintf("%+v", childIndexTrace), reflect.TypeOf(n).String(), n)
		isFaulty(n, childIndexTrace, t)
		return true
	})

}
