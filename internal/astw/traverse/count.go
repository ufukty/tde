package traverse

import "go/ast"

func CountNonNilNodes(node ast.Node) int {
	counter := 0
	Traverse(GetTraversableNodeForASTNode(node), func(tNodePtr *TraversableNode) bool {
		if node != nil {
			counter++
		}
		return true
	})
	return counter
}
