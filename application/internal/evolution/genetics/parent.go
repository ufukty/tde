package genetics

import (
	"fmt"
	"go/ast"
	"tde/internal/astw/astwutl"
)

func findParent(r ast.Node, n any) (p ast.Node) {
	ast.Inspect(r, func(m ast.Node) bool {
		if r == m || m == nil || astwutl.IsNodeNil(m) {
			return p == nil
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
			if n, ok := n.([]*ast.Comment); ok && &n == &m.List { // list comparison by pointers
				p = m
			}

		case *ast.Field:
			if n, ok := n.(*ast.CommentGroup); ok && n == m.Doc {
				p = m
			}
			if n, ok := n.([]*ast.Ident); ok && &n == &m.Names { // list comparison by pointers
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Type {
				p = m
			}
			if n, ok := n.(*ast.BasicLit); ok && n == m.Tag {
				p = m
			}
			if n, ok := n.(*ast.CommentGroup); ok && n == m.Comment {
				p = m
			}

		case *ast.FieldList:
			if n, ok := n.([]*ast.Field); ok && &n == &m.List { // list comparison by pointers
				p = m
			}

		// Expressions

		case *ast.Ellipsis:
			if n, ok := n.(ast.Expr); ok && n == m.Elt {
				p = m
			}

		case *ast.FuncLit:
			if n, ok := n.(*ast.FuncType); ok && n == m.Type {
				p = m
			}
			if n, ok := n.(*ast.BlockStmt); ok && n == m.Body {
				p = m
			}

		case *ast.CompositeLit:
			if n, ok := n.(ast.Expr); ok && n == m.Type {
				p = m
			}
			if n, ok := n.([]ast.Expr); ok && &n == &m.Elts { // list comparison by pointers
				p = m
			}

		case *ast.ParenExpr:
			if n, ok := n.(ast.Expr); ok && n == m.X {
				p = m
			}

		case *ast.SelectorExpr:
			if n, ok := n.(ast.Expr); ok && n == m.X {
				p = m
			}
			if n, ok := n.(*ast.Ident); ok && n == m.Sel {
				p = m
			}

		case *ast.IndexExpr:
			if n, ok := n.(ast.Expr); ok && n == m.X {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Index {
				p = m
			}

		case *ast.IndexListExpr:
			if n, ok := n.(ast.Expr); ok && n == m.X {
				p = m
			}
			if n, ok := n.([]ast.Expr); ok && &n == &m.Indices { // list comparison by pointers
				p = m
			}

		case *ast.SliceExpr:
			if n, ok := n.(ast.Expr); ok && n == m.X {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Low {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.High {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Max {
				p = m
			}

		case *ast.TypeAssertExpr:
			if n, ok := n.(ast.Expr); ok && n == m.X {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Type {
				p = m
			}

		case *ast.CallExpr:
			if n, ok := n.(ast.Expr); ok && n == m.Fun {
				p = m
			}
			if n, ok := n.([]ast.Expr); ok && &n == &m.Args { // list comparison by pointers
				p = m
			}

		case *ast.StarExpr:
			if n, ok := n.(ast.Expr); ok && n == m.X {
				p = m
			}

		case *ast.UnaryExpr:
			if n, ok := n.(ast.Expr); ok && n == m.X {
				p = m
			}

		case *ast.BinaryExpr:
			if n, ok := n.(ast.Expr); ok && n == m.X {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Y {
				p = m
			}

		case *ast.KeyValueExpr:
			if n, ok := n.(ast.Expr); ok && n == m.Key {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Value {
				p = m
			}

		// Types
		case *ast.ArrayType:
			if n, ok := n.(ast.Expr); ok && n == m.Len {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Elt {
				p = m
			}

		case *ast.StructType:
			if n, ok := n.(*ast.FieldList); ok && n == m.Fields {
				p = m
			}

		case *ast.FuncType:
			if n, ok := n.(*ast.FieldList); ok && n == m.TypeParams { // fields
				p = m
			}
			if n, ok := n.(*ast.FieldList); ok && n == m.Params {
				p = m
			}
			if n, ok := n.(*ast.FieldList); ok && n == m.Results {
				p = m
			}

		case *ast.InterfaceType:
			if n, ok := n.(*ast.FieldList); ok && n == m.Methods {
				p = m
			}

		case *ast.MapType:
			if n, ok := n.(ast.Expr); ok && n == m.Key {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Value {
				p = m
			}

		case *ast.ChanType:
			if n, ok := n.(ast.Expr); ok && n == m.Value {
				p = m
			}

		// Statements

		case *ast.DeclStmt:
			if n, ok := n.(ast.Decl); ok && n == m.Decl {
				p = m
			}

		case *ast.LabeledStmt:
			if n, ok := n.(*ast.Ident); ok && n == m.Label {
				p = m
			}
			if n, ok := n.(ast.Stmt); ok && n == m.Stmt {
				p = m
			}

		case *ast.ExprStmt:
			if n, ok := n.(ast.Expr); ok && n == m.X {
				p = m
			}

		case *ast.SendStmt:
			if n, ok := n.(ast.Expr); ok && n == m.Chan {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Value {
				p = m
			}

		case *ast.IncDecStmt:
			if n, ok := n.(ast.Expr); ok && n == m.X {
				p = m
			}

		case *ast.AssignStmt:
			if n, ok := n.([]ast.Expr); ok && &n == &m.Lhs { // list comparison by pointers
				p = m
			}
			if n, ok := n.([]ast.Expr); ok && &n == &m.Rhs { // list comparison by pointers
				p = m
			}

		case *ast.GoStmt:
			if n, ok := n.(*ast.CallExpr); ok && n == m.Call {
				p = m
			}

		case *ast.DeferStmt:
			if n, ok := n.(*ast.CallExpr); ok && n == m.Call {
				p = m
			}

		case *ast.ReturnStmt:
			if n, ok := n.([]ast.Expr); ok && &n == &m.Results { // list comparison by pointers
				p = m
			}

		case *ast.BranchStmt:
			if n, ok := n.(*ast.Ident); ok && n == m.Label {
				p = m
			}

		case *ast.BlockStmt:
			if n, ok := n.([]ast.Stmt); ok && &n == &m.List { // list comparison by pointers
				p = m
			}

		case *ast.IfStmt:
			if n, ok := n.(ast.Stmt); ok && n == m.Init {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Cond {
				p = m
			}
			if n, ok := n.(*ast.BlockStmt); ok && n == m.Body {
				p = m
			}
			if n, ok := n.(ast.Stmt); ok && n == m.Else {
				p = m
			}

		case *ast.CaseClause:
			if n, ok := n.([]ast.Expr); ok && &n == &m.List { // list comparison by pointers
				p = m
			}
			if n, ok := n.([]ast.Stmt); ok && &n == &m.Body { // list comparison by pointers
				p = m
			}

		case *ast.SwitchStmt:
			if n, ok := n.(ast.Stmt); ok && n == m.Init {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Tag {
				p = m
			}
			if n, ok := n.(*ast.BlockStmt); ok && n == m.Body {
				p = m
			}

		case *ast.TypeSwitchStmt:
			if n, ok := n.(ast.Stmt); ok && n == m.Init {
				p = m
			}
			if n, ok := n.(ast.Stmt); ok && n == m.Assign {
				p = m
			}
			if n, ok := n.(*ast.BlockStmt); ok && n == m.Body {
				p = m
			}

		case *ast.CommClause:
			if n, ok := n.(ast.Stmt); ok && n == m.Comm {
				p = m
			}
			if n, ok := n.([]ast.Stmt); ok && &n == &m.Body { // list comparison by pointers
				p = m
			}

		case *ast.SelectStmt:
			if n, ok := n.(*ast.BlockStmt); ok && n == m.Body {
				p = m
			}

		case *ast.ForStmt:
			if n, ok := n.(ast.Stmt); ok && n == m.Init {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Cond {
				p = m
			}
			if n, ok := n.(ast.Stmt); ok && n == m.Post {
				p = m
			}
			if n, ok := n.(*ast.BlockStmt); ok && n == m.Body {
				p = m
			}

		case *ast.RangeStmt:
			if n, ok := n.(ast.Expr); ok && n == m.Key {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Value {
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.X {
				p = m
			}
			if n, ok := n.(*ast.BlockStmt); ok && n == m.Body {
				p = m
			}

		// Declarations
		case *ast.ImportSpec:
			if n, ok := n.(*ast.CommentGroup); ok && n == m.Doc {
				p = m
			}
			if n, ok := n.(*ast.Ident); ok && n == m.Name {
				p = m
			}
			if n, ok := n.(*ast.BasicLit); ok && n == m.Path {
				p = m
			}
			if n, ok := n.(*ast.CommentGroup); ok && n == m.Comment {
				p = m
			}

		case *ast.ValueSpec:
			if n, ok := n.(*ast.CommentGroup); ok && n == m.Doc {
				p = m
			}
			if n, ok := n.([]*ast.Ident); ok && &n == &m.Names { // list comparison by pointers
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Type {
				p = m
			}
			if n, ok := n.([]ast.Expr); ok && &n == &m.Values { // list comparison by pointers
				p = m
			}
			if n, ok := n.(*ast.CommentGroup); ok && n == m.Comment {
				p = m
			}

		case *ast.TypeSpec:
			if n, ok := n.(*ast.CommentGroup); ok && n == m.Doc {
				p = m
			}
			if n, ok := n.(*ast.Ident); ok && n == m.Name {
				p = m
			}
			if n, ok := n.(*ast.FieldList); ok && n == m.TypeParams { // fields
				p = m
			}
			if n, ok := n.(ast.Expr); ok && n == m.Type {
				p = m
			}
			if n, ok := n.(*ast.CommentGroup); ok && n == m.Comment {
				p = m
			}

		case *ast.GenDecl:
			if n, ok := n.(*ast.CommentGroup); ok && n == m.Doc {
				p = m
			}
			if n, ok := n.([]ast.Spec); ok && &n == &m.Specs { // list comparison by pointers
				p = m
			}

		case *ast.FuncDecl:
			if n, ok := n.(*ast.CommentGroup); ok && n == m.Doc {
				p = m
			}
			if n, ok := n.(*ast.FieldList); ok && n == m.Recv {
				p = m
			}
			if n, ok := n.(*ast.Ident); ok && n == m.Name {
				p = m
			}
			if n, ok := n.(*ast.FuncType); ok && n == m.Type {
				p = m
			}
			if n, ok := n.(*ast.BlockStmt); ok && n == m.Body {
				p = m
			}

		// Files and packages

		case *ast.File:
			if n, ok := n.(*ast.CommentGroup); ok && n == m.Doc {
				p = m
			}
			if n, ok := n.(*ast.Ident); ok && n == m.Name {
				p = m
			}
			if n, ok := n.([]ast.Decl); ok && &n == &m.Decls { // list comparison by pointers
				p = m
			}

		default:
			panic(fmt.Sprintf("unexpected node type %T", n))
		}

		return p == nil
	})
	return
}
