package ast_wrapper

import "go/ast"

// One TraversableNode's TraversableSubnodes are
//  1. If the TraversableNode is Node-Slice: Items of the slice
//  2. If the TraversableNode is Node-like: Fields of the struct that are either Node-like or Node-slice
func (tNode *TraversableNode) GetTraversableSubnodes() []TraversableNode {
	if tNode.ExpectedType.IsInterfaceType() {
		return TraversableNodesFromInterfaceTypeNode(tNode) // Use TraversableNodesFromConcreteTypeNode after checking actually assigned node's type

	} else if tNode.ExpectedType.IsSliceType() {
		return TraversableNodesFromSliceTypeNode(tNode)

	} else if tNode.ExpectedType.IsConcreteType() {
		return TraversableNodesFromConcreteTypeNode(tNode)

	}
	panic("Failed on GetTraversableSubnodes. There should be any ExpectedTypes don't fall into any of above cases")
	// return []TraversableNode{}
}

type TraverseTraceItem struct {
	parent *TraversableNode
	index  int
}

func traverseHelper(parent TraversableNode, trace []TraverseTraceItem, callback func(node TraversableNode, trace []TraverseTraceItem) bool) {
	if cont := callback(parent, trace); !cont {
		return
	}
	for i, field := range parent.GetTraversableSubnodes() {
		traverseHelper(field, append(trace, TraverseTraceItem{&parent, i}), callback)
	}
}

// There are 2 differences compared to WalkWithNil:
//  1. Sends the information of expected type for nil values
//  2. Threats slices are individual nodes, their items will have isolated indices from sibling nodes.
func Traverse(n ast.Node, callback func(n TraversableNode, trace []TraverseTraceItem) bool) {
	traverseHelper(GetTraversableNodeForASTNode(n), []TraverseTraceItem{}, callback)
}
