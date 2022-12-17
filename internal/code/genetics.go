package code

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
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

// func PickCrossOverPoint(fn Function, nodeClass NodeClass) (parent, child *ast.Node) {
// 	linearized := ListSubNodesThatConforms(fn.Root.Body)
// 	childRoot := **utilities.Pick(linearized)
// 	var (
// 		parentOfChild ast.Node
// 	)
// 	astutil.Apply(fn.Root.Body, func(c *astutil.Cursor) bool {
// 		if c.Node() == childRoot {

// 			c.Index()
// 			parentOfChild = c.Parent()
// 		}
// 		return parentOfChild == nil
// 	}, nil)
// 	return &childRoot, &parentOfChild
// }

func CrossOver(parentA, parentB *Function, fsetA, fsetB *token.FileSet) bool {
	// targetNodeClass := *utilities.Pick([]NodeClass{Statement, Expression})

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
		if reflect.TypeOf(subA) == reflect.TypeOf(subB) {
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

	fmt.Println("+++")
	// printer.Fprint(os.Stdout, fsetA, parentA.Root.Body)
	fmt.Println("+++")

	fmt.Println(&subA, &subB)
	// switch subA := subA.(type) {
	// case ast.Decl:
	// 	if subB, ok := subB.(ast.Decl); ok {
	// 		sub
	// 	}
	// case ast.Expr:
	// 	if subB, ok := subB.(ast.Expr); ok {
	// 		sub
	// 	}
	// case ast.Stmt:
	// 	if subB, ok := subB.(ast.Stmt); ok {
	// 		sub
	// 	}
	// case ast.Spec:
	// 	if subB, ok := subB.(ast.Spec); ok {
	// 		sub
	// 	}
	// }

	// switch targetNodeClass {
	// case Statement:
	// 	crsOvrParentA, okA := (*crsOvrParentA).(ast.Stmt)
	// 	crsOvrParentB, okB := (*crsOvrParentB).(ast.Stmt)
	// 	if okA && okB {
	// 		break

	// 	}
	// case Expression:
	// 	// crsOvrParentA, okA := (*crsOvrParentA).(ast.Expr)
	// 	// crsOvrParentB, okB := (*crsOvrParentB).(ast.Expr)
	// 	// if okA && okB {

	// 	// }
	// }

	return replacedA //&& replacedB
}
