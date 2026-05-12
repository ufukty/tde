package lines

import (
	"fmt"
	"go/ast"
	"tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/utilities/pick"
	"tde/internal/utilities/randoms"
)

func SwapLines(ctx *models.MutationParameters) error {
	blockStmts := listBlockStmts(ctx.FuncDecl.Body, 2)
	if len(blockStmts) == 0 {
		return models.ErrUnsupportedMutation
	}
	choosenNode, err := pick.Pick(blockStmts)
	if err != nil {
		return fmt.Errorf("picking one out of many block statements: %w", err)
	}

	switch choosenNode := choosenNode.(type) {
	case *ast.BlockStmt:
		cutPoint := randoms.UniformIntN(len(choosenNode.List) - 1)
		choosenNode.List[cutPoint], choosenNode.List[cutPoint+1] = choosenNode.List[cutPoint+1], choosenNode.List[cutPoint]
	case *ast.CommClause:
		cutPoint := randoms.UniformIntN(len(choosenNode.Body) - 1)
		choosenNode.Body[cutPoint], choosenNode.Body[cutPoint+1] = choosenNode.Body[cutPoint+1], choosenNode.Body[cutPoint]
	case *ast.CaseClause:
		cutPoint := randoms.UniformIntN(len(choosenNode.Body) - 1)
		choosenNode.Body[cutPoint], choosenNode.Body[cutPoint+1] = choosenNode.Body[cutPoint+1], choosenNode.Body[cutPoint]
	}
	return nil
}
