package utilities

import (
	"go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func ListSubtree(root ast.Node) (subnodes []ast.Node) {
	ast.Inspect(root, func(n ast.Node) bool {
		if n != nil {
			subnodes = append(subnodes, n)
		}
		return true
	})
	return
}

func ListSubtreeWithFilter(root ast.Node, filter func(n ast.Node) bool) (filteredNodes []ast.Node) {
	list := []ast.Node{}
	ast.Inspect(root, func(n ast.Node) bool {
		if filter(n) {
			list = append(list, n)
		}
		return true
	})
	return list
}

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
