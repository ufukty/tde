package test_package

import (
	utl "tde/internal/utilities"

	"go/ast"
)

// Sends traces of parents and childIndices to callback for each node.
//
// NOTE:
//   - Won't visit nil valued nodes. If visiting nil valued nodes is necessary
//     then use WalkWithNils() instead.
//   - Indices are not stable for Package fields because maps are not ordered.
func InspectWithTrace(node ast.Node, callback func(node ast.Node, parents []ast.Node, indices []int) bool) {
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
	ast.Inspect(node, func(currentNode ast.Node) bool {
		var recurse = false
		if currentNode != nil {
			recurse = callback(currentNode, parents, indices)
			if recurse {
				updateParents(currentNode)
				updateIndices(currentNode)
			} else {
				updateLastIndex()
			}
		} else {
			updateParents(currentNode)
			updateIndices(currentNode)
			callback(currentNode, parents, indices)
			updateLastIndex()
		}
		return recurse
	})
}

// The Pre function will be called before subnodes are visited and the Post
// function will be called for nodes after it is called for all of the subnodes.
//
// NOTE:
//   - Won't visit nil valued nodes. If visiting nil valued nodes is necessary
//     then use WalkWithNils() instead.
//   - Indices are not stable for Package fields because maps are not ordered.
func InspectTwiceWithTrace(
	node ast.Node,
	pre func(node ast.Node, parents []ast.Node, indices []int) bool,
	post func(node ast.Node, parents []ast.Node, indices []int),
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
	ast.Inspect(node, func(currentNode ast.Node) bool {
		var recurse = true
		if currentNode != nil {
			if pre != nil {
				recurse = pre(currentNode, parents, indices)
			}
			if recurse {
				updateParents(currentNode)
				updateIndices(currentNode)
			} else {
				updateLastIndex()
			}
		} else {
			ntemp := utl.SliceLast(parents)
			updateParents(currentNode)
			updateIndices(currentNode)
			if post != nil {
				post(ntemp, parents, indices)
			}
			updateLastIndex()
		}
		return recurse
	})
}

// Calls the callback for every child of the node, ignores itself of the node
// and all of the nodes deeper than 1 level.
//
// NOTE:
//   - Won't visit nil valued nodes. If visiting nil valued nodes is necessary
//     then use WalkWithNils() instead.
//   - Indices are not stable for Package fields because maps are not ordered.
func InspectChildren(node ast.Node, callback func(node ast.Node, indices int)) {
	InspectWithTrace(node, func(currentNode ast.Node, parents []ast.Node, indices []int) bool {
		switch len(parents) {
		case 0:
			return true
		case 1:
			callback(currentNode, indices[len(indices)-1])
			return false
		default:
			return false
		}
	})
}

// Calls the callback for every child of the node, ignores itself of the node and
// all of the nodes deeper than 1 level.
//
// NOTE:
//   - Won't visit nil valued nodes. If visiting nil valued nodes is necessary
//     then use WalkWithNils() instead.
//   - Indices are not stable for Package fields because maps are not ordered.
func InspectChildrenTwice(node ast.Node, pre func(node ast.Node, indices int), post func(node ast.Node, indices int)) {
	InspectTwiceWithTrace(node,
		func(currentNode ast.Node, parents []ast.Node, indices []int) bool {
			switch len(parents) {
			case 0:
				return true
			case 1:
				if pre != nil {
					pre(currentNode, indices[len(indices)-1])
				}
				return false
			default:
				return false
			}
		}, func(currentNode ast.Node, parents []ast.Node, indices []int) {
			if len(parents) == 1 {
				if post != nil {
					post(currentNode, indices[len(indices)-1])
				}
			}
		},
	)
}
