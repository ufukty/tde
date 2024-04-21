package genetics

import (
	"fmt"
	"go/ast"
	"slices"
	"tde/internal/astw/types"
	"tde/internal/utilities/pick"
)

func pickInTypes(fd *ast.FuncDecl, ts []types.NodeType) (cursor, error) {
	cs := []cursor{}
	inspect(fd.Body, func(c *cursor) bool {
		if slices.Index(ts, c.field.expected) != -1 {
			cs = append(cs, cursor{
				field:  c.field,
				parent: c.parent,
				fi:     c.fi,
			})
		}
		return true
	})

	if c, err := pick.Pick(cs); err != nil {
		return c, fmt.Errorf("picking the cursor: %w", err)
	} else {
		return c, nil
	}
}

func Crossover(fd1, fd2 *ast.FuncDecl) error {
	mutuals, err := mutualFieldTypes(fd1, fd2)
	if err != nil {
		return fmt.Errorf("mutualtypes: %w", err)
	}

	c1, err := pickInTypes(fd1, mutuals)
	if err != nil {
		return fmt.Errorf("pickInTypes 1: %w", err)
	}

	c2, err := pickInTypes(fd2, mutuals)
	if err != nil {
		return fmt.Errorf("pickInTypes 2: %w", err)
	}

	// swap
	replaceOnParent(c1.parent, c1.field.ptr, c2.field.ptr)
	replaceOnParent(c2.parent, c2.field.ptr, c1.field.ptr)

	return nil
}
