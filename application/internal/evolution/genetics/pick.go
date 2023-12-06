package genetics

import (
	"fmt"
	"go/ast"
	"math/rand"
)

var (
	ErrEmptySlice = fmt.Errorf("empty slice")
)

func pickf(r ast.Node, f func(*cursor) bool) (cursor, error) {
	cs := []cursor{}
	inspect(r, func(c *cursor) bool {
		if f(c) {
			cs = append(cs, cursor{
				field:  c.field,
				parent: c.parent,
				fi:     c.fi,
			})
		}
		return true
	})
	if len(cs) == 0 {
		return cursor{}, ErrEmptySlice
	}
	return cs[rand.Intn(len(cs))], nil
}
