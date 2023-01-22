package ast_wrapper

import "go/ast"

// Use to represent nullable, Node-like (ast.Node etc.) or Node-slice ([]ast.Node etc.) components of the AST.
type TraversableNode struct {
	Value        any
	ExpectedType NodeType // This usually will be filled while visiting the parent node by leveraging the struct definitions instead of type checking on value because it can be nil too. Also, Traverse() callbacks can leverage for learning which type of node is needed to construct.
	IsNil        bool
}

func GetTraversableNodeForASTNode(node ast.Node) TraversableNode {
	return TraversableNode{
		Value:        node,
		ExpectedType: GetNodeTypeForASTNode(node),
		IsNil:        isNodeNil(node),
	}
}

func TraversableNodesFromSliceTypeNode(tNode *TraversableNode) (tNodes []TraversableNode) {
	if tNode == nil {
		return
	}

	switch slice := tNode.Value.(type) {
	case []*ast.Comment:
		for _, item := range slice {
			tNodes = append(tNodes, TraversableNode{item, Comment, item == nil})
		}
	case []*ast.CommentGroup:
		for _, item := range slice {
			tNodes = append(tNodes, TraversableNode{item, CommentGroup, item == nil})
		}
	case []*ast.ImportSpec:
		for _, item := range slice {
			tNodes = append(tNodes, TraversableNode{item, ImportSpec, item == nil})
		}
	case []*ast.Ident:
		for _, item := range slice {
			tNodes = append(tNodes, TraversableNode{item, Ident, item == nil})
		}
	case []*ast.Field:
		for _, item := range slice {
			tNodes = append(tNodes, TraversableNode{item, Field, item == nil})
		}
	case []ast.Stmt:
		for _, item := range slice {
			tNodes = append(tNodes, TraversableNode{item, GetNodeTypeForASTNode(item), item == nil})
		}
	case []ast.Decl:
		for _, item := range slice {
			tNodes = append(tNodes, TraversableNode{item, GetNodeTypeForASTNode(item), item == nil})
		}
	case []ast.Spec:
		for _, item := range slice {
			tNodes = append(tNodes, TraversableNode{item, GetNodeTypeForASTNode(item), item == nil})
		}
	case []ast.Expr:
		for _, item := range slice {
			tNodes = append(tNodes, TraversableNode{item, GetNodeTypeForASTNode(item), item == nil})
		}
	case []ast.Node:
		for _, item := range slice {
			tNodes = append(tNodes, TraversableNode{item, GetNodeTypeForASTNode(item), item == nil})
		}
	}
	return
}

func TraversableNodesFromInterfaceTypeNode(tNode *TraversableNode) []TraversableNode {
	if tNode.IsNil {
		return []TraversableNode{}
	}
	switch (*tNode).Value.(type) {
	case ast.Expr:
		return TraversableNodesFromConcreteTypeNode(tNode)
	case ast.Stmt:
		return TraversableNodesFromConcreteTypeNode(tNode)
	case ast.Decl:
		return TraversableNodesFromConcreteTypeNode(tNode)
	case ast.Spec:
		return TraversableNodesFromConcreteTypeNode(tNode)
	}
	return []TraversableNode{}
}

func TraversableNodesFromConcreteTypeNode(tNode *TraversableNode) []TraversableNode {
	if (*tNode).IsNil {
		return []TraversableNode{}
	}

	switch n := (*tNode).Value.(type) {

	case *ast.Comment:
		return []TraversableNode{}

	case *ast.CommentGroup:
		return []TraversableNode{}

	case *ast.Field:
		return []TraversableNode{
			{n.Names, IdentSlice, n.Names == nil},
			{n.Type, Expr, n.Type == nil},
			{n.Tag, BasicLit, n.Tag == nil},
		}

	case *ast.FieldList:
		return []TraversableNode{
			{n.List, FieldSlice, n.List == nil},
		}

	case *ast.BadExpr:
		return []TraversableNode{}

	case *ast.Ident:
		return []TraversableNode{}

	case *ast.Ellipsis:
		return []TraversableNode{
			{n.Elt, Expr, n.Elt == nil},
		}

	case *ast.BasicLit:
		return []TraversableNode{}

	case *ast.FuncLit:
		return []TraversableNode{
			{n.Type, FuncType, n.Type == nil},
			{n.Body, BlockStmt, n.Body == nil},
		}

	case *ast.CompositeLit:
		return []TraversableNode{
			{n.Type, Expr, n.Type == nil},
			{n.Elts, ExprSlice, n.Elts == nil},
		}

	case *ast.ParenExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil},
		}

	case *ast.SelectorExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil},
			{n.Sel, Ident, n.Sel == nil},
		}

	case *ast.IndexExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil},
			{n.Index, Expr, n.Index == nil},
		}

	case *ast.IndexListExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil},
			{n.Indices, ExprSlice, n.Indices == nil},
		}

	case *ast.SliceExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil},
			{n.Low, Expr, n.Low == nil},
			{n.High, Expr, n.High == nil},
			{n.Max, Expr, n.Max == nil},
		}

	case *ast.TypeAssertExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil},
			{n.Type, Expr, n.Type == nil},
		}

	case *ast.CallExpr:
		return []TraversableNode{
			{n.Fun, Expr, n.Fun == nil},
			{n.Args, ExprSlice, n.Args == nil},
		}

	case *ast.StarExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil},
		}

	case *ast.UnaryExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil},
		}

	case *ast.BinaryExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil},
			{n.Y, Expr, n.Y == nil},
		}

	case *ast.KeyValueExpr:
		return []TraversableNode{
			{n.Key, Expr, n.Key == nil},
			{n.Value, Expr, n.Value == nil},
		}

	case *ast.ArrayType:
		return []TraversableNode{
			{n.Len, Expr, n.Len == nil},
			{n.Elt, Expr, n.Elt == nil},
		}

	case *ast.StructType:
		return []TraversableNode{
			{n.Fields, FieldList, n.Fields == nil},
		}

	case *ast.FuncType:
		return []TraversableNode{
			{n.TypeParams, FieldList, n.TypeParams == nil},
			{n.Params, FieldList, n.Params == nil},
			{n.Results, FieldList, n.Results == nil},
		}

	case *ast.InterfaceType:
		return []TraversableNode{
			{n.Methods, FieldList, n.Methods == nil},
		}

	case *ast.MapType:
		return []TraversableNode{
			{n.Key, Expr, n.Key == nil},
			{n.Value, Expr, n.Value == nil},
		}

	case *ast.ChanType:
		return []TraversableNode{
			{n.Value, Expr, n.Value == nil},
		}

	case *ast.BadStmt:
		return []TraversableNode{}

	case *ast.DeclStmt:
		return []TraversableNode{
			{n.Decl, Decl, n.Decl == nil},
		}

	case *ast.EmptyStmt:
		return []TraversableNode{}

	case *ast.LabeledStmt:
		return []TraversableNode{
			{n.Label, Ident, n.Label == nil},
			{n.Stmt, Stmt, n.Stmt == nil},
		}

	case *ast.ExprStmt:
		return []TraversableNode{
			{n.X, Expr, n.X == nil},
		}

	case *ast.SendStmt:
		return []TraversableNode{
			{n.Chan, Expr, n.Chan == nil},
			{n.Value, Expr, n.Value == nil},
		}

	case *ast.IncDecStmt:
		return []TraversableNode{
			{n.X, Expr, n.X == nil},
		}

	case *ast.AssignStmt:
		return []TraversableNode{
			{n.Lhs, ExprSlice, n.Lhs == nil},
			{n.Rhs, ExprSlice, n.Rhs == nil},
		}

	case *ast.GoStmt:
		return []TraversableNode{
			{n.Call, CallExpr, n.Call == nil},
		}

	case *ast.DeferStmt:
		return []TraversableNode{
			{n.Call, CallExpr, n.Call == nil},
		}

	case *ast.ReturnStmt:
		return []TraversableNode{
			{n.Results, ExprSlice, n.Results == nil},
		}

	case *ast.BranchStmt:
		return []TraversableNode{
			{n.Label, Ident, n.Label == nil},
		}

	case *ast.BlockStmt:
		return []TraversableNode{
			{n.List, StmtSlice, n.List == nil},
		}

	case *ast.IfStmt:
		return []TraversableNode{
			{n.Init, Stmt, n.Init == nil},
			{n.Cond, Expr, n.Cond == nil},
			{n.Body, BlockStmt, n.Body == nil},
			{n.Else, Stmt, n.Else == nil},
		}

	case *ast.CaseClause:
		return []TraversableNode{
			{n.List, ExprSlice, n.List == nil},
			{n.Body, StmtSlice, n.Body == nil},
		}

	case *ast.SwitchStmt:
		return []TraversableNode{
			{n.Init, Stmt, n.Init == nil},
			{n.Tag, Expr, n.Tag == nil},
			{n.Body, BlockStmt, n.Body == nil},
		}

	case *ast.TypeSwitchStmt:
		return []TraversableNode{
			{n.Init, Stmt, n.Init == nil},
			{n.Assign, Stmt, n.Assign == nil},
			{n.Body, BlockStmt, n.Body == nil},
		}

	case *ast.CommClause:
		return []TraversableNode{
			{n.Comm, Stmt, n.Comm == nil},
			{n.Body, StmtSlice, n.Body == nil},
		}

	case *ast.SelectStmt:
		return []TraversableNode{
			{n.Body, BlockStmt, n.Body == nil},
		}

	case *ast.ForStmt:
		return []TraversableNode{
			{n.Init, Stmt, n.Init == nil},
			{n.Cond, Expr, n.Cond == nil},
			{n.Post, Stmt, n.Post == nil},
			{n.Body, BlockStmt, n.Body == nil},
		}

	case *ast.RangeStmt:
		return []TraversableNode{
			{n.Key, Expr, n.Key == nil},
			{n.Value, Expr, n.Value == nil},
			{n.X, Expr, n.X == nil},
			{n.Body, BlockStmt, n.Body == nil},
		}

	case *ast.ImportSpec:
		return []TraversableNode{
			{n.Name, Ident, n.Name == nil},
			{n.Path, BasicLit, n.Path == nil},
		}

	case *ast.ValueSpec:
		return []TraversableNode{
			{n.Names, IdentSlice, n.Names == nil},
			{n.Type, Expr, n.Type == nil},
			{n.Values, ExprSlice, n.Values == nil},
		}

	case *ast.TypeSpec:
		return []TraversableNode{
			{n.Name, Ident, n.Name == nil},
			{n.TypeParams, FieldList, n.TypeParams == nil},
			{n.Type, Expr, n.Type == nil},
		}

	case *ast.BadDecl:
		return []TraversableNode{}

	case *ast.GenDecl:
		return []TraversableNode{
			{n.Specs, SpecSlice, n.Specs == nil},
		}

	case *ast.FuncDecl:
		return []TraversableNode{
			{n.Recv, FieldList, n.Recv == nil},
			{n.Name, Ident, n.Name == nil},
			{n.Type, FuncType, n.Type == nil},
			{n.Body, BlockStmt, n.Body == nil},
		}

	case *ast.File:
		return []TraversableNode{
			{n.Name, Ident, n.Name == nil},
			{n.Decls, DeclSlice, n.Decls == nil},
			{n.Imports, ImportSpecSlice, n.Imports == nil},
			{n.Unresolved, IdentSlice, n.Unresolved == nil},
		}

		// case ast*.*ast.Package:
		// 	return []TraversableNode{
		// 	{n.Files, map[string]*File},
		// }

	}
	return []TraversableNode{}
}
