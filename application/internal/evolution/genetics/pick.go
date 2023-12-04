package genetics

import (
	"fmt"
	"go/ast"
	"math/rand"
	"tde/internal/astw/types"
)

func pick[T any](s []T) (T, bool) {
	if len(s) > 0 {
		return s[rand.Intn(len(s))], true
	}
	return *new(T), false
}

func pickInType(r ast.Node, t types.NodeType) (any, bool) {
	switch t {
	case types.CommentSlice:
		return pick(linearizefiltered[*[]*ast.Comment](r))
	case types.FieldSlice:
		return pick(linearizefiltered[*[]*ast.Field](r))
	case types.IdentSlice:
		return pick(linearizefiltered[*[]*ast.Ident](r))

	case types.BasicLit:
		return pick(linearizefiltered[*ast.BasicLit](r))
	case types.BlockStmt:
		return pick(linearizefiltered[*ast.BlockStmt](r))
	case types.CallExpr:
		return pick(linearizefiltered[*ast.CallExpr](r))
	case types.CommentGroup:
		return pick(linearizefiltered[*ast.CommentGroup](r))
	case types.FieldList:
		return pick(linearizefiltered[*ast.FieldList](r))
	case types.FuncType:
		return pick(linearizefiltered[*ast.FuncType](r))
	case types.Ident:
		return pick(linearizefiltered[*ast.Ident](r))

	case types.DeclSlice:
		return pick(linearizefiltered[*[]ast.Decl](r))
	case types.ExprSlice:
		return pick(linearizefiltered[*[]ast.Expr](r))
	case types.SpecSlice:
		return pick(linearizefiltered[*[]ast.Spec](r))
	case types.StmtSlice:
		return pick(linearizefiltered[*[]ast.Stmt](r))

	case types.Decl:
		return pick(linearizefiltered[ast.Decl](r))
	case types.Expr:
		return pick(linearizefiltered[ast.Expr](r))
	case types.Stmt:
		return pick(linearizefiltered[ast.Stmt](r))
	}
	panic(fmt.Sprintf("unexpected type %q", t))
}
