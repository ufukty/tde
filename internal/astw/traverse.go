package astw

// One TraversableNode's TraversableSubnodes are
//  1. If the TraversableNode is Node-Slice: Items of the slice
//  2. If the TraversableNode is Node-like: Fields of the struct that are either Node-like or Node-slice
func (tNode *TraversableNode) GetTraversableSubnodes() []*TraversableNode {
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

func traverseHelper(tNodePtr *TraversableNode, callback func(tNodePtr *TraversableNode) bool) {
	if cont := callback(tNodePtr); !cont {
		return
	}
	for _, field := range tNodePtr.GetTraversableSubnodes() {
		traverseHelper(field, callback)
	}
}

// There are 2 differences compared to WalkWithNil:
//  1. Sends the information of expected type for nil values
//  2. Threats slices are individual nodes, their items will have isolated indices from sibling nodes.
func Traverse(tNode TraversableNode, callback func(tNodePtr *TraversableNode) bool) {
	traverseHelper(&tNode, callback)
}
