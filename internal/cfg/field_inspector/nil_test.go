package field_inspector

import (
	"fmt"
	"go/ast"
	"reflect"
	"tde/internal/cfg/astcfg/context"
	"testing"
)

func Test_GetConstructor(t *testing.T) {
	s := struct {
		Spec ast.Spec
		Decl ast.Decl
		Expr ast.Expr
		Stmt ast.Stmt
	}{}

	xSpec, _ := reflect.TypeOf(s).FieldByName("Spec")
	xDecl, _ := reflect.TypeOf(s).FieldByName("Decl")
	xExpr, _ := reflect.TypeOf(s).FieldByName("Expr")
	xStmt, _ := reflect.TypeOf(s).FieldByName("Stmt")

	testCases := []reflect.Type{
		xSpec.Type,
		xDecl.Type,
		xExpr.Type,
		xStmt.Type,
		// reflect.TypeOf(&ast.Comment{}),
		// reflect.TypeOf(&ast.CommentGroup{}),
		reflect.TypeOf(&ast.Field{}),
		reflect.TypeOf(&ast.FieldList{}),
		// reflect.TypeOf(&ast.BadExpr{}),
		reflect.TypeOf(&ast.Ident{}),
		reflect.TypeOf(&ast.Ellipsis{}),
		reflect.TypeOf(&ast.BasicLit{}),
		reflect.TypeOf(&ast.FuncLit{}),
		reflect.TypeOf(&ast.CompositeLit{}),
		reflect.TypeOf(&ast.ParenExpr{}),
		reflect.TypeOf(&ast.SelectorExpr{}),
		reflect.TypeOf(&ast.IndexExpr{}),
		reflect.TypeOf(&ast.IndexListExpr{}),
		reflect.TypeOf(&ast.SliceExpr{}),
		reflect.TypeOf(&ast.TypeAssertExpr{}),
		reflect.TypeOf(&ast.CallExpr{}),
		reflect.TypeOf(&ast.StarExpr{}),
		reflect.TypeOf(&ast.UnaryExpr{}),
		reflect.TypeOf(&ast.BinaryExpr{}),
		reflect.TypeOf(&ast.KeyValueExpr{}),
		reflect.TypeOf(&ast.ArrayType{}),
		reflect.TypeOf(&ast.StructType{}),
		reflect.TypeOf(&ast.FuncType{}),
		reflect.TypeOf(&ast.InterfaceType{}),
		reflect.TypeOf(&ast.MapType{}),
		reflect.TypeOf(&ast.ChanType{}),
		// reflect.TypeOf(&ast.BadStmt{}),
		reflect.TypeOf(&ast.DeclStmt{}),
		reflect.TypeOf(&ast.EmptyStmt{}),
		reflect.TypeOf(&ast.LabeledStmt{}),
		reflect.TypeOf(&ast.ExprStmt{}),
		reflect.TypeOf(&ast.SendStmt{}),
		reflect.TypeOf(&ast.IncDecStmt{}),
		reflect.TypeOf(&ast.AssignStmt{}),
		reflect.TypeOf(&ast.GoStmt{}),
		reflect.TypeOf(&ast.DeferStmt{}),
		reflect.TypeOf(&ast.ReturnStmt{}),
		reflect.TypeOf(&ast.BranchStmt{}),
		reflect.TypeOf(&ast.BlockStmt{}),
		reflect.TypeOf(&ast.IfStmt{}),
		reflect.TypeOf(&ast.CaseClause{}),
		reflect.TypeOf(&ast.SwitchStmt{}),
		reflect.TypeOf(&ast.TypeSwitchStmt{}),
		reflect.TypeOf(&ast.CommClause{}),
		reflect.TypeOf(&ast.SelectStmt{}),
		reflect.TypeOf(&ast.ForStmt{}),
		reflect.TypeOf(&ast.RangeStmt{}),
		reflect.TypeOf(&ast.ImportSpec{}),
		reflect.TypeOf(&ast.ValueSpec{}),
		reflect.TypeOf(&ast.TypeSpec{}),
		// reflect.TypeOf(&ast.BadDecl{}),
		reflect.TypeOf(&ast.GenDecl{}),
		reflect.TypeOf(&ast.FuncDecl{}),
		// reflect.TypeOf(&ast.File{}),
		// reflect.TypeOf(&ast.Package{}),
	}

	for _, want := range testCases {
		// fmt.Printf("input: %s, want: %p", input.String(), &want)
		constructor := GetConstructor(want)
		if constructor == nil {
			t.Errorf("failed, could not find the constructor for %s", want.String())
			continue
		}
		node := constructor(context.NewContext(), 1)
		got := reflect.TypeOf(node)
		if got != want && !got.Implements(want) {
			t.Errorf("failed for reflect type: %s, got a constructor that returns a node with reflect type %s", want.String(), got.String())
		}
	}
}

func Test_GetListOfAvailableSpots(t *testing.T) {
	// ast.Print(token.NewFileSet(), TEST_TREE)
	availableSpots := GetListOfAvailableSpots(TEST_TREE)
	fmt.Println(availableSpots)
	// ast.Print(token.NewFileSet(), TEST_TREE)
}
