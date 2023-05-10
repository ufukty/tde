package node_constructor

import (
	"tde/internal/cfg/context-resolution/context"
	"tde/internal/tokenw"
	utl "tde/internal/utilities"

	"go/ast"
)

var basicLitGenerators = []func() *ast.BasicLit{
	basicIntegerLiteral,
	basicStringLiteral,
	basicFloatLiteral,
	basicCharacterLiteral,
}

func BasicLit(ctx *context.Context, limit int) *ast.BasicLit {
	if limit == 0 {
		return nil
	}
	return (*utl.Pick(basicLitGenerators))()
}

func BinaryExpr(ctx *context.Context, limit int) *ast.BinaryExpr {
	if limit == 0 {
		return nil
	}
	return &ast.BinaryExpr{
		// OpPos: token.NoPos,
		X:  Expr(ctx, limit-1),
		Y:  Expr(ctx, limit-1),
		Op: *utl.Pick(tokenw.AcceptedByBinaryExpr),
	}
}

func CallExpr(ctx *context.Context, limit int) *ast.CallExpr {
	// TODO: function calls with more than 1 arguments
	if limit == 0 {
		return nil
	}
	return &ast.CallExpr{
		// Lparen:   token.NoPos,
		// Rparen:   token.NoPos,
		// Ellipsis: token.NoPos, // FIXME: variadic parameter support
		Fun:  Expr(ctx, limit-1),
		Args: []ast.Expr{Expr(ctx, limit-1)},
	}
}

func CompositeLit(ctx *context.Context, limit int) *ast.CompositeLit {
	// TODO: check Incomplete property
	if limit == 0 {
		return nil
	}
	return &ast.CompositeLit{
		// Lbrace:     token.NoPos,
		// Rbrace:     token.NoPos,
		Type:       Type(ctx, limit-1),
		Elts:       []ast.Expr{Expr(ctx, limit-1)},
		Incomplete: false,
	}
}

func Ellipsis(ctx *context.Context, limit int) *ast.Ellipsis {
	if limit == 0 {
		return nil
	}
	return &ast.Ellipsis{
		// Ellipsis: token.NoPos,
		Elt: Expr(ctx, limit-1),
	}
}

func FuncLit(ctx *context.Context, limit int) *ast.FuncLit {
	// TODO:
	if limit == 0 {
		return nil
	}
	return &ast.FuncLit{
		Type: FuncType(ctx, limit-1),
		Body: BlockStmt(ctx, limit-1),
	}
}

func Ident(ctx *context.Context, limit int) *ast.Ident {
	if limit == 0 {
		return nil
	}
	return generateNewIdent()
}

func IndexExpr(ctx *context.Context, limit int) *ast.IndexExpr {
	if limit == 0 {
		return nil
	}
	return &ast.IndexExpr{
		// Lbrack: token.NoPos,
		// Rbrack: token.NoPos,
		X:     Expr(ctx, limit-1),
		Index: Expr(ctx, limit-1),
	}
}

func IndexListExpr(ctx *context.Context, limit int) *ast.IndexListExpr {
	// TODO: Multi-dimensional arrays
	if limit == 0 {
		return nil
	}
	return &ast.IndexListExpr{
		// Lbrack:  token.NoPos,
		// Rbrack:  token.NoPos,
		X:       Expr(ctx, limit-1),
		Indices: []ast.Expr{Expr(ctx, limit-1)},
	}
}

func KeyValueExpr(ctx *context.Context, limit int) *ast.KeyValueExpr {
	if limit == 0 {
		return nil
	}
	return &ast.KeyValueExpr{
		// Colon: token.NoPos,
		Key:   Expr(ctx, limit-1),
		Value: Expr(ctx, limit-1),
	}
}

func ParenExpr(ctx *context.Context, limit int) *ast.ParenExpr {
	if limit == 0 {
		return nil
	}
	return &ast.ParenExpr{
		// Lparen: token.NoPos,
		// Rparen: token.NoPos,
		X: Expr(ctx, limit-1),
	}
}

func SelectorExpr(ctx *context.Context, limit int) *ast.SelectorExpr {
	// FIXME: randomly produced X and Sel values will never work, maybe choose from imported libraries' exported functions, or previously declared struct instances that has methods
	if limit == 0 {
		return nil
	}
	return &ast.SelectorExpr{
		X:   Expr(ctx, limit-1),
		Sel: &ast.Ident{},
	}
}

func SliceExpr(ctx *context.Context, limit int) *ast.SliceExpr {
	if limit == 0 {
		return nil
	}
	return &ast.SliceExpr{
		// Lbrack: token.NoPos,
		// Rbrack: token.NoPos,
		X:      Expr(ctx, limit-1),
		Low:    Expr(ctx, limit-1),
		High:   Expr(ctx, limit-1),
		Max:    nil,
		Slice3: false,
	}
}

func StarExpr(ctx *context.Context, limit int) *ast.StarExpr {
	if limit == 0 {
		return nil
	}
	return &ast.StarExpr{
		// Star: token.NoPos,
		X: Expr(ctx, limit-1),
	}
}

func TypeAssertExpr(ctx *context.Context, limit int) *ast.TypeAssertExpr {
	if limit == 0 {
		return nil
	}
	return &ast.TypeAssertExpr{
		// Lparen: token.NoPos,
		// Rparen: token.NoPos,
		X:    Expr(ctx, limit-1),
		Type: InterfaceType(ctx, limit-1),
	}
}

func UnaryExpr(ctx *context.Context, limit int) *ast.UnaryExpr {
	if limit == 0 {
		return nil
	}
	return &ast.UnaryExpr{
		// OpPos: token.NoPos,
		X:  Expr(ctx, limit-1),
		Op: *utl.Pick(tokenw.AcceptedByUnaryExpr),
	}
}
