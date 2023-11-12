package lines

import (
	"go/ast"
	"tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/utilities"
)

func SwapLines(ctx *models.MutationParameters) (ok bool) {
	blockStmts := listBlockStmts(ctx.FuncDecl.Body, 2)
	if len(blockStmts) == 0 {
		return false
	}
	choosenNode := *utilities.Pick(blockStmts)

	switch choosenNode := choosenNode.(type) {
	case *ast.BlockStmt:
		cutPoint := utilities.URandIntN(len(choosenNode.List) - 1)
		choosenNode.List[cutPoint], choosenNode.List[cutPoint+1] = choosenNode.List[cutPoint+1], choosenNode.List[cutPoint]
	case *ast.CommClause:
		cutPoint := utilities.URandIntN(len(choosenNode.Body) - 1)
		choosenNode.Body[cutPoint], choosenNode.Body[cutPoint+1] = choosenNode.Body[cutPoint+1], choosenNode.Body[cutPoint]
	case *ast.CaseClause:
		cutPoint := utilities.URandIntN(len(choosenNode.Body) - 1)
		choosenNode.Body[cutPoint], choosenNode.Body[cutPoint+1] = choosenNode.Body[cutPoint+1], choosenNode.Body[cutPoint]
	}
	return true
}
