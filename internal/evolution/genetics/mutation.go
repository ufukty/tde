package genetics

import (
	"fmt"
	"go/ast"
	"math/rand"
	"slices"
	"tde/internal/astw/types"
	"tde/internal/evolution/genetics/nodes"
)

var (
	ErrNoPicks                   = fmt.Errorf("could not pick a node")
	ErrFieldAndValueTypeMismatch = fmt.Errorf("failed on asserting the 'any' typed value to the expected field type")
)

// Subtree generation
// TODO: capability to add nodes to slices
func Grow(nc *nodes.Creator, fd *ast.FuncDecl) error {
	c, err := pickf(fd.Body, func(c *cursor) bool { return isNil(c.field) })
	if err != nil {
		return fmt.Errorf("pickf: %w", ErrNoPicks)
	}
	n, err := nc.InType(c.field.expected, 1)
	if err != nil {
		return fmt.Errorf("generating a node in type %s: %w", c.field.expected, err)
	}
	if err := replaceOnParentWithCursor(c, n); err != nil {
		return fmt.Errorf("replacing the empty field: %w", err)
	}
	return nil
}

func Regenerate(nc *nodes.Creator, fd *ast.FuncDecl) error {
	c, err := pickf(fd.Body, func(c *cursor) bool { return !isNil(c.field) })
	if err != nil {
		return fmt.Errorf("pickf: %w", ErrNoPicks)
	}
	if err := regenerateField(nc, c); err != nil {
		return fmt.Errorf("replacing the empty field: %w", err)
	}
	return nil
}

// TODO: populate the context with in-scope symbols (also DeleteLine, SwapLines)
func AddLine(nc *nodes.Creator, fd *ast.FuncDecl) error {
	csr, err := pickf(fd.Body, func(c *cursor) bool { return c.field.expected == types.BlockStmt })
	if err != nil {
		return fmt.Errorf("pickf: %w", ErrNoPicks)
	}

	bs, ok := csr.field.ptr.(*ast.BlockStmt)
	if !ok {
		return ErrFieldAndValueTypeMismatch
	}

	i := rand.Intn(len(bs.List) + 1)

	l, err := nc.Stmt(1)
	if err != nil {
		return fmt.Errorf("generating line: %w", err)
	}

	bs.List = slices.Insert(bs.List, i, l)

	return nil
}

// FIXME:
func DeleteLine(fd *ast.FuncDecl) error {
	csr, err := pickf(fd.Body, func(c *cursor) bool { return c.field.expected == types.BlockStmt })
	if err != nil {
		return fmt.Errorf("pickf: %w", ErrNoPicks)
	}

	bs, ok := csr.field.ptr.(*ast.BlockStmt)
	if !ok {
		return ErrFieldAndValueTypeMismatch
	}

	i := rand.Intn(len(bs.List) + 1)
	bs.List = append(bs.List[:i], bs.List[i+1:]...)

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

// TODO:
// func ShuffleLiterals(nc *nodes.Creator, fd *ast.FuncDecl) error {
// 	c, err := pickf(fd.Body, func(c *cursor) bool { return c.field.expected == types.Ident })
// 	if err != nil {
// 		return fmt.Errorf("pickf: %w", ErrNoPicks)
// 	}

// 	return nil
// }
