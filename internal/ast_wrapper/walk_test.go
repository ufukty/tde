package ast_wrapper

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_WalkPersistentChildIndexTraces(t *testing.T) {

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

	WalkWithNils(TEST_TREE, func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		// fmt.Printf("%-60s %-20s %v %#v\n", fmt.Sprintf("%+v", parentTrace), fmt.Sprintf("%+v", childIndexTrace), reflect.TypeOf(n).String(), n)
		isFaulty(n, childIndexTrace, t)
		return true
	})

}

func Test_WalkCoveringTypes(t *testing.T) {
	_, astPkgs, _ := LoadDir(".")

	defer func() {
		if r := recover(); r != nil {
			t.Error(fmt.Sprint("Failed on walking syntax trees of entire package: ", r))
		}
	}()

	for _, pkg := range astPkgs {
		WalkWithNils(pkg, func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
			if n != nil {
				fmt.Printf("%-20s %v\n", reflect.TypeOf(n).String(), childIndexTrace)
			} else {
				fmt.Printf("%-20s %v\n", "nil", childIndexTrace)
			}
			return true
		})
	}
}

func Test_WalkListLeaves(t *testing.T) {
	_, astFile, _ := LoadFile("./walk.go")

	var indices = [][]int{}
	var parents = [][]ast.Node{}

	WalkWithNils(astFile, func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		if n == nil {
			indices = append(indices, slices.Clone(childIndexTrace))
			parents = append(parents, slices.Clone(parentTrace))
		}
		return true
	})

	for i := 0; i < len(indices); i++ {
		fmt.Printf("%p, %s, %v\n", parents[i][len(parents[i])-1], reflect.TypeOf(parents[i][len(parents[i])-1]).String(), indices[i])
	}

	for i := 0; i < len(indices); i++ {
		ast.Print(token.NewFileSet(), parents[i][len(parents[i])-1])
	}
}
