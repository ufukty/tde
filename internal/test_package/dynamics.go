package test_package

import (
	"go/ast"
)

func IsAppendable(node ast.Node) bool {
	switch node.(type) {
	case
		*ast.ArrayType,
		*ast.BinaryExpr,
		*ast.BranchStmt,
		*ast.ChanType,
		*ast.DeclStmt,
		*ast.DeferStmt,
		*ast.Ellipsis,
		*ast.EmptyStmt,
		*ast.ExprStmt,
		*ast.ForStmt,
		*ast.FuncDecl,
		*ast.FuncLit,
		*ast.FuncType,
		*ast.GoStmt,
		*ast.IfStmt,
		*ast.ImportSpec,
		*ast.IncDecStmt,
		*ast.IndexExpr,
		*ast.InterfaceType,
		*ast.KeyValueExpr,
		*ast.LabeledStmt,
		*ast.MapType,
		*ast.ParenExpr,
		*ast.RangeStmt,
		*ast.SelectorExpr,
		*ast.SelectStmt,
		*ast.SendStmt,
		*ast.SliceExpr,
		*ast.StarExpr,
		*ast.StructType,
		*ast.SwitchStmt,
		*ast.TypeAssertExpr,
		*ast.TypeSpec,
		*ast.TypeSwitchStmt,
		*ast.UnaryExpr:
		return false
	case
		*ast.AssignStmt,
		*ast.BlockStmt,
		*ast.CallExpr,
		*ast.CaseClause,
		*ast.CommClause,
		*ast.CompositeLit,
		*ast.Field,
		*ast.FieldList,
		*ast.GenDecl,
		*ast.IndexListExpr,
		*ast.ReturnStmt,
		*ast.ValueSpec:
		return true
	}
	return false
}
