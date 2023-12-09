package nodes

import (
	"fmt"
	"go/ast"
	"go/token"
)

// TODO: Consider adding receiver functions (methods)
func (c *Creator) FuncDecl(l int) (*ast.FuncDecl, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Type, err := c.FuncType(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating FuncDecl.Type: %w", err)
	}
	Body, err := c.BlockStmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating FuncDecl.Body: %w", err)
	}

	return &ast.FuncDecl{
		Name: generateFunctionName(),
		Type: Type,
		Body: Body,
	}, nil
}

// FIXME: Produces only "variable" declarations. "import", "constant", "type" declarations are ignored.
func (c *Creator) GenDecl(l int) (*ast.GenDecl, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Specs, err := c.ValueSpec(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating GenDecl.Specs: %w", err)
	}

	return &ast.GenDecl{
		Tok:   token.VAR,
		Specs: []ast.Spec{Specs},
	}, nil
}
