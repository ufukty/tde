package genetics

import (
	"fmt"
	"go/ast"
)

func isNil(n any) bool {
	switch n := n.(type) {
	case []*ast.Field:
		return n == nil
	case []*ast.Ident:
		return n == nil
	case []ast.Decl:
		return n == nil
	case []ast.Expr:
		return n == nil
	case []ast.Spec:
		return n == nil
	case []ast.Stmt:
		return n == nil
	case *ast.BasicLit:
		return n == nil
	case *ast.BlockStmt:
		return n == nil
	case *ast.CallExpr:
		return n == nil
	case *ast.FieldList:
		return n == nil
	case *ast.FuncType:
		return n == nil
	case *ast.Ident:
		return n == nil
	case ast.Decl:
		return n == nil
	case ast.Expr:
		return n == nil
	case ast.Stmt:
		return n == nil
	default:
		panic(fmt.Sprintf("unexpected type: %T", n))
	}
}
