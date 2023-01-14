package node_constructor

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
	"tde/internal/utilities"
)

var CreatedVariables []*ast.Ident
var DeclaredFunctionNames []*ast.Ident
var GeneratedBranchLabels []*ast.Ident

func init() {

}

func generateVariableName() *ast.Ident {
	ident := ast.NewIdent(fmt.Sprintf("var%d", len(CreatedVariables)+1))
	CreatedVariables = append(CreatedVariables, ident)
	return ident
}

func generateFunctionName() *ast.Ident {
	ident := ast.NewIdent(fmt.Sprintf("HelperFunction%d", len(DeclaredFunctionNames)+1))
	DeclaredFunctionNames = append(DeclaredFunctionNames, ident)
	return ident
}

func generateBranchLabel() *ast.Ident {
	ident := ast.NewIdent(fmt.Sprintf("BranchLabel%d", len(GeneratedBranchLabels)+1))
	GeneratedBranchLabels = append(GeneratedBranchLabels, ident)
	return ident
}

func basicIntegerLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.INT, Value: fmt.Sprint(*utilities.Pick([]int{0, 1}))}
}

func basicFloatLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.FLOAT, Value: fmt.Sprint(utilities.URandFloatForCrypto()), ValuePos: token.NoPos}
}

func basicStringLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.STRING, Value: "", ValuePos: token.NoPos}
}

func basicCharacterLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.CHAR, Value: *utilities.Pick(strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789()[]{}_.=&!+-*/%:; ", "")), ValuePos: token.NoPos}
}
