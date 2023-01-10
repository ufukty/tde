package astcfg

import (
	"tde/internal/ast_wrapper"
	"tde/internal/utilities"

	"go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

type ASTCFG struct{}

func (c *ASTCFG) Develop(funcDeclaration *ast.FuncDecl) {
	// TODO: Pick a valid insertion point: []BlockStatement etc.
	// TODO: Pick compatible node class for child node according to insertion point's type: Statement, Expression etc.
	// TODO: Pick a node type from node class: IfStatement, ForStatement etc.
	// TODO: Initialize an instance with the node type: IfStatement{}
	// TODO: Append the instance to insertion point
}

func (c *ASTCFG) PickRandomCrossOverPoint() {}

func (c *ASTCFG) PickCompatibleCrossOverPoint(nodeTypeClass NodeTypeClass) {

}

func (c *ASTCFG) PickCutPoint() ast.Node {
	return &ast.BadExpr{}
}

func GenerateRandomSubtree(nodeTypeClass NodeTypeClass) ast.Node {
	var (
		node ast.Node
		// isGenerated = false
	)
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
