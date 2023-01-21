package ast_wrapper

import (
	"go/ast"
)

func Fields(n ast.Node) ([]any, []NodeType) {

	switch n := n.(type) {

	// Comments and fields

	case *ast.Comment:
		// nothing to do
	case *ast.CommentGroup:
		return []any{n.List}, []NodeType{CommentSlice}
	case *ast.Field:
		return []any{n.Names, n.Type, n.Tag}, []NodeType{IdentSlice, Expr, BasicLit}
	case *ast.FieldList:
		return []any{n.List}, []NodeType{FieldSlice}

	// Expressions

	case *ast.BadExpr, *ast.Ident, *ast.BasicLit:
		// nothing to do
	case *ast.Ellipsis:
		return []any{n.Elt}, []NodeType{Expr}
	case *ast.FuncLit:
		return []any{n.Type, n.Body}, []NodeType{FuncType, BlockStmt}
	case *ast.CompositeLit:
		return []any{n.Type, n.Elts}, []NodeType{Expr, ExprSlice}
	case *ast.ParenExpr:
		return []any{n.X}, []NodeType{Expr}
	case *ast.SelectorExpr:
		return []any{n.X, n.Sel}, []NodeType{Expr, Ident}
	case *ast.IndexExpr:
		return []any{n.X, n.Index}, []NodeType{Expr, Expr}
	case *ast.IndexListExpr:
		return []any{n.X, n.Indices}, []NodeType{Expr, ExprSlice}
	case *ast.SliceExpr:
		return []any{n.X, n.Low, n.High, n.Max}, []NodeType{Expr, Expr, Expr, Expr}
	case *ast.TypeAssertExpr:
		return []any{n.X, n.Type}, []NodeType{Expr, Expr}
	case *ast.CallExpr:
		return []any{n.Fun, n.Args}, []NodeType{Expr, ExprSlice}
	case *ast.StarExpr:
		return []any{n.X}, []NodeType{Expr}
	case *ast.UnaryExpr:
		return []any{n.X}, []NodeType{Expr}
	case *ast.BinaryExpr:
		return []any{n.X, n.Y}, []NodeType{Expr, Expr}
	case *ast.KeyValueExpr:
		return []any{n.Key, n.Value}, []NodeType{Expr, Expr}

	// Types

	case *ast.ArrayType:
		return []any{n.Len, n.Elt}, []NodeType{Expr, Expr}
	case *ast.StructType:
		return []any{n.Fields}, []NodeType{FieldList}
	case *ast.FuncType:
		return []any{n.TypeParams, n.Params, n.Results}, []NodeType{FieldList, FieldList, FieldList}
	case *ast.InterfaceType:
		return []any{n.Methods}, []NodeType{FieldList}
	case *ast.MapType:
		return []any{n.Key, n.Value}, []NodeType{Expr, Expr}
	case *ast.ChanType:
		return []any{n.Value}, []NodeType{Expr}

	// Statements

	case *ast.BadStmt:
		// nothing to do
	case *ast.DeclStmt:
		return []any{n.Decl}, []NodeType{Decl}
	case *ast.EmptyStmt:
		// nothing to do
	case *ast.LabeledStmt:
		return []any{n.Label, n.Stmt}, []NodeType{Ident, Stmt}
	case *ast.ExprStmt:
		return []any{n.X}, []NodeType{Expr}
	case *ast.SendStmt:
		return []any{n.Chan, n.Value}, []NodeType{Expr, Expr}
	case *ast.IncDecStmt:
		return []any{n.X}, []NodeType{IncDecStmt}
	case *ast.AssignStmt:
		return []any{n.Lhs, n.Rhs}, []NodeType{AssignStmt, ExprSlice}
	case *ast.GoStmt:
		return []any{n.Call}, []NodeType{CallExpr}
	case *ast.DeferStmt:
		return []any{n.Call}, []NodeType{CallExpr}
	case *ast.ReturnStmt:
		return []any{n.Results}, []NodeType{ExprSlice}
	case *ast.BranchStmt:
		return []any{n.Label}, []NodeType{Ident}
	case *ast.BlockStmt:
		return []any{n.List}, []NodeType{Stmt}
	case *ast.IfStmt:
		return []any{n.Init, n.Cond, n.Body, n.Else}, []NodeType{Stmt, Expr, BlockStmt, Stmt}
	case *ast.CaseClause:
		return []any{n.List, n.Body}, []NodeType{ExprSlice, StmtSlice}
	case *ast.SwitchStmt:
		return []any{n.Init, n.Tag, n.Body}, []NodeType{Stmt, Expr, BlockStmt}
	case *ast.TypeSwitchStmt:
		return []any{n.Init, n.Assign, n.Body}, []NodeType{Stmt, Stmt, BlockStmt}
	case *ast.CommClause:
		return []any{n.Comm, n.Body}, []NodeType{Stmt, StmtSlice}
	case *ast.SelectStmt:
		return []any{n.Body}, []NodeType{BlockStmt}
	case *ast.ForStmt:
		return []any{n.Init, n.Cond, n.Post, n.Body}, []NodeType{Stmt, Expr, Stmt, ForStmt}
	case *ast.RangeStmt:
		return []any{n.Key, n.Value, n.X, n.Body}, []NodeType{Expr, Expr, Expr, BlockStmt}

	// Specifications

	case *ast.ImportSpec:
		return []any{n.Name, n.Path}, []NodeType{Ident, BasicLit}
	case *ast.ValueSpec:
		return []any{n.Names, n.Type, n.Values}, []NodeType{IdentSlice, Expr, ExprSlice}
	case *ast.TypeSpec:
		return []any{n.Name, n.TypeParams, n.Type}, []NodeType{Ident, FieldList, Expr}
	case *ast.BadDecl:
		// nothing to do

	// Declarations

	case *ast.GenDecl:
		return []any{n.Specs}, []NodeType{SpecSlice}
	case *ast.FuncDecl:
		return []any{n.Recv, n.Name, n.Type, n.Body}, []NodeType{FieldList, Ident, FuncType, BlockStmt}

	// File & Package

	case *ast.File:
		return []any{n.Name, n.Decls}, []NodeType{Ident, DeclSlice}
	case *ast.Package:
		for _, f := range n.Files {
			return []any{f.Name, f.Decls}, []NodeType{Ident, DeclSlice}
		}
	}

	return nil, nil
}
