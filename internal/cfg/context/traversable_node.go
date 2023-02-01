package context

import (
	"go/ast"
	trav "tde/internal/astw/traverse"
)

func GetContextForSpot(pkg *ast.Package, tFuncDecl, tSpot *trav.TraversableNode) (*Context, error) {
	ctx := NewContext()

	Package(ctx, pkg)
	FillContextForFunctionDeclaration(ctx, tFuncDecl, tSpot)

	return ctx, nil
}
