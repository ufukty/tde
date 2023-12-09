package nodes

import (
	"go/ast"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"
	"tde/internal/utilities/pick"
)

// only valid values are types such int, float, string, bool
func IdentType(ctx *context.Context, limit int) *ast.Ident {
	if limit == 0 {
		return nil
	}
	return ast.NewIdent(*pick.Pick([]string{"int", "float", "string", "bool"}))
}

func ArrayType(ctx *context.Context, limit int) *ast.ArrayType {
	// FIXME: // is there any usecase thar is not achievable with a slice but only with a ...T array
	if limit == 0 {
		return nil
	}
	return &ast.ArrayType{
		// Lbrack: token.NoPos,
		Len: nil,
		Elt: Type(ctx, limit-1),
	}
}

func ChanType(ctx *context.Context, limit int) *ast.ChanType {
	if limit == 0 {
		return nil
	}
	return &ast.ChanType{
		// Begin: token.NoPos,
		// Arrow: token.NoPos,
		Dir:   *pick.Pick([]ast.ChanDir{ast.SEND, ast.RECV}),
		Value: Type(ctx, limit-1),
	}
}

func FuncType(ctx *context.Context, limit int) *ast.FuncType {
	// FIXME:
	if limit == 0 {
		return nil
	}
	return &ast.FuncType{
		// Func:       token.NoPos,
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
		// Interface:  token.NoPos,
		Incomplete: false,
		Methods:    FieldList(ctx, limit-1),
	}
}

func MapType(ctx *context.Context, limit int) *ast.MapType {
	if limit == 0 {
		return nil
	}
	return &ast.MapType{
		// Map:   token.NoPos,
		Key:   Type(ctx, limit-1),
		Value: Type(ctx, limit-1),
	}
}

func StructType(ctx *context.Context, limit int) *ast.StructType {
	if limit == 0 {
		return nil
	}
	return &ast.StructType{
		// Struct:     token.NoPos,
		Incomplete: false,
		Fields:     FieldList(ctx, limit-1),
	}
}

func ParenExprForType(ctx *context.Context, limit int) *ast.ParenExpr {
	if limit == 0 {
		return nil
	}
	return &ast.ParenExpr{
		// Lparen: 0,
		// Rparen: 0,
		X: Type(ctx, limit-1),
	}
}

// should cover needs for:
//   - package.Type such as: types.NodeType
//     -
func SelectorExprForType(ctx *context.Context, limit int) *ast.SelectorExpr {
	if limit == 0 {
		return nil
	}
	return &ast.SelectorExpr{
		X:   Expr(ctx, limit-1),
		Sel: Ident(ctx, limit-1),
	}
}

func StarExprForType(ctx *context.Context, limit int) *ast.StarExpr {
	if limit == 0 {
		return nil
	}
	return &ast.StarExpr{
		// Star: 0,
		X: Type(ctx, limit-1),
	}
}
