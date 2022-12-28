package astcfg

import (
	"go/ast"
	"go/token"
	"tde/internal/utilities"
)

var NodeConstructors = map[NodeType]func() ast.Node{
	// Comment:        func() ast.Node { return &ast.Comment{} },
	// CommentGroup:   func() ast.Node { return &ast.CommentGroup{} },
	// Field:          func() ast.Node { return &ast.Field{} },
	// FieldList:      func() ast.Node { return &ast.FieldList{} },
	// File:           func() ast.Node { return &ast.File{} },
	// Package:        func() ast.Node { return &ast.Package{} },
	ArrayType: func() ast.Node { return &ast.ArrayType{} },
	AssignStmt: func() ast.Node {
		return &ast.AssignStmt{
			Lhs:    []ast.Expr{},
			TokPos: 0,
			Tok:    0,
			Rhs:    []ast.Expr{},
		}
	},
	BadDecl: func() ast.Node { return &ast.BadDecl{} },
	BadExpr: func() ast.Node { return &ast.BadExpr{} },
	BadStmt: func() ast.Node { return &ast.BadStmt{} },
	BasicLit: func() ast.Node { // DONE:
		literalKind := *utilities.Pick([]token.Token{token.INT, token.FLOAT, token.IMAG, token.CHAR, token.STRING})
		literalValue := map[token.Token]string{
			token.INT:    "0",
			token.FLOAT:  "0.0",
			token.IMAG:   "0.0i",
			token.CHAR:   "'0'",
			token.STRING: `""`,
		}[literalKind]
		return &ast.BasicLit{
			Kind:  literalKind,
			Value: literalValue,
		}
	},
	BinaryExpr: func() ast.Node {
		return &ast.BinaryExpr{
			X:     nil,
			OpPos: 0,
			Op:    0,
			Y:     nil,
		}
	},
	BlockStmt:      func() ast.Node { return &ast.BlockStmt{} },
	BranchStmt:     func() ast.Node { return &ast.BranchStmt{} },
	CallExpr:       func() ast.Node { return &ast.CallExpr{} },
	CaseClause:     func() ast.Node { return &ast.CaseClause{} },
	ChanType:       func() ast.Node { return &ast.ChanType{} },
	CommClause:     func() ast.Node { return &ast.CommClause{} },
	CompositeLit:   func() ast.Node { return &ast.CompositeLit{} },
	DeclStmt:       func() ast.Node { return &ast.DeclStmt{} },
	DeferStmt:      func() ast.Node { return &ast.DeferStmt{} },
	Ellipsis:       func() ast.Node { return &ast.Ellipsis{} },
	EmptyStmt:      func() ast.Node { return &ast.EmptyStmt{} },
	ExprStmt:       func() ast.Node { return &ast.ExprStmt{} },
	ForStmt:        func() ast.Node { return &ast.ForStmt{} },
	FuncDecl:       func() ast.Node { return &ast.FuncDecl{} },
	FuncLit:        func() ast.Node { return &ast.FuncLit{} },
	FuncType:       func() ast.Node { return &ast.FuncType{} },
	GenDecl:        func() ast.Node { return &ast.GenDecl{} },
	GoStmt:         func() ast.Node { return &ast.GoStmt{} },
	Ident:          func() ast.Node { return &ast.Ident{} },
	IfStmt:         func() ast.Node { return &ast.IfStmt{} },
	ImportSpec:     func() ast.Node { return &ast.ImportSpec{} },
	IncDecStmt:     func() ast.Node { return &ast.IncDecStmt{} },
	IndexExpr:      func() ast.Node { return &ast.IndexExpr{} },
	IndexListExpr:  func() ast.Node { return &ast.IndexListExpr{} },
	InterfaceType:  func() ast.Node { return &ast.InterfaceType{} },
	KeyValueExpr:   func() ast.Node { return &ast.KeyValueExpr{} },
	LabeledStmt:    func() ast.Node { return &ast.LabeledStmt{} },
	MapType:        func() ast.Node { return &ast.MapType{} },
	ParenExpr:      func() ast.Node { return &ast.ParenExpr{} },
	RangeStmt:      func() ast.Node { return &ast.RangeStmt{} },
	ReturnStmt:     func() ast.Node { return &ast.ReturnStmt{} },
	SelectorExpr:   func() ast.Node { return &ast.SelectorExpr{} },
	SelectStmt:     func() ast.Node { return &ast.SelectStmt{} },
	SendStmt:       func() ast.Node { return &ast.SendStmt{} },
	SliceExpr:      func() ast.Node { return &ast.SliceExpr{} },
	StarExpr:       func() ast.Node { return &ast.StarExpr{} },
	StructType:     func() ast.Node { return &ast.StructType{} },
	SwitchStmt:     func() ast.Node { return &ast.SwitchStmt{} },
	TypeAssertExpr: func() ast.Node { return &ast.TypeAssertExpr{} },
	TypeSpec:       func() ast.Node { return &ast.TypeSpec{} },
	TypeSwitchStmt: func() ast.Node { return &ast.TypeSwitchStmt{} },
	UnaryExpr:      func() ast.Node { return &ast.UnaryExpr{} },
	ValueSpec:      func() ast.Node { return &ast.ValueSpec{} },
}
