package astwutl

import (
	"go/ast"
)

func children(root ast.Node) []ast.Node {
	list := []ast.Node{}
	ast.Inspect(root, func(n ast.Node) bool {
		if n == root {
			return true
		}
		list = append(list, n)
		return false
	})
	return list
}
