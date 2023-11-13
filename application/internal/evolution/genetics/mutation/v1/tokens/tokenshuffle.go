package tokens

import (
	"tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/utilities"

	"go/ast"
	"go/token"
)

func listTokenContainingNodes(n ast.Node) (tokenContainingNodes []ast.Node) {
	ast.Inspect(n, func(n ast.Node) bool {
		switch n := n.(type) {
		case
			*ast.BasicLit,
			*ast.UnaryExpr,
			*ast.BinaryExpr,
			*ast.IncDecStmt,
			*ast.AssignStmt,
			*ast.BranchStmt,
			*ast.RangeStmt,
			*ast.GenDecl:
			tokenContainingNodes = append(tokenContainingNodes, n)
		}
		return true
	})
	return
}

func chooseNewTokenAndAssign(n ast.Node) (newToken token.Token) {
	switch n := n.(type) {
	case *ast.BasicLit:
		newToken = *utilities.PickExcept(AcceptedByBasicLit, []token.Token{n.Kind})
		n.Kind = newToken
	case *ast.UnaryExpr:
		newToken = *utilities.PickExcept(AcceptedByUnaryExpr, []token.Token{n.Op})
		n.Op = newToken
	case *ast.BinaryExpr:
		newToken = *utilities.PickExcept(AcceptedByBinaryExpr, []token.Token{n.Op})
		n.Op = newToken
	case *ast.IncDecStmt:
		newToken = *utilities.PickExcept(AcceptedByIncDecStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.AssignStmt:
		newToken = *utilities.PickExcept(AcceptedByAssignStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.BranchStmt:
		newToken = *utilities.PickExcept(AcceptedByBranchStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.RangeStmt:
		newToken = *utilities.PickExcept(AcceptedByRangeStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.GenDecl: // what are the chances for an import statement can work as a type declaration?
		newToken = *utilities.PickExcept(AcceptedByGenDecl, []token.Token{n.Tok})
		n.Tok = newToken
	}
	return
}

func tokenShuffle(n ast.Node) (changedNode ast.Node, newToken token.Token, err error) {
	// list available
	tokenContainingNodes := listTokenContainingNodes(n)
	if len(tokenContainingNodes) == 0 {
		return changedNode, newToken, models.ErrUnsupportedMutation
	}
	changedNode = *utilities.Pick(tokenContainingNodes)
	newToken = chooseNewTokenAndAssign(changedNode)
	return changedNode, newToken, nil
}

func TokenShuffle(ctx *models.MutationParameters) error {
	_, _, err := tokenShuffle(ctx.FuncDecl.Body)
	return err
}
