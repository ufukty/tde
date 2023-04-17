package clean_clone

import "go/ast"

func Expr(src ast.Expr) ast.Expr {
	if src == nil {
		return nil
	}

	switch src := src.(type) {
	case *ast.BadExpr:
		return BadExpr(src)
	case *ast.Ident:
		return Ident(src)
	case *ast.Ellipsis:
		return Ellipsis(src)
	case *ast.BasicLit:
		return BasicLit(src)
	case *ast.FuncLit:
		return FuncLit(src)
	case *ast.CompositeLit:
		return CompositeLit(src)
	case *ast.ParenExpr:
		return ParenExpr(src)
	case *ast.SelectorExpr:
		return SelectorExpr(src)
	case *ast.IndexExpr:
		return IndexExpr(src)
	case *ast.IndexListExpr:
		return IndexListExpr(src)
	case *ast.SliceExpr:
		return SliceExpr(src)
	case *ast.TypeAssertExpr:
		return TypeAssertExpr(src)
	case *ast.CallExpr:
		return CallExpr(src)
	case *ast.StarExpr:
		return StarExpr(src)
	case *ast.UnaryExpr:
		return UnaryExpr(src)
	case *ast.BinaryExpr:
		return BinaryExpr(src)
	case *ast.KeyValueExpr:
		return KeyValueExpr(src)
	case *ast.ArrayType:
		return ArrayType(src)
	case *ast.StructType:
		return StructType(src)
	case *ast.FuncType:
		return FuncType(src)
	case *ast.InterfaceType:
		return InterfaceType(src)
	case *ast.MapType:
		return MapType(src)
	case *ast.ChanType:
		return ChanType(src)
	}

	return nil
}

func Stmt(src ast.Stmt) ast.Stmt {
	if src == nil {
		return nil
	}

	switch src := src.(type) {
	case *ast.BadStmt:
		return BadStmt(src)
	case *ast.DeclStmt:
		return DeclStmt(src)
	case *ast.EmptyStmt:
		return EmptyStmt(src)
	case *ast.LabeledStmt:
		return LabeledStmt(src)
	case *ast.ExprStmt:
		return ExprStmt(src)
	case *ast.SendStmt:
		return SendStmt(src)
	case *ast.IncDecStmt:
		return IncDecStmt(src)
	case *ast.AssignStmt:
		return AssignStmt(src)
	case *ast.GoStmt:
		return GoStmt(src)
	case *ast.DeferStmt:
		return DeferStmt(src)
	case *ast.ReturnStmt:
		return ReturnStmt(src)
	case *ast.BranchStmt:
		return BranchStmt(src)
	case *ast.BlockStmt:
		return BlockStmt(src)
	case *ast.IfStmt:
		return IfStmt(src)
	case *ast.CaseClause:
		return CaseClause(src)
	case *ast.SwitchStmt:
		return SwitchStmt(src)
	case *ast.TypeSwitchStmt:
		return TypeSwitchStmt(src)
	case *ast.CommClause:
		return CommClause(src)
	case *ast.SelectStmt:
		return SelectStmt(src)
	case *ast.ForStmt:
		return ForStmt(src)
	case *ast.RangeStmt:
		return RangeStmt(src)
	}

	return nil
}

func Decl(src ast.Decl) ast.Decl {
	if src == nil {
		return nil
	}

	switch src := src.(type) {
	case *ast.BadDecl:
		return BadDecl(src)
	case *ast.GenDecl:
		return GenDecl(src)
	case *ast.FuncDecl:
		return FuncDecl(src)
	}

	return nil
}

func Spec(src ast.Spec) ast.Spec {
	if src == nil {
		return nil
	}

	switch src := src.(type) {
	case *ast.ImportSpec:
		return ImportSpec(src)
	case *ast.TypeSpec:
		return TypeSpec(src)
	case *ast.ValueSpec:
		return ValueSpec(src)
	}

	return nil
}
