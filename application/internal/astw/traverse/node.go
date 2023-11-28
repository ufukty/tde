package traverse

import (
	"tde/internal/astw/astwutl"
	"tde/internal/astw/types"

	"go/ast"
)

type Spot int

const (
	Before = Spot(iota)
	Item
	After
)

// Use to represent nullable, Node-like (ast.Node etc.) or Node-slice ([]ast.Node etc.) components of the AST.
type Node struct {
	ref
	IsNil       bool
	Expected    types.NodeType // Set by Node.Subnodes(). Can be used from Once() and Twice() when the Value is nil
	Constraints constraints
	Parent      *Node
}

func NewNode(node ast.Node) *Node {
	return &Node{
		ref:      newFieldRef(&node),
		Expected: types.TypeFor(node),
		IsNil:    astwutl.IsNodeNil(node),
		Parent:   nil,
	}
}

// One TraversableNode's TraversableSubnodes are
//  1. If the TraversableNode is Node-Slice: Items of the slice
//  2. If the TraversableNode is Node-like: Fields of the struct that are either Node-like or Node-slice
func (n *Node) Subnodes() []*Node {
	c := constraints{}
	if n.Expected.IsInterfaceType() {
		return subnodesForIface(n, c) // Use TraversableNodesFromConcreteTypeNode after checking actually assigned node's type
	} else if n.Expected.IsSliceType() {
		return subnodesForSliceField(n, c)
	} else if n.Expected.IsConcreteType() {
		return subnodesForConcrete(n, c)
	}
	panic("Failed on GetTraversableSubnodes. There should be any ExpectedTypes don't fall into any of above cases")
	// return []TraversableNode{}
}
