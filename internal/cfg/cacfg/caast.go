package cacfg

import (
	"go/ast"

	"tde/internal/ast_wrapper"
	"tde/internal/cfg/astcfg"
	"tde/internal/utilities"

	"golang.org/x/tools/go/ast/astutil"
)

type ContextAwareCFG struct{}

func (c *ContextAwareCFG) Develop() {
	// TODO: Pick a valid insertion point: []BlockStatement etc.
	// TODO: Trace back from insertion point up until to the root, to find declaration statements of variables, functions, imports etc.
	// TODO: Pick compatible node class for child node according to insertion point's type: Statement, Expression etc.
	// TODO: Pick a node type from node class: IfStatement, ForStatement, StartExpression, Literal, FunctionCall etc.
	// TODO: If a literal or ident has choosen; do either:
	//       TODO: create new entry for the new symbol (variable, function, import);
	//       TODO: pick appropriate value from the stack of declarations.
	// TODO: Initialize an instance with the node type: IfStatement{}
	// TODO: Append the instance to insertion point
}

func GenerateRandomSubtree(nodeTypeClass astcfg.NodeTypeClass) ast.Node {
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
				newStmt := GenerateRandomSubtree(astcfg.Statement)
				c.InsertBefore(newStmt)
				isInserted = true
			}
		}
		return !isInserted
	}, nil)
}

func NewVar(stack *Stack) (declaration *ast.DeclStmt, access *ast.Ident) {
	// TODO: Find available variable name (not in use currently)
	// TODO: Prepare declaration and access statements

	return
}
