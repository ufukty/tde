package lines

import (
	"go/ast"
	"tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/utilities"
)

func listBlockStmts(n ast.Node, subnodes int) (blockStmts []ast.Node) {
	ast.Inspect(n, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.BlockStmt:
			if len(n.List) >= subnodes {
				blockStmts = append(blockStmts, n)
			}
		case *ast.CommClause:
			if len(n.Body) >= subnodes {
				blockStmts = append(blockStmts, n)
			}
		case *ast.CaseClause:
			if len(n.Body) >= subnodes {
				blockStmts = append(blockStmts, n)
			}
		}

		return true
	})
	return
}

func RemoveLine(ctx *models.MutationParameters) error {
	blockstmts := listBlockStmts(ctx.FuncDecl.Body, 1)
	if len(blockstmts) == 0 {
		return models.ErrUnsupportedMutation
	}
	choosenNode := *utilities.Pick(blockstmts)
	switch choosenNode := choosenNode.(type) {
	case *ast.BlockStmt:
		cutPoint := utilities.URandIntN(len(choosenNode.List))
		choosenNode.List = append(choosenNode.List[:cutPoint], choosenNode.List[cutPoint+1:]...)
	case *ast.CommClause:
		cutPoint := utilities.URandIntN(len(choosenNode.Body))
		choosenNode.Body = append(choosenNode.Body[:cutPoint], choosenNode.Body[cutPoint+1:]...)
	case *ast.CaseClause:
		cutPoint := utilities.URandIntN(len(choosenNode.Body))
		choosenNode.Body = append(choosenNode.Body[:cutPoint], choosenNode.Body[cutPoint+1:]...)
	}
	return nil
}
