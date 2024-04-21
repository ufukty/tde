package genetics

import "go/ast"

type cursor struct {
	field  field // always a pointer. could be a ast.Node or a slice them.
	parent ast.Node
	fi     int // field index in parent
}

func inspectrec(n ast.Node, csr *cursor, c func(*cursor) bool) {
	for i, child := range children(n) {
		csr.field = child
		csr.fi = i
		csr.parent = n
		if c(csr) {
			if child, ok := child.ptr.(ast.Node); ok {
				inspectrec(child, csr, c)
			}
		}
	}
}

func inspect(n ast.Node, c func(c *cursor) bool) {
	csr := &cursor{
		field: field{
			ptr:      n,
			expected: typeOf(n),
		},
		parent: nil,
		fi:     -1,
	}
	c(csr)
	inspectrec(n, csr, c)
}
