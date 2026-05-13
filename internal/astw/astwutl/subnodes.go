package astwutl

import (
	"go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func ChildNodes(root ast.Node) (children []ast.Node) {
	list := []ast.Node{}
	astutil.Apply(root, func(c *astutil.Cursor) bool {
		if c.Node() == nil {
			return false
		}
		if c.Parent() == root {
			list = append(list, c.Node())
		}
		if c.Node() == root {
			return true
		} else {
			return false
		}
	}, nil)
	return list
}
