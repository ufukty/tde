package test_package

import (
	"fmt"
	"go/ast"
	"reflect"
	"testing"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Test_InspectWithTrace(t *testing.T) {
	// NOTE: Expected output when commented lines are enable de-commented
	// *ast.File            &{<nil> 0 main [0x12fb3c0] scope 0x1303cd0 { func Addition } [] [] []} [] []
	// *ast.Ident           main                                               [0]                  [0x12ff5e0]
	// <nil>                <invalid reflect.Value>                            [0]                  [0x12ff5e0]
	// *ast.FuncDecl        &{<nil> <nil> Addition 0x12fb2a0 0x12fb380}        [1]                  [0x12ff5e0]
	// *ast.Ident           Addition                                           [1 0]                [0x12ff5e0 0x12fb3c0]
	// <nil>                <invalid reflect.Value>                            [1 0]                [0x12ff5e0 0x12fb3c0]
	// *ast.FuncType        &{0 <nil> 0x12fb400 0x12fb340}                     [1 1]                [0x12ff5e0 0x12fb3c0]
	// *ast.FieldList       &{0 [0x12ff2c0] 0}                                 [1 1 0]              [0x12ff5e0 0x12fb3c0 0x12fb2a0]
	// *ast.Field           &{<nil> [a b] int <nil> <nil>}                     [1 1 0 0]            [0x12ff5e0 0x12fb3c0 0x12fb2a0 0x12fb400]
	// *ast.Ident           a                                                  [1 1 0 0 0]          [0x12ff5e0 0x12fb3c0 0x12fb2a0 0x12fb400 0x12ff2c0]
	// <nil>                <invalid reflect.Value>                            [1 1 0 0 0]          [0x12ff5e0 0x12fb3c0 0x12fb2a0 0x12fb400 0x12ff2c0]
	// *ast.Ident           b                                                  [1 1 0 0 1]          [0x12ff5e0 0x12fb3c0 0x12fb2a0 0x12fb400 0x12ff2c0]
	// <nil>                <invalid reflect.Value>                            [1 1 0 0 1]          [0x12ff5e0 0x12fb3c0 0x12fb2a0 0x12fb400 0x12ff2c0]
	// *ast.Ident           int                                                [1 1 0 0 2]          [0x12ff5e0 0x12fb3c0 0x12fb2a0 0x12fb400 0x12ff2c0]
	// <nil>                <invalid reflect.Value>                            [1 1 0 0 2]          [0x12ff5e0 0x12fb3c0 0x12fb2a0 0x12fb400 0x12ff2c0]
	// <nil>                <invalid reflect.Value>                            [1 1 0 0]            [0x12ff5e0 0x12fb3c0 0x12fb2a0 0x12fb400]
	// <nil>                <invalid reflect.Value>                            [1 1 0]              [0x12ff5e0 0x12fb3c0 0x12fb2a0]
	// *ast.FieldList       &{0 [0x12ff280] 0}                                 [1 1 1]              [0x12ff5e0 0x12fb3c0 0x12fb2a0]
	// *ast.Field           &{<nil> [] int <nil> <nil>}                        [1 1 1 0]            [0x12ff5e0 0x12fb3c0 0x12fb2a0 0x12fb340]
	// *ast.Ident           int                                                [1 1 1 0 0]          [0x12ff5e0 0x12fb3c0 0x12fb2a0 0x12fb340 0x12ff280]
	// <nil>                <invalid reflect.Value>                            [1 1 1 0 0]          [0x12ff5e0 0x12fb3c0 0x12fb2a0 0x12fb340 0x12ff280]
	// <nil>                <invalid reflect.Value>                            [1 1 1 0]            [0x12ff5e0 0x12fb3c0 0x12fb2a0 0x12fb340]
	// <nil>                <invalid reflect.Value>                            [1 1 1]              [0x12ff5e0 0x12fb3c0 0x12fb2a0]
	// <nil>                <invalid reflect.Value>                            [1 1]                [0x12ff5e0 0x12fb3c0]
	// *ast.BlockStmt       &{0 [0x12fb220] 0}                                 [1 2]                [0x12ff5e0 0x12fb3c0]
	// *ast.ReturnStmt      &{0 [0x12fb5c0]}                                   [1 2 0]              [0x12ff5e0 0x12fb3c0 0x12fb380]
	// *ast.BinaryExpr      &{a 0 + b}                                         [1 2 0 0]            [0x12ff5e0 0x12fb3c0 0x12fb380 0x12fb220]
	// *ast.Ident           a                                                  [1 2 0 0 0]          [0x12ff5e0 0x12fb3c0 0x12fb380 0x12fb220 0x12fb5c0]
	// <nil>                <invalid reflect.Value>                            [1 2 0 0 0]          [0x12ff5e0 0x12fb3c0 0x12fb380 0x12fb220 0x12fb5c0]
	// *ast.Ident           b                                                  [1 2 0 0 1]          [0x12ff5e0 0x12fb3c0 0x12fb380 0x12fb220 0x12fb5c0]
	// <nil>                <invalid reflect.Value>                            [1 2 0 0 1]          [0x12ff5e0 0x12fb3c0 0x12fb380 0x12fb220 0x12fb5c0]
	// <nil>                <invalid reflect.Value>                            [1 2 0 0]            [0x12ff5e0 0x12fb3c0 0x12fb380 0x12fb220]
	// <nil>                <invalid reflect.Value>                            [1 2 0]              [0x12ff5e0 0x12fb3c0 0x12fb380]
	// <nil>                <invalid reflect.Value>                            [1 2]                [0x12ff5e0 0x12fb3c0]
	// <nil>                <invalid reflect.Value>                            [1]                  [0x12ff5e0]
	// <nil>                <invalid reflect.Value>                            []                   []

	testCases := map[string][]int{
		reflect.TypeOf(&ast.FuncDecl{}).String():   {1},
		reflect.TypeOf(&ast.FuncType{}).String():   {1, 1},
		reflect.TypeOf(&ast.BlockStmt{}).String():  {1, 2},
		reflect.TypeOf(&ast.ReturnStmt{}).String(): {1, 2, 0},
		reflect.TypeOf(&ast.BinaryExpr{}).String(): {1, 2, 0, 0},
	}
	testCasesKeysOnly := maps.Keys(testCases)

	InspectWithTrace(TEST_TREE, func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		fmt.Printf(
			"%-20s %-50s %-20s %-60s\n",
			fmt.Sprint(reflect.TypeOf(currentNode)),
			fmt.Sprint(reflect.ValueOf(currentNode)),
			fmt.Sprint(childIndexTrace),
			fmt.Sprint(parentTrace),
		)

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
