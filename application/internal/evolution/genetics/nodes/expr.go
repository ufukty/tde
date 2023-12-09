package nodes

import (
	"go/ast"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"
	"tde/internal/evolution/genetics/nodes/tokens"
	"tde/internal/utilities/pick"
)

var basicLitGenerators = []func() *ast.BasicLit{
	basicIntegerLiteral,
	basicStringLiteral,
	basicFloatLiteral,
	basicCharacterLiteral,
}

func BasicLit(ctx *context.Context, limit int) (*ast.BasicLit, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return (*pick.Pick(basicLitGenerators))()
}

func BinaryExpr(ctx *context.Context, limit int) (*ast.BinaryExpr, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.BinaryExpr{
		// OpPos: token.NoPos,
		X:  Expr(ctx, limit-1),
		Y:  Expr(ctx, limit-1),
		Op: *pick.Pick(tokens.AcceptedByBinaryExpr),
	}, nil
}

func CallExpr(ctx *context.Context, limit int) (*ast.CallExpr, error) {
	// TODO: function calls with more than 1 arguments
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.CallExpr{
		// Lparen:   token.NoPos,
		// Rparen:   token.NoPos,
		// Ellipsis: token.NoPos, // FIXME: variadic parameter support
		Fun:  Expr(ctx, limit-1),
		Args: []ast.Expr{Expr(ctx, limit-1)},
	}, nil
}

func CompositeLit(ctx *context.Context, limit int) (*ast.CompositeLit, error) {
	// TODO: check Incomplete property
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.CompositeLit{
		// Lbrace:     token.NoPos,
		// Rbrace:     token.NoPos,
		Type:       Type(ctx, limit-1),
		Elts:       []ast.Expr{Expr(ctx, limit-1)},
		Incomplete: false,
	}, nil
}

func Ellipsis(ctx *context.Context, limit int) (*ast.Ellipsis, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.Ellipsis{
		// Ellipsis: token.NoPos,
		Elt: Expr(ctx, limit-1),
	}, nil
}

func FuncLit(ctx *context.Context, limit int) (*ast.FuncLit, error) {
	// TODO:
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.FuncLit{
		Type: FuncType(ctx, limit-1),
		Body: BlockStmt(ctx, limit-1),
	}, nil
}

func Ident(ctx *context.Context, limit int) (*ast.Ident, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return generateNewIdent(), nil
}

func IndexExpr(ctx *context.Context, limit int) (*ast.IndexExpr, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.IndexExpr{
		// Lbrack: token.NoPos,
		// Rbrack: token.NoPos,
		X:     Expr(ctx, limit-1),
		Index: Expr(ctx, limit-1),
	}, nil
}

func IndexListExpr(ctx *context.Context, limit int) (*ast.IndexListExpr, error) {
	// TODO: Multi-dimensional arrays
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.IndexListExpr{
		// Lbrack:  token.NoPos,
		// Rbrack:  token.NoPos,
		X:       Expr(ctx, limit-1),
		Indices: []ast.Expr{Expr(ctx, limit-1)},
	}, nil
}

func KeyValueExpr(ctx *context.Context, limit int) (*ast.KeyValueExpr, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.KeyValueExpr{
		// Colon: token.NoPos,
		Key:   Expr(ctx, limit-1),
		Value: Expr(ctx, limit-1),
	}, nil
}

func ParenExpr(ctx *context.Context, limit int) (*ast.ParenExpr, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.ParenExpr{
		// Lparen: token.NoPos,
		// Rparen: token.NoPos,
		X: Expr(ctx, limit-1),
	}, nil
}

func SelectorExpr(ctx *context.Context, limit int) (*ast.SelectorExpr, error) {
	// FIXME: randomly produced X and Sel values will never work, maybe choose from imported libraries' exported functions, or previously declared struct instances that has methods
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.SelectorExpr{
		X:   Expr(ctx, limit-1),
		Sel: &ast.Ident{},
	}, nil
}

func SliceExpr(ctx *context.Context, limit int) (*ast.SliceExpr, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.SliceExpr{
		// Lbrack: token.NoPos,
		// Rbrack: token.NoPos,
		X:      Expr(ctx, limit-1),
		Low:    Expr(ctx, limit-1),
		High:   Expr(ctx, limit-1),
		Max:    nil,
		Slice3: false,
	}, nil
}

func StarExpr(ctx *context.Context, limit int) (*ast.StarExpr, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.StarExpr{
		// Star: token.NoPos,
		X: Expr(ctx, limit-1),
	}, nil
}

func TypeAssertExpr(ctx *context.Context, limit int) (*ast.TypeAssertExpr, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.TypeAssertExpr{
		// Lparen: token.NoPos,
		// Rparen: token.NoPos,
		X:    Expr(ctx, limit-1),
		Type: InterfaceType(ctx, limit-1),
	}, nil
}

func UnaryExpr(ctx *context.Context, limit int) (*ast.UnaryExpr, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.UnaryExpr{
		// OpPos: token.NoPos,
		X:  Expr(ctx, limit-1),
		Op: *pick.Pick(tokens.AcceptedByUnaryExpr),
	}, nil
}
