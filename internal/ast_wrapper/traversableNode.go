package ast_wrapper

import "go/ast"

// Use to represent nullable, Node-like (ast.Node etc.) or Node-slice ([]ast.Node etc.) components of the AST.
type TraversableNode struct {
	Value        any
	ExpectedType NodeType // This usually will be filled while visiting the parent node by leveraging the struct definitions instead of type checking on value because it can be nil too. Also, Traverse() callbacks can leverage for learning which type of node is needed to construct.
	IsNil        bool
	Parent       *TraversableNode
}

func GetTraversableNodeForASTNode(node ast.Node) TraversableNode {
	return TraversableNode{
		Value:        node,
		ExpectedType: GetNodeTypeForASTNode(node),
		IsNil:        isNodeNil(node),
		Parent:       nil,
	}
}

func TraversableNodesFromSliceTypeNode(tNode *TraversableNode) (tSubnodes []TraversableNode) {
	if tNode == nil {
		return
	}

	switch slice := tNode.Value.(type) {
	case []*ast.Comment:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, TraversableNode{item, Comment, item == nil, tNode})
		}
	case []*ast.CommentGroup:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, TraversableNode{item, CommentGroup, item == nil, tNode})
		}
	case []*ast.ImportSpec:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, TraversableNode{item, ImportSpec, item == nil, tNode})
		}
	case []*ast.Ident:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, TraversableNode{item, Ident, item == nil, tNode})
		}
	case []*ast.Field:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, TraversableNode{item, Field, item == nil, tNode})
		}
	case []ast.Stmt:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, TraversableNode{item, GetNodeTypeForASTNode(item), item == nil, tNode})
		}
	case []ast.Decl:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, TraversableNode{item, GetNodeTypeForASTNode(item), item == nil, tNode})
		}
	case []ast.Spec:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, TraversableNode{item, GetNodeTypeForASTNode(item), item == nil, tNode})
		}
	case []ast.Expr:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, TraversableNode{item, GetNodeTypeForASTNode(item), item == nil, tNode})
		}
		// case []ast.Node:
		// 	for _, item := range slice {
		// 		tSubnodes = append(tSubnodes, TraversableNode{item, GetNodeTypeForASTNode(item), item == nil, tNode})
		// 	}
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
			{n.Names, IdentSlice, n.Names == nil, tNode},
			{n.Type, Expr, n.Type == nil, tNode},
			{n.Tag, BasicLit, n.Tag == nil, tNode},
		}

	case *ast.FieldList:
		return []TraversableNode{
			{n.List, FieldSlice, n.List == nil, tNode},
		}

	case *ast.BadExpr:
		return []TraversableNode{}

	case *ast.Ident:
		return []TraversableNode{}

	case *ast.Ellipsis:
		return []TraversableNode{
			{n.Elt, Expr, n.Elt == nil, tNode},
		}

	case *ast.BasicLit:
		return []TraversableNode{}

	case *ast.FuncLit:
		return []TraversableNode{
			{n.Type, FuncType, n.Type == nil, tNode},
			{n.Body, BlockStmt, n.Body == nil, tNode},
		}

	case *ast.CompositeLit:
		return []TraversableNode{
			{n.Type, Expr, n.Type == nil, tNode},
			{n.Elts, ExprSlice, n.Elts == nil, tNode},
		}

	case *ast.ParenExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil, tNode},
		}

	case *ast.SelectorExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil, tNode},
			{n.Sel, Ident, n.Sel == nil, tNode},
		}

	case *ast.IndexExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil, tNode},
			{n.Index, Expr, n.Index == nil, tNode},
		}

	case *ast.IndexListExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil, tNode},
			{n.Indices, ExprSlice, n.Indices == nil, tNode},
		}

	case *ast.SliceExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil, tNode},
			{n.Low, Expr, n.Low == nil, tNode},
			{n.High, Expr, n.High == nil, tNode},
			{n.Max, Expr, n.Max == nil, tNode},
		}

	case *ast.TypeAssertExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil, tNode},
			{n.Type, Expr, n.Type == nil, tNode},
		}

	case *ast.CallExpr:
		return []TraversableNode{
			{n.Fun, Expr, n.Fun == nil, tNode},
			{n.Args, ExprSlice, n.Args == nil, tNode},
		}

	case *ast.StarExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil, tNode},
		}

	case *ast.UnaryExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil, tNode},
		}

	case *ast.BinaryExpr:
		return []TraversableNode{
			{n.X, Expr, n.X == nil, tNode},
			{n.Y, Expr, n.Y == nil, tNode},
		}

	case *ast.KeyValueExpr:
		return []TraversableNode{
			{n.Key, Expr, n.Key == nil, tNode},
			{n.Value, Expr, n.Value == nil, tNode},
		}

	case *ast.ArrayType:
		return []TraversableNode{
			{n.Len, Expr, n.Len == nil, tNode},
			{n.Elt, Expr, n.Elt == nil, tNode},
		}

	case *ast.StructType:
		return []TraversableNode{
			{n.Fields, FieldList, n.Fields == nil, tNode},
		}

	case *ast.FuncType:
		return []TraversableNode{
			{n.TypeParams, FieldList, n.TypeParams == nil, tNode},
			{n.Params, FieldList, n.Params == nil, tNode},
			{n.Results, FieldList, n.Results == nil, tNode},
		}

	case *ast.InterfaceType:
		return []TraversableNode{
			{n.Methods, FieldList, n.Methods == nil, tNode},
		}

	case *ast.MapType:
		return []TraversableNode{
			{n.Key, Expr, n.Key == nil, tNode},
			{n.Value, Expr, n.Value == nil, tNode},
		}

	case *ast.ChanType:
		return []TraversableNode{
			{n.Value, Expr, n.Value == nil, tNode},
		}

	case *ast.BadStmt:
		return []TraversableNode{}

	case *ast.DeclStmt:
		return []TraversableNode{
			{n.Decl, Decl, n.Decl == nil, tNode},
		}

	case *ast.EmptyStmt:
		return []TraversableNode{}

	case *ast.LabeledStmt:
		return []TraversableNode{
			{n.Label, Ident, n.Label == nil, tNode},
			{n.Stmt, Stmt, n.Stmt == nil, tNode},
		}

	case *ast.ExprStmt:
		return []TraversableNode{
			{n.X, Expr, n.X == nil, tNode},
		}

	case *ast.SendStmt:
		return []TraversableNode{
			{n.Chan, Expr, n.Chan == nil, tNode},
			{n.Value, Expr, n.Value == nil, tNode},
		}

	case *ast.IncDecStmt:
		return []TraversableNode{
			{n.X, Expr, n.X == nil, tNode},
		}

	case *ast.AssignStmt:
		return []TraversableNode{
			{n.Lhs, ExprSlice, n.Lhs == nil, tNode},
			{n.Rhs, ExprSlice, n.Rhs == nil, tNode},
		}

	case *ast.GoStmt:
		return []TraversableNode{
			{n.Call, CallExpr, n.Call == nil, tNode},
		}

	case *ast.DeferStmt:
		return []TraversableNode{
			{n.Call, CallExpr, n.Call == nil, tNode},
		}

	case *ast.ReturnStmt:
		return []TraversableNode{
			{n.Results, ExprSlice, n.Results == nil, tNode},
		}

	case *ast.BranchStmt:
		return []TraversableNode{
			{n.Label, Ident, n.Label == nil, tNode},
		}

	case *ast.BlockStmt:
		return []TraversableNode{
			{n.List, StmtSlice, n.List == nil, tNode},
		}

	case *ast.IfStmt:
		return []TraversableNode{
			{n.Init, Stmt, n.Init == nil, tNode},
			{n.Cond, Expr, n.Cond == nil, tNode},
			{n.Body, BlockStmt, n.Body == nil, tNode},
			{n.Else, Stmt, n.Else == nil, tNode},
		}

	case *ast.CaseClause:
		return []TraversableNode{
			{n.List, ExprSlice, n.List == nil, tNode},
			{n.Body, StmtSlice, n.Body == nil, tNode},
		}

	case *ast.SwitchStmt:
		return []TraversableNode{
			{n.Init, Stmt, n.Init == nil, tNode},
			{n.Tag, Expr, n.Tag == nil, tNode},
			{n.Body, BlockStmt, n.Body == nil, tNode},
		}

	case *ast.TypeSwitchStmt:
		return []TraversableNode{
			{n.Init, Stmt, n.Init == nil, tNode},
			{n.Assign, Stmt, n.Assign == nil, tNode},
			{n.Body, BlockStmt, n.Body == nil, tNode},
		}

	case *ast.CommClause:
		return []TraversableNode{
			{n.Comm, Stmt, n.Comm == nil, tNode},
			{n.Body, StmtSlice, n.Body == nil, tNode},
		}

	case *ast.SelectStmt:
		return []TraversableNode{
			{n.Body, BlockStmt, n.Body == nil, tNode},
		}

	case *ast.ForStmt:
		return []TraversableNode{
			{n.Init, Stmt, n.Init == nil, tNode},
			{n.Cond, Expr, n.Cond == nil, tNode},
			{n.Post, Stmt, n.Post == nil, tNode},
			{n.Body, BlockStmt, n.Body == nil, tNode},
		}

	case *ast.RangeStmt:
		return []TraversableNode{
			{n.Key, Expr, n.Key == nil, tNode},
			{n.Value, Expr, n.Value == nil, tNode},
			{n.X, Expr, n.X == nil, tNode},
			{n.Body, BlockStmt, n.Body == nil, tNode},
		}

	case *ast.ImportSpec:
		return []TraversableNode{
			{n.Name, Ident, n.Name == nil, tNode},
			{n.Path, BasicLit, n.Path == nil, tNode},
		}

	case *ast.ValueSpec:
		return []TraversableNode{
			{n.Names, IdentSlice, n.Names == nil, tNode},
			{n.Type, Expr, n.Type == nil, tNode},
			{n.Values, ExprSlice, n.Values == nil, tNode},
		}

	case *ast.TypeSpec:
		return []TraversableNode{
			{n.Name, Ident, n.Name == nil, tNode},
			{n.TypeParams, FieldList, n.TypeParams == nil, tNode},
			{n.Type, Expr, n.Type == nil, tNode},
		}

	case *ast.BadDecl:
		return []TraversableNode{}

	case *ast.GenDecl:
		return []TraversableNode{
			{n.Specs, SpecSlice, n.Specs == nil, tNode},
		}

	case *ast.FuncDecl:
		return []TraversableNode{
			{n.Recv, FieldList, n.Recv == nil, tNode},
			{n.Name, Ident, n.Name == nil, tNode},
			{n.Type, FuncType, n.Type == nil, tNode},
			{n.Body, BlockStmt, n.Body == nil, tNode},
		}

	case *ast.File:
		return []TraversableNode{
			{n.Name, Ident, n.Name == nil, tNode},
			{n.Decls, DeclSlice, n.Decls == nil, tNode},
			{n.Imports, ImportSpecSlice, n.Imports == nil, tNode},
			{n.Unresolved, IdentSlice, n.Unresolved == nil, tNode},
		}

		// case ast*.*ast.Package:
		// 	return []TraversableNode{
		// 	{n.Files, map[string]*File},
		// }

	}
	return []TraversableNode{}
}
