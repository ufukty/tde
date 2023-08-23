package switch_lines

import (
	"go/ast"
	"tde/internal/evolution/genetics/mutation/common"
	"tde/internal/utilities"
)

func list(n ast.Node) (stmtSliceContainingSubnodes []ast.Node) {
	ast.Inspect(n, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.BlockStmt:
			if len(n.List) > 1 {
				stmtSliceContainingSubnodes = append(stmtSliceContainingSubnodes, n)
			}
		case *ast.CommClause:
			if len(n.Body) > 1 {
				stmtSliceContainingSubnodes = append(stmtSliceContainingSubnodes, n)
			}
		case *ast.CaseClause:
			if len(n.Body) > 1 {
				stmtSliceContainingSubnodes = append(stmtSliceContainingSubnodes, n)
			}
		}

		return true
	})
	return
}

func swap(n ast.Node) {
	switch n := n.(type) {
	case *ast.BlockStmt:
		cutPoint := utilities.URandIntN(len(n.List) - 1)
		n.List[cutPoint], n.List[cutPoint+1] = n.List[cutPoint+1], n.List[cutPoint]
	case *ast.CommClause:
		cutPoint := utilities.URandIntN(len(n.Body) - 1)
		n.Body[cutPoint], n.Body[cutPoint+1] = n.Body[cutPoint+1], n.Body[cutPoint]
	case *ast.CaseClause:
		cutPoint := utilities.URandIntN(len(n.Body) - 1)
		n.Body[cutPoint], n.Body[cutPoint+1] = n.Body[cutPoint+1], n.Body[cutPoint]
	}
}

func SiblingSwap(n ast.Node) (ok bool) {
	stmtSliceContainingSubnodes := list(n)
	if len(stmtSliceContainingSubnodes) == 0 {
		return false
	}
	choosenNode := *utilities.Pick(stmtSliceContainingSubnodes)
	swap(choosenNode)
	return true
}

func GeneticOperation(ctx *common.GeneticOperationContext) bool {
	return SiblingSwap(ctx.FuncDecl.Body)
}
