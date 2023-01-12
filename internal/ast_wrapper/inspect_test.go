package ast_wrapper

import (
	"go/ast"
	"reflect"
	"testing"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Test_InspectWithTrace(t *testing.T) {
	// NOTE: Expected output when commented lines are enable de-commented
	// *ast.File         []                                                     []
	// *ast.Ident        [0x12d1600]                                            [0]
	// <nil>             [0x12d1600]                                            [0]
	// *ast.FuncDecl     [0x12d1600]                                            [1]
	// *ast.Ident        [0x12d1600 0x12cd360]                                  [1 0]
	// <nil>             [0x12d1600 0x12cd360]                                  [1 0]
	// *ast.FuncType     [0x12d1600 0x12cd360]                                  [1 1]
	// *ast.FieldList    [0x12d1600 0x12cd360 0x12cd240]                        [1 1 0]
	// *ast.Field        [0x12d1600 0x12cd360 0x12cd240 0x12cd3a0]              [1 1 0 0]
	// *ast.Ident        [0x12d1600 0x12cd360 0x12cd240 0x12cd3a0 0x12d1260]    [1 1 0 0 0]
	// <nil>             [0x12d1600 0x12cd360 0x12cd240 0x12cd3a0 0x12d1260]    [1 1 0 0 0]
	// *ast.Ident        [0x12d1600 0x12cd360 0x12cd240 0x12cd3a0 0x12d1260]    [1 1 0 0 1]
	// <nil>             [0x12d1600 0x12cd360 0x12cd240 0x12cd3a0 0x12d1260]    [1 1 0 0 1]
	// *ast.Ident        [0x12d1600 0x12cd360 0x12cd240 0x12cd3a0 0x12d1260]    [1 1 0 0 2]
	// <nil>             [0x12d1600 0x12cd360 0x12cd240 0x12cd3a0 0x12d1260]    [1 1 0 0 2]
	// <nil>             [0x12d1600 0x12cd360 0x12cd240 0x12cd3a0]              [1 1 0 0]
	// <nil>             [0x12d1600 0x12cd360 0x12cd240]                        [1 1 0]
	// *ast.FieldList    [0x12d1600 0x12cd360 0x12cd240]                        [1 1 1]
	// *ast.Field        [0x12d1600 0x12cd360 0x12cd240 0x12cd2e0]              [1 1 1 0]
	// *ast.Ident        [0x12d1600 0x12cd360 0x12cd240 0x12cd2e0 0x12d1220]    [1 1 1 0 0]
	// <nil>             [0x12d1600 0x12cd360 0x12cd240 0x12cd2e0 0x12d1220]    [1 1 1 0 0]
	// <nil>             [0x12d1600 0x12cd360 0x12cd240 0x12cd2e0]              [1 1 1 0]
	// <nil>             [0x12d1600 0x12cd360 0x12cd240]                        [1 1 1]
	// <nil>             [0x12d1600 0x12cd360]                                  [1 1]
	// *ast.BlockStmt    [0x12d1600 0x12cd360]                                  [1 2]
	// *ast.ReturnStmt   [0x12d1600 0x12cd360 0x12cd320]                        [1 2 0]
	// *ast.BinaryExpr   [0x12d1600 0x12cd360 0x12cd320 0x12cd1c0]              [1 2 0 0]
	// *ast.Ident        [0x12d1600 0x12cd360 0x12cd320 0x12cd1c0 0x12cd560]    [1 2 0 0 0]
	// <nil>             [0x12d1600 0x12cd360 0x12cd320 0x12cd1c0 0x12cd560]    [1 2 0 0 0]
	// *ast.Ident        [0x12d1600 0x12cd360 0x12cd320 0x12cd1c0 0x12cd560]    [1 2 0 0 1]
	// <nil>             [0x12d1600 0x12cd360 0x12cd320 0x12cd1c0 0x12cd560]    [1 2 0 0 1]
	// <nil>             [0x12d1600 0x12cd360 0x12cd320 0x12cd1c0]              [1 2 0 0]
	// <nil>             [0x12d1600 0x12cd360 0x12cd320]                        [1 2 0]
	// <nil>             [0x12d1600 0x12cd360]                                  [1 2]
	// <nil>             [0x12d1600]                                            [1]
	// <nil>             []                                                     []

	testCases := map[string][]int{
		reflect.TypeOf(&ast.FuncDecl{}).String():   {1},
		reflect.TypeOf(&ast.FuncType{}).String():   {1, 1},
		reflect.TypeOf(&ast.BlockStmt{}).String():  {1, 2},
		reflect.TypeOf(&ast.ReturnStmt{}).String(): {1, 2, 0},
		reflect.TypeOf(&ast.BinaryExpr{}).String(): {1, 2, 0, 0},
	}
	testCasesKeysOnly := maps.Keys(testCases)

	InspectWithTrace(TEST_TREE, func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		// fmt.Printf("%-30s %-60s %-100s\n", fmt.Sprint(reflect.TypeOf(currentNode)), fmt.Sprint(parentTrace), fmt.Sprint(childIndexTrace))

		if currentNode != nil {
			typeGot := reflect.TypeOf(currentNode).String()
			if idx := slices.Index(testCasesKeysOnly, typeGot); idx != -1 {
				if traceWant := testCases[testCasesKeysOnly[idx]]; slices.Compare(traceWant, childIndexTrace) != 0 {
					t.Errorf("failed. want: %v got: %v", traceWant, childIndexTrace)
				}
			}
		}

		return true
	})
}

func Test_InspectTwiceWithTrace(t *testing.T) {
	// NOTE: Expected output when commented lines are enable de-commented
	// 1st *ast.File          []                                                     []
	// 1st *ast.Ident         [0x12d4580]                                            [0]
	// 2nd *ast.Ident         [0x12d4580]                                            [0]
	// 1st *ast.FuncDecl      [0x12d4580]                                            [1]
	// 1st *ast.Ident         [0x12d4580 0x12d0360]                                  [1 0]
	// 2nd *ast.Ident         [0x12d4580 0x12d0360]                                  [1 0]
	// 1st *ast.FuncType      [0x12d4580 0x12d0360]                                  [1 1]
	// 1st *ast.FieldList     [0x12d4580 0x12d0360 0x12d0240]                        [1 1 0]
	// 1st *ast.Field         [0x12d4580 0x12d0360 0x12d0240 0x12d03a0]              [1 1 0 0]
	// 1st *ast.Ident         [0x12d4580 0x12d0360 0x12d0240 0x12d03a0 0x12d4260]    [1 1 0 0 0]
	// 2nd *ast.Ident         [0x12d4580 0x12d0360 0x12d0240 0x12d03a0 0x12d4260]    [1 1 0 0 0]
	// 1st *ast.Ident         [0x12d4580 0x12d0360 0x12d0240 0x12d03a0 0x12d4260]    [1 1 0 0 1]
	// 2nd *ast.Ident         [0x12d4580 0x12d0360 0x12d0240 0x12d03a0 0x12d4260]    [1 1 0 0 1]
	// 1st *ast.Ident         [0x12d4580 0x12d0360 0x12d0240 0x12d03a0 0x12d4260]    [1 1 0 0 2]
	// 2nd *ast.Ident         [0x12d4580 0x12d0360 0x12d0240 0x12d03a0 0x12d4260]    [1 1 0 0 2]
	// 2nd *ast.Field         [0x12d4580 0x12d0360 0x12d0240 0x12d03a0]              [1 1 0 0]
	// 2nd *ast.FieldList     [0x12d4580 0x12d0360 0x12d0240]                        [1 1 0]
	// 1st *ast.FieldList     [0x12d4580 0x12d0360 0x12d0240]                        [1 1 1]
	// 1st *ast.Field         [0x12d4580 0x12d0360 0x12d0240 0x12d02e0]              [1 1 1 0]
	// 1st *ast.Ident         [0x12d4580 0x12d0360 0x12d0240 0x12d02e0 0x12d4220]    [1 1 1 0 0]
	// 2nd *ast.Ident         [0x12d4580 0x12d0360 0x12d0240 0x12d02e0 0x12d4220]    [1 1 1 0 0]
	// 2nd *ast.Field         [0x12d4580 0x12d0360 0x12d0240 0x12d02e0]              [1 1 1 0]
	// 2nd *ast.FieldList     [0x12d4580 0x12d0360 0x12d0240]                        [1 1 1]
	// 2nd *ast.FuncType      [0x12d4580 0x12d0360]                                  [1 1]
	// 1st *ast.BlockStmt     [0x12d4580 0x12d0360]                                  [1 2]
	// 1st *ast.ReturnStmt    [0x12d4580 0x12d0360 0x12d0320]                        [1 2 0]
	// 1st *ast.BinaryExpr    [0x12d4580 0x12d0360 0x12d0320 0x12d01c0]              [1 2 0 0]
	// 1st *ast.Ident         [0x12d4580 0x12d0360 0x12d0320 0x12d01c0 0x12d0560]    [1 2 0 0 0]
	// 2nd *ast.Ident         [0x12d4580 0x12d0360 0x12d0320 0x12d01c0 0x12d0560]    [1 2 0 0 0]
	// 1st *ast.Ident         [0x12d4580 0x12d0360 0x12d0320 0x12d01c0 0x12d0560]    [1 2 0 0 1]
	// 2nd *ast.Ident         [0x12d4580 0x12d0360 0x12d0320 0x12d01c0 0x12d0560]    [1 2 0 0 1]
	// 2nd *ast.BinaryExpr    [0x12d4580 0x12d0360 0x12d0320 0x12d01c0]              [1 2 0 0]
	// 2nd *ast.ReturnStmt    [0x12d4580 0x12d0360 0x12d0320]                        [1 2 0]
	// 2nd *ast.BlockStmt     [0x12d4580 0x12d0360]                                  [1 2]
	// 2nd *ast.FuncDecl      [0x12d4580]                                            [1]
	// 2nd *ast.File          []                                                     []

	testCases := map[string][]int{
		reflect.TypeOf(&ast.FuncDecl{}).String():   {1},
		reflect.TypeOf(&ast.FuncType{}).String():   {1, 1},
		reflect.TypeOf(&ast.BlockStmt{}).String():  {1, 2},
		reflect.TypeOf(&ast.ReturnStmt{}).String(): {1, 2, 0},
		reflect.TypeOf(&ast.BinaryExpr{}).String(): {1, 2, 0, 0},
	}
	testCasesKeysOnly := maps.Keys(testCases)

	tester := func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) {
		if currentNode != nil {
			typeGot := reflect.TypeOf(currentNode).String()
			if idx := slices.Index(testCasesKeysOnly, typeGot); idx != -1 {
				if traceWant := testCases[testCasesKeysOnly[idx]]; slices.Compare(traceWant, childIndexTrace) != 0 {
					t.Errorf("failed. want: %v got: %v", traceWant, childIndexTrace)
				}
			}
		}
	}

	InspectTwiceWithTrace(TEST_TREE, func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		// fmt.Printf("1st %-30s %-60s %-100s\n", fmt.Sprint(reflect.TypeOf(currentNode)), fmt.Sprint(parentTrace), fmt.Sprint(childIndexTrace))
		tester(currentNode, parentTrace, childIndexTrace)
		return true
	}, func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) {
		// fmt.Printf("2nd %-30s %-60s %-100s\n", fmt.Sprint(reflect.TypeOf(currentNode)), fmt.Sprint(parentTrace), fmt.Sprint(childIndexTrace))
		tester(currentNode, parentTrace, childIndexTrace)
	})
}
