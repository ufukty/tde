package traced

import (
	"go/ast"
	"reflect"
	"testing"
)

func Test_Retrace(t *testing.T) {
	testCase := []int{1, 1, 0, 0, 2}
	expectedMidNodes := []ast.Node{
		TEST_TREE,
		TEST_TREE.Decls[0],
		TEST_TREE.Decls[0].(*ast.FuncDecl).Type,
		TEST_TREE.Decls[0].(*ast.FuncDecl).Type.Params,
		TEST_TREE.Decls[0].(*ast.FuncDecl).Type.Params.List[0],
		TEST_TREE.Decls[0].(*ast.FuncDecl).Type.Params.List[0].Type,
	}
	Retrace(TEST_TREE, testCase, func(node ast.Node, depth int) {
		if expectedMidNodes[depth] != node {
			t.Errorf("Failed on Retrace depth: %d expected node: (value: %s, type: %s), got (value: %s, type: %s)",
				depth,
				reflect.ValueOf(node),
				reflect.TypeOf(node),
				reflect.ValueOf(expectedMidNodes[depth]),
				reflect.TypeOf(expectedMidNodes[depth]),
			)
		}
	})
}
