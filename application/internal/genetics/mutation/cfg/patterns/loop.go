package patterns

import (
	"tde/internal/genetics/mutation/cfg/ctxres/context"
	"tde/internal/genetics/mutation/cfg/nodes"

	"go/ast"
	"go/token"
)

func LoopOverSlice(ctx *context.Context, limit int, slice *ast.Ident) *ast.RangeStmt {
	if limit == 0 {
		return nil
	}
	return &ast.RangeStmt{
		Key:   nil,
		Value: nil,
		X:     nil,
		Body:  nodes.BlockStmt(ctx, limit-1),
	}
}

//	for i := 0; i < min(len(a), len(b)); i++ {
//	  a_, b_ := a[i], b[i]
//	  ...
//	}
func LoopOverSlices(ctx *context.Context, limit int, sliceA, sliceB *ast.Ident) *ast.ForStmt {
	if limit == 0 {
		return nil
	}
	return &ast.ForStmt{
		Init: nil,
		Cond: &ast.BinaryExpr{},
		Post: nil,
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.DEFINE,
					Lhs: []ast.Expr{
						&ast.Ident{Name: "valueA"},
						&ast.Ident{Name: "valueB"},
					},
					Rhs: []ast.Expr{
						&ast.IndexExpr{
							X:     sliceA,
							Index: &ast.Ident{Name: "i"},
						},
						&ast.IndexExpr{
							X:     sliceB,
							Index: &ast.Ident{Name: "i"},
						},
					},
				},
			},
		},
	}
}
