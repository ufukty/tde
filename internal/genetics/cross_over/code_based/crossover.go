package code_based

import (
	ast_utl "tde/internal/astw/utilities"
	"tde/internal/utilities"

	"go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func CrossOver(parentA, parentB *ast.FuncDecl) bool {

	var (
		nodesA   = ast_utl.ListSubtree(parentA)[1:]
		nodesB   = ast_utl.ListSubtree(parentB)[1:]
		selected = false
		subA     ast.Node
		subB     ast.Node
		// supA     ast.Node
		// supB     ast.Node
		// indA     int
		// indB int
	)

	for !selected {
		subA = *utilities.Pick(nodesA)
		subB = *utilities.Pick(nodesB)
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
