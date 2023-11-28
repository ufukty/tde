package traverse

import (
	"go/ast"
	"tde/internal/astw/types"
)

func subnode(parent *Node, rf ref, isNil bool, expected types.NodeType, c constraint) *Node {
	return &Node{
		ref:         rf,
		IsNil:       isNil,
		Expected:    expected,
		Constraints: parent.Constraints.Clone(c),
		Parent:      parent,
	}
}

func subnodesForConcrete(n *Node, c constraints) []*Node {
	if n.IsNil {
		return []*Node{}
	}

	switch v := n.ref.Get().(type) {

	case *ast.Comment:
		if v == nil {
			return []*Node{}
		}
		return []*Node{}

	case *ast.CommentGroup:
		if v == nil {
			return []*Node{}
		}
		return []*Node{}

	case *ast.Field:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Names), v.Names == nil, types.IdentSlice, 0),
			subnode(n, newFieldRef(&v.Type), v.Type == nil, types.TypeExpr, 0),
			subnode(n, newFieldRef(&v.Tag), v.Tag == nil, types.BasicLit, 0),
		}

	case *ast.FieldList:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.List), v.List == nil, types.FieldSlice, 0),
		}

	case *ast.BadExpr:
		if v == nil {
			return []*Node{}
		}
		return []*Node{}

	case *ast.Ident:
		if v == nil {
			return []*Node{}
		}
		return []*Node{}

	case *ast.Ellipsis:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Elt), v.Elt == nil, types.Expr, 0),
		}

	case *ast.BasicLit:
		if v == nil {
			return []*Node{}
		}
		return []*Node{}

	case *ast.FuncLit:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Type), v.Type == nil, types.FuncType, 0),
			subnode(n, newFieldRef(&v.Body), v.Body == nil, types.BlockStmt, 0),
		}

	case *ast.CompositeLit:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Type), v.Type == nil, types.TypeExpr, 0),
			subnode(n, newFieldRef(&v.Elts), v.Elts == nil, types.ExprSlice, 0),
		}

	case *ast.ParenExpr:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.X), v.X == nil, types.Expr, 0),
		}

	case *ast.SelectorExpr:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.X), v.X == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Sel), v.Sel == nil, types.Ident, 0),
		}

	case *ast.IndexExpr:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.X), v.X == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Index), v.Index == nil, types.Expr, 0),
		}

	case *ast.IndexListExpr:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.X), v.X == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Indices), v.Indices == nil, types.ExprSlice, 0),
		}

	case *ast.SliceExpr:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.X), v.X == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Low), v.Low == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.High), v.High == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Max), v.Max == nil, types.Expr, 0),
		}

	case *ast.TypeAssertExpr:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.X), v.X == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Type), v.Type == nil, types.TypeExpr, 0),
		}

	case *ast.CallExpr:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Fun), v.Fun == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Args), v.Args == nil, types.ExprSlice, 0),
		}

	case *ast.StarExpr:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.X), v.X == nil, types.Expr, 0),
		}

	case *ast.UnaryExpr:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.X), v.X == nil, types.Expr, 0),
		}

	case *ast.BinaryExpr:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.X), v.X == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Y), v.Y == nil, types.Expr, 0),
		}

	case *ast.KeyValueExpr:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Key), v.Key == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Value), v.Value == nil, types.Expr, 0),
		}

	case *ast.ArrayType:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Len), v.Len == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Elt), v.Elt == nil, types.Expr, 0),
		}

	case *ast.StructType:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Fields), v.Fields == nil, types.FieldList, 0),
		}

	case *ast.FuncType:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.TypeParams), v.TypeParams == nil, types.FieldList, 0),
			subnode(n, newFieldRef(&v.Params), v.Params == nil, types.FieldList, 0),
			subnode(n, newFieldRef(&v.Results), v.Results == nil, types.FieldList, 0),
		}

	case *ast.InterfaceType:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Methods), v.Methods == nil, types.FieldList, 0),
		}

	case *ast.MapType:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Key), v.Key == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Value), v.Value == nil, types.Expr, 0),
		}

	case *ast.ChanType:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Value), v.Value == nil, types.Expr, 0),
		}

	case *ast.BadStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{}

	case *ast.DeclStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Decl), v.Decl == nil, types.Decl, 0),
		}

	case *ast.EmptyStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{}

	case *ast.LabeledStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Label), v.Label == nil, types.Ident, 0),
			subnode(n, newFieldRef(&v.Stmt), v.Stmt == nil, types.Stmt, 0),
		}

	case *ast.ExprStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.X), v.X == nil, types.Expr, 0),
		}

	case *ast.SendStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Chan), v.Chan == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Value), v.Value == nil, types.Expr, 0),
		}

	case *ast.IncDecStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.X), v.X == nil, types.Expr, 0),
		}

	case *ast.AssignStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Lhs), v.Lhs == nil, types.ExprSlice, 0),
			subnode(n, newFieldRef(&v.Rhs), v.Rhs == nil, types.ExprSlice, 0),
		}

	case *ast.GoStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Call), v.Call == nil, types.CallExpr, 0),
		}

	case *ast.DeferStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Call), v.Call == nil, types.CallExpr, 0),
		}

	case *ast.ReturnStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Results), v.Results == nil, types.ExprSlice, 0),
		}

	case *ast.BranchStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Label), v.Label == nil, types.Ident, 0),
		}

	case *ast.BlockStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.List), v.List == nil, types.StmtSlice, 0),
		}

	case *ast.IfStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Init), v.Init == nil, types.Stmt, controlFlowInit),
			subnode(n, newFieldRef(&v.Cond), v.Cond == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Body), v.Body == nil, types.BlockStmt, 0),
			subnode(n, newFieldRef(&v.Else), v.Else == nil, types.Stmt, 0),
		}

	case *ast.CaseClause:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.List), v.List == nil, types.ExprSlice, 0),
			subnode(n, newFieldRef(&v.Body), v.Body == nil, types.StmtSlice, 0),
		}

	case *ast.SwitchStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Init), v.Init == nil, types.Stmt, 0),
			subnode(n, newFieldRef(&v.Tag), v.Tag == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Body), v.Body == nil, types.BlockStmt, 0),
		}

	case *ast.TypeSwitchStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Init), v.Init == nil, types.Stmt, 0),
			subnode(n, newFieldRef(&v.Assign), v.Assign == nil, types.Stmt, 0),
			subnode(n, newFieldRef(&v.Body), v.Body == nil, types.BlockStmt, 0),
		}

	case *ast.CommClause:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Comm), v.Comm == nil, types.Stmt, 0),
			subnode(n, newFieldRef(&v.Body), v.Body == nil, types.StmtSlice, 0),
		}

	case *ast.SelectStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Body), v.Body == nil, types.BlockStmt, 0),
		}

	case *ast.ForStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Init), v.Init == nil, types.Stmt, 0),
			subnode(n, newFieldRef(&v.Cond), v.Cond == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Post), v.Post == nil, types.Stmt, 0),
			subnode(n, newFieldRef(&v.Body), v.Body == nil, types.BlockStmt, 0),
		}

	case *ast.RangeStmt:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Key), v.Key == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Value), v.Value == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.X), v.X == nil, types.Expr, 0),
			subnode(n, newFieldRef(&v.Body), v.Body == nil, types.BlockStmt, 0),
		}

	case *ast.ImportSpec:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Name), v.Name == nil, types.Ident, 0),
			subnode(n, newFieldRef(&v.Path), v.Path == nil, types.BasicLit, 0),
		}

	case *ast.ValueSpec:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Names), v.Names == nil, types.IdentSlice, 0),
			subnode(n, newFieldRef(&v.Type), v.Type == nil, types.TypeExpr, 0),
			subnode(n, newFieldRef(&v.Values), v.Values == nil, types.ExprSlice, 0),
		}

	case *ast.TypeSpec:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Name), v.Name == nil, types.Ident, 0),
			subnode(n, newFieldRef(&v.TypeParams), v.TypeParams == nil, types.FieldList, 0),
			subnode(n, newFieldRef(&v.Type), v.Type == nil, types.TypeExpr, 0),
		}

	case *ast.BadDecl:
		if v == nil {
			return []*Node{}
		}
		return []*Node{}

	case *ast.GenDecl:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Specs), v.Specs == nil, types.SpecSlice, 0),
		}

	case *ast.FuncDecl:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Recv), v.Recv == nil, types.FieldList, 0),
			subnode(n, newFieldRef(&v.Name), v.Name == nil, types.Ident, 0),
			subnode(n, newFieldRef(&v.Type), v.Type == nil, types.TypeExpr, 0),
			subnode(n, newFieldRef(&v.Body), v.Body == nil, types.BlockStmt, 0),
		}

	case *ast.File:
		if v == nil {
			return []*Node{}
		}
		return []*Node{
			subnode(n, newFieldRef(&v.Name), v.Name == nil, types.Ident, 0),
			subnode(n, newFieldRef(&v.Decls), v.Decls == nil, types.DeclSlice, 0),
			subnode(n, newFieldRef(&v.Imports), v.Imports == nil, types.ImportSpecSlice, 0),
			subnode(n, newFieldRef(&v.Unresolved), v.Unresolved == nil, types.IdentSlice, 0),
		}

	case *ast.Package:
		if v == nil {
			return []*Node{}
		}
		ns := []*Node{}
		for _, f := range (*v).Files {
			subnode(n, newFieldRef(f), f == nil, types.File, 0)
		}
		return ns

	}

	return []*Node{}
}
