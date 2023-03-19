package literals

import (
	"go/ast"
	"go/token"
	"tde/internal/genetics/common"
	"tde/internal/utilities"
)

func listApplicableNodes(n ast.Node) (applicableNodes []*ast.BasicLit) {
	ast.Inspect(n, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		switch n := n.(type) {
		case *ast.BasicLit:
			applicableNodes = append(applicableNodes, n)
		}
		return true
	})
	return
}

func mutateStringLiteral(s string) string {
	// partially regenerate
	// one char mutate
	return s
}

func mutateLiteralInteger(i string) string {
	// random assignment

	// 10% + -
	return i
}

func perform(choosenNode *ast.BasicLit) {
	switch choosenNode.Kind {
	case token.STRING:
		choosenNode.Value = mutateStringLiteral(choosenNode.Value)
	case token.INT:
		choosenNode.Value = mutateLiteralInteger(choosenNode.Value)
	}
}

func GeneticOperation(ctx *common.GeneticOperationContext) bool {
	applicableNodes := listApplicableNodes(ctx.FuncDecl.Body)
	choosenNode := *utilities.Pick(applicableNodes)
	perform(choosenNode)
	return true
}
