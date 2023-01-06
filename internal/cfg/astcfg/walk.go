package astcfg

import (
	"go/ast"
)

var NodeSlice ast.Node // Used by "walkEachNd" in parentTrace to imply traversed struct field is a slice type eg: []Expr, []Stmt

func walkEachNode[T ast.Node](nodeSlice []T, parentTrace []ast.Node, childIndexTrace []int, callback func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int)) {
	parentTrace = append(parentTrace, NodeSlice)
	for i, v := range nodeSlice {
		walkHelper(v, parentTrace, append(childIndexTrace, i), callback)
	}
}

func walkHelper(n ast.Node, parentTrace []ast.Node, childIndexTrace []int, callback func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int)) {
	callback(n, parentTrace, childIndexTrace)

	if n == nil {
		return
	}

	parentTrace = append(parentTrace, n)
	switch n := n.(type) {

	// MARK: Comments and fields

	case *ast.Comment:
		// nothing to do

	case *ast.CommentGroup:
		walkEachNode(n.List, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.Field:
		walkEachNode(n.Names, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Type, parentTrace, append(childIndexTrace, 1), callback)
		walkHelper(n.Tag, parentTrace, append(childIndexTrace, 2), callback)

	case *ast.FieldList:
		walkEachNode(n.List, parentTrace, append(childIndexTrace, 0), callback)

	// MARK: Expressions

	case *ast.BadExpr, *ast.Ident, *ast.BasicLit:
		// nothing to do

	case *ast.Ellipsis:
		walkHelper(n.Elt, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.FuncLit:
		walkHelper(n.Type, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Body, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.CompositeLit:
		walkHelper(n.Type, parentTrace, append(childIndexTrace, 0), callback)
		walkEachNode(n.Elts, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.ParenExpr:
		walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.SelectorExpr:
		walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Sel, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.IndexExpr:
		walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Index, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.IndexListExpr:
		walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		walkEachNode(n.Indices, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.SliceExpr:
		walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Low, parentTrace, append(childIndexTrace, 1), callback)
		walkHelper(n.High, parentTrace, append(childIndexTrace, 2), callback)
		walkHelper(n.Max, parentTrace, append(childIndexTrace, 3), callback)

	case *ast.TypeAssertExpr:
		walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Type, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.CallExpr:
		walkHelper(n.Fun, parentTrace, append(childIndexTrace, 0), callback)
		walkEachNode(n.Args, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.StarExpr:
		walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.UnaryExpr:
		walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.BinaryExpr:
		walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Y, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.KeyValueExpr:
		walkHelper(n.Key, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Value, parentTrace, append(childIndexTrace, 1), callback)

	// MARK: Types

	case *ast.ArrayType:
		walkHelper(n.Len, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Elt, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.StructType:
		walkHelper(n.Fields, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.FuncType:
		walkHelper(n.TypeParams, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Params, parentTrace, append(childIndexTrace, 1), callback)
		walkHelper(n.Results, parentTrace, append(childIndexTrace, 2), callback)

	case *ast.InterfaceType:
		walkHelper(n.Methods, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.MapType:
		walkHelper(n.Key, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Value, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.ChanType:
		walkHelper(n.Value, parentTrace, append(childIndexTrace, 0), callback)

	// MARK: Statements

	case *ast.BadStmt:
		// nothing to do

	case *ast.DeclStmt:
		walkHelper(n.Decl, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.EmptyStmt:
		// nothing to do

	case *ast.LabeledStmt:
		walkHelper(n.Label, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Stmt, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.ExprStmt:
		walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.SendStmt:
		walkHelper(n.Chan, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Value, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.IncDecStmt:
		walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.AssignStmt:
		walkEachNode(n.Lhs, parentTrace, append(childIndexTrace, 0), callback)
		walkEachNode(n.Rhs, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.GoStmt:
		walkHelper(n.Call, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.DeferStmt:
		walkHelper(n.Call, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.ReturnStmt:
		walkEachNode(n.Results, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.BranchStmt:
		walkHelper(n.Label, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.BlockStmt:
		walkEachNode(n.List, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.IfStmt:
		walkHelper(n.Init, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Cond, parentTrace, append(childIndexTrace, 1), callback)
		walkHelper(n.Body, parentTrace, append(childIndexTrace, 2), callback)
		walkHelper(n.Else, parentTrace, append(childIndexTrace, 3), callback)

	case *ast.CaseClause:
		walkEachNode(n.List, parentTrace, append(childIndexTrace, 0), callback)
		walkEachNode(n.Body, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.SwitchStmt:
		walkHelper(n.Init, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Tag, parentTrace, append(childIndexTrace, 1), callback)
		walkHelper(n.Body, parentTrace, append(childIndexTrace, 2), callback)

	case *ast.TypeSwitchStmt:
		walkHelper(n.Init, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Assign, parentTrace, append(childIndexTrace, 1), callback)
		walkHelper(n.Body, parentTrace, append(childIndexTrace, 2), callback)

	case *ast.CommClause:
		walkHelper(n.Comm, parentTrace, append(childIndexTrace, 0), callback)
		walkEachNode(n.Body, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.SelectStmt:
		walkHelper(n.Body, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.ForStmt:
		walkHelper(n.Init, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Cond, parentTrace, append(childIndexTrace, 1), callback)
		walkHelper(n.Post, parentTrace, append(childIndexTrace, 2), callback)
		walkHelper(n.Body, parentTrace, append(childIndexTrace, 3), callback)

	case *ast.RangeStmt:
		walkHelper(n.Key, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Value, parentTrace, append(childIndexTrace, 1), callback)
		walkHelper(n.X, parentTrace, append(childIndexTrace, 2), callback)
		walkHelper(n.Body, parentTrace, append(childIndexTrace, 3), callback)

	// MARK: Declarations

	case *ast.ImportSpec:
		walkHelper(n.Name, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Path, parentTrace, append(childIndexTrace, 1), callback)

	case *ast.ValueSpec:
		walkEachNode(n.Names, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Type, parentTrace, append(childIndexTrace, 1), callback)
		walkEachNode(n.Values, parentTrace, append(childIndexTrace, 2), callback)

	case *ast.TypeSpec:
		walkHelper(n.Name, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.TypeParams, parentTrace, append(childIndexTrace, 1), callback)
		walkHelper(n.Type, parentTrace, append(childIndexTrace, 2), callback)

	case *ast.BadDecl:
		// nothing to do

	case *ast.GenDecl:
		walkEachNode(n.Specs, parentTrace, append(childIndexTrace, 0), callback)

	case *ast.FuncDecl:
		walkHelper(n.Recv, parentTrace, append(childIndexTrace, 0), callback)
		walkHelper(n.Name, parentTrace, append(childIndexTrace, 1), callback)
		walkHelper(n.Type, parentTrace, append(childIndexTrace, 2), callback)
		walkHelper(n.Body, parentTrace, append(childIndexTrace, 3), callback)

		// case *ast.File:
		// 	walkHelper(n.Name, parentTrace, append(childIndexTrace, 0), callback)
		// 	walkEachNd(n.Decls, parentTrace, append(childIndexTrace, 1), callback)

		// case *ast.Package:
		// 	for _, f := range n.Files {
		// 		walkHelper(f, parentTrace, append(childIndexTrace, 0), callback)
		// 	}

		// }
	}
}

// walk calls the callback function for every joint (eg. []Stmt..., []Expr..., Stmt, Expr, Token)
func Walk(root ast.Node, callback func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int)) {
	walkHelper(root, []ast.Node{}, []int{}, callback)
}

// Inspects (Walks) the partial-AST of root; chooses one insertion point and make an appropriate type Node append
func Enrich(root ast.Node) {

}
