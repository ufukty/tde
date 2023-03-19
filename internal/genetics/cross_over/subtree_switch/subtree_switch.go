package subtree_switch

import (
	"go/ast"
	"tde/internal/astw/traced"
	"tde/internal/astw/traverse"
	"tde/internal/genetics/mutation/common"
	"tde/internal/utilities"

	"golang.org/x/tools/go/ast/astutil"
)

type ReverseTreeNode struct {
	Node   ast.Node
	Parent ast.Node
	Index  int
}

func listSubtreeWithParents(root ast.Node) (list []ReverseTreeNode) {
	traced.InspectWithTrace(root, func(node ast.Node, parents []ast.Node, indices []int) bool {
		list = append(list, ReverseTreeNode{
			Node:   node,
			Parent: parents[len(parents)-1],
		})
		return true
	})
	return
}

func listNonNilTraversableNodes(root ast.Node) {
	_ = traverse.GetTraversableNodeForASTNode(root)
}

func ModernizedCrossOver(offspring1, offspring2 *ast.FuncDecl) (ok bool) {
	// child1, child2 = clone.FuncDecl(parent1), clone.FuncDecl(parent2)
	subtree1, subtree2 := listSubtreeWithParents(offspring1.Body), listSubtreeWithParents(offspring2.Body)
	cutPoint1, cutPoint2 := *utilities.Pick(subtree1), *utilities.Pick(subtree2)

	ok1 := false
	cutPoint1.Parent = astutil.Apply(cutPoint1.Parent, func(c *astutil.Cursor) bool {
		if c.Node() == cutPoint1.Node {
			c.Replace(cutPoint2.Node)
			ok1 = true
		}
		return !ok1
	}, nil)

	ok2 := false
	cutPoint2.Parent = astutil.Apply(cutPoint2.Parent, func(c *astutil.Cursor) bool {
		// if c.Parent() == supB && ((c.Index() < 0 && c.Index() == indB) || (c.Index() == indB)) {
		if c.Node() == cutPoint2.Node {
			c.Replace(cutPoint1.Node)
			ok2 = true
		}
		return !ok2
	}, nil)

	return ok1 && ok2
}

func GeneticOperation(ctx *common.GeneticOperationContext) bool {

}
