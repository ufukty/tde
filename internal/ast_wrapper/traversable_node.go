package ast_wrapper

import (
	"go/ast"
)

// Use to represent nullable, Node-like (ast.Node etc.) or Node-slice ([]ast.Node etc.) components of the AST.
type TraversableNode struct {
	Value        any
	ExpectedType NodeType // This usually will be filled while visiting the parent node by leveraging the struct definitions instead of type checking on value because it can be nil too. Also, Traverse() callbacks can leverage for learning which type of node is needed to construct.
	IsNil        bool
	Parent       *TraversableNode
	Ref          Ref
}

func GetTraversableNodeForASTNode(node ast.Node) TraversableNode {
	return TraversableNode{
		Value:        node,
		ExpectedType: GetNodeTypeForASTNode(node),
		IsNil:        isNodeNil(node),
		Parent:       nil,
		Ref:          NewStructRef(&node),
	}
}

func TraversableNodesFromSliceTypeNode(tNode *TraversableNode) (tSubnodes []*TraversableNode) {
	if tNode.IsNil {
		return
	}

	switch slice := tNode.Value.(type) {
	case []*ast.Comment:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, Comment, item == nil, tNode, NewSliceRef(&slice)})
		}
	case []*ast.CommentGroup:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, CommentGroup, item == nil, tNode, NewSliceRef(&slice)})
		}
	case []*ast.ImportSpec:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, ImportSpec, item == nil, tNode, NewSliceRef(&slice)})
		}
	case []*ast.Ident:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, Ident, item == nil, tNode, NewSliceRef(&slice)})
		}
	case []*ast.Field:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, Field, item == nil, tNode, NewSliceRef(&slice)})
		}
	case []ast.Stmt:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, GetNodeTypeForASTNode(item), item == nil, tNode, NewSliceRef(&slice)})
		}
	case []ast.Decl:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, GetNodeTypeForASTNode(item), item == nil, tNode, NewSliceRef(&slice)})
		}
	case []ast.Spec:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, GetNodeTypeForASTNode(item), item == nil, tNode, NewSliceRef(&slice)})
		}
	case []ast.Expr:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, GetNodeTypeForASTNode(item), item == nil, tNode, NewSliceRef(&slice)})
		}
		// case []ast.Node:
		// 	for _, item := range slice {
		// 		tSubnodes = append(tSubnodes, &TraversableNode{item, GetNodeTypeForASTNode(item), item == nil, tNode, NewSliceRef(&slice)})
		// 	}
	}
	return
}

func TraversableNodesFromInterfaceTypeNode(tNode *TraversableNode) []*TraversableNode {
	if tNode.IsNil {
		return []*TraversableNode{}
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
	return []*TraversableNode{}
}

func TraversableNodesFromConcreteTypeNode(tNode *TraversableNode) []*TraversableNode {
	if tNode.IsNil {
		return []*TraversableNode{}
	}

	switch n := (*tNode).Value.(type) {

	case *ast.Comment:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{}

	case *ast.CommentGroup:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{}

	case *ast.Field:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Names, IdentSlice, n.Names == nil, tNode, NewSliceRef(&n.Names)},
			{n.Type, Expr, n.Type == nil, tNode, NewStructRef(&n.Type)},
			{n.Tag, BasicLit, n.Tag == nil, tNode, NewStructRef(&n.Tag)},
		}

	case *ast.FieldList:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.List, FieldSlice, n.List == nil, tNode, NewSliceRef(&n.List)},
		}

	case *ast.BadExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{}

	case *ast.Ident:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{}

	case *ast.Ellipsis:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Elt, Expr, n.Elt == nil, tNode, NewStructRef(&n.Elt)},
		}

	case *ast.BasicLit:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{}

	case *ast.FuncLit:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Type, FuncType, n.Type == nil, tNode, NewStructRef(&n.Type)},
			{n.Body, BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.CompositeLit:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Type, Expr, n.Type == nil, tNode, NewStructRef(&n.Type)},
			{n.Elts, ExprSlice, n.Elts == nil, tNode, NewSliceRef(&n.Elts)},
		}

	case *ast.ParenExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, Expr, n.X == nil, tNode, NewStructRef(&n.X)},
		}

	case *ast.SelectorExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Sel, Ident, n.Sel == nil, tNode, NewStructRef(&n.Sel)},
		}

	case *ast.IndexExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Index, Expr, n.Index == nil, tNode, NewStructRef(&n.Index)},
		}

	case *ast.IndexListExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Indices, ExprSlice, n.Indices == nil, tNode, NewSliceRef(&n.Indices)},
		}

	case *ast.SliceExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Low, Expr, n.Low == nil, tNode, NewStructRef(&n.Low)},
			{n.High, Expr, n.High == nil, tNode, NewStructRef(&n.High)},
			{n.Max, Expr, n.Max == nil, tNode, NewStructRef(&n.Max)},
		}

	case *ast.TypeAssertExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Type, Expr, n.Type == nil, tNode, NewStructRef(&n.Type)},
		}

	case *ast.CallExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Fun, Expr, n.Fun == nil, tNode, NewStructRef(&n.Fun)},
			{n.Args, ExprSlice, n.Args == nil, tNode, NewSliceRef(&n.Args)},
		}

	case *ast.StarExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, Expr, n.X == nil, tNode, NewStructRef(&n.X)},
		}

	case *ast.UnaryExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, Expr, n.X == nil, tNode, NewStructRef(&n.X)},
		}

	case *ast.BinaryExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Y, Expr, n.Y == nil, tNode, NewStructRef(&n.Y)},
		}

	case *ast.KeyValueExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Key, Expr, n.Key == nil, tNode, NewStructRef(&n.Key)},
			{n.Value, Expr, n.Value == nil, tNode, NewStructRef(&n.Value)},
		}

	case *ast.ArrayType:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Len, Expr, n.Len == nil, tNode, NewStructRef(&n.Len)},
			{n.Elt, Expr, n.Elt == nil, tNode, NewStructRef(&n.Elt)},
		}

	case *ast.StructType:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Fields, FieldList, n.Fields == nil, tNode, NewStructRef(&n.Fields)},
		}

	case *ast.FuncType:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.TypeParams, FieldList, n.TypeParams == nil, tNode, NewStructRef(&n.TypeParams)},
			{n.Params, FieldList, n.Params == nil, tNode, NewStructRef(&n.Params)},
			{n.Results, FieldList, n.Results == nil, tNode, NewStructRef(&n.Results)},
		}

	case *ast.InterfaceType:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Methods, FieldList, n.Methods == nil, tNode, NewStructRef(&n.Methods)},
		}

	case *ast.MapType:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Key, Expr, n.Key == nil, tNode, NewStructRef(&n.Key)},
			{n.Value, Expr, n.Value == nil, tNode, NewStructRef(&n.Value)},
		}

	case *ast.ChanType:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Value, Expr, n.Value == nil, tNode, NewStructRef(&n.Value)},
		}

	case *ast.BadStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{}

	case *ast.DeclStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Decl, Decl, n.Decl == nil, tNode, NewStructRef(&n.Decl)},
		}

	case *ast.EmptyStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{}

	case *ast.LabeledStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Label, Ident, n.Label == nil, tNode, NewStructRef(&n.Label)},
			{n.Stmt, Stmt, n.Stmt == nil, tNode, NewStructRef(&n.Stmt)},
		}

	case *ast.ExprStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, Expr, n.X == nil, tNode, NewStructRef(&n.X)},
		}

	case *ast.SendStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Chan, Expr, n.Chan == nil, tNode, NewStructRef(&n.Chan)},
			{n.Value, Expr, n.Value == nil, tNode, NewStructRef(&n.Value)},
		}

	case *ast.IncDecStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, Expr, n.X == nil, tNode, NewStructRef(&n.X)},
		}

	case *ast.AssignStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Lhs, ExprSlice, n.Lhs == nil, tNode, NewSliceRef(&n.Lhs)},
			{n.Rhs, ExprSlice, n.Rhs == nil, tNode, NewSliceRef(&n.Rhs)},
		}

	case *ast.GoStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Call, CallExpr, n.Call == nil, tNode, NewStructRef(&n.Call)},
		}

	case *ast.DeferStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Call, CallExpr, n.Call == nil, tNode, NewStructRef(&n.Call)},
		}

	case *ast.ReturnStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Results, ExprSlice, n.Results == nil, tNode, NewSliceRef(&n.Results)},
		}

	case *ast.BranchStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Label, Ident, n.Label == nil, tNode, NewStructRef(&n.Label)},
		}

	case *ast.BlockStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.List, StmtSlice, n.List == nil, tNode, NewSliceRef(&n.List)},
		}

	case *ast.IfStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Init, Stmt, n.Init == nil, tNode, NewStructRef(&n.Init)},
			{n.Cond, Expr, n.Cond == nil, tNode, NewStructRef(&n.Cond)},
			{n.Body, BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
			{n.Else, Stmt, n.Else == nil, tNode, NewStructRef(&n.Else)},
		}

	case *ast.CaseClause:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.List, ExprSlice, n.List == nil, tNode, NewSliceRef(&n.List)},
			{n.Body, StmtSlice, n.Body == nil, tNode, NewSliceRef(&n.Body)},
		}

	case *ast.SwitchStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Init, Stmt, n.Init == nil, tNode, NewStructRef(&n.Init)},
			{n.Tag, Expr, n.Tag == nil, tNode, NewStructRef(&n.Tag)},
			{n.Body, BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.TypeSwitchStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Init, Stmt, n.Init == nil, tNode, NewStructRef(&n.Init)},
			{n.Assign, Stmt, n.Assign == nil, tNode, NewStructRef(&n.Assign)},
			{n.Body, BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.CommClause:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Comm, Stmt, n.Comm == nil, tNode, NewStructRef(&n.Comm)},
			{n.Body, StmtSlice, n.Body == nil, tNode, NewSliceRef(&n.Body)},
		}

	case *ast.SelectStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Body, BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.ForStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Init, Stmt, n.Init == nil, tNode, NewStructRef(&n.Init)},
			{n.Cond, Expr, n.Cond == nil, tNode, NewStructRef(&n.Cond)},
			{n.Post, Stmt, n.Post == nil, tNode, NewStructRef(&n.Post)},
			{n.Body, BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.RangeStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Key, Expr, n.Key == nil, tNode, NewStructRef(&n.Key)},
			{n.Value, Expr, n.Value == nil, tNode, NewStructRef(&n.Value)},
			{n.X, Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Body, BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.ImportSpec:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Name, Ident, n.Name == nil, tNode, NewStructRef(&n.Name)},
			{n.Path, BasicLit, n.Path == nil, tNode, NewStructRef(&n.Path)},
		}

	case *ast.ValueSpec:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Names, IdentSlice, n.Names == nil, tNode, NewSliceRef(&n.Names)},
			{n.Type, Expr, n.Type == nil, tNode, NewStructRef(&n.Type)},
			{n.Values, ExprSlice, n.Values == nil, tNode, NewSliceRef(&n.Values)},
		}

	case *ast.TypeSpec:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Name, Ident, n.Name == nil, tNode, NewStructRef(&n.Name)},
			{n.TypeParams, FieldList, n.TypeParams == nil, tNode, NewStructRef(&n.TypeParams)},
			{n.Type, Expr, n.Type == nil, tNode, NewStructRef(&n.Type)},
		}

	case *ast.BadDecl:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{}

	case *ast.GenDecl:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Specs, SpecSlice, n.Specs == nil, tNode, NewSliceRef(&n.Specs)},
		}

	case *ast.FuncDecl:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Recv, FieldList, n.Recv == nil, tNode, NewStructRef(&n.Recv)},
			{n.Name, Ident, n.Name == nil, tNode, NewStructRef(&n.Name)},
			{n.Type, FuncType, n.Type == nil, tNode, NewStructRef(&n.Type)},
			{n.Body, BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.File:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Name, Ident, n.Name == nil, tNode, NewStructRef(&n.Name)},
			{n.Decls, DeclSlice, n.Decls == nil, tNode, NewSliceRef(&n.Decls)},
			{n.Imports, ImportSpecSlice, n.Imports == nil, tNode, NewSliceRef(&n.Imports)},
			{n.Unresolved, IdentSlice, n.Unresolved == nil, tNode, NewSliceRef(&n.Unresolved)},
		}

		// case ast*.*ast.Package:
		// if n == nil {
		// 	return []*TraversableNode{}
		// }
		// 	return []*TraversableNode{
		// 	{n.Files, map[string]*File},
		// }

	}

	return []*TraversableNode{}
}
