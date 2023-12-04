package genetics

import (
	"go/ast"
)

// subtree linearization
func linearize(r ast.Node) (list []any) {
	inspect(r, func(c *cursor) bool {
		list = append(list, c.node)
		return true
	})
	return
}

// subtree linearization
func linearizefiltered[FILTER any](r ast.Node) (list []FILTER) {
	inspect(r, func(c *cursor) bool {
		if n, ok := (c.node.ptr).(FILTER); ok {
			list = append(list, n)
		}
		return true
	})
	return
}

// subtree linearization
func linearizeCursorsForNilValuedFields(r ast.Node) []cursor {
	cs := []cursor{}
	inspect(r, func(c *cursor) bool {
		if isNil(c.node) {
			cs = append(cs, cursor{
				node:   c.node,
				parent: c.parent,
				field:  c.field,
			})
		}
		return true
	})
	return cs
}

// use pointers with FILTER eg. `*ast.BlockStmt` or `*[]ast.Stmt`
func filter[FILTER any](r []any) (list []FILTER) {
	for _, v := range r {
		if v, ok := v.(FILTER); ok {
			list = append(list, v)
		}
	}
	return
}
