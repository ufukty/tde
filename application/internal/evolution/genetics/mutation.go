package genetics

import (
	"fmt"
	"go/ast"
	"math/rand"
	"slices"
	"tde/internal/astw/types"
	"tde/internal/evolution/genetics/mutation/v1/stg/nodes"
)

var (
	ErrNoPicks                   = fmt.Errorf("could not pick a node")
	ErrFieldAndValueTypeMismatch = fmt.Errorf("failed on asserting the 'any' typed value to the expected field type")
)

// Subtree generation
// TODO: capability to add nodes to slices
// TODO: pick nil holding parents
func Grow(fd *ast.FuncDecl) error {
	c, err := pickf(fd.Body, func(c *cursor) bool { return isNil(c.field) })
	if err != nil {
		return fmt.Errorf("pickf: %w", ErrNoPicks)
	}

	n := nodes.InType(c.field.expected, ctx, 1)
	if err := replaceOnParentWithCursor(c, n); err != nil {
		return fmt.Errorf("replacing the empty field: %w", err)
	}

	return nil
}

// TODO: populate the context with in-scope symbols (also DeleteLine, SwapLines)
func AddLine(fd *ast.FuncDecl) error {
	csr, err := pickf(fd.Body, func(c *cursor) bool { return c.field.expected == types.BlockStmt })
	if err != nil {
		return fmt.Errorf("pickf: %w", ErrNoPicks)
	}

	bs, ok := csr.field.ptr.(*ast.BlockStmt)
	if !ok {
		return ErrFieldAndValueTypeMismatch
	}

	l := nodes.Stmt(nil, 1)
	i := rand.Intn(len(bs.List) + 1)
	bs.List = slices.Insert(bs.List, i, l)

	return nil
}

func DeleteLine(fd *ast.FuncDecl) error {
	csr, err := pickf(fd.Body, func(c *cursor) bool { return c.field.expected == types.BlockStmt })
	if err != nil {
		return fmt.Errorf("pickf: %w", ErrNoPicks)
	}

	bs, ok := csr.field.ptr.(*ast.BlockStmt)
	if !ok {
		return ErrFieldAndValueTypeMismatch
	}

	l := nodes.Stmt(nil, 1)
	i := rand.Intn(len(bs.List) + 1)
	bs.List = slices.Insert(bs.List, i, l)

	return nil
}

func SwapLines(fd *ast.FuncDecl) error {
	csr, err := pickf(fd.Body, func(c *cursor) bool {
		if c.field.expected == types.BlockStmt {
			if bs, ok := c.field.ptr.(*ast.BlockStmt); ok {
				return len(bs.List) >= 2
			}
		}
		return false
	})
	if err != nil {
		return fmt.Errorf("pickf: %w", ErrNoPicks)
	}

	bs, ok := csr.field.ptr.(*ast.BlockStmt)
	if !ok {
		return ErrFieldAndValueTypeMismatch
	}

	i := rand.Intn(len(bs.List) + 1)
	var j int
	for i != j {
		j = rand.Intn(len(bs.List) + 1)
	}

	bs.List[i], bs.List[j] = bs.List[j], bs.List[i]

	return nil
}

func ShuffleLiterals(fd *ast.FuncDecl) error {
	c, err := pickf(fd.Body, func(c *cursor) bool { return c.field.expected == types.Ident })
	if err != nil {
		return fmt.Errorf("pickf: %w", ErrNoPicks)
	}

	n := nodes.InType(c.field.expected, ctx, 1)
	if err := replaceOnParentWithCursor(c, n); err != nil {
		return fmt.Errorf("replacing the empty field: %w", err)
	}

	return nil
}
