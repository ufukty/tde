package caast

import (
	"fmt"
	"tde/internal/utilities"

	"go/ast"
	"go/token"
)

func GenerateRandomLiteral() (kind token.Token, value string) {
	return (*utilities.Pick([]func() (token.Token, string){
		func() (token.Token, string) {
			return token.INT, "0"
		},
		func() (token.Token, string) {
			return token.STRING, ""
		},
		func() (token.Token, string) {
			return token.FLOAT, fmt.Sprint(utilities.URandFloatForCrypto())
		},
	}))()
}

func GenerateRandomNumberOfInstance(nodeTypeClass NodeTypeClass) []ast.Node {
	list := []ast.Node{}
	switch nodeTypeClass {
	case Statement:
		return nil
	}
	return list
}

func GenerateRandomNumberOfStatements() []ast.Stmt {
	stmts := []ast.Stmt{}
	for _, n := range GenerateRandomNumberOfInstance(Statement) {
		if n, ok := n.(ast.Stmt); ok {
			stmts = append(stmts, n)
		} else {
			panic("looks like GenerateRandomNumberOfInstance() returned another type of node then given (Statement)")
		}
	}
	return stmts
}

func GenerateReturnStatement() []ast.Expr {
	return []ast.Expr{}
}

func GenerateBooleanExpression() ast.Expr {
	return &ast.BadExpr{}
}

func GenerateBlockStatement() *ast.BlockStmt {
	return &ast.BlockStmt{}
}
