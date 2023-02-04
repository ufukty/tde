package traverse

import (
	ast_types "tde/internal/astw/types"
	ast_utl "tde/internal/astw/utilities"

	"go/ast"
)

// Use to represent nullable, Node-like (ast.Node etc.) or Node-slice ([]ast.Node etc.) components of the AST.
type TraversableNode struct {
	Value           any
	ExpectedType    ast_types.NodeType // This usually will be filled while visiting the parent node by leveraging the struct definitions instead of type checking on value because it can be nil too. Also, Traverse() callbacks can leverage for learning which type of node is needed to construct.
	PointsToNilSpot bool
	Parent          *TraversableNode
	Ref             Ref
}

func GetTraversableNodeForASTNode(node ast.Node) *TraversableNode {
	return &TraversableNode{
		Value:           node,
		ExpectedType:    ast_types.GetNodeTypeForASTNode(node),
		PointsToNilSpot: ast_utl.IsNodeNil(node),
		Parent:          nil,
		Ref:             NewStructRef(&node),
	}
}

func TraversableNodesFromSliceTypeNode(tNode *TraversableNode) (tSubnodes []*TraversableNode) {
	if tNode.PointsToNilSpot {
		return
	}

	switch slice := tNode.Value.(type) {
	case []*ast.Comment:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, ast_types.Comment, item == nil, tNode, NewStructRef(&item)})
		}
	case []*ast.CommentGroup:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, ast_types.CommentGroup, item == nil, tNode, NewStructRef(&item)})
		}
	case []*ast.ImportSpec:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, ast_types.ImportSpec, item == nil, tNode, NewStructRef(&item)})
		}
	case []*ast.Ident:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, ast_types.Ident, item == nil, tNode, NewStructRef(&item)})
		}
	case []*ast.Field:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, ast_types.Field, item == nil, tNode, NewStructRef(&item)})
		}
	case []ast.Stmt:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, ast_types.GetNodeTypeForASTNode(item), item == nil, tNode, NewStructRef(&item)})
		}
	case []ast.Decl:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, ast_types.GetNodeTypeForASTNode(item), item == nil, tNode, NewStructRef(&item)})
		}
	case []ast.Spec:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, ast_types.GetNodeTypeForASTNode(item), item == nil, tNode, NewStructRef(&item)})
		}
	case []ast.Expr:
		for _, item := range slice {
			tSubnodes = append(tSubnodes, &TraversableNode{item, ast_types.GetNodeTypeForASTNode(item), item == nil, tNode, NewStructRef(&item)})
		}
		// case []ast.Node:
		// 	for _, item := range slice {
		// 		tSubnodes = append(tSubnodes, &TraversableNode{item, types.GetNodeTypeForASTNode(item), item == nil, tNode, NewStructRef(&item)})
		// 	}
	}
	return
}

func TraversableNodesFromInterfaceTypeNode(tNode *TraversableNode) []*TraversableNode {
	if tNode.PointsToNilSpot {
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
	if tNode.PointsToNilSpot {
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
			{n.Names, ast_types.IdentSlice, n.Names == nil, tNode, NewSliceRef(&n.Names)},
			{n.Type, ast_types.Expr, n.Type == nil, tNode, NewStructRef(&n.Type)},
			{n.Tag, ast_types.BasicLit, n.Tag == nil, tNode, NewStructRef(&n.Tag)},
		}

	case *ast.FieldList:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.List, ast_types.FieldSlice, n.List == nil, tNode, NewSliceRef(&n.List)},
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
			{n.Elt, ast_types.Expr, n.Elt == nil, tNode, NewStructRef(&n.Elt)},
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
			{n.Type, ast_types.FuncType, n.Type == nil, tNode, NewStructRef(&n.Type)},
			{n.Body, ast_types.BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.CompositeLit:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Type, ast_types.Expr, n.Type == nil, tNode, NewStructRef(&n.Type)},
			{n.Elts, ast_types.ExprSlice, n.Elts == nil, tNode, NewSliceRef(&n.Elts)},
		}

	case *ast.ParenExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, ast_types.Expr, n.X == nil, tNode, NewStructRef(&n.X)},
		}

	case *ast.SelectorExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, ast_types.Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Sel, ast_types.Ident, n.Sel == nil, tNode, NewStructRef(&n.Sel)},
		}

	case *ast.IndexExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, ast_types.Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Index, ast_types.Expr, n.Index == nil, tNode, NewStructRef(&n.Index)},
		}

	case *ast.IndexListExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, ast_types.Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Indices, ast_types.ExprSlice, n.Indices == nil, tNode, NewSliceRef(&n.Indices)},
		}

	case *ast.SliceExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, ast_types.Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Low, ast_types.Expr, n.Low == nil, tNode, NewStructRef(&n.Low)},
			{n.High, ast_types.Expr, n.High == nil, tNode, NewStructRef(&n.High)},
			{n.Max, ast_types.Expr, n.Max == nil, tNode, NewStructRef(&n.Max)},
		}

	case *ast.TypeAssertExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, ast_types.Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Type, ast_types.Expr, n.Type == nil, tNode, NewStructRef(&n.Type)},
		}

	case *ast.CallExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Fun, ast_types.Expr, n.Fun == nil, tNode, NewStructRef(&n.Fun)},
			{n.Args, ast_types.ExprSlice, n.Args == nil, tNode, NewSliceRef(&n.Args)},
		}

	case *ast.StarExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, ast_types.Expr, n.X == nil, tNode, NewStructRef(&n.X)},
		}

	case *ast.UnaryExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, ast_types.Expr, n.X == nil, tNode, NewStructRef(&n.X)},
		}

	case *ast.BinaryExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, ast_types.Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Y, ast_types.Expr, n.Y == nil, tNode, NewStructRef(&n.Y)},
		}

	case *ast.KeyValueExpr:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Key, ast_types.Expr, n.Key == nil, tNode, NewStructRef(&n.Key)},
			{n.Value, ast_types.Expr, n.Value == nil, tNode, NewStructRef(&n.Value)},
		}

	case *ast.ArrayType:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Len, ast_types.Expr, n.Len == nil, tNode, NewStructRef(&n.Len)},
			{n.Elt, ast_types.Expr, n.Elt == nil, tNode, NewStructRef(&n.Elt)},
		}

	case *ast.StructType:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Fields, ast_types.FieldList, n.Fields == nil, tNode, NewStructRef(&n.Fields)},
		}

	case *ast.FuncType:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.TypeParams, ast_types.FieldList, n.TypeParams == nil, tNode, NewStructRef(&n.TypeParams)},
			{n.Params, ast_types.FieldList, n.Params == nil, tNode, NewStructRef(&n.Params)},
			{n.Results, ast_types.FieldList, n.Results == nil, tNode, NewStructRef(&n.Results)},
		}

	case *ast.InterfaceType:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Methods, ast_types.FieldList, n.Methods == nil, tNode, NewStructRef(&n.Methods)},
		}

	case *ast.MapType:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Key, ast_types.Expr, n.Key == nil, tNode, NewStructRef(&n.Key)},
			{n.Value, ast_types.Expr, n.Value == nil, tNode, NewStructRef(&n.Value)},
		}

	case *ast.ChanType:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Value, ast_types.Expr, n.Value == nil, tNode, NewStructRef(&n.Value)},
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
			{n.Decl, ast_types.Decl, n.Decl == nil, tNode, NewStructRef(&n.Decl)},
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
			{n.Label, ast_types.Ident, n.Label == nil, tNode, NewStructRef(&n.Label)},
			{n.Stmt, ast_types.Stmt, n.Stmt == nil, tNode, NewStructRef(&n.Stmt)},
		}

	case *ast.ExprStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, ast_types.Expr, n.X == nil, tNode, NewStructRef(&n.X)},
		}

	case *ast.SendStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Chan, ast_types.Expr, n.Chan == nil, tNode, NewStructRef(&n.Chan)},
			{n.Value, ast_types.Expr, n.Value == nil, tNode, NewStructRef(&n.Value)},
		}

	case *ast.IncDecStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.X, ast_types.Expr, n.X == nil, tNode, NewStructRef(&n.X)},
		}

	case *ast.AssignStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Lhs, ast_types.ExprSlice, n.Lhs == nil, tNode, NewSliceRef(&n.Lhs)},
			{n.Rhs, ast_types.ExprSlice, n.Rhs == nil, tNode, NewSliceRef(&n.Rhs)},
		}

	case *ast.GoStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Call, ast_types.CallExpr, n.Call == nil, tNode, NewStructRef(&n.Call)},
		}

	case *ast.DeferStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Call, ast_types.CallExpr, n.Call == nil, tNode, NewStructRef(&n.Call)},
		}

	case *ast.ReturnStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Results, ast_types.ExprSlice, n.Results == nil, tNode, NewSliceRef(&n.Results)},
		}

	case *ast.BranchStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Label, ast_types.Ident, n.Label == nil, tNode, NewStructRef(&n.Label)},
		}

	case *ast.BlockStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.List, ast_types.StmtSlice, n.List == nil, tNode, NewSliceRef(&n.List)},
		}

	case *ast.IfStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Init, ast_types.Stmt, n.Init == nil, tNode, NewStructRef(&n.Init)},
			{n.Cond, ast_types.Expr, n.Cond == nil, tNode, NewStructRef(&n.Cond)},
			{n.Body, ast_types.BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
			{n.Else, ast_types.Stmt, n.Else == nil, tNode, NewStructRef(&n.Else)},
		}

	case *ast.CaseClause:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.List, ast_types.ExprSlice, n.List == nil, tNode, NewSliceRef(&n.List)},
			{n.Body, ast_types.StmtSlice, n.Body == nil, tNode, NewSliceRef(&n.Body)},
		}

	case *ast.SwitchStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Init, ast_types.Stmt, n.Init == nil, tNode, NewStructRef(&n.Init)},
			{n.Tag, ast_types.Expr, n.Tag == nil, tNode, NewStructRef(&n.Tag)},
			{n.Body, ast_types.BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.TypeSwitchStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Init, ast_types.Stmt, n.Init == nil, tNode, NewStructRef(&n.Init)},
			{n.Assign, ast_types.Stmt, n.Assign == nil, tNode, NewStructRef(&n.Assign)},
			{n.Body, ast_types.BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.CommClause:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Comm, ast_types.Stmt, n.Comm == nil, tNode, NewStructRef(&n.Comm)},
			{n.Body, ast_types.StmtSlice, n.Body == nil, tNode, NewSliceRef(&n.Body)},
		}

	case *ast.SelectStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Body, ast_types.BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.ForStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Init, ast_types.Stmt, n.Init == nil, tNode, NewStructRef(&n.Init)},
			{n.Cond, ast_types.Expr, n.Cond == nil, tNode, NewStructRef(&n.Cond)},
			{n.Post, ast_types.Stmt, n.Post == nil, tNode, NewStructRef(&n.Post)},
			{n.Body, ast_types.BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.RangeStmt:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Key, ast_types.Expr, n.Key == nil, tNode, NewStructRef(&n.Key)},
			{n.Value, ast_types.Expr, n.Value == nil, tNode, NewStructRef(&n.Value)},
			{n.X, ast_types.Expr, n.X == nil, tNode, NewStructRef(&n.X)},
			{n.Body, ast_types.BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.ImportSpec:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Name, ast_types.Ident, n.Name == nil, tNode, NewStructRef(&n.Name)},
			{n.Path, ast_types.BasicLit, n.Path == nil, tNode, NewStructRef(&n.Path)},
		}

	case *ast.ValueSpec:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Names, ast_types.IdentSlice, n.Names == nil, tNode, NewSliceRef(&n.Names)},
			{n.Type, ast_types.Expr, n.Type == nil, tNode, NewStructRef(&n.Type)},
			{n.Values, ast_types.ExprSlice, n.Values == nil, tNode, NewSliceRef(&n.Values)},
		}

	case *ast.TypeSpec:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Name, ast_types.Ident, n.Name == nil, tNode, NewStructRef(&n.Name)},
			{n.TypeParams, ast_types.FieldList, n.TypeParams == nil, tNode, NewStructRef(&n.TypeParams)},
			{n.Type, ast_types.Expr, n.Type == nil, tNode, NewStructRef(&n.Type)},
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
			{n.Specs, ast_types.SpecSlice, n.Specs == nil, tNode, NewSliceRef(&n.Specs)},
		}

	case *ast.FuncDecl:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Recv, ast_types.FieldList, n.Recv == nil, tNode, NewStructRef(&n.Recv)},
			{n.Name, ast_types.Ident, n.Name == nil, tNode, NewStructRef(&n.Name)},
			{n.Type, ast_types.FuncType, n.Type == nil, tNode, NewStructRef(&n.Type)},
			{n.Body, ast_types.BlockStmt, n.Body == nil, tNode, NewStructRef(&n.Body)},
		}

	case *ast.File:
		if n == nil {
			return []*TraversableNode{}
		}
		return []*TraversableNode{
			{n.Name, ast_types.Ident, n.Name == nil, tNode, NewStructRef(&n.Name)},
			{n.Decls, ast_types.DeclSlice, n.Decls == nil, tNode, NewSliceRef(&n.Decls)},
			{n.Imports, ast_types.ImportSpecSlice, n.Imports == nil, tNode, NewSliceRef(&n.Imports)},
			{n.Unresolved, ast_types.IdentSlice, n.Unresolved == nil, tNode, NewSliceRef(&n.Unresolved)},
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
