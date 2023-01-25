package astw

import (
	"go/ast"
	utl "tde/internal/utilities"
)

// Simple ast.Inspect() wrapper.
//
// Additionally, sends traces of parents and childIndices to callback for each node.
//
// If visiting nil valued nodes is necessary then use WalkWithNils() instead.
// FIXME: Change value parameters with pointers; to let callback compare currentNode based on address instead of values
func InspectWithTrace(start ast.Node, callback func(node ast.Node, parents []ast.Node, indices []int) bool) {
	var (
		parents []ast.Node
		indices []int
	)
	var (
		updateParents = func(n ast.Node) {
			if n != nil {
				parents = append(parents, n)
			} else {
				parents = utl.SliceRemoveLast(parents)
			}
		}
		updateIndices = func(n ast.Node) {
			if n != nil {
				indices = append(indices, 0)
			} else {
				indices = utl.SliceRemoveLast(indices)
			}
		}
		updateLastIndex = func() {
			if l := len(indices); l > 0 {
				indices[l-1]++
			}
		}
	)
	ast.Inspect(start, func(n ast.Node) bool {
		var recurse = false
		if n != nil {
			recurse = callback(n, parents, indices)
			if recurse {
				updateParents(n)
				updateIndices(n)
			} else {
				updateLastIndex()
			}
		} else {
			updateParents(n)
			updateIndices(n)
			callback(n, parents, indices)
			updateLastIndex()
		}
		return recurse
	})
}

// The Pre function will be called before subnodes are visited and the Post function will be called for nodes after it is called for all of the subnodes.
//
// Won't visit nil valued nodes.
func InspectTwiceWithTrace(
	startNode ast.Node,
	pre func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool,
	post func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int),
) {
	var (
		parents []ast.Node
		indices []int
	)
	var (
		updateParents = func(n ast.Node) {
			if n != nil {
				parents = append(parents, n)
			} else {
				parents = utl.SliceRemoveLast(parents)
			}
		}
		updateIndices = func(n ast.Node) {
			if n != nil {
				indices = append(indices, 0)
			} else {
				indices = utl.SliceRemoveLast(indices)
			}
		}
		updateLastIndex = func() {
			if l := len(indices); l > 0 {
				indices[l-1]++
			}
		}
	)
	ast.Inspect(startNode, func(n ast.Node) bool {
		var recurse = true
		if n != nil {
			if pre != nil {
				recurse = pre(n, parents, indices)
			}
			if recurse {
				updateParents(n)
				updateIndices(n)
			} else {
				updateLastIndex()
			}
		} else {
			ntemp := utl.SliceLast(parents)
			updateParents(n)
			updateIndices(n)
			if post != nil {
				post(ntemp, parents, indices)
			}
			updateLastIndex()
		}
		return recurse

	})
}

// Calls the callback for every child of the node, ignores itself of the node and all of the nodes deeper than 1 level.
func InspectChildren(node ast.Node, callback func(node ast.Node, indices int)) {
	InspectWithTrace(node, func(n ast.Node, parents []ast.Node, indices []int) bool {
		switch len(parents) {
		case 0:
			return true
		case 1:
			callback(n, indices[len(indices)-1])
			return false
		default:
			return false
		}
	})
}
