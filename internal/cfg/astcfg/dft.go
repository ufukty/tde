package astcfg

import "go/ast"

func dft(node ast.Node, depth int, maxDepth int, callback func(node ast.Node, ancestry []ast.Node, childIndexes []int)) {
	if depth == maxDepth {
		return
	}

	switch node := node.(type) {
		case *ast.ArrayType:
		case *ast.AssignStmt:
		case *ast.BasicLit:
		case *ast.BinaryExpr:
		case *ast.BlockStmt:
		case *ast.BranchStmt:
		case *ast.CallExpr:
		case *ast.CaseClause:
		case *ast.ChanType:
		case *ast.CommClause:
		case *ast.CompositeLit:
		case *ast.DeclStmt:
		case *ast.DeferStmt:
		case *ast.Ellipsis:
		case *ast.EmptyStmt:
		case *ast.ExprStmt:
		case *ast.ForStmt:
		case *ast.FuncDecl:
		case *ast.FuncLit:
		case *ast.FuncType:
		case *ast.GenDecl:
		case *ast.GoStmt:
		case *ast.Ident:
		case *ast.IfStmt:
		case *ast.ImportSpec:
		case *ast.IncDecStmt:
		case *ast.IndexExpr:
		case *ast.IndexListExpr:
		case *ast.InterfaceType:
		case *ast.KeyValueExpr:
		case *ast.LabeledStmt:
		case *ast.MapType:
		case *ast.ParenExpr:
		case *ast.RangeStmt:
		case *ast.ReturnStmt:
		case *ast.SelectorExpr:
		case *ast.SelectStmt:
		case *ast.SendStmt:
		case *ast.SliceExpr:
		case *ast.StarExpr:
		case *ast.StructType:
		case *ast.SwitchStmt:
		case *ast.TypeAssertExpr:
		case *ast.TypeSpec:
		case *ast.TypeSwitchStmt:
		case *ast.UnaryExpr:
		case *ast.ValueSpec:
	}
}

func DFT(root ast.Node, maxDepth int, callback func(node ast.Node, ancestry []ast.Node, childIndexes []int)) {
	dft(root, 0, maxDepth, callback)
}
