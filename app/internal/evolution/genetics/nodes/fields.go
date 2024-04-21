package nodes

import (
	"fmt"
	"go/ast"
)

func (c *Creator) Field(l int) (*ast.Field, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	n, err := c.Ident(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating Ident for Field.Names: %w", err)
	}
	t, err := c.Type(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating Field.Type: %w", err)
	}
	return &ast.Field{
		Names: []*ast.Ident{n},
		Type:  t,
		Tag:   nil,
	}, nil
}

func (c *Creator) FieldList(l int) (*ast.FieldList, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	f, err := c.Field(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating Field for FieldList.List: %w", err)
	}
	return &ast.FieldList{
		List: []*ast.Field{f},
	}, nil
}
