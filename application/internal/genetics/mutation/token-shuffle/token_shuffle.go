package token_shuffle

import (
	"go/ast"
	"go/token"
	"tde/internal/genetics/mutation/common"
	"tde/internal/tokenw"
	utl "tde/internal/utilities"
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
		newToken = *utl.PickExcept(tokenw.AcceptedByBasicLit, []token.Token{n.Kind})
		n.Kind = newToken
	case *ast.UnaryExpr:
		newToken = *utl.PickExcept(tokenw.AcceptedByUnaryExpr, []token.Token{n.Op})
		n.Op = newToken
	case *ast.BinaryExpr:
		newToken = *utl.PickExcept(tokenw.AcceptedByBinaryExpr, []token.Token{n.Op})
		n.Op = newToken
	case *ast.IncDecStmt:
		newToken = *utl.PickExcept(tokenw.AcceptedByIncDecStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.AssignStmt:
		newToken = *utl.PickExcept(tokenw.AcceptedByAssignStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.BranchStmt:
		newToken = *utl.PickExcept(tokenw.AcceptedByBranchStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.RangeStmt:
		newToken = *utl.PickExcept(tokenw.AcceptedByRangeStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.GenDecl: // what are the chances for an import statement can work as a type declaration?
		newToken = *utl.PickExcept(tokenw.AcceptedByGenDecl, []token.Token{n.Tok})
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
	changedNode = *utl.Pick(tokenContainingNodes)
	newToken = chooseNewTokenAndAssign(changedNode)
	return changedNode, newToken, true
}

func GeneticOperation(ctx *common.GeneticOperationContext) bool {
	_, _, ok := Perform(ctx.FuncDecl.Body)
	return ok
}
