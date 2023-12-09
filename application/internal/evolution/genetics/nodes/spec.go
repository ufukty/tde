package nodes

import (
	"go/ast"
	"go/token"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"
	"tde/internal/utilities/pick"
)

func ImportSpec(ctx *context.Context, limit int) (*ast.ImportSpec, error) {
	// TODO: Store imported packages for later use
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.ImportSpec{
		// EndPos: token.NoPos,
		Name: nil,
		Path: &ast.BasicLit{
			// ValuePos: token.NoPos,
			Kind:  token.STRING,
			Value: *pick.Pick(AllowedPackagesToImport),
		},
	}, nil
}

// rel. type keyword
func TypeSpec(ctx *context.Context, limit int) (*ast.TypeSpec, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.TypeSpec{
		// Assign:     token.NoPos,
		Doc:        nil,
		Name:       Ident(ctx, limit-1),
		TypeParams: FieldList(ctx, limit-1),
		Type:       nil,
		Comment:    nil,
	}, nil
}

// Returns ValueSpec which is list of pairs of variable names and values to assign
func ValueSpec(ctx *context.Context, limit int) (*ast.ValueSpec, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.ValueSpec{
		Names: []*ast.Ident{
			generateNewIdent(),
		},
		Values: []ast.Expr{Expr(ctx, limit-1)},
	}, nil
}
