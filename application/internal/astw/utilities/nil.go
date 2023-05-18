package utilities

import "go/ast"

// This mess is because simple comparison of n == nil doesn't work.
func IsNodeNil(n ast.Node) bool {
	switch n := n.(type) {
	case *ast.ArrayType:
		return n == nil
	case *ast.AssignStmt:
		return n == nil
	case *ast.BadDecl:
		return n == nil
	case *ast.BadExpr:
		return n == nil
	case *ast.BadStmt:
		return n == nil
	case *ast.BasicLit:
		return n == nil
	case *ast.BinaryExpr:
		return n == nil
	case *ast.BlockStmt:
		return n == nil
	case *ast.BranchStmt:
		return n == nil
	case *ast.CallExpr:
		return n == nil
	case *ast.CaseClause:
		return n == nil
	case *ast.ChanType:
		return n == nil
	case *ast.CommClause:
		return n == nil
	case *ast.Comment:
		return n == nil
	case *ast.CommentGroup:
		return n == nil
	case *ast.CompositeLit:
		return n == nil
	case *ast.DeclStmt:
		return n == nil
	case *ast.DeferStmt:
		return n == nil
	case *ast.Ellipsis:
		return n == nil
	case *ast.EmptyStmt:
		return n == nil
	case *ast.ExprStmt:
		return n == nil
	case *ast.Field:
		return n == nil
	case *ast.FieldList:
		return n == nil
	case *ast.File:
		return n == nil
	case *ast.ForStmt:
		return n == nil
	case *ast.FuncDecl:
		return n == nil
	case *ast.FuncLit:
		return n == nil
	case *ast.FuncType:
		return n == nil
	case *ast.GenDecl:
		return n == nil
	case *ast.GoStmt:
		return n == nil
	case *ast.Ident:
		return n == nil
	case *ast.IfStmt:
		return n == nil
	case *ast.ImportSpec:
		return n == nil
	case *ast.IncDecStmt:
		return n == nil
	case *ast.IndexExpr:
		return n == nil
	case *ast.IndexListExpr:
		return n == nil
	case *ast.InterfaceType:
		return n == nil
	case *ast.KeyValueExpr:
		return n == nil
	case *ast.LabeledStmt:
		return n == nil
	case *ast.MapType:
		return n == nil
	case *ast.Package:
		return n == nil
	case *ast.ParenExpr:
		return n == nil
	case *ast.RangeStmt:
		return n == nil
	case *ast.ReturnStmt:
		return n == nil
	case *ast.SelectorExpr:
		return n == nil
	case *ast.SelectStmt:
		return n == nil
	case *ast.SendStmt:
		return n == nil
	case *ast.SliceExpr:
		return n == nil
	case *ast.StarExpr:
		return n == nil
	case *ast.StructType:
		return n == nil
	case *ast.SwitchStmt:
		return n == nil
	case *ast.TypeAssertExpr:
		return n == nil
	case *ast.TypeSpec:
		return n == nil
	case *ast.TypeSwitchStmt:
		return n == nil
	case *ast.UnaryExpr:
		return n == nil
	case *ast.ValueSpec:
		return n == nil

	case ast.Expr:
		return n == nil
	case ast.Stmt:
		return n == nil
	case ast.Decl:
		return n == nil
	case ast.Spec:
		return n == nil

	case ast.Node:
		return n == nil
	}
	return true
}
