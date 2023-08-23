package token_shuffle

import (
	"tde/internal/evolution/genetics/mutation/common"
	"tde/internal/evolution/genetics/mutation/tokens"
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
		newToken = *utilities.PickExcept(tokens.AcceptedByBasicLit, []token.Token{n.Kind})
		n.Kind = newToken
	case *ast.UnaryExpr:
		newToken = *utilities.PickExcept(tokens.AcceptedByUnaryExpr, []token.Token{n.Op})
		n.Op = newToken
	case *ast.BinaryExpr:
		newToken = *utilities.PickExcept(tokens.AcceptedByBinaryExpr, []token.Token{n.Op})
		n.Op = newToken
	case *ast.IncDecStmt:
		newToken = *utilities.PickExcept(tokens.AcceptedByIncDecStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.AssignStmt:
		newToken = *utilities.PickExcept(tokens.AcceptedByAssignStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.BranchStmt:
		newToken = *utilities.PickExcept(tokens.AcceptedByBranchStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.RangeStmt:
		newToken = *utilities.PickExcept(tokens.AcceptedByRangeStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.GenDecl: // what are the chances for an import statement can work as a type declaration?
		newToken = *utilities.PickExcept(tokens.AcceptedByGenDecl, []token.Token{n.Tok})
		n.Tok = newToken
	}
	return
}

func Perform(n ast.Node) (changedNode ast.Node, newToken token.Token, ok bool) {
	// list available
	tokenContainingNodes := listTokenContainingNodes(n)
	if len(tokenContainingNodes) == 0 {
		return changedNode, newToken, false
	}
	changedNode = *utilities.Pick(tokenContainingNodes)
	newToken = chooseNewTokenAndAssign(changedNode)
	return changedNode, newToken, true
}

func GeneticOperation(ctx *common.GeneticOperationContext) bool {
	_, _, ok := Perform(ctx.FuncDecl.Body)
	return ok
}
