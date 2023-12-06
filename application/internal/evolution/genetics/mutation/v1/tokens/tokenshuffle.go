package tokens

import (
	"fmt"
	"go/ast"
	"go/token"
	"tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/utilities/pick"
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

func chooseNewTokenAndAssign(n ast.Node) (newToken token.Token, err error) {
	switch n := n.(type) {
	case *ast.BasicLit:
		newToken, err = pick.Except(AcceptedByBasicLit, []token.Token{n.Kind})
		if err != nil {
			return 0, fmt.Errorf("picking one out of many tokens: %w", err)
		}
		n.Kind = newToken
	case *ast.UnaryExpr:
		newToken, err = pick.Except(AcceptedByUnaryExpr, []token.Token{n.Op})
		if err != nil {
			return 0, fmt.Errorf("picking one out of many tokens: %w", err)
		}
		n.Op = newToken
	case *ast.BinaryExpr:
		newToken, err = pick.Except(AcceptedByBinaryExpr, []token.Token{n.Op})
		if err != nil {
			return 0, fmt.Errorf("picking one out of many tokens: %w", err)
		}
		n.Op = newToken
	case *ast.IncDecStmt:
		newToken, err = pick.Except(AcceptedByIncDecStmt, []token.Token{n.Tok})
		if err != nil {
			return 0, fmt.Errorf("picking one out of many tokens: %w", err)
		}
		n.Tok = newToken
	case *ast.AssignStmt:
		newToken, err = pick.Except(AcceptedByAssignStmt, []token.Token{n.Tok})
		if err != nil {
			return 0, fmt.Errorf("picking one out of many tokens: %w", err)
		}
		n.Tok = newToken
	case *ast.BranchStmt:
		newToken, err = pick.Except(AcceptedByBranchStmt, []token.Token{n.Tok})
		if err != nil {
			return 0, fmt.Errorf("picking one out of many tokens: %w", err)
		}
		n.Tok = newToken
	case *ast.RangeStmt:
		newToken, err = pick.Except(AcceptedByRangeStmt, []token.Token{n.Tok})
		if err != nil {
			return 0, fmt.Errorf("picking one out of many tokens: %w", err)
		}
		n.Tok = newToken
	case *ast.GenDecl: // what are the chances for an import statement can work as a type declaration?
		newToken, err = pick.Except(AcceptedByGenDecl, []token.Token{n.Tok})
		if err != nil {
			return 0, fmt.Errorf("picking one out of many tokens: %w", err)
		}
		n.Tok = newToken
	}
	return newToken, nil
}

func tokenShuffle(n ast.Node) (changedNode ast.Node, newToken token.Token, err error) {
	// list available
	tokenContainingNodes := listTokenContainingNodes(n)
	if len(tokenContainingNodes) == 0 {
		return changedNode, newToken, models.ErrUnsupportedMutation
	}
	changedNode, err = pick.Pick(tokenContainingNodes)
	if err != nil {
		return nil, 0, fmt.Errorf("picking the mutation node: %w", err)
	}
	newToken, err = chooseNewTokenAndAssign(changedNode)
	if err != nil {
		return nil, 0, fmt.Errorf("chooseNewTokenAndAssign: %w", err)
	}
	return changedNode, newToken, nil
}

func TokenShuffle(ctx *models.MutationParameters) error {
	_, _, err := tokenShuffle(ctx.FuncDecl.Body)
	return err
}
