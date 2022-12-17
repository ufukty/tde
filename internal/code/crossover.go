package code

import (
	"fmt"
	"go/ast"
	"tde/internal/utilities"

	"golang.org/x/tools/go/ast/astutil"
)

func ListSubNodes(root ast.Node) []*ast.Node {
	nodes := []*ast.Node{}
	ast.Inspect(root, func(n ast.Node) bool {
		if n != nil && n != root {
			nodes = append(nodes, &n)
		}
		return true
	})
	return nodes
}

func FindParentNodeAndChildIndex(root ast.Node, child ast.Node) (parent ast.Node, childIndex int) {
	var (
		parentTrace     []ast.Node
		childIndexTrace []int
		found           = false
	)
	var (
		updateParentTrace = func(n ast.Node) {
			if n != nil {
				parentTrace = append(parentTrace, n)
			} else {
				parentTrace = parentTrace[:len(parentTrace)-1]
			}
		}
		updateChildIndexTrace = func(n ast.Node) {
			if n != nil {
				childIndexTrace = append(childIndexTrace, 0)
			} else {
				l := len(childIndexTrace)
				childIndexTrace = childIndexTrace[:l-1]
				childIndexTrace[l-2]++
			}
		}
	)
	ast.Inspect(root, func(n ast.Node) bool {
		if !found {
			updateParentTrace(n)
			updateChildIndexTrace(n)
		}

		if n != nil && n == child {
			found = true
		}

		return !found
	})

	if found {
		return parentTrace[len(parentTrace)-2], childIndexTrace[len(childIndexTrace)-2]
	}
	return nil, -1
}

func AreSameNodeType(l, r ast.Node) bool {
	switch l.(type) {
	case ast.Decl:
		if _, ok := r.(ast.Decl); ok {
			return true
		}
	case ast.Expr:
		if _, ok := r.(ast.Expr); ok {
			return true
		}
	case ast.Stmt:
		if _, ok := r.(ast.Stmt); ok {
			return true
		}
	case ast.Spec:
		if _, ok := r.(ast.Spec); ok {
			return true
		}
	}
	return false
}

func CrossOver(parentA, parentB *Function) bool {

	var (
		nodesA   = ListSubNodes(parentA.Root)
		nodesB   = ListSubNodes(parentB.Root)
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

	fmt.Println(&subA, &subB)

	// supA, _ = FindParentNodeAndChildIndex(parentA.Root, subA)
	// supB, indB = FindParentNodeAndChildIndex(parentB.Root, subB)

	replacedA := false
	parentA.Root.Body = astutil.Apply(parentA.Root.Body, func(c *astutil.Cursor) bool {
		if c.Node() == subA {
			fmt.Println("a")
			c.Replace(subB)
			replacedA = true
		}
		return !replacedA
	}, nil).(*ast.BlockStmt)

	replacedB := false
	parentB.Root.Body = astutil.Apply(parentB.Root.Body, func(c *astutil.Cursor) bool {
		// if c.Parent() == supB && ((c.Index() < 0 && c.Index() == indB) || (c.Index() == indB)) {
		if c.Node() == subB {
			fmt.Println("b")
			c.Replace(subA)
			replacedB = true
		}
		return !replacedB
	}, nil).(*ast.BlockStmt)

	return replacedA && replacedB
}
