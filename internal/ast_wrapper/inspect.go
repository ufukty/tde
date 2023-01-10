package ast_wrapper

import "go/ast"

func Inspect(node ast.Node, callback func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool) {
	var (
		parentTrace     []ast.Node
		childIndexTrace []int
	)
	var (
		updateParentTrace = func(n ast.Node) {
			if n != nil {
				parentTrace = append(parentTrace, n)
			} else {
				parentTrace = parentTrace[:len(parentTrace)-1]
			}
		}
		updateChildIndexTrace = func(n ast.Node) {
			if n != nil {
				childIndexTrace = append(childIndexTrace, 0)
			} else {
				l := len(childIndexTrace)
				childIndexTrace = childIndexTrace[:l-1]
				childIndexTrace[l-2]++
			}
		}
	)
	ast.Inspect(node, func(n ast.Node) bool {
		updateParentTrace(n)
		updateChildIndexTrace(n)
		return callback(n, parentTrace[:len(parentTrace)-1], childIndexTrace[:len(childIndexTrace)-1])
	})
}
