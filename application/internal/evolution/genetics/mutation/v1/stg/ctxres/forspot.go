package ctxres

import (
	"tde/internal/astw/traverse"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/examiner"

	"go/ast"
)

func GetContextForSpot(pkg *ast.Package, tFuncDecl, tSpot *traverse.TraversableNode) (*context.Context, error) {
	ctx := context.NewContext()

	examiner.Pkg(ctx, pkg)
	ctx.ScopeIn()
	examiner.FuncDecl(ctx, tFuncDecl, tSpot)

	return ctx, nil
}
