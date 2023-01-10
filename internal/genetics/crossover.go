package genetics

import (
	"go/ast"
	"tde/internal/ast_wrapper"
	"tde/internal/utilities"

	"golang.org/x/tools/go/ast/astutil"
)

func CrossOver(parentA, parentB *ast.FuncDecl) bool {

	var (
		nodesA   = ast_wrapper.ListSubnodes(parentA)[1:]
		nodesB   = ast_wrapper.ListSubnodes(parentB)[1:]
		selected = false
		subA     ast.Node
		subB     ast.Node
		// supA     ast.Node
		// supB     ast.Node
		// indA     int
		// indB int
	)

	for !selected {
		subA = **utilities.Pick(nodesA)
		subB = **utilities.Pick(nodesB)
		if AreSameNodeType(subA, subB) {
			selected = true
		}
	}

	// supA, _ = FindParentNodeAndChildIndex(parentA, subA)
	// supB, indB = FindParentNodeAndChildIndex(parentB, subB)

	replacedA := false
	parentA.Body = astutil.Apply(parentA.Body, func(c *astutil.Cursor) bool {
		if c.Node() == subA {
			c.Replace(subB)
			replacedA = true
		}
		return !replacedA
	}, nil).(*ast.BlockStmt)

	replacedB := false
	parentB.Body = astutil.Apply(parentB.Body, func(c *astutil.Cursor) bool {
		// if c.Parent() == supB && ((c.Index() < 0 && c.Index() == indB) || (c.Index() == indB)) {
		if c.Node() == subB {
			c.Replace(subA)
			replacedB = true
		}
		return !replacedB
	}, nil).(*ast.BlockStmt)

	return replacedA && replacedB
}
