package genetics

import (
	"fmt"
	"go/ast"
	"tde/internal/astw/types"
)

// (p)arent -> *(f)ield
func fieldType(p ast.Node, f any) types.NodeType {
	switch p := p.(type) {

	// case
	//  *ast.Package,
	// 	*ast.Comment,
	// 	*ast.BadExpr,
	// 	*ast.Ident,
	// 	*ast.BasicLit,
	// 	*ast.BadStmt,
	// 	*ast.EmptyStmt,
	// 	*ast.BadDecl:

	case *ast.CommentGroup:
		if c, ok := f.([]*ast.Comment); ok && &c == &p.List {
			return types.CommentSlice
		}

	case *ast.Field:
		if c, ok := f.(*ast.CommentGroup); ok && c == p.Doc {
			return types.CommentGroup
		}
		if c, ok := f.([]*ast.Ident); ok && &c == &p.Names {
			return types.IdentSlice
		}
		if c, ok := f.(ast.Expr); ok && c == p.Type {
			return types.Expr
		}
		if c, ok := f.(*ast.BasicLit); ok && c == p.Tag {
			return types.BasicLit
		}
		if c, ok := f.(*ast.CommentGroup); ok && c == p.Comment {
			return types.CommentGroup
		}

	case *ast.FieldList:
		if c, ok := f.([]*ast.Field); ok && &c == &p.List {
			return types.FieldSlice
		}

	// Expressions

	case *ast.Ellipsis:
		if c, ok := f.(ast.Expr); ok && c == p.Elt {
			return types.Expr
		}

	case *ast.FuncLit:
		if c, ok := f.(*ast.FuncType); ok && c == p.Type {
			return types.FuncType
		}
		if c, ok := f.(*ast.BlockStmt); ok && c == p.Body {
			return types.BlockStmt
		}

	case *ast.CompositeLit:
		if c, ok := f.(ast.Expr); ok && c == p.Type {
			return types.Expr
		}
		if c, ok := f.([]ast.Expr); ok && &c == &p.Elts {
			return types.ExprSlice
		}

	case *ast.ParenExpr:
		if c, ok := f.(ast.Expr); ok && c == p.X {
			return types.Expr
		}

	case *ast.SelectorExpr:
		if c, ok := f.(ast.Expr); ok && c == p.X {
			return types.Expr
		}
		if c, ok := f.(*ast.Ident); ok && c == p.Sel {
			return types.Ident
		}

	case *ast.IndexExpr:
		if c, ok := f.(ast.Expr); ok && c == p.X {
			return types.Expr
		}
		if c, ok := f.(ast.Expr); ok && c == p.Index {
			return types.Expr
		}

	case *ast.IndexListExpr:
		if c, ok := f.(ast.Expr); ok && c == p.X {
			return types.Expr
		}
		if c, ok := f.([]ast.Expr); ok && &c == &p.Indices {
			return types.ExprSlice
		}

	case *ast.SliceExpr:
		if c, ok := f.(ast.Expr); ok && c == p.X {
			return types.Expr
		}
		if c, ok := f.(ast.Expr); ok && c == p.Low {
			return types.Expr
		}
		if c, ok := f.(ast.Expr); ok && c == p.High {
			return types.Expr
		}
		if c, ok := f.(ast.Expr); ok && c == p.Max {
			return types.Expr
		}

	case *ast.TypeAssertExpr:
		if c, ok := f.(ast.Expr); ok && c == p.X {
			return types.Expr
		}
		if c, ok := f.(ast.Expr); ok && c == p.Type {
			return types.Expr
		}

	case *ast.CallExpr:
		if c, ok := f.(ast.Expr); ok && c == p.Fun {
			return types.Expr
		}
		if c, ok := f.([]ast.Expr); ok && &c == &p.Args {
			return types.ExprSlice
		}

	case *ast.StarExpr:
		if c, ok := f.(ast.Expr); ok && c == p.X {
			return types.Expr
		}

	case *ast.UnaryExpr:
		if c, ok := f.(ast.Expr); ok && c == p.X {
			return types.Expr
		}

	case *ast.BinaryExpr:
		if c, ok := f.(ast.Expr); ok && c == p.X {
			return types.Expr
		}
		if c, ok := f.(ast.Expr); ok && c == p.Y {
			return types.Expr
		}

	case *ast.KeyValueExpr:
		if c, ok := f.(ast.Expr); ok && c == p.Key {
			return types.Expr
		}
		if c, ok := f.(ast.Expr); ok && c == p.Value {
			return types.Expr
		}

	// Types
	case *ast.ArrayType:
		if c, ok := f.(ast.Expr); ok && c == p.Len {
			return types.Expr
		}
		if c, ok := f.(ast.Expr); ok && c == p.Elt {
			return types.Expr
		}

	case *ast.StructType:
		if c, ok := f.(*ast.FieldList); ok && c == p.Fields {
			return types.FieldList
		}

	case *ast.FuncType:
		if c, ok := f.(*ast.FieldList); ok && c == p.TypeParams { // fields
			return types.FieldList
		}
		if c, ok := f.(*ast.FieldList); ok && c == p.Params {
			return types.FieldList
		}
		if c, ok := f.(*ast.FieldList); ok && c == p.Results {
			return types.FieldList
		}

	case *ast.InterfaceType:
		if c, ok := f.(*ast.FieldList); ok && c == p.Methods {
			return types.FieldList
		}

	case *ast.MapType:
		if c, ok := f.(ast.Expr); ok && c == p.Key {
			return types.Expr
		}
		if c, ok := f.(ast.Expr); ok && c == p.Value {
			return types.Expr
		}

	case *ast.ChanType:
		if c, ok := f.(ast.Expr); ok && c == p.Value {
			return types.Expr
		}

	// Statements

	case *ast.DeclStmt:
		if c, ok := f.(ast.Decl); ok && c == p.Decl {
			return types.Decl
		}

	case *ast.LabeledStmt:
		if c, ok := f.(*ast.Ident); ok && c == p.Label {
			return types.Ident
		}
		if c, ok := f.(ast.Stmt); ok && c == p.Stmt {
			return types.Stmt
		}

	case *ast.ExprStmt:
		if c, ok := f.(ast.Expr); ok && c == p.X {
			return types.Expr
		}

	case *ast.SendStmt:
		if c, ok := f.(ast.Expr); ok && c == p.Chan {
			return types.Expr
		}
		if c, ok := f.(ast.Expr); ok && c == p.Value {
			return types.Expr
		}

	case *ast.IncDecStmt:
		if c, ok := f.(ast.Expr); ok && c == p.X {
			return types.Expr
		}

	case *ast.AssignStmt:
		if c, ok := f.([]ast.Expr); ok && &c == &p.Lhs {
			return types.ExprSlice
		}
		if c, ok := f.([]ast.Expr); ok && &c == &p.Rhs {
			return types.ExprSlice
		}

	case *ast.GoStmt:
		if c, ok := f.(*ast.CallExpr); ok && c == p.Call {
			return types.CallExpr
		}

	case *ast.DeferStmt:
		if c, ok := f.(*ast.CallExpr); ok && c == p.Call {
			return types.CallExpr
		}

	case *ast.ReturnStmt:
		if c, ok := f.([]ast.Expr); ok && &c == &p.Results {
			return types.ExprSlice
		}

	case *ast.BranchStmt:
		if c, ok := f.(*ast.Ident); ok && c == p.Label {
			return types.Ident
		}

	case *ast.BlockStmt:
		if c, ok := f.([]ast.Stmt); ok && &c == &p.List {
			return types.StmtSlice
		}

	case *ast.IfStmt:
		if c, ok := f.(ast.Stmt); ok && c == p.Init {
			return types.Stmt
		}
		if c, ok := f.(ast.Expr); ok && c == p.Cond {
			return types.Expr
		}
		if c, ok := f.(*ast.BlockStmt); ok && c == p.Body {
			return types.BlockStmt
		}
		if c, ok := f.(ast.Stmt); ok && c == p.Else {
			return types.Stmt
		}

	case *ast.CaseClause:
		if c, ok := f.([]ast.Expr); ok && &c == &p.List {
			return types.ExprSlice
		}
		if c, ok := f.([]ast.Stmt); ok && &c == &p.Body {
			return types.StmtSlice
		}

	case *ast.SwitchStmt:
		if c, ok := f.(ast.Stmt); ok && c == p.Init {
			return types.Stmt
		}
		if c, ok := f.(ast.Expr); ok && c == p.Tag {
			return types.Expr
		}
		if c, ok := f.(*ast.BlockStmt); ok && c == p.Body {
			return types.BlockStmt
		}

	case *ast.TypeSwitchStmt:
		if c, ok := f.(ast.Stmt); ok && c == p.Init {
			return types.Stmt
		}
		if c, ok := f.(ast.Stmt); ok && c == p.Assign {
			return types.Stmt
		}
		if c, ok := f.(*ast.BlockStmt); ok && c == p.Body {
			return types.BlockStmt
		}

	case *ast.CommClause:
		if c, ok := f.(ast.Stmt); ok && c == p.Comm {
			return types.Stmt
		}
		if c, ok := f.([]ast.Stmt); ok && &c == &p.Body {
			return types.StmtSlice
		}

	case *ast.SelectStmt:
		if c, ok := f.(*ast.BlockStmt); ok && c == p.Body {
			return types.BlockStmt
		}

	case *ast.ForStmt:
		if c, ok := f.(ast.Stmt); ok && c == p.Init {
			return types.Stmt
		}
		if c, ok := f.(ast.Expr); ok && c == p.Cond {
			return types.Expr
		}
		if c, ok := f.(ast.Stmt); ok && c == p.Post {
			return types.Stmt
		}
		if c, ok := f.(*ast.BlockStmt); ok && c == p.Body {
			return types.BlockStmt
		}

	case *ast.RangeStmt:
		if c, ok := f.(ast.Expr); ok && c == p.Key {
			return types.Expr
		}
		if c, ok := f.(ast.Expr); ok && c == p.Value {
			return types.Expr
		}
		if c, ok := f.(ast.Expr); ok && c == p.X {
			return types.Expr
		}
		if c, ok := f.(*ast.BlockStmt); ok && c == p.Body {
			return types.BlockStmt
		}

	// Declarations
	case *ast.ImportSpec:
		if c, ok := f.(*ast.CommentGroup); ok && c == p.Doc {
			return types.CommentGroup
		}
		if c, ok := f.(*ast.Ident); ok && c == p.Name {
			return types.Ident
		}
		if c, ok := f.(*ast.BasicLit); ok && c == p.Path {
			return types.BasicLit
		}
		if c, ok := f.(*ast.CommentGroup); ok && c == p.Comment {
			return types.CommentGroup
		}

	case *ast.ValueSpec:
		if c, ok := f.(*ast.CommentGroup); ok && c == p.Doc {
			return types.CommentGroup
		}
		if c, ok := f.([]*ast.Ident); ok && &c == &p.Names {
			return types.IdentSlice
		}
		if c, ok := f.(ast.Expr); ok && c == p.Type {
			return types.Expr
		}
		if c, ok := f.([]ast.Expr); ok && &c == &p.Values {
			return types.ExprSlice
		}
		if c, ok := f.(*ast.CommentGroup); ok && c == p.Comment {
			return types.CommentGroup
		}

	case *ast.TypeSpec:
		if c, ok := f.(*ast.CommentGroup); ok && c == p.Doc {
			return types.CommentGroup
		}
		if c, ok := f.(*ast.Ident); ok && c == p.Name {
			return types.Ident
		}
		if c, ok := f.(*ast.FieldList); ok && c == p.TypeParams { // fields
			return types.FieldList
		}
		if c, ok := f.(ast.Expr); ok && c == p.Type {
			return types.Expr
		}
		if c, ok := f.(*ast.CommentGroup); ok && c == p.Comment {
			return types.CommentGroup
		}

	case *ast.GenDecl:
		if c, ok := f.(*ast.CommentGroup); ok && c == p.Doc {
			return types.CommentGroup
		}
		if c, ok := f.([]ast.Spec); ok && &c == &p.Specs {
			return types.SpecSlice
		}

	case *ast.FuncDecl:
		if c, ok := f.(*ast.CommentGroup); ok && c == p.Doc {
			return types.CommentGroup
		}
		if c, ok := f.(*ast.FieldList); ok && c == p.Recv {
			return types.FieldList
		}
		if c, ok := f.(*ast.Ident); ok && c == p.Name {
			return types.Ident
		}
		if c, ok := f.(*ast.FuncType); ok && c == p.Type {
			return types.FuncType
		}
		if c, ok := f.(*ast.BlockStmt); ok && c == p.Body {
			return types.BlockStmt
		}

	// Files and packages

	case *ast.File:
		if c, ok := f.(*ast.CommentGroup); ok && c == p.Doc {
			return types.CommentGroup
		}
		if c, ok := f.(*ast.Ident); ok && c == p.Name {
			return types.Ident
		}
		if c, ok := f.([]ast.Decl); ok && &c == &p.Decls {
			return types.DeclSlice
		}
	}

	panic(fmt.Sprintf("unexpected node type %T", f))
}
