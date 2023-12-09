package nodes

import (
	"fmt"
	"go/ast"
	"go/token"
	"tde/internal/utilities/pick"
)

// TODO: Store imported packages for later use
func (c *Creator) ImportSpec(l int) (*ast.ImportSpec, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Value, err := pick.Pick(AllowedPackagesToImport)
	if err != nil {
		return nil, fmt.Errorf("generating ImportSpec.Value: %w", err)
	}

	return &ast.ImportSpec{
		Name: nil,
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: Value,
		},
	}, nil
}

// rel. type keyword
func (c *Creator) TypeSpec(l int) (*ast.TypeSpec, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Name, err := c.Ident(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating Name: %w", err)
	}
	TypeParams, err := c.FieldList(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating TypeParams: %w", err)
	}

	return &ast.TypeSpec{
		Doc:        nil,
		Name:       Name,
		TypeParams: TypeParams,
		Type:       nil,
	}, nil
}

// Returns ValueSpec which is list of pairs of variable names and values to assign
func (c *Creator) ValueSpec(l int) (*ast.ValueSpec, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Value, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating ValueSpec.Value: %w", err)
	}

	return &ast.ValueSpec{
		Names:  []*ast.Ident{generateNewIdent()},
		Values: []ast.Expr{Value},
	}, nil
}
