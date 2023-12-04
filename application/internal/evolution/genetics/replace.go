package genetics

import (
	"fmt"
	"go/ast"
	"tde/internal/astw/astwutl"
)

// Replaces the (c)urrent node with (n)ext node by finding the "field" on its parent, not directly.
// Returns false if the replacement has failed (either because type-incomplient or c doesn't exist).
// r: root
// NOTE: Will not work for assigning to nil fields
func replaceOnParent(root ast.Node, c, n any) bool {
	found := false
	replaced := false

	ast.Inspect(root, func(m ast.Node) bool {
		if root == m || m == nil || astwutl.IsNodeNil(m) {
			return !found
		}

		switch m := m.(type) {

		// case *ast.Package,
		// 	*ast.Comment,
		// 	*ast.BadExpr,
		// 	*ast.Ident,
		// 	*ast.BasicLit,
		// 	*ast.BadStmt,
		// 	*ast.EmptyStmt,
		// 	*ast.BadDecl:

		case *ast.CommentGroup:
			if c, ok := c.([]*ast.Comment); ok && &c == &m.List { // list comparison by pointers
				found = true
				if c, ok := n.([]*ast.Comment); ok {
					m.List = c
					replaced = true
				}
			}

		case *ast.Field:
			if c, ok := c.(*ast.CommentGroup); ok && c == m.Doc {
				found = true
				if n, ok := n.(*ast.CommentGroup); ok {
					m.Doc = n
					replaced = true
				}
			}
			if c, ok := c.([]*ast.Ident); ok && &c == &m.Names { // list comparison by pointers
				found = true
				if n, ok := n.([]*ast.Ident); ok {
					m.Names = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Type {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Type = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.BasicLit); ok && c == m.Tag {
				found = true
				if n, ok := n.(*ast.BasicLit); ok {
					m.Tag = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.CommentGroup); ok && c == m.Comment {
				found = true
				if n, ok := n.(*ast.CommentGroup); ok {
					m.Comment = n
					replaced = true
				}
			}

		case *ast.FieldList:
			if c, ok := c.([]*ast.Field); ok && &c == &m.List { // list comparison by pointers
				found = true
				if n, ok := n.([]*ast.Field); ok {
					m.List = n
					replaced = true
				}
			}

		// Expressions

		case *ast.Ellipsis:
			if c, ok := c.(ast.Expr); ok && c == m.Elt {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Elt = n
					replaced = true
				}
			}

		case *ast.FuncLit:
			if c, ok := c.(*ast.FuncType); ok && c == m.Type {
				found = true
				if n, ok := n.(*ast.FuncType); ok {
					m.Type = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.BlockStmt); ok && c == m.Body {
				found = true
				if n, ok := n.(*ast.BlockStmt); ok {
					m.Body = n
					replaced = true
				}
			}

		case *ast.CompositeLit:
			if c, ok := c.(ast.Expr); ok && c == m.Type {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Type = n
					replaced = true
				}
			}
			if c, ok := c.([]ast.Expr); ok && &c == &m.Elts { // list comparison by pointers
				found = true
				if n, ok := n.([]ast.Expr); ok {
					m.Elts = n
					replaced = true
				}
			}

		case *ast.ParenExpr:
			if c, ok := c.(ast.Expr); ok && c == m.X {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.X = n
					replaced = true
				}
			}

		case *ast.SelectorExpr:
			if c, ok := c.(ast.Expr); ok && c == m.X {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.X = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.Ident); ok && c == m.Sel {
				found = true
				if n, ok := n.(*ast.Ident); ok {
					m.Sel = n
					replaced = true
				}
			}

		case *ast.IndexExpr:
			if c, ok := c.(ast.Expr); ok && c == m.X {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.X = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Index {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Index = n
					replaced = true
				}
			}

		case *ast.IndexListExpr:
			if c, ok := c.(ast.Expr); ok && c == m.X {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.X = n
					replaced = true
				}
			}
			if c, ok := c.([]ast.Expr); ok && &c == &m.Indices { // list comparison by pointers
				found = true
				if n, ok := n.([]ast.Expr); ok {
					m.Indices = n
					replaced = true
				}
			}

		case *ast.SliceExpr:
			if c, ok := c.(ast.Expr); ok && c == m.X {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.X = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Low {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Low = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.High {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.High = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Max {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Max = n
					replaced = true
				}
			}

		case *ast.TypeAssertExpr:
			if c, ok := c.(ast.Expr); ok && c == m.X {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.X = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Type {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Type = n
					replaced = true
				}
			}

		case *ast.CallExpr:
			if c, ok := c.(ast.Expr); ok && c == m.Fun {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Fun = n
					replaced = true
				}
			}
			if c, ok := c.([]ast.Expr); ok && &c == &m.Args { // list comparison by pointers
				found = true
				if n, ok := n.([]ast.Expr); ok {
					m.Args = n
					replaced = true
				}
			}

		case *ast.StarExpr:
			if c, ok := c.(ast.Expr); ok && c == m.X {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.X = n
					replaced = true
				}
			}

		case *ast.UnaryExpr:
			if c, ok := c.(ast.Expr); ok && c == m.X {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.X = n
					replaced = true
				}
			}

		case *ast.BinaryExpr:
			if c, ok := c.(ast.Expr); ok && c == m.X {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.X = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Y {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Y = n
					replaced = true
				}
			}

		case *ast.KeyValueExpr:
			if c, ok := c.(ast.Expr); ok && c == m.Key {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Key = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Value {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Value = n
					replaced = true
				}
			}

		// Types
		case *ast.ArrayType:
			if c, ok := c.(ast.Expr); ok && c == m.Len {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Len = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Elt {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Elt = n
					replaced = true
				}
			}

		case *ast.StructType:
			if c, ok := c.(*ast.FieldList); ok && c == m.Fields {
				found = true
				if n, ok := n.(*ast.FieldList); ok {
					m.Fields = n
					replaced = true
				}
			}

		case *ast.FuncType:
			if c, ok := c.(*ast.FieldList); ok && c == m.TypeParams { // fields
				found = true
				if n, ok := n.(*ast.FieldList); ok {
					m.TypeParams = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.FieldList); ok && c == m.Params {
				found = true
				if n, ok := n.(*ast.FieldList); ok {
					m.Params = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.FieldList); ok && c == m.Results {
				found = true
				if n, ok := n.(*ast.FieldList); ok {
					m.Results = n
					replaced = true
				}
			}

		case *ast.InterfaceType:
			if c, ok := c.(*ast.FieldList); ok && c == m.Methods {
				found = true
				if n, ok := n.(*ast.FieldList); ok {
					m.Methods = n
					replaced = true
				}
			}

		case *ast.MapType:
			if c, ok := c.(ast.Expr); ok && c == m.Key {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Key = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Value {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Value = n
					replaced = true
				}
			}

		case *ast.ChanType:
			if c, ok := c.(ast.Expr); ok && c == m.Value {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Value = n
					replaced = true
				}
			}

		// Statements

		case *ast.DeclStmt:
			if c, ok := c.(ast.Decl); ok && c == m.Decl {
				found = true
				if n, ok := n.(ast.Decl); ok {
					m.Decl = n
					replaced = true
				}
			}

		case *ast.LabeledStmt:
			if c, ok := c.(*ast.Ident); ok && c == m.Label {
				found = true
				if n, ok := n.(*ast.Ident); ok {
					m.Label = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Stmt); ok && c == m.Stmt {
				found = true
				if n, ok := n.(ast.Stmt); ok {
					m.Stmt = n
					replaced = true
				}
			}

		case *ast.ExprStmt:
			if c, ok := c.(ast.Expr); ok && c == m.X {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.X = n
					replaced = true
				}
			}

		case *ast.SendStmt:
			if c, ok := c.(ast.Expr); ok && c == m.Chan {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Chan = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Value {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Value = n
					replaced = true
				}
			}

		case *ast.IncDecStmt:
			if c, ok := c.(ast.Expr); ok && c == m.X {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.X = n
					replaced = true
				}
			}

		case *ast.AssignStmt:
			if c, ok := c.([]ast.Expr); ok && &c == &m.Lhs { // list comparison by pointers
				found = true
				if n, ok := n.([]ast.Expr); ok {
					m.Lhs = n
					replaced = true
				}
			}
			if c, ok := c.([]ast.Expr); ok && &c == &m.Rhs { // list comparison by pointers
				found = true
				if n, ok := n.([]ast.Expr); ok {
					m.Rhs = n
					replaced = true
				}
			}

		case *ast.GoStmt:
			if c, ok := c.(*ast.CallExpr); ok && c == m.Call {
				found = true
				if n, ok := n.(*ast.CallExpr); ok {
					m.Call = n
					replaced = true
				}
			}

		case *ast.DeferStmt:
			if c, ok := c.(*ast.CallExpr); ok && c == m.Call {
				found = true
				if n, ok := n.(*ast.CallExpr); ok {
					m.Call = n
					replaced = true
				}
			}

		case *ast.ReturnStmt:
			if c, ok := c.([]ast.Expr); ok && &c == &m.Results { // list comparison by pointers
				found = true
				if n, ok := n.([]ast.Expr); ok {
					m.Results = n
					replaced = true
				}
			}

		case *ast.BranchStmt:
			if c, ok := c.(*ast.Ident); ok && c == m.Label {
				found = true
				if n, ok := n.(*ast.Ident); ok {
					m.Label = n
					replaced = true
				}
			}

		case *ast.BlockStmt:
			if c, ok := c.([]ast.Stmt); ok && &c == &m.List { // list comparison by pointers
				found = true
				if n, ok := n.([]ast.Stmt); ok {
					m.List = n
					replaced = true
				}
			}

		case *ast.IfStmt:
			if c, ok := c.(ast.Stmt); ok && c == m.Init {
				found = true
				if n, ok := n.(ast.Stmt); ok {
					m.Init = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Cond {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Cond = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.BlockStmt); ok && c == m.Body {
				found = true
				if n, ok := n.(*ast.BlockStmt); ok {
					m.Body = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Stmt); ok && c == m.Else {
				found = true
				if n, ok := n.(ast.Stmt); ok {
					m.Else = n
					replaced = true
				}
			}

		case *ast.CaseClause:
			if c, ok := c.([]ast.Expr); ok && &c == &m.List { // list comparison by pointers
				found = true
				if n, ok := n.([]ast.Expr); ok {
					m.List = n
					replaced = true
				}
			}
			if c, ok := c.([]ast.Stmt); ok && &c == &m.Body { // list comparison by pointers
				found = true
				if n, ok := n.([]ast.Stmt); ok {
					m.Body = n
					replaced = true
				}
			}

		case *ast.SwitchStmt:
			if c, ok := c.(ast.Stmt); ok && c == m.Init {
				found = true
				if n, ok := n.(ast.Stmt); ok {
					m.Init = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Tag {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Tag = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.BlockStmt); ok && c == m.Body {
				found = true
				if n, ok := n.(*ast.BlockStmt); ok {
					m.Body = n
					replaced = true
				}
			}

		case *ast.TypeSwitchStmt:
			if c, ok := c.(ast.Stmt); ok && c == m.Init {
				found = true
				if n, ok := n.(ast.Stmt); ok {
					m.Init = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Stmt); ok && c == m.Assign {
				found = true
				if n, ok := n.(ast.Stmt); ok {
					m.Assign = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.BlockStmt); ok && c == m.Body {
				found = true
				if n, ok := n.(*ast.BlockStmt); ok {
					m.Body = n
					replaced = true
				}
			}

		case *ast.CommClause:
			if c, ok := c.(ast.Stmt); ok && c == m.Comm {
				found = true
				if n, ok := n.(ast.Stmt); ok {
					m.Comm = n
					replaced = true
				}
			}
			if c, ok := c.([]ast.Stmt); ok && &c == &m.Body { // list comparison by pointers
				found = true
				if n, ok := n.([]ast.Stmt); ok {
					m.Body = n
					replaced = true
				}
			}

		case *ast.SelectStmt:
			if c, ok := c.(*ast.BlockStmt); ok && c == m.Body {
				found = true
				if n, ok := n.(*ast.BlockStmt); ok {
					m.Body = n
					replaced = true
				}
			}

		case *ast.ForStmt:
			if c, ok := c.(ast.Stmt); ok && c == m.Init {
				found = true
				if n, ok := n.(ast.Stmt); ok {
					m.Init = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Cond {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Cond = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Stmt); ok && c == m.Post {
				found = true
				if n, ok := n.(ast.Stmt); ok {
					m.Post = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.BlockStmt); ok && c == m.Body {
				found = true
				if n, ok := n.(*ast.BlockStmt); ok {
					m.Body = n
					replaced = true
				}
			}

		case *ast.RangeStmt:
			if c, ok := c.(ast.Expr); ok && c == m.Key {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Key = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Value {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Value = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.X {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.X = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.BlockStmt); ok && c == m.Body {
				found = true
				if n, ok := n.(*ast.BlockStmt); ok {
					m.Body = n
					replaced = true
				}
			}

		// Declarations
		case *ast.ImportSpec:
			if c, ok := c.(*ast.CommentGroup); ok && c == m.Doc {
				found = true
				if n, ok := n.(*ast.CommentGroup); ok {
					m.Doc = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.Ident); ok && c == m.Name {
				found = true
				if n, ok := n.(*ast.Ident); ok {
					m.Name = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.BasicLit); ok && c == m.Path {
				found = true
				if n, ok := n.(*ast.BasicLit); ok {
					m.Path = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.CommentGroup); ok && c == m.Comment {
				found = true
				if n, ok := n.(*ast.CommentGroup); ok {
					m.Comment = n
					replaced = true
				}
			}

		case *ast.ValueSpec:
			if c, ok := c.(*ast.CommentGroup); ok && c == m.Doc {
				found = true
				if n, ok := n.(*ast.CommentGroup); ok {
					m.Doc = n
					replaced = true
				}
			}
			if c, ok := c.([]*ast.Ident); ok && &c == &m.Names { // list comparison by pointers
				found = true
				if n, ok := n.([]*ast.Ident); ok {
					m.Names = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Type {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Type = n
					replaced = true
				}
			}
			if c, ok := c.([]ast.Expr); ok && &c == &m.Values { // list comparison by pointers
				found = true
				if n, ok := n.([]ast.Expr); ok {
					m.Values = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.CommentGroup); ok && c == m.Comment {
				found = true
				if n, ok := n.(*ast.CommentGroup); ok {
					m.Comment = n
					replaced = true
				}
			}

		case *ast.TypeSpec:
			if c, ok := c.(*ast.CommentGroup); ok && c == m.Doc {
				found = true
				if n, ok := n.(*ast.CommentGroup); ok {
					m.Doc = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.Ident); ok && c == m.Name {
				found = true
				if n, ok := n.(*ast.Ident); ok {
					m.Name = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.FieldList); ok && c == m.TypeParams { // fields
				found = true
				if n, ok := n.(*ast.FieldList); ok {
					m.TypeParams = n
					replaced = true
				}
			}
			if c, ok := c.(ast.Expr); ok && c == m.Type {
				found = true
				if n, ok := n.(ast.Expr); ok {
					m.Type = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.CommentGroup); ok && c == m.Comment {
				found = true
				if n, ok := n.(*ast.CommentGroup); ok {
					m.Comment = n
					replaced = true
				}
			}

		case *ast.GenDecl:
			if c, ok := c.(*ast.CommentGroup); ok && c == m.Doc {
				found = true
				if n, ok := n.(*ast.CommentGroup); ok {
					m.Doc = n
					replaced = true
				}
			}
			if c, ok := c.([]ast.Spec); ok && &c == &m.Specs { // list comparison by pointers
				found = true
				if n, ok := n.([]ast.Spec); ok {
					m.Specs = n
					replaced = true
				}
			}

		case *ast.FuncDecl:
			if c, ok := c.(*ast.CommentGroup); ok && c == m.Doc {
				found = true
				if n, ok := n.(*ast.CommentGroup); ok {
					m.Doc = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.FieldList); ok && c == m.Recv {
				found = true
				if n, ok := n.(*ast.FieldList); ok {
					m.Recv = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.Ident); ok && c == m.Name {
				found = true
				if n, ok := n.(*ast.Ident); ok {
					m.Name = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.FuncType); ok && c == m.Type {
				found = true
				if n, ok := n.(*ast.FuncType); ok {
					m.Type = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.BlockStmt); ok && c == m.Body {
				found = true
				if n, ok := n.(*ast.BlockStmt); ok {
					m.Body = n
					replaced = true
				}
			}

		// Files and packages

		case *ast.File:
			if c, ok := c.(*ast.CommentGroup); ok && c == m.Doc {
				found = true
				if n, ok := n.(*ast.CommentGroup); ok {
					m.Doc = n
					replaced = true
				}
			}
			if c, ok := c.(*ast.Ident); ok && c == m.Name {
				found = true
				if n, ok := n.(*ast.Ident); ok {
					m.Name = n
					replaced = true
				}
			}
			if c, ok := c.([]ast.Decl); ok && &c == &m.Decls { // list comparison by pointers
				found = true
				if n, ok := n.([]ast.Decl); ok {
					m.Decls = n
					replaced = true
				}
			}

		default:
			panic(fmt.Sprintf("Apply: unexpected node type %T", n))
		}

		return !found

	})
	return replaced
}
