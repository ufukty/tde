package nodes

import (
	"fmt"
	"go/ast"
	"tde/internal/evolution/genetics/nodes/tokens"
	"tde/internal/utilities/pick"
)

var basicLitGenerators = []func() *ast.BasicLit{
	basicIntegerLiteral,
	basicStringLiteral,
	basicFloatLiteral,
	basicCharacterLiteral,
}

func (c *Creator) BasicLit(l int) (*ast.BasicLit, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	generator, err := pick.Pick(basicLitGenerators)
	if err != nil {
		return nil, fmt.Errorf("picking amongst different type literal generators: %w", err)
	}
	return generator(), nil
}

func (c *Creator) BinaryExpr(l int) (*ast.BinaryExpr, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	op, err := pick.Pick(tokens.AcceptedByBinaryExpr)
	if err != nil {
		return nil, fmt.Errorf("generating BinaryExpr.Op: %w", err)
	}
	X, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating BinaryExpr.X: %w", err)
	}
	Y, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating BinaryExpr.Y: %w", err)
	}
	return &ast.BinaryExpr{
		// OpPos: token.NoPos,
		X:  X,
		Y:  Y,
		Op: op,
	}, nil
}

// TODO: function calls with more than 1 arguments
// TODO: variadic parameter support
func (c *Creator) CallExpr(l int) (*ast.CallExpr, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	fun, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating CallExpr.Fun: %w", err)
	}
	args, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating CallExpr.Args: %w", err)
	}
	return &ast.CallExpr{
		// Ellipsis: token.NoPos,
		Fun:  fun,
		Args: []ast.Expr{args},
	}, nil
}

func (c *Creator) CompositeLit(l int) (*ast.CompositeLit, error) {
	// TODO: check Incomplete property
	if l == 0 {
		return nil, ErrLimitReached
	}
	t, err := c.Type(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating CompositeLit.Type: %w", err)
	}
	elts, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating CompositeLit.Elts: %w", err)
	}

	return &ast.CompositeLit{
		// Lbrace:     token.NoPos,
		// Rbrace:     token.NoPos,
		Type:       t,
		Elts:       []ast.Expr{elts},
		Incomplete: false,
	}, nil
}

func (c *Creator) Ellipsis(l int) (*ast.Ellipsis, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	elt, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating Ellipsis.Elt: %w", err)
	}

	return &ast.Ellipsis{
		Elt: elt,
	}, nil
}

// TODO:
func (c *Creator) FuncLit(l int) (*ast.FuncLit, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	typ, err := c.FuncType(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating FuncLit.Type: %w", err)
	}
	bs, err := c.BlockStmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating FuncLit.BlockStmt: %w", err)
	}

	return &ast.FuncLit{
		Type: typ,
		Body: bs,
	}, nil
}

func (c *Creator) Ident(l int) (*ast.Ident, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	return generateNewIdent(), nil
}

func (c *Creator) IndexExpr(l int) (*ast.IndexExpr, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	X, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating IndexExpr.X: %w", err)
	}
	Index, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating IndexExpr.Index: %w", err)
	}

	return &ast.IndexExpr{
		X:     X,
		Index: Index,
	}, nil
}

// TODO: Multi-dimensional arrays
func (c *Creator) IndexListExpr(l int) (*ast.IndexListExpr, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	X, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating IndexListExpr.X: %w", err)
	}
	Indices, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating IndexListExpr.Indices: %w", err)
	}

	return &ast.IndexListExpr{
		X:       X,
		Indices: []ast.Expr{Indices},
	}, nil
}

func (c *Creator) KeyValueExpr(l int) (*ast.KeyValueExpr, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Key, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating KeyValueExpr.Key: %w", err)
	}
	Value, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating KeyValueExpr.Value: %w", err)
	}

	return &ast.KeyValueExpr{
		Key:   Key,
		Value: Value,
	}, nil
}

func (c *Creator) ParenExpr(l int) (*ast.ParenExpr, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	X, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating ParenExpr.X: %w", err)
	}

	return &ast.ParenExpr{
		X: X,
	}, nil
}

// FIXME: randomly produced X and Sel values will never work, maybe choose from
// imported libraries' exported functions, or previously declared struct instances that has methods
func (c *Creator) SelectorExpr(l int) (*ast.SelectorExpr, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	X, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating SelectorExpr.X: %w", err)
	}
	Sel, err := c.Ident(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating SelectorExpr.Sel: %w", err)
	}

	return &ast.SelectorExpr{
		X:   X,
		Sel: Sel,
	}, nil
}

func (c *Creator) SliceExpr(l int) (*ast.SliceExpr, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	X, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating SliceExpr.X: %w", err)
	}
	Low, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating SliceExpr.Low: %w", err)
	}
	High, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating SliceExpr.High: %w", err)
	}
	return &ast.SliceExpr{
		X:      X,
		Low:    Low,
		High:   High,
		Max:    nil,
		Slice3: false,
	}, nil
}

func (c *Creator) StarExpr(l int) (*ast.StarExpr, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	X, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating StarExpr.X: %w", err)
	}

	return &ast.StarExpr{
		X: X,
	}, nil
}

func (c *Creator) TypeAssertExpr(l int) (*ast.TypeAssertExpr, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	X, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating TypeAssertExpr.X: %w", err)
	}
	Type, err := c.InterfaceType(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating TypeAssertExpr.Type: %w", err)
	}

	return &ast.TypeAssertExpr{
		X:    X,
		Type: Type,
	}, nil
}

func (c *Creator) UnaryExpr(l int) (*ast.UnaryExpr, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	X, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating UnaryExpr.X: %w", err)
	}
	Op, err := pick.Pick(tokens.AcceptedByUnaryExpr)
	if err != nil {
		return nil, fmt.Errorf("generating UnaryExpr.Op: %w", err)
	}

	return &ast.UnaryExpr{
		X:  X,
		Op: Op,
	}, nil
}
