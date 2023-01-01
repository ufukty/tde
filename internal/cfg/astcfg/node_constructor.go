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
	CreatedVariables      []*ast.Ident
	DeclaredFunctionNames []*ast.Ident
	GeneratedBranchLabels []*ast.Ident
	Dictionary            map[NodeType]func() ast.Node
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
		AssignStmt:     nc.AssignStmt, // DONE:
		BasicLit:       nc.BasicLit,   // DONE:
		BinaryExpr:     nc.BinaryExpr, // DONE:
		BlockStmt:      nc.BlockStmt,  // DONE:
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
		GenDecl:        nc.GenDecl, // DONE:
		GoStmt:         nc.GoStmt,
		Ident:          nc.Ident,  // DONE:
		IfStmt:         nc.IfStmt, // DONE:
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
		ValueSpec:      nc.ValueSpec, // DONE:
	}
	return &nc
}

func (nc *NodeConstructor) Construct(nodeType NodeType) ast.Node {
	return nc.Dictionary[nodeType]()
}

// MARK: Compliying Interface instance

// Chooses a node type that confirms ast.Spec interface; initializes and returns
func (nc *NodeConstructor) Spec() ast.Spec {
	return nc.Construct(*utilities.Pick(Dict_NodeTypeClassToNodeType[Spec])).(ast.Spec)
}

// Chooses a node type that confirms ast.Decl interface; initializes and returns
func (nc *NodeConstructor) Decl() ast.Decl {
	return nc.Construct(*utilities.Pick(Dict_NodeTypeClassToNodeType[Declaration])).(ast.Decl)
}

// Chooses a node type that confirms ast.Expr interface; initializes and returns
func (nc *NodeConstructor) Expr() ast.Expr {
	return nc.Construct(*utilities.Pick(Dict_NodeTypeClassToNodeType[Expression])).(ast.Expr)
}

// Chooses a node type that confirms ast.Stmt interface; initializes and returns
func (nc *NodeConstructor) Stmt() ast.Stmt {
	return nc.Construct(*utilities.Pick(Dict_NodeTypeClassToNodeType[Statement])).(ast.Stmt)
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
	return &ast.BasicLit{Kind: token.INT, Value: string(*utilities.Pick([]int{0, 1}))}
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
	return &ast.BlockStmt{List: []ast.Stmt{}, Lbrace: token.NoPos, Rbrace: token.NoPos}
}

func (nc *NodeConstructor) BranchStmt() ast.Node {
	return &ast.BranchStmt{TokPos: token.NoPos, Tok: *utilities.Pick(tokenConstructor.AcceptedByBranchStmt), Label: nc.generateBranchLabel()}
}

func (nc *NodeConstructor) CallExpr() ast.Node { // FIXME: Consider function calls with more than 1 arguments
	return &ast.CallExpr{Fun: nc.Expr(), Lparen: token.NoPos, Args: []ast.Expr{nc.Expr()}, Ellipsis: token.NoPos, Rparen: token.NoPos}
}

func (nc *NodeConstructor) CaseClause() ast.Node {
	return &ast.CaseClause{
		Case:  token.NoPos,
		List:  []ast.Expr{nc.Expr()},
		Colon: token.NoPos,
		Body:  []ast.Stmt{nc.Stmt()},
	}
}

func (nc *NodeConstructor) ChanType() ast.Node {
	return &ast.ChanType{}
}

func (nc *NodeConstructor) CommClause() ast.Node {
	return &ast.CommClause{}
}

func (nc *NodeConstructor) CompositeLit() ast.Node {
	return &ast.CompositeLit{}
}

func (nc *NodeConstructor) DeclStmt() ast.Node { // either with initial value assignment or declaration only
	return &ast.DeclStmt{Decl: nc.GenDecl().(*ast.GenDecl)}
}

func (nc *NodeConstructor) DeferStmt() ast.Node {
	return &ast.DeferStmt{}
}

func (nc *NodeConstructor) Ellipsis() ast.Node {
	return &ast.Ellipsis{}
}

func (nc *NodeConstructor) EmptyStmt() ast.Node {
	return &ast.EmptyStmt{}
}

func (nc *NodeConstructor) ExprStmt() ast.Node {
	return &ast.ExprStmt{}
}

func (nc *NodeConstructor) ForStmt() ast.Node {
	return &ast.ForStmt{}
}

func (nc *NodeConstructor) FuncDeclAsMethod() ast.Node {
	return &ast.FuncDecl{}
}

func (nc *NodeConstructor) FuncDecl() ast.Node {
	return &ast.FuncDecl{Name: nc.generateFunctionName(), Type: nc.FuncType().(*ast.FuncType), Body: nc.BlockStmt().(*ast.BlockStmt)}
}

func (nc *NodeConstructor) FuncLit() ast.Node {
	return &ast.FuncLit{}
}

func (nc *NodeConstructor) FuncType() ast.Node {
	return &ast.FuncType{}
}

func (nc *NodeConstructor) GenDecl() ast.Node {
	// Produces only "variable" declarations. "import", "constant", "type" declarations are ignored.
	return &ast.GenDecl{TokPos: token.NoPos, Tok: token.VAR, Lparen: token.NoPos, Rparen: token.NoPos, Specs: []ast.Spec{nc.ValueSpec().(*ast.ValueSpec)}}
}

func (nc *NodeConstructor) GoStmt() ast.Node {
	return &ast.GoStmt{}
}

func (nc *NodeConstructor) Ident() ast.Node {
	return nc.generateVariableName()
}

func (nc *NodeConstructor) IfStmt() ast.Node {
	return &ast.IfStmt{If: token.NoPos, Init: nil, Cond: nc.Expr(), Body: &ast.BlockStmt{Lbrace: token.NoPos, List: []ast.Stmt{nc.Stmt()}, Rbrace: token.NoPos}, Else: nil}
}

func (nc *NodeConstructor) ImportSpec() ast.Node {
	return &ast.ImportSpec{}
}

func (nc *NodeConstructor) IncDecStmt() ast.Node {
	return &ast.IncDecStmt{}
}

func (nc *NodeConstructor) IndexExpr() ast.Node {
	return &ast.IndexExpr{}
}

func (nc *NodeConstructor) IndexListExpr() ast.Node {
	return &ast.IndexListExpr{}
}

func (nc *NodeConstructor) InterfaceType() ast.Node {
	return &ast.InterfaceType{}
}

func (nc *NodeConstructor) KeyValueExpr() ast.Node {
	return &ast.KeyValueExpr{}
}

func (nc *NodeConstructor) LabeledStmt() ast.Node {
	return &ast.LabeledStmt{}
}

func (nc *NodeConstructor) MapType() ast.Node {
	return &ast.MapType{}
}

func (nc *NodeConstructor) ParenExpr() ast.Node {
	return &ast.ParenExpr{}
}

func (nc *NodeConstructor) RangeStmt() ast.Node {
	return &ast.RangeStmt{}
}

func (nc *NodeConstructor) ReturnStmt() ast.Node {
	return &ast.ReturnStmt{}
}

func (nc *NodeConstructor) SelectorExpr() ast.Node {
	return &ast.SelectorExpr{}
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
