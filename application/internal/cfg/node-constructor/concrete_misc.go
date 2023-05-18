package node_constructor

import (
	"go/ast"
	"tde/internal/cfg/context-resolution/context"
)

func Field(ctx *context.Context, limit int) *ast.Field {
	if limit == 0 {
		return nil
	}
	return &ast.Field{
		Names: []*ast.Ident{
			Ident(ctx, limit-1),
		},
		Type: Type(ctx, limit-1),
		Tag:  nil,
	}
}

func FieldList(ctx *context.Context, limit int) *ast.FieldList {
	if limit == 0 {
		return nil
	}
	return &ast.FieldList{
		// Opening: token.NoPos,
		// Closing: token.NoPos,
		List: []*ast.Field{
			Field(ctx, limit-1),
		},
	}
}
