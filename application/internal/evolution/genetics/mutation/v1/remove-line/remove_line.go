package remove_line

import (
	"go/ast"
	"tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/utilities"
)

func list(n ast.Node) (stmtSliceContainingSubnodes []ast.Node) {
	ast.Inspect(n, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.BlockStmt:
			if len(n.List) > 0 {
				stmtSliceContainingSubnodes = append(stmtSliceContainingSubnodes, n)
			}
		case *ast.CommClause:
			if len(n.Body) > 0 {
				stmtSliceContainingSubnodes = append(stmtSliceContainingSubnodes, n)
			}
		case *ast.CaseClause:
			if len(n.Body) > 0 {
				stmtSliceContainingSubnodes = append(stmtSliceContainingSubnodes, n)
			}
		}

		return true
	})
	return
}

func removeOneLine(n ast.Node) {
	switch n := n.(type) {
	case *ast.BlockStmt:
		cutPoint := utilities.URandIntN(len(n.List))
		n.List = append(n.List[:cutPoint], n.List[cutPoint+1:]...)
	case *ast.CommClause:
		cutPoint := utilities.URandIntN(len(n.Body))
		n.Body = append(n.Body[:cutPoint], n.Body[cutPoint+1:]...)
	case *ast.CaseClause:
		cutPoint := utilities.URandIntN(len(n.Body))
		n.Body = append(n.Body[:cutPoint], n.Body[cutPoint+1:]...)
	}
}

func RemoveLine(n ast.Node) (ok bool) {
	stmtSliceContainingSubnodes := list(n)
	if len(stmtSliceContainingSubnodes) == 0 {
		return false
	}
	choosenNode := *utilities.Pick(stmtSliceContainingSubnodes)
	removeOneLine(choosenNode)
	return true
}

func GeneticOperation(ctx *models.GeneticOperationContext) bool {
	return RemoveLine(ctx.FuncDecl.Body)
}
