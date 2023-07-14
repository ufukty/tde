package context_resolution

import (
	"tde/internal/astw/traverse"
	"tde/internal/cfg/context-resolution/context"
	"tde/internal/cfg/context-resolution/examiner/function_examiner"
	"tde/internal/cfg/context-resolution/examiner/package_examiner"

	"go/ast"
)

func GetContextForSpot(pkg *ast.Package, tFuncDecl, tSpot *traverse.TraversableNode) (*context.Context, error) {
	ctx := context.NewContext()

	package_examiner.Examine(ctx, pkg)
	ctx.ScopeIn()
	function_examiner.Examine(ctx, tFuncDecl, tSpot)

	return ctx, nil
}
