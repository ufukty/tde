package node_constructor

import (
	"tde/internal/cfg/context"
	utl "tde/internal/utilities"

	"go/ast"
	"go/token"
)

func BasicLit(ctx *context.Context, limit int) *ast.BasicLit {
	if limit == 0 {
		return nil
	}
	return (*utl.Pick([]func() *ast.BasicLit{
		basicIntegerLiteral,
		basicStringLiteral,
		basicFloatLiteral,
		basicCharacterLiteral,
	}))()
}

func BinaryExpr(ctx *context.Context, limit int) *ast.BinaryExpr {
	if limit == 0 {
		return nil
	}
	return &ast.BinaryExpr{
		X:     Expr(ctx, limit-1),
		OpPos: token.NoPos,
		Op:    *utl.Pick(tokenConstructor.AcceptedByBinaryExpr),
		Y:     Expr(ctx, limit-1),
	}
}

func CallExpr(ctx *context.Context, limit int) *ast.CallExpr {
	// TODO: function calls with more than 1 arguments
	if limit == 0 {
		return nil
	}
	return &ast.CallExpr{
		Fun:      Expr(ctx, limit-1),
		Lparen:   token.NoPos,
		Args:     []ast.Expr{Expr(ctx, limit-1)},
		Ellipsis: token.NoPos,
		Rparen:   token.NoPos,
	}
}

func CompositeLit(ctx *context.Context, limit int) *ast.CompositeLit {
	// TODO: check Incomplete property
	if limit == 0 {
		return nil
	}
	return &ast.CompositeLit{
		Type:       Type(ctx, limit-1),
		Lbrace:     token.NoPos,
		Elts:       []ast.Expr{Expr(ctx, limit-1)},
		Rbrace:     token.NoPos,
		Incomplete: false,
	}
}

func Ellipsis(ctx *context.Context, limit int) *ast.Ellipsis {
	if limit == 0 {
		return nil
	}
	return &ast.Ellipsis{
		Ellipsis: token.NoPos,
		Elt:      Expr(ctx, limit-1),
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
		X:      Expr(ctx, limit-1),
		Lbrack: token.NoPos,
		Index:  Expr(ctx, limit-1),
		Rbrack: token.NoPos,
	}
}

func IndexListExpr(ctx *context.Context, limit int) *ast.IndexListExpr {
	// TODO: Multi-dimensional arrays
	if limit == 0 {
		return nil
	}
	return &ast.IndexListExpr{
		X:       Expr(ctx, limit-1),
		Lbrack:  token.NoPos,
		Indices: []ast.Expr{Expr(ctx, limit-1)},
		Rbrack:  token.NoPos,
	}
}

func KeyValueExpr(ctx *context.Context, limit int) *ast.KeyValueExpr {
	if limit == 0 {
		return nil
	}
	return &ast.KeyValueExpr{
		Key:   Expr(ctx, limit-1),
		Colon: token.NoPos,
		Value: Expr(ctx, limit-1),
	}
}

func ParenExpr(ctx *context.Context, limit int) *ast.ParenExpr {
	if limit == 0 {
		return nil
	}
	return &ast.ParenExpr{
		Lparen: token.NoPos,
		X:      Expr(ctx, limit-1),
		Rparen: token.NoPos,
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
		X:      Expr(ctx, limit-1),
		Lbrack: token.NoPos,
		Low:    Expr(ctx, limit-1),
		High:   Expr(ctx, limit-1),
		Max:    nil,
		Slice3: false,
		Rbrack: token.NoPos,
	}
}

func StarExpr(ctx *context.Context, limit int) *ast.StarExpr {
	if limit == 0 {
		return nil
	}
	return &ast.StarExpr{
		Star: token.NoPos,
		X:    Expr(ctx, limit-1),
	}
}

func TypeAssertExpr(ctx *context.Context, limit int) *ast.TypeAssertExpr {
	if limit == 0 {
		return nil
	}
	return &ast.TypeAssertExpr{
		X:      Expr(ctx, limit-1),
		Lparen: token.NoPos,
		Type:   InterfaceType(ctx, limit-1),
		Rparen: token.NoPos,
	}
}

func UnaryExpr(ctx *context.Context, limit int) *ast.UnaryExpr {
	if limit == 0 {
		return nil
	}
	return &ast.UnaryExpr{
		OpPos: token.NoPos,
		Op:    *utl.Pick(tokenConstructor.AcceptedByUnaryExpr),
		X:     Expr(ctx, limit-1),
	}
}
