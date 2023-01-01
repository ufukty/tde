package astcfg

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
	"tde/internal/utilities"
)

// struct is used to avoid cluttering the namespace with the names of node constructor functions
type NodeConstructor struct {
	CreatedVariables  []*ast.Ident
	DeclaredFunctions []*ast.FuncDecl
	Dictionary        map[NodeType]func() ast.Node
}

func NewNodeConstructor() *NodeConstructor {
	nc := NodeConstructor{}
	nc.CreatedVariables = []*ast.Ident{}
	nc.Dictionary = map[NodeType]func() ast.Node{
		// BECAUSE: Those are currently available structures in ast package and their usage is not planned.
		// Comment:        func() ast.Node {},
		// CommentGroup:   func() ast.Node {},
		// Field:          func() ast.Node {},
		// FieldList:      func() ast.Node {},
		// File:           func() ast.Node {},
		// Package:        func() ast.Node {},
		// BadDecl:        func() ast.Node {},
		// BadExpr:        func() ast.Node {},
		// BadStmt:        func() ast.Node {},
		ArrayType:      nc.ArrayType,
		AssignStmt:     nc.AssignStmt,
		BasicLit:       nc.BasicLit, // DONE:
		BinaryExpr:     nc.BinaryExpr,
		BlockStmt:      nc.BlockStmt, // DONE:
		BranchStmt:     nc.BranchStmt,
		CallExpr:       nc.CallExpr,
		CaseClause:     nc.CaseClause,
		ChanType:       nc.ChanType,
		CommClause:     nc.CommClause,
		CompositeLit:   nc.CompositeLit,
		DeclStmt:       nc.DeclStmt, // DONE:
		DeferStmt:      nc.DeferStmt,
		Ellipsis:       nc.Ellipsis,
		EmptyStmt:      nc.EmptyStmt,
		ExprStmt:       nc.ExprStmt,
		ForStmt:        nc.ForStmt,
		FuncDecl:       nc.FuncDecl,
		FuncLit:        nc.FuncLit,
		FuncType:       nc.FuncType,
		GenDecl:        nc.GenDecl,
		GoStmt:         nc.GoStmt,
		Ident:          nc.Ident,
		IfStmt:         nc.IfStmt,
		ImportSpec:     nc.ImportSpec,
		IncDecStmt:     nc.IncDecStmt,
		IndexExpr:      nc.IndexExpr,
		IndexListExpr:  nc.IndexListExpr,
		InterfaceType:  nc.InterfaceType,
		KeyValueExpr:   nc.KeyValueExpr,
		LabeledStmt:    nc.LabeledStmt,
		MapType:        nc.MapType,
		ParenExpr:      nc.ParenExpr,
		RangeStmt:      nc.RangeStmt,
		ReturnStmt:     nc.ReturnStmt,
		SelectorExpr:   nc.SelectorExpr,
		SelectStmt:     nc.SelectStmt,
		SendStmt:       nc.SendStmt,
		SliceExpr:      nc.SliceExpr,
		StarExpr:       nc.StarExpr,
		StructType:     nc.StructType,
		SwitchStmt:     nc.SwitchStmt,
		TypeAssertExpr: nc.TypeAssertExpr,
		TypeSpec:       nc.TypeSpec,
		TypeSwitchStmt: nc.TypeSwitchStmt,
		UnaryExpr:      nc.UnaryExpr,
		ValueSpec:      nc.ValueSpec,
	}
	return &nc
}

func (nc *NodeConstructor) Construct(nodeType NodeType) ast.Node {
	return nc.Dictionary[nodeType]()
}

func (*NodeConstructor) ArrayType() ast.Node {
	return &ast.ArrayType{}
}

func (*NodeConstructor) AssignStmt() ast.Node {
	return &ast.AssignStmt{}
}

func (*NodeConstructor) basicIntegerLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.INT, Value: string(*utilities.Pick([]int{0, 1}))}
}

func (*NodeConstructor) basicFloatLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.FLOAT, Value: fmt.Sprint(utilities.URandFloatForCrypto()), ValuePos: token.NoPos}
}

func (*NodeConstructor) basicStringLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.STRING, Value: "", ValuePos: token.NoPos}
}

func (*NodeConstructor) basicCharacterLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.CHAR, Value: *utilities.Pick(strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789()[]{}_.=&!+-*/%:; ", "")), ValuePos: token.NoPos}
}

func (nc *NodeConstructor) BasicLit() *ast.BasicLit {
	return (*utilities.Pick([]func() *ast.BasicLit{nc.basicIntegerLiteral, nc.basicStringLiteral, nc.basicFloatLiteral, nc.basicCharacterLiteral}))()
}

func (*NodeConstructor) BinaryExpr() ast.Node {
	return &ast.BinaryExpr{}
}

func (*NodeConstructor) BlockStmt() ast.Node {
	return &ast.BlockStmt{List: []ast.Stmt{}, Lbrace: token.NoPos, Rbrace: token.NoPos}
}

func (*NodeConstructor) BranchStmt() ast.Node {
	return &ast.BranchStmt{}
}

func (*NodeConstructor) CallExpr() ast.Node {
	return &ast.CallExpr{}
}

func (*NodeConstructor) CaseClause() ast.Node {
	return &ast.CaseClause{}
}

func (*NodeConstructor) ChanType() ast.Node {
	return &ast.ChanType{}
}

func (*NodeConstructor) CommClause() ast.Node {
	return &ast.CommClause{}
}

func (*NodeConstructor) CompositeLit() ast.Node {
	return &ast.CompositeLit{}
}

func (nc *NodeConstructor) generateVariableName() *ast.Ident {
	ident := ast.NewIdent(fmt.Sprintf("var%d", len(nc.CreatedVariables)+1))
	nc.CreatedVariables = append(nc.CreatedVariables, ident)
	return ident
}

func (nc *NodeConstructor) DeclStmt() ast.Node { // either with initial value assignment or declaration only
	ident, value := nc.generateVariableName(), nc.BasicLit()
	return &ast.DeclStmt{
		Decl: &ast.GenDecl{
			TokPos: token.NoPos,
			Tok:    token.VAR,
			Lparen: token.NoPos,
			Rparen: token.NoPos,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names:  []*ast.Ident{ident},
					Values: []ast.Expr{value},
				},
			},
		},
	}
}

func (*NodeConstructor) DeferStmt() ast.Node {
	return &ast.DeferStmt{}
}

func (*NodeConstructor) Ellipsis() ast.Node {
	return &ast.Ellipsis{}
}

func (*NodeConstructor) EmptyStmt() ast.Node {
	return &ast.EmptyStmt{}
}

func (*NodeConstructor) ExprStmt() ast.Node {
	return &ast.ExprStmt{}
}

func (*NodeConstructor) ForStmt() ast.Node {
	return &ast.ForStmt{}
}

func (nc *NodeConstructor) generateFunctionName() *ast.Ident {
	ident := ast.NewIdent(fmt.Sprintf("HelperFunction%d", len(nc.CreatedVariables)+1))
	nc.CreatedVariables = append(nc.CreatedVariables, ident)
	return ident
}

func (*NodeConstructor) FuncDeclAsMethod() ast.Node {
	return &ast.FuncDecl{}
}

func (nc *NodeConstructor) FuncDecl() ast.Node {
	ident := nc.generateFunctionName()
	return &ast.FuncDecl{Name: ident}
}

func (*NodeConstructor) FuncLit() ast.Node {
	return &ast.FuncLit{}
}

func (*NodeConstructor) FuncType() ast.Node {
	return &ast.FuncType{}
}

func (*NodeConstructor) GenDecl() ast.Node {
	return &ast.GenDecl{}
}

func (*NodeConstructor) GoStmt() ast.Node {
	return &ast.GoStmt{}
}

func (*NodeConstructor) Ident() ast.Node {
	return &ast.Ident{}
}

func (*NodeConstructor) IfStmt() ast.Node {
	return &ast.IfStmt{
		If:   0,
		Init: nil,
		Cond: nil,
		Body: &ast.BlockStmt{},
		Else: nil,
	}
}

func (*NodeConstructor) ImportSpec() ast.Node {
	return &ast.ImportSpec{}
}

func (*NodeConstructor) IncDecStmt() ast.Node {
	return &ast.IncDecStmt{}
}

func (*NodeConstructor) IndexExpr() ast.Node {
	return &ast.IndexExpr{}
}

func (*NodeConstructor) IndexListExpr() ast.Node {
	return &ast.IndexListExpr{}
}

func (*NodeConstructor) InterfaceType() ast.Node {
	return &ast.InterfaceType{}
}

func (*NodeConstructor) KeyValueExpr() ast.Node {
	return &ast.KeyValueExpr{}
}

func (*NodeConstructor) LabeledStmt() ast.Node {
	return &ast.LabeledStmt{}
}

func (*NodeConstructor) MapType() ast.Node {
	return &ast.MapType{}
}

func (*NodeConstructor) ParenExpr() ast.Node {
	return &ast.ParenExpr{}
}

func (*NodeConstructor) RangeStmt() ast.Node {
	return &ast.RangeStmt{}
}

func (*NodeConstructor) ReturnStmt() ast.Node {
	return &ast.ReturnStmt{}
}

func (*NodeConstructor) SelectorExpr() ast.Node {
	return &ast.SelectorExpr{}
}

func (*NodeConstructor) SelectStmt() ast.Node {
	return &ast.SelectStmt{}
}

func (*NodeConstructor) SendStmt() ast.Node {
	return &ast.SendStmt{}
}

func (*NodeConstructor) SliceExpr() ast.Node {
	return &ast.SliceExpr{}
}

func (*NodeConstructor) StarExpr() ast.Node {
	return &ast.StarExpr{}
}

func (*NodeConstructor) StructType() ast.Node {
	return &ast.StructType{}
}

func (*NodeConstructor) SwitchStmt() ast.Node {
	return &ast.SwitchStmt{}
}

func (*NodeConstructor) TypeAssertExpr() ast.Node {
	return &ast.TypeAssertExpr{}
}

func (*NodeConstructor) TypeSpec() ast.Node {
	return &ast.TypeSpec{}
}

func (*NodeConstructor) TypeSwitchStmt() ast.Node {
	return &ast.TypeSwitchStmt{}
}

func (*NodeConstructor) UnaryExpr() ast.Node {
	return &ast.UnaryExpr{}
}

func (*NodeConstructor) ValueSpec() ast.Node {
	// Also created by GenDecl method
	return &ast.ValueSpec{}
}

var nodeConstructor = NewNodeConstructor()
