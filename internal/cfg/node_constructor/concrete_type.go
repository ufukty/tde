package node_constructor

import (
	"go/ast"
	"go/token"
	"tde/internal/cfg/context_resolution/context"
	utl "tde/internal/utilities"
)

// only valid values are types such int, float, string, bool
func IdentType(ctx *context.Context, limit int) *ast.Ident {
	if limit == 0 {
		return nil
	}
	return ast.NewIdent(*utl.Pick([]string{"int", "float", "string", "bool"}))
}

func ArrayType(ctx *context.Context, limit int) *ast.ArrayType {
	// FIXME: // is there any usecase thar is not achievable with a slice but only with a ...T array
	if limit == 0 {
		return nil
	}
	return &ast.ArrayType{
		Lbrack: token.NoPos,
		Len:    nil,
		Elt:    Expr(ctx, limit-1),
	}
}

func ChanType(ctx *context.Context, limit int) *ast.ChanType {
	if limit == 0 {
		return nil
	}
	return &ast.ChanType{
		Begin: token.NoPos,
		Arrow: token.NoPos,
		Dir:   *utl.Pick([]ast.ChanDir{ast.SEND, ast.RECV}),
		Value: Type(ctx, limit-1),
	}
}

func FuncType(ctx *context.Context, limit int) *ast.FuncType {
	// FIXME:
	if limit == 0 {
		return nil
	}
	return &ast.FuncType{
		Func:       token.NoPos,
		TypeParams: FieldList(ctx, limit-1),
		Params:     FieldList(ctx, limit-1),
		Results:    FieldList(ctx, limit-1),
	}
}

func InterfaceType(ctx *context.Context, limit int) *ast.InterfaceType {
	if limit == 0 {
		return nil
	}
	return &ast.InterfaceType{
		Interface:  token.NoPos,
		Methods:    FieldList(ctx, limit-1),
		Incomplete: false,
	}
}

func MapType(ctx *context.Context, limit int) *ast.MapType {
	if limit == 0 {
		return nil
	}
	return &ast.MapType{
		Map:   token.NoPos,
		Key:   Type(ctx, limit-1),
		Value: Type(ctx, limit-1),
	}
}

func StructType(ctx *context.Context, limit int) *ast.StructType {
	if limit == 0 {
		return nil
	}
	return &ast.StructType{
		Struct:     token.NoPos,
		Fields:     FieldList(ctx, limit-1),
		Incomplete: false,
	}
}
