package astcfg

import (
	"fmt"
	"go/ast"
	"go/token"
	"tde/internal/utilities"
)

type NodeConstructor struct {
	Dictionary map[NodeType]func() ast.Node
}

func NewNodeConstructor() *NodeConstructor {
	nc := NodeConstructor{}
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
		BasicLit:       nc.BasicLit,
		BinaryExpr:     nc.BinaryExpr,
		BlockStmt:      nc.BlockStmt,
		BranchStmt:     nc.BranchStmt,
		CallExpr:       nc.CallExpr,
		CaseClause:     nc.CaseClause,
		ChanType:       nc.ChanType,
		CommClause:     nc.CommClause,
		CompositeLit:   nc.CompositeLit,
		DeclStmt:       nc.DeclStmt,
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

func (*NodeConstructor) BasicLit() ast.Node {
	litKind := *utilities.Pick([]token.Token{token.INT, token.FLOAT, token.IMAG, token.CHAR, token.STRING})
	litValue := map[token.Token]string{
		token.INT:    "0",
		token.FLOAT:  fmt.Sprint(utilities.URandFloatForCrypto()),
		token.IMAG:   "0.0i",
		token.CHAR:   "'0'",
		token.STRING: `""`,
	}[litKind]
	return &ast.BasicLit{
		Kind:  litKind,
		Value: litValue,
	}
}

func (*NodeConstructor) BinaryExpr() ast.Node {
	return &ast.BinaryExpr{}
}

func (*NodeConstructor) BlockStmt() ast.Node {
	return &ast.BlockStmt{}
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

func (*NodeConstructor) DeclStmt() ast.Node {
	return &ast.DeclStmt{}
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

func (*NodeConstructor) FuncDecl() ast.Node {
	return &ast.FuncDecl{}
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
	return &ast.IfStmt{}
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
	return &ast.ValueSpec{}
}

var nodeConstructor = NewNodeConstructor()
