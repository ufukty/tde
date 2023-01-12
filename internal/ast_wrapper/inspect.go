package ast_wrapper

import (
	"go/ast"
	utl "tde/internal/utilities"
)

func InspectWithTrace(startNode ast.Node, callback func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool) {
	var (
		parentTrace     []ast.Node
		childIndexTrace []int
	)
	var (
		updateParentTrace = func(n ast.Node) {
			if n != nil {
				parentTrace = append(parentTrace, n)
			} else {
				parentTrace = utl.SliceRemoveLast(parentTrace)
			}
		}
		updateChildIndexTraceItems = func(n ast.Node) {
			if n != nil {
				childIndexTrace = append(childIndexTrace, 0)
			} else {
				childIndexTrace = utl.SliceRemoveLast(childIndexTrace)
			}
		}
		updateChildIndexTraceLastCounter = func(n ast.Node) {
			if l := len(childIndexTrace); (n == nil) && (l > 0) {
				childIndexTrace[l-1]++
			}
		}
	)
	ast.Inspect(startNode, func(n ast.Node) bool {
		var ret bool
		if n != nil {
			ret = callback(n, parentTrace, childIndexTrace)
			updateParentTrace(n)
			updateChildIndexTraceItems(n)
		} else {
			updateParentTrace(n)
			updateChildIndexTraceItems(n)
			ret = callback(n, parentTrace, childIndexTrace)
			updateChildIndexTraceLastCounter(n)
		}
		return ret
	})
}

func InspectTwiceWithTrace(
	startNode ast.Node,
	pre func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool,
	post func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int),
) {
	var (
		parentTrace     []ast.Node
		childIndexTrace []int
	)
	var (
		updateParentTrace = func(n ast.Node) {
			if n != nil {
				parentTrace = append(parentTrace, n)
			} else {
				parentTrace = utl.SliceRemoveLast(parentTrace)
			}
		}
		updateChildIndexTraceItems = func(n ast.Node) {
			if n != nil {
				childIndexTrace = append(childIndexTrace, 0)
			} else {
				childIndexTrace = utl.SliceRemoveLast(childIndexTrace)
			}
		}
		updateChildIndexTraceLastCounter = func(n ast.Node) {
			if l := len(childIndexTrace); (n == nil) && (l > 0) {
				childIndexTrace[l-1]++
			}
		}
	)
	ast.Inspect(startNode, func(n ast.Node) bool {
		var ret = false
		if n != nil {
			ret = pre(n, parentTrace, childIndexTrace)
			updateParentTrace(n)
			updateChildIndexTraceItems(n)
		} else {
			ntemp := utl.SliceLast(parentTrace)
			updateParentTrace(n)
			updateChildIndexTraceItems(n)
			post(ntemp, parentTrace, childIndexTrace)
			updateChildIndexTraceLastCounter(n)
		}
		return ret

	})
}

// Different than Inspect function, ignores subnodes far than 1 depth.
func InspectChildren(startNode ast.Node, callback func(currentNode ast.Node, childIndex int)) {
	InspectWithTrace(startNode, func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		switch len(parentTrace) {
		case 0:
			return true
		case 1:
			callback(n, childIndexTrace[len(childIndexTrace)-1])
			return false
		default:
			return false
		}
	})
}

func FindNodeWithChildIndexTrace(childIndexTrace []int) ast.Node {
	return &ast.AssignStmt{}
}
