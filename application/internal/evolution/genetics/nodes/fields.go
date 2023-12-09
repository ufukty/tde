package nodes

import (
	"go/ast"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"
)

func Field(ctx *context.Context, limit int) (*ast.Field, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.Field{
		Names: []*ast.Ident{
			Ident(ctx, limit-1),
		},
		Type: Type(ctx, limit-1),
		Tag:  nil,
	}, nil
}

func FieldList(ctx *context.Context, limit int) (*ast.FieldList, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.FieldList{
		// Opening: token.NoPos,
		// Closing: token.NoPos,
		List: []*ast.Field{
			Field(ctx, limit-1),
		},
	}, nil
}
