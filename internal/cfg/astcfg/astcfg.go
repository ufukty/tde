package astcfg

import (
	"tde/internal/ast_wrapper"
	"tde/internal/utilities"

	"go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func FindTypeOfFieldWithIndexTrace(indexTrace []int) NodeType {
	return BlockStmt
}

func MakeAddition(funcDeclaration *ast.FuncDecl) {
	// TODO: Pick a valid insertion point: []BlockStatement etc.
	// TODO: Pick compatible node class for child node according to insertion point's type: Statement, Expression etc.
	// TODO: Pick a node type from node class: IfStatement, ForStatement etc.
	// TODO: Initialize an instance with the node type: IfStatement{}
	// TODO: Append the instance to insertion point

	type Candidate struct {
		parent     ast.Node
		childIndex int
	}
	candidates := []Candidate{}
	ast_wrapper.WalkWithNils(funcDeclaration, func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		if n == nil {
			candidates = append(candidates, Candidate{
				parentTrace[len(parentTrace)-1], 
				childIndexTrace[len(childIndexTrace)-1],
			})
		}
		return true
	})
	choosenCandidate := utilities.Pick(candidates)
	ast_wrapper.WalkWithNils(choosenCandidate.parent, func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		if len(childIndexTrace) == 0 {
			return true
		}
		if childIndexTrace[0] == choosenCandidate.childIndex {
			if n == nil {
				switch n.(type) {
				case ast.Node:

				}
			}

		}
		return false
	})

}

func PickRandomCrossOverPoint() {}

func PickCompatibleCrossOverPoint(nodeTypeClass NodeTypeClass) {

}

func PickCutPoint() ast.Node {
	return &ast.BadExpr{}
}

func GenerateRandomSubtree(nodeTypeClass NodeTypeClass) ast.Node {
	var node ast.Node
	// var isGenerated = false

	// for !isGenerated {
	// 	node = GenerateInstance(utilities.Pick())
	// }
	return node
}

// It picks random BlockStmt amongst all subnodes, generates new Statement
func NewLine(f *ast.FuncDecl) {
	body := f.Body

	parentNode := *utilities.Pick(ast_wrapper.FilterSubNodes(body, func(n ast.Node) bool {
		if _, ok := n.(*ast.BlockStmt); ok {
			return true
		}
		return false
	}))
	siblingNode := *utilities.Pick(ast_wrapper.ChildNodes(parentNode))
	isInserted := false
	astutil.Apply(parentNode, func(c *astutil.Cursor) bool {
		if !isInserted {
			if c.Node() == siblingNode {
				newStmt := GenerateRandomSubtree(Statement)
				c.InsertBefore(newStmt)
				isInserted = true
			}
		}
		return !isInserted
	}, nil)
}
