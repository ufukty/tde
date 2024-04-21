package ctxres

import (
	"tde/internal/astw/traverse"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"

	"go/ast"
)

func GetContextForSpot(pkg *ast.Package, tFuncDecl, tSpot *traverse.TraversableNode) (*context.Context, error) {
	ctx := context.NewContext()

	ctx.ExeminePkg(pkg)
	ctx.ScopeIn()
	ctx.ExamineFuncDecl(tFuncDecl, tSpot)

	return ctx, nil
}
