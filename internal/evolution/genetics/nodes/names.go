package nodes

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
	"tde/internal/utilities/pick"
	"tde/internal/utilities/randoms"
)

var CreatedVariables []*ast.Ident
var DeclaredFunctionNames []*ast.Ident
var GeneratedBranchLabels []*ast.Ident

func init() {
	CreatedVariables = []*ast.Ident{}
	DeclaredFunctionNames = []*ast.Ident{}
	GeneratedBranchLabels = []*ast.Ident{}
}

func generateNewIdent() *ast.Ident {
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
	l, _ := pick.Pick([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	return &ast.BasicLit{Kind: token.INT, Value: fmt.Sprint(l)}
}

func basicFloatLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.FLOAT, Value: fmt.Sprint(randoms.UniformCryptoFloat()), ValuePos: token.NoPos}
}

func basicStringLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.STRING, Value: "", ValuePos: token.NoPos}
}

func basicCharacterLiteral() *ast.BasicLit {
	v, _ := pick.Pick(strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789()[]{}_.=&!+-*/%:; \\", ""))
	return &ast.BasicLit{Kind: token.CHAR, Value: v, ValuePos: token.NoPos}
}
