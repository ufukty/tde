package ctxres

import (
	"go/ast"

	"tde/internal/ast/traverse"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"
)

func GetContextForSpot(pkg *ast.Package, tFuncDecl, tSpot *traverse.TraversableNode) (*context.Context, error) {
	ctx := context.NewContext()

	ctx.ExeminePkg(pkg)
	ctx.ScopeIn()
	ctx.ExamineFuncDecl(tFuncDecl, tSpot)

	return ctx, nil
}
