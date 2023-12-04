package genetics

import "go/ast"

type cursor struct {
	node   field // always a pointer. could be a ast.Node or a slice them.
	parent ast.Node
	field  int // field index in parent
}

func inspectrec(n ast.Node, csr *cursor, c func(*cursor) bool) {
	for i, child := range children(n) {
		csr.node = child
		csr.field = i
		csr.parent = n
		if c(csr) {
			if child, ok := child.ptr.(ast.Node); ok {
				inspectrec(child, csr, c)
			}
		}
	}
}

func inspect(n ast.Node, c func(c *cursor) bool) {
	inspectrec(n, &cursor{}, c)
}
