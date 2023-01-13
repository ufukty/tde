package ast_wrapper

import (
	"go/ast"
	"reflect"
	"testing"
)

func Test_IsAppendable(t *testing.T) {
	testCases := []ast.Node{&ast.ArrayType{}, &ast.AssignStmt{}, &ast.BinaryExpr{},
		&ast.BlockStmt{}, &ast.BranchStmt{}, &ast.CallExpr{}, &ast.CaseClause{},
		&ast.ChanType{}, &ast.CommClause{}, &ast.CompositeLit{}, &ast.DeclStmt{},
		&ast.DeferStmt{}, &ast.Ellipsis{}, &ast.EmptyStmt{}, &ast.ExprStmt{},
		&ast.Field{}, &ast.FieldList{}, &ast.ForStmt{}, &ast.FuncDecl{}, &ast.FuncLit{},
		&ast.FuncType{}, &ast.GenDecl{}, &ast.GoStmt{}, &ast.IfStmt{}, &ast.ImportSpec{},
		&ast.IncDecStmt{}, &ast.IndexExpr{}, &ast.IndexListExpr{}, &ast.InterfaceType{},
		&ast.KeyValueExpr{}, &ast.LabeledStmt{}, &ast.MapType{}, &ast.ParenExpr{},
		&ast.RangeStmt{}, &ast.ReturnStmt{}, &ast.SelectorExpr{}, &ast.SelectStmt{},
		&ast.SendStmt{}, &ast.SliceExpr{}, &ast.StarExpr{}, &ast.StructType{},
		&ast.SwitchStmt{}, &ast.TypeAssertExpr{}, &ast.TypeSpec{}, &ast.TypeSwitchStmt{},
		&ast.UnaryExpr{}, &ast.ValueSpec{}}

	IsAppendableWithReflect := func(node ast.Node) bool {
		typeOf := reflect.TypeOf(node).Elem()
		numberOfFields := typeOf.NumField()
		for i := 0; i < numberOfFields; i++ {
			if typeOf.Field(i).Type.Kind() == reflect.Slice {
				return true
			}
		}
		return false
	}

	for _, node := range testCases {
		got := IsAppendable(node)
		want := IsAppendableWithReflect(node)
		if want != got {
			t.Errorf("failed, for %v want: %t got: %t", reflect.TypeOf(node).String(), want, got)
		}

	}
}
