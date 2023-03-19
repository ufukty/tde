package code_based

import "go/ast"

func FindParentNodeAndChildIndex(root ast.Node, child ast.Node) (parent ast.Node, childIndex int) {
	var (
		parentTrace     []ast.Node
		childIndexTrace []int
		found           = false
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
	ast.Inspect(root, func(n ast.Node) bool {
		if !found {
			updateParentTrace(n)
			updateChildIndexTrace(n)
		}

		if n != nil && n == child {
			found = true
		}

		return !found
	})

	if found {
		return parentTrace[len(parentTrace)-2], childIndexTrace[len(childIndexTrace)-2]
	}
	return nil, -1
}

func AreSameNodeType(l, r ast.Node) bool {
	switch l.(type) {
	case ast.Decl:
		if _, ok := r.(ast.Decl); ok {
			return true
		}
	case ast.Expr:
		if _, ok := r.(ast.Expr); ok {
			return true
		}
	case ast.Stmt:
		if _, ok := r.(ast.Stmt); ok {
			return true
		}
	case ast.Spec:
		if _, ok := r.(ast.Spec); ok {
			return true
		}
	}
	return false
}
