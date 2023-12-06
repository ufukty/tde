package subtreeswitch

import (
	"fmt"
	"go/ast"
	"tde/internal/astw/traced"
	"tde/internal/utilities/pick"

	"golang.org/x/tools/go/ast/astutil"
)

type reverseTreeNode struct {
	Node   ast.Node
	Parent ast.Node
	Index  int
}

func listSubtreeWithParents(root ast.Node) (stmts []reverseTreeNode, exprs []reverseTreeNode) {
	traced.InspectWithTrace(root, func(node ast.Node, parents []ast.Node, indices []int) bool {
		if node == root || node == nil {
			return true
		}
		switch node := node.(type) {
		case ast.Stmt:
			stmts = append(stmts, reverseTreeNode{
				Node:   node,
				Parent: parents[len(parents)-1],
			})
		case ast.Expr:
			stmts = append(stmts, reverseTreeNode{
				Node:   node,
				Parent: parents[len(parents)-1],
			})
		}
		return true
	})
	return
}

type cutPointType int

const (
	ctpStmt = cutPointType(iota)
	ctpExpr
)

func pickCutPoints(offspring1, offspring2 *ast.FuncDecl) (cutPoint1, cutPoint2 *reverseTreeNode, err error) {
	subtree1stmts, subtree1exprs := listSubtreeWithParents(offspring1.Body)
	subtree2stmts, subtree2exprs := listSubtreeWithParents(offspring2.Body)

	availableCutPointTypes := []cutPointType{}
	if len(subtree1stmts) > 0 && len(subtree2stmts) > 0 {
		availableCutPointTypes = append(availableCutPointTypes, ctpStmt)
	}
	if len(subtree1exprs) > 0 && len(subtree2exprs) > 0 {
		availableCutPointTypes = append(availableCutPointTypes, ctpExpr)
	}

	choosenCutPointType, err := pick.Pick(availableCutPointTypes)
	if err != nil {
		return nil, nil, fmt.Errorf("choosing cut point type: %w", err)
	}

	var cp1, cp2 reverseTreeNode
	switch choosenCutPointType {
	case ctpStmt:
		cp1, err = pick.Pick(subtree1stmts)
		if err != nil {
			return nil, nil, fmt.Errorf("picking cut point 1: %w", err)
		}
		cp2, err = pick.Pick(subtree2stmts)
		if err != nil {
			return nil, nil, fmt.Errorf("picking cut point 2: %w", err)
		}

	case ctpExpr:
		cp1, err := pick.Pick(subtree1exprs)
		if err != nil {
			return nil, nil, fmt.Errorf("picking cut point 1: %w", err)
		}
		cutPoint1 = &cp1
		cp2, err = pick.Pick(subtree2exprs)
		if err != nil {
			return nil, nil, fmt.Errorf("picking cut point 2: %w", err)
		}
	}
	return &cp1, &cp2, nil
}

func attempt(offspring1, offspring2 *ast.FuncDecl) (err error) {
	cutPoint1, cutPoint2, err := pickCutPoints(offspring1, offspring2)
	if err != nil {
		return fmt.Errorf("picking cut points: %w", err)
	}

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

	return nil
}

func SubtreeSwitch(offspring1, offspring2 *ast.FuncDecl) (ok bool) {
	for counter := 0; counter < 50; counter++ {
		// fmt.Println("run:", counter)
		status := func() (status bool) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
					status = false
				} else {
					status = true
				}
			}()
			attempt(offspring1, offspring2)
			return
		}()
		if status {
			return true
		}
	}
	return false
}
