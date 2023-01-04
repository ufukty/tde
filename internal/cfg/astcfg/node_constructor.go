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
	CreatedVariables        []*ast.Ident
	DeclaredFunctionNames   []*ast.Ident
	GeneratedBranchLabels   []*ast.Ident
	AllowedPackagesToImport []string
	Classes                 map[NodeTypeClass][]NodeType
	Dictionary              map[NodeType]func() ast.Node
}

// TODO: Consider adding those:
// Comment
// CommentGroup
// File
// Package
// BadDecl
// BadExpr
// BadStmt
// BECAUSE: Those are currently available structures in ast package and their usage is not planned.
func NewNodeConstructor() *NodeConstructor {
	nc := NodeConstructor{}
	nc.CreatedVariables = []*ast.Ident{}
	nc.Classes = map[NodeTypeClass][]NodeType{
		Expression: {
			ArrayType, BasicLit, BinaryExpr, CallExpr, ChanType,
			CompositeLit, Ellipsis, FuncLit, FuncType, Ident, IndexExpr,
			IndexListExpr, InterfaceType, KeyValueExpr, MapType, ParenExpr,
			SelectorExpr, SliceExpr, StarExpr, StructType, TypeAssertExpr,
			UnaryExpr,
		},
		Statement: {
			AssignStmt, BlockStmt, BranchStmt, CaseClause,
			CommClause, DeclStmt, DeferStmt, EmptyStmt, ExprStmt, ForStmt,
			GoStmt, IfStmt, IncDecStmt, LabeledStmt, RangeStmt, ReturnStmt,
			SelectStmt, SendStmt, SwitchStmt, TypeSwitchStmt,
		},
		Declaration: {
			FuncDecl, GenDecl,
		},
		Specification: {
			ImportSpec, TypeSpec, ValueSpec,
		},
		TypeDeclaration: {
			ArrayType, ChanType, FuncType, InterfaceType, MapType, StructType,
			Ident, ParenExpr, SelectorExpr, StarExpr,
		},
	}
	nc.AllowedPackagesToImport = []string{
		"fmt", "strings", "math",
	}
	nc.Dictionary = map[NodeType]func() ast.Node{
		Ident: nc.Ident, // DONE:

		ArrayType:     nc.ArrayType,
		ChanType:      nc.ChanType,
		FuncType:      nc.FuncType,
		InterfaceType: nc.InterfaceType,
		MapType:       nc.MapType,
		StructType:    nc.StructType,

		AssignStmt:     nc.AssignStmt, // DONE:
		BlockStmt:      nc.BlockStmt,  // DONE:
		BranchStmt:     nc.BranchStmt,
		DeclStmt:       nc.DeclStmt, // DONE:
		DeferStmt:      nc.DeferStmt,
		EmptyStmt:      nc.EmptyStmt,
		ExprStmt:       nc.ExprStmt,
		ForStmt:        nc.ForStmt,
		GoStmt:         nc.GoStmt,
		IfStmt:         nc.IfStmt, // DONE:
		IncDecStmt:     nc.IncDecStmt,
		LabeledStmt:    nc.LabeledStmt,
		RangeStmt:      nc.RangeStmt,
		ReturnStmt:     nc.ReturnStmt,
		SelectStmt:     nc.SelectStmt,
		SendStmt:       nc.SendStmt,
		SwitchStmt:     nc.SwitchStmt,
		TypeSwitchStmt: nc.TypeSwitchStmt,

		BasicLit:     nc.BasicLit, // DONE:
		CompositeLit: nc.CompositeLit,
		FuncLit:      nc.FuncLit,

		BinaryExpr:     nc.BinaryExpr, // DONE:
		CallExpr:       nc.CallExpr,
		IndexExpr:      nc.IndexExpr,
		IndexListExpr:  nc.IndexListExpr,
		KeyValueExpr:   nc.KeyValueExpr,
		ParenExpr:      nc.ParenExpr,
		SelectorExpr:   nc.SelectorExpr,
		SliceExpr:      nc.SliceExpr,
		StarExpr:       nc.StarExpr,
		TypeAssertExpr: nc.TypeAssertExpr,
		UnaryExpr:      nc.UnaryExpr,

		FuncDecl: nc.FuncDecl,
		GenDecl:  nc.GenDecl, // DONE:

		CaseClause: nc.CaseClause,
		CommClause: nc.CommClause,
		Ellipsis:   nc.Ellipsis,

		ImportSpec: nc.ImportSpec,
		TypeSpec:   nc.TypeSpec,
		ValueSpec:  nc.ValueSpec, // DONE:

		Field:     nc.Field,
		FieldList: nc.FieldList,
	}
	return &nc
}

func (nc *NodeConstructor) Construct(nodeType NodeType) ast.Node {
	return nc.Dictionary[nodeType]()
}

// MARK: Compliying Interface instance

// Chooses a node type that confirms ast.Spec interface; initializes and returns
func (nc *NodeConstructor) Spec() ast.Spec {
	return nc.Construct(*utilities.Pick(nc.Classes[Specification])).(ast.Spec)
}

// Chooses a node type that confirms ast.Decl interface; initializes and returns
func (nc *NodeConstructor) Decl() ast.Decl {
	return nc.Construct(*utilities.Pick(nc.Classes[Declaration])).(ast.Decl)
}

// Chooses a node type that confirms ast.Expr interface; initializes and returns
func (nc *NodeConstructor) Expr() ast.Expr {
	return nc.Construct(*utilities.Pick(nc.Classes[Expression])).(ast.Expr)
}

// Chooses a node type that confirms ast.Stmt interface; initializes and returns
func (nc *NodeConstructor) Stmt() ast.Stmt {
	return nc.Construct(*utilities.Pick(nc.Classes[Statement])).(ast.Stmt)
}

// Chooses, initialized and returns an instance of random Type Declaration (Expression)
func (nc *NodeConstructor) Type() ast.Expr {
	return nc.Construct(*utilities.Pick(nc.Classes[TypeDeclaration])).(ast.Expr)
}

// MARK: Helper methods

func (nc *NodeConstructor) generateVariableName() *ast.Ident {
	ident := ast.NewIdent(fmt.Sprintf("var%d", len(nc.CreatedVariables)+1))
	nc.CreatedVariables = append(nc.CreatedVariables, ident)
	return ident
}

func (nc *NodeConstructor) generateFunctionName() *ast.Ident {
	ident := ast.NewIdent(fmt.Sprintf("HelperFunction%d", len(nc.DeclaredFunctionNames)+1))
	nc.DeclaredFunctionNames = append(nc.DeclaredFunctionNames, ident)
	return ident
}

func (nc *NodeConstructor) generateBranchLabel() *ast.Ident {
	ident := ast.NewIdent(fmt.Sprintf("BranchLabel%d", len(nc.GeneratedBranchLabels)+1))
	nc.GeneratedBranchLabels = append(nc.GeneratedBranchLabels, ident)
	return ident
}

func (nc *NodeConstructor) basicIntegerLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.INT, Value: fmt.Sprint(*utilities.Pick([]int{0, 1}))}
}

func (nc *NodeConstructor) basicFloatLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.FLOAT, Value: fmt.Sprint(utilities.URandFloatForCrypto()), ValuePos: token.NoPos}
}

func (nc *NodeConstructor) basicStringLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.STRING, Value: "", ValuePos: token.NoPos}
}

func (nc *NodeConstructor) basicCharacterLiteral() *ast.BasicLit {
	return &ast.BasicLit{Kind: token.CHAR, Value: *utilities.Pick(strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789()[]{}_.=&!+-*/%:; ", "")), ValuePos: token.NoPos}
}

// MARK: Node types

func (nc *NodeConstructor) ArrayType() ast.Node { // FIXME: // is there any usecase thar is not achievable with a slice but only with a ...T array
	return &ast.ArrayType{Lbrack: token.NoPos, Len: nil, Elt: nc.Expr()}
}

func (nc *NodeConstructor) AssignStmt() ast.Node {
	return &ast.AssignStmt{Lhs: []ast.Expr{nc.Expr()}, TokPos: token.NoPos, Tok: *utilities.Pick(tokenConstructor.AccepetedByAssignStmt), Rhs: []ast.Expr{nc.Expr()}}
}

func (nc *NodeConstructor) BasicLit() ast.Node {
	return (*utilities.Pick([]func() *ast.BasicLit{nc.basicIntegerLiteral, nc.basicStringLiteral, nc.basicFloatLiteral, nc.basicCharacterLiteral}))()
}

func (nc *NodeConstructor) BinaryExpr() ast.Node {
	return &ast.BinaryExpr{X: nc.Expr(), OpPos: token.NoPos, Op: *utilities.Pick(tokenConstructor.AcceptedByBinaryExpr), Y: nc.Expr()}
}

func (nc *NodeConstructor) BlockStmt() ast.Node {
	return &ast.BlockStmt{List: []ast.Stmt{nc.Stmt()}, Lbrace: token.NoPos, Rbrace: token.NoPos}
}

func (nc *NodeConstructor) BranchStmt() ast.Node {
	return &ast.BranchStmt{TokPos: token.NoPos, Tok: *utilities.Pick(tokenConstructor.AcceptedByBranchStmt), Label: *utilities.Pick(nc.GeneratedBranchLabels)}
}

func (nc *NodeConstructor) CallExpr() ast.Node { // TODO: function calls with more than 1 arguments
	return &ast.CallExpr{Fun: nc.Expr(), Lparen: token.NoPos, Args: []ast.Expr{nc.Expr()}, Ellipsis: token.NoPos, Rparen: token.NoPos}
}

func (nc *NodeConstructor) CaseClause() ast.Node {
	return &ast.CaseClause{Case: token.NoPos, List: []ast.Expr{nc.Expr()}, Colon: token.NoPos, Body: []ast.Stmt{nc.Stmt()}}
}

func (nc *NodeConstructor) ChanType() ast.Node {
	return &ast.ChanType{Begin: token.NoPos, Arrow: token.NoPos, Dir: *utilities.Pick([]ast.ChanDir{ast.SEND, ast.RECV}), Value: nc.Type()}
}

func (nc *NodeConstructor) CommClause() ast.Node {
	return &ast.CommClause{Case: token.NoPos, Colon: token.NoPos, Body: []ast.Stmt{nc.Stmt()}}
}

func (nc *NodeConstructor) CompositeLit() ast.Node { // TODO: check Incomplete property
	return &ast.CompositeLit{Type: nc.Type(), Lbrace: token.NoPos, Elts: []ast.Expr{nc.Expr()}, Rbrace: token.NoPos, Incomplete: false}
}

func (nc *NodeConstructor) DeclStmt() ast.Node { // either with initial value assignment or declaration only
	return &ast.DeclStmt{Decl: nc.GenDecl().(*ast.GenDecl)}
}

func (nc *NodeConstructor) DeferStmt() ast.Node {
	return &ast.DeferStmt{Defer: token.NoPos, Call: nc.CallExpr().(*ast.CallExpr)}
}

func (nc *NodeConstructor) Ellipsis() ast.Node {
	return &ast.Ellipsis{Ellipsis: token.NoPos, Elt: nc.Expr()}
}

func (nc *NodeConstructor) EmptyStmt() ast.Node {
	return &ast.EmptyStmt{Semicolon: token.NoPos, Implicit: false}
}

func (nc *NodeConstructor) ExprStmt() ast.Node {
	return &ast.ExprStmt{X: nc.Expr()}
}

func (nc *NodeConstructor) Field() ast.Node {
	return &ast.Field{Names: []*ast.Ident{nc.Ident().(*ast.Ident)}, Type: nc.Type(), Tag: nil}
}

func (nc *NodeConstructor) FieldList() ast.Node {
	return &ast.FieldList{Opening: token.NoPos, List: []*ast.Field{nc.Field().(*ast.Field)}, Closing: token.NoPos}
}

func (nc *NodeConstructor) ForStmt() ast.Node {
	return &ast.ForStmt{For: token.NoPos, Init: nc.Stmt(), Cond: nc.Expr(), Post: nc.Stmt(), Body: nc.BlockStmt().(*ast.BlockStmt)}
}

func (nc *NodeConstructor) FuncDecl() ast.Node { // TODO: Consider adding receiver functions (methods)
	return &ast.FuncDecl{Name: nc.generateFunctionName(), Type: nc.FuncType().(*ast.FuncType), Body: nc.BlockStmt().(*ast.BlockStmt)}
}

func (nc *NodeConstructor) FuncLit() ast.Node { // TODO:
	return &ast.FuncLit{Type: nc.FuncType().(*ast.FuncType), Body: nc.BlockStmt().(*ast.BlockStmt)}
}

func (nc *NodeConstructor) FuncType() ast.Node { // FIXME:
	return &ast.FuncType{Func: token.NoPos, TypeParams: nc.FieldList().(*ast.FieldList), Params: nc.FieldList().(*ast.FieldList), Results: nc.FieldList().(*ast.FieldList)}
}

// Produces only "variable" declarations. "import", "constant", "type" declarations are ignored.
func (nc *NodeConstructor) GenDecl() ast.Node {
	return &ast.GenDecl{TokPos: token.NoPos, Tok: token.VAR, Lparen: token.NoPos, Rparen: token.NoPos, Specs: []ast.Spec{nc.ValueSpec().(*ast.ValueSpec)}}
}

func (nc *NodeConstructor) GoStmt() ast.Node {
	return &ast.GoStmt{Go: token.NoPos, Call: nc.CallExpr().(*ast.CallExpr)}
}

func (nc *NodeConstructor) Ident() ast.Node {
	return nc.generateVariableName()
}

func (nc *NodeConstructor) IfStmt() ast.Node {
	return &ast.IfStmt{If: token.NoPos, Init: nil, Cond: nc.Expr(), Body: &ast.BlockStmt{Lbrace: token.NoPos, List: []ast.Stmt{nc.Stmt()}, Rbrace: token.NoPos}, Else: nil}
}

func (nc *NodeConstructor) ImportSpec() ast.Node { // TODO: Store imported packages for later use
	return &ast.ImportSpec{Name: nil, Path: &ast.BasicLit{ValuePos: token.NoPos, Kind: token.STRING, Value: *utilities.Pick(nc.AllowedPackagesToImport)}, EndPos: token.NoPos}
}

func (nc *NodeConstructor) IncDecStmt() ast.Node {
	return &ast.IncDecStmt{X: nc.Expr(), TokPos: token.NoPos, Tok: *utilities.Pick(tokenConstructor.AccepetedByIncDecStmt)}
}

func (nc *NodeConstructor) IndexExpr() ast.Node {
	return &ast.IndexExpr{X: nc.Expr(), Lbrack: token.NoPos, Index: nc.Expr(), Rbrack: token.NoPos}
}

func (nc *NodeConstructor) IndexListExpr() ast.Node { // TODO: Multi-dimensional arrays
	return &ast.IndexListExpr{X: nc.Expr(), Lbrack: token.NoPos, Indices: []ast.Expr{nc.Expr()}, Rbrack: token.NoPos}
}

func (nc *NodeConstructor) InterfaceType() ast.Node {
	return &ast.InterfaceType{Interface: token.NoPos, Methods: nc.FieldList().(*ast.FieldList), Incomplete: false}
}

func (nc *NodeConstructor) KeyValueExpr() ast.Node {
	return &ast.KeyValueExpr{Key: nc.Expr(), Colon: token.NoPos, Value: nc.Expr()}
}

func (nc *NodeConstructor) LabeledStmt() ast.Node {
	return &ast.LabeledStmt{Label: nc.generateBranchLabel(), Colon: token.NoPos, Stmt: nc.Stmt()}
}

func (nc *NodeConstructor) MapType() ast.Node {
	return &ast.MapType{Map: token.NoPos, Key: nc.Type(), Value: nc.Type()}
}

func (nc *NodeConstructor) ParenExpr() ast.Node {
	return &ast.ParenExpr{Lparen: token.NoPos, X: nc.Expr(), Rparen: token.NoPos}
}

func (nc *NodeConstructor) RangeStmt() ast.Node {
	return &ast.RangeStmt{For: token.NoPos, Key: nc.Expr(), Value: nc.Expr(), TokPos: token.NoPos, Tok: *utilities.Pick(tokenConstructor.AcceptedByRangeStmt), X: nc.Expr(), Body: nc.BlockStmt().(*ast.BlockStmt)}
}

func (nc *NodeConstructor) ReturnStmt() ast.Node { // TODO: multiple return values
	return &ast.ReturnStmt{Return: token.NoPos, Results: []ast.Expr{nc.Expr()}}
}

func (nc *NodeConstructor) SelectorExpr() ast.Node { // FIXME: randomly produced X and Sel values will never work, maybe choose from imported libraries' exported functions, or previously declared struct instances that has methods
	return &ast.SelectorExpr{X: nc.Expr(), Sel: &ast.Ident{}}
}

func (nc *NodeConstructor) SelectStmt() ast.Node {
	return &ast.SelectStmt{}
}

func (nc *NodeConstructor) SendStmt() ast.Node {
	return &ast.SendStmt{}
}

func (nc *NodeConstructor) SliceExpr() ast.Node {
	return &ast.SliceExpr{}
}

func (nc *NodeConstructor) StarExpr() ast.Node {
	return &ast.StarExpr{}
}

func (nc *NodeConstructor) StructType() ast.Node {
	return &ast.StructType{}
}

func (nc *NodeConstructor) SwitchStmt() ast.Node {
	return &ast.SwitchStmt{}
}

func (nc *NodeConstructor) TypeAssertExpr() ast.Node {
	return &ast.TypeAssertExpr{}
}

func (nc *NodeConstructor) TypeSpec() ast.Node {
	return &ast.TypeSpec{}
}

func (nc *NodeConstructor) TypeSwitchStmt() ast.Node {
	return &ast.TypeSwitchStmt{}
}

func (nc *NodeConstructor) UnaryExpr() ast.Node {
	return &ast.UnaryExpr{}
}

// Returns ValueSpec which is list of pairs of variable names and values to assign
func (nc *NodeConstructor) ValueSpec() ast.Node {
	return &ast.ValueSpec{Names: []*ast.Ident{nc.generateVariableName()}, Values: []ast.Expr{nc.Expr()}}
}

var nodeConstructor = NewNodeConstructor()
