package context

import (
	"go/ast"
	"tde/internal/astw"
)

func GetContextForSpot(pkg *ast.Package, tFuncDecl, tSpot *astw.TraversableNode) (*Context, error) {
	ctx := NewContext()

	Package(ctx, pkg)
	FillContextForFunctionDeclaration(ctx, tFuncDecl, tSpot)

	return ctx, nil
}
