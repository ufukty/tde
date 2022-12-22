package caast

import (
	"go/ast"
	"tde/internal/utilities"

	"golang.org/x/tools/go/ast/astutil"
)

type CAAST struct{}

func (c *CAAST) Develop() {

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

	parentNode := *utilities.Pick(utilities.FilterSubNodes(body, func(n ast.Node) bool {
		if _, ok := n.(*ast.BlockStmt); ok {
			return true
		}
		return false
	}))
	siblingNode := *utilities.Pick(utilities.ChildNodes(parentNode))
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
