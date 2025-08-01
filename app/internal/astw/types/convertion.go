package types

import "go/ast"

func GetNodeTypeForASTNode(n ast.Node) NodeType {
	switch n.(type) {
	case *ast.ArrayType:
		return ArrayType
	case *ast.AssignStmt:
		return AssignStmt
	case *ast.BadDecl:
		return BadDecl
	case *ast.BadExpr:
		return BadExpr
	case *ast.BadStmt:
		return BadStmt
	case *ast.BasicLit:
		return BasicLit
	case *ast.BinaryExpr:
		return BinaryExpr
	case *ast.BlockStmt:
		return BlockStmt
	case *ast.BranchStmt:
		return BranchStmt
	case *ast.CallExpr:
		return CallExpr
	case *ast.CaseClause:
		return CaseClause
	case *ast.ChanType:
		return ChanType
	case *ast.CommClause:
		return CommClause
	case *ast.Comment:
		return Comment
	case *ast.CommentGroup:
		return CommentGroup
	case *ast.CompositeLit:
		return CompositeLit
	case *ast.DeclStmt:
		return DeclStmt
	case *ast.DeferStmt:
		return DeferStmt
	case *ast.Ellipsis:
		return Ellipsis
	case *ast.EmptyStmt:
		return EmptyStmt
	case *ast.ExprStmt:
		return ExprStmt
	case *ast.Field:
		return Field
	case *ast.FieldList:
		return FieldList
	case *ast.File:
		return File
	case *ast.ForStmt:
		return ForStmt
	case *ast.FuncDecl:
		return FuncDecl
	case *ast.FuncLit:
		return FuncLit
	case *ast.FuncType:
		return FuncType
	case *ast.GenDecl:
		return GenDecl
	case *ast.GoStmt:
		return GoStmt
	case *ast.Ident:
		return Ident
	case *ast.IfStmt:
		return IfStmt
	case *ast.ImportSpec:
		return ImportSpec
	case *ast.IncDecStmt:
		return IncDecStmt
	case *ast.IndexExpr:
		return IndexExpr
	case *ast.IndexListExpr:
		return IndexListExpr
	case *ast.InterfaceType:
		return InterfaceType
	case *ast.KeyValueExpr:
		return KeyValueExpr
	case *ast.LabeledStmt:
		return LabeledStmt
	case *ast.MapType:
		return MapType
	case *ast.Package:
		return Package
	case *ast.ParenExpr:
		return ParenExpr
	case *ast.RangeStmt:
		return RangeStmt
	case *ast.ReturnStmt:
		return ReturnStmt
	case *ast.SelectorExpr:
		return SelectorExpr
	case *ast.SelectStmt:
		return SelectStmt
	case *ast.SendStmt:
		return SendStmt
	case *ast.SliceExpr:
		return SliceExpr
	case *ast.StarExpr:
		return StarExpr
	case *ast.StructType:
		return StructType
	case *ast.SwitchStmt:
		return SwitchStmt
	case *ast.TypeAssertExpr:
		return TypeAssertExpr
	case *ast.TypeSpec:
		return TypeSpec
	case *ast.TypeSwitchStmt:
		return TypeSwitchStmt
	case *ast.UnaryExpr:
		return UnaryExpr
	case *ast.ValueSpec:
		return ValueSpec
	default:
		panic("Unhandled condition GetNodeTypeForASTNode")
	}
}
