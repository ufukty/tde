package genetics

import (
	"fmt"
	"go/ast"
	"tde/internal/astw/types"
)

func typeOf(n any) types.NodeType {
	switch n := n.(type) {
	case *[]*ast.Comment:
		return types.CommentSlice
	case *[]*ast.Field:
		return types.FieldSlice
	case *[]*ast.Ident:
		return types.IdentSlice
	case *[]ast.Decl:
		return types.DeclSlice
	case *[]ast.Expr:
		return types.ExprSlice
	case *[]ast.Spec:
		return types.SpecSlice
	case *[]ast.Stmt:
		return types.StmtSlice
	case *ast.BasicLit:
		return types.BasicLit
	case *ast.BlockStmt:
		return types.BlockStmt
	case *ast.CallExpr:
		return types.CallExpr
	case *ast.CommentGroup:
		return types.CommentGroup
	case *ast.FieldList:
		return types.FieldList
	case *ast.FuncType:
		return types.FuncType
	case *ast.Ident:
		return types.Ident
	case ast.Decl:
		return types.Decl
	case ast.Expr:
		return types.Expr
	case ast.Stmt:
		return types.Stmt
	default:
		panic(fmt.Sprintf("unexpected type: %T", n))
	}
}

func isInType(o any, t types.NodeType) bool {
	switch t {
	case types.CommentSlice:
		_, ok := o.(*[]*ast.Comment)
		return ok
	case types.FieldSlice:
		_, ok := o.(*[]*ast.Field)
		return ok
	case types.IdentSlice:
		_, ok := o.(*[]*ast.Ident)
		return ok
	case types.DeclSlice:
		_, ok := o.(*[]ast.Decl)
		return ok
	case types.ExprSlice:
		_, ok := o.(*[]ast.Expr)
		return ok
	case types.SpecSlice:
		_, ok := o.(*[]ast.Spec)
		return ok
	case types.StmtSlice:
		_, ok := o.(*[]ast.Stmt)
		return ok
	case types.BasicLit:
		_, ok := o.(*ast.BasicLit)
		return ok
	case types.BlockStmt:
		_, ok := o.(*ast.BlockStmt)
		return ok
	case types.CallExpr:
		_, ok := o.(*ast.CallExpr)
		return ok
	case types.CommentGroup:
		_, ok := o.(*ast.CommentGroup)
		return ok
	case types.FieldList:
		_, ok := o.(*ast.FieldList)
		return ok
	case types.FuncType:
		_, ok := o.(*ast.FuncType)
		return ok
	case types.Ident:
		_, ok := o.(*ast.Ident)
		return ok
	case types.Decl:
		_, ok := o.(ast.Decl)
		return ok
	case types.Expr:
		_, ok := o.(ast.Expr)
		return ok
	case types.Stmt:
		_, ok := o.(ast.Stmt)
		return ok
	default:
		panic(fmt.Sprintf("unexpected value: %q", t))
	}
}
