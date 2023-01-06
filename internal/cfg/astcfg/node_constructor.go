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
	Dictionary              map[NodeType]any
	Dictionaries            struct {
		Statement       map[NodeType]func() ast.Stmt
		Expression      map[NodeType]func() ast.Expr
		Declaration     map[NodeType]func() ast.Decl
		Specification   map[NodeType]func() ast.Spec
		TypeDeclaration map[NodeType]func() ast.Expr
	}
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
	nc.Dictionaries = struct {
		Statement       map[NodeType]func() ast.Stmt
		Expression      map[NodeType]func() ast.Expr
		Declaration     map[NodeType]func() ast.Decl
		Specification   map[NodeType]func() ast.Spec
		TypeDeclaration map[NodeType]func() ast.Expr
	}{
		Statement: map[NodeType]func() ast.Stmt{
			AssignStmt:     func() ast.Stmt { return nc.AssignStmt() },
			BlockStmt:      func() ast.Stmt { return nc.BlockStmt() },
			BranchStmt:     func() ast.Stmt { return nc.BranchStmt() },
			CaseClause:     func() ast.Stmt { return nc.CaseClause() },
			CommClause:     func() ast.Stmt { return nc.CommClause() },
			DeclStmt:       func() ast.Stmt { return nc.DeclStmt() },
			DeferStmt:      func() ast.Stmt { return nc.DeferStmt() },
			EmptyStmt:      func() ast.Stmt { return nc.EmptyStmt() },
			ExprStmt:       func() ast.Stmt { return nc.ExprStmt() },
			ForStmt:        func() ast.Stmt { return nc.ForStmt() },
			GoStmt:         func() ast.Stmt { return nc.GoStmt() },
			IfStmt:         func() ast.Stmt { return nc.IfStmt() },
			IncDecStmt:     func() ast.Stmt { return nc.IncDecStmt() },
			LabeledStmt:    func() ast.Stmt { return nc.LabeledStmt() },
			RangeStmt:      func() ast.Stmt { return nc.RangeStmt() },
			ReturnStmt:     func() ast.Stmt { return nc.ReturnStmt() },
			SelectStmt:     func() ast.Stmt { return nc.SelectStmt() },
			SendStmt:       func() ast.Stmt { return nc.SendStmt() },
			SwitchStmt:     func() ast.Stmt { return nc.SwitchStmt() },
			TypeSwitchStmt: func() ast.Stmt { return nc.TypeSwitchStmt() },
		},
		Expression: map[NodeType]func() ast.Expr{
			ArrayType:      func() ast.Expr { return nc.ArrayType() },
			BasicLit:       func() ast.Expr { return nc.BasicLit() },
			BinaryExpr:     func() ast.Expr { return nc.BinaryExpr() },
			CallExpr:       func() ast.Expr { return nc.CallExpr() },
			ChanType:       func() ast.Expr { return nc.ChanType() },
			CompositeLit:   func() ast.Expr { return nc.CompositeLit() },
			Ellipsis:       func() ast.Expr { return nc.Ellipsis() },
			FuncLit:        func() ast.Expr { return nc.FuncLit() },
			FuncType:       func() ast.Expr { return nc.FuncType() },
			Ident:          func() ast.Expr { return nc.Ident() },
			IndexExpr:      func() ast.Expr { return nc.IndexExpr() },
			IndexListExpr:  func() ast.Expr { return nc.IndexListExpr() },
			InterfaceType:  func() ast.Expr { return nc.InterfaceType() },
			KeyValueExpr:   func() ast.Expr { return nc.KeyValueExpr() },
			MapType:        func() ast.Expr { return nc.MapType() },
			ParenExpr:      func() ast.Expr { return nc.ParenExpr() },
			SelectorExpr:   func() ast.Expr { return nc.SelectorExpr() },
			SliceExpr:      func() ast.Expr { return nc.SliceExpr() },
			StarExpr:       func() ast.Expr { return nc.StarExpr() },
			StructType:     func() ast.Expr { return nc.StructType() },
			TypeAssertExpr: func() ast.Expr { return nc.TypeAssertExpr() },
			UnaryExpr:      func() ast.Expr { return nc.UnaryExpr() },
		},
		Specification: map[NodeType]func() ast.Spec{
			ImportSpec: func() ast.Spec { return nc.ImportSpec() },
			TypeSpec:   func() ast.Spec { return nc.TypeSpec() },
			ValueSpec:  func() ast.Spec { return nc.ValueSpec() },
		},
		Declaration: map[NodeType]func() ast.Decl{
			FuncDecl: func() ast.Decl { return nc.FuncDecl() },
			GenDecl:  func() ast.Decl { return nc.GenDecl() },
		},
		TypeDeclaration: map[NodeType]func() ast.Expr{
			ArrayType:     func() ast.Expr { return nc.ArrayType() },
			ChanType:      func() ast.Expr { return nc.ChanType() },
			FuncType:      func() ast.Expr { return nc.FuncType() },
			InterfaceType: func() ast.Expr { return nc.InterfaceType() },
			MapType:       func() ast.Expr { return nc.MapType() },
			StructType:    func() ast.Expr { return nc.StructType() },
			Ident:         func() ast.Expr { return nc.Ident() },
			ParenExpr:     func() ast.Expr { return nc.ParenExpr() },
			SelectorExpr:  func() ast.Expr { return nc.SelectorExpr() },
			StarExpr:      func() ast.Expr { return nc.StarExpr() },
		},
	}

	nc.Dictionary = map[NodeType]any{
		Ident: nc.Ident,

		ArrayType:     nc.ArrayType,
		ChanType:      nc.ChanType,
		FuncType:      nc.FuncType,
		InterfaceType: nc.InterfaceType,
		MapType:       nc.MapType,
		StructType:    nc.StructType,

		AssignStmt:     nc.AssignStmt,
		BlockStmt:      nc.BlockStmt,
		BranchStmt:     nc.BranchStmt,
		DeclStmt:       nc.DeclStmt,
		DeferStmt:      nc.DeferStmt,
		EmptyStmt:      nc.EmptyStmt,
		ExprStmt:       nc.ExprStmt,
		ForStmt:        nc.ForStmt,
		GoStmt:         nc.GoStmt,
		IfStmt:         nc.IfStmt,
		IncDecStmt:     nc.IncDecStmt,
		LabeledStmt:    nc.LabeledStmt,
		RangeStmt:      nc.RangeStmt,
		ReturnStmt:     nc.ReturnStmt,
		SelectStmt:     nc.SelectStmt,
		SendStmt:       nc.SendStmt,
		SwitchStmt:     nc.SwitchStmt,
		TypeSwitchStmt: nc.TypeSwitchStmt,

		BasicLit:     nc.BasicLit,
		CompositeLit: nc.CompositeLit,
		FuncLit:      nc.FuncLit,

		BinaryExpr:     nc.BinaryExpr,
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
		GenDecl:  nc.GenDecl,

		CaseClause: nc.CaseClause,
		CommClause: nc.CommClause,
		Ellipsis:   nc.Ellipsis,

		ImportSpec: nc.ImportSpec,
		TypeSpec:   nc.TypeSpec,
		ValueSpec:  nc.ValueSpec,

		Field:     nc.Field,
		FieldList: nc.FieldList,
	}
	return &nc
}

func (nc *NodeConstructor) Construct(nodeType NodeType) ast.Node {
	return nc.Dictionary[nodeType].(func() any)().(ast.Node)
}

// MARK: Compliying Interface instance

// Chooses a node type that confirms ast.Spec interface; initializes and returns
func (nc *NodeConstructor) Spec() ast.Spec {
	nodeType := *utilities.Pick(nc.Classes[Specification])
	return nc.Dictionaries.Specification[nodeType]()
}

// Chooses a node type that confirms ast.Decl interface; initializes and returns
func (nc *NodeConstructor) Decl() ast.Decl {
	nodeType := *utilities.Pick(nc.Classes[Declaration])
	return nc.Dictionaries.Declaration[nodeType]()
}

// Chooses a node type that confirms ast.Expr interface; initializes and returns
func (nc *NodeConstructor) Expr() ast.Expr {
	nodeType := *utilities.Pick(nc.Classes[Expression])
	return nc.Dictionaries.Expression[nodeType]()
}

// Chooses a node type that confirms ast.Stmt interface; initializes and returns
func (nc *NodeConstructor) Stmt() ast.Stmt {
	nodeType := *utilities.Pick(nc.Classes[Statement])
	return nc.Dictionaries.Statement[nodeType]()
}

// Chooses, initialized and returns an instance of random Type Declaration (Expression)
func (nc *NodeConstructor) Type() ast.Expr {
	nodeType := *utilities.Pick(nc.Classes[TypeDeclaration])
	return nc.Dictionaries.TypeDeclaration[nodeType]()
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

func (nc *NodeConstructor) ArrayType() *ast.ArrayType { // FIXME: // is there any usecase thar is not achievable with a slice but only with a ...T array
	return &ast.ArrayType{Lbrack: token.NoPos, Len: nil, Elt: nc.Expr()}
}

func (nc *NodeConstructor) AssignStmt() *ast.AssignStmt {
	return &ast.AssignStmt{Lhs: []ast.Expr{nc.Expr()}, TokPos: token.NoPos, Tok: *utilities.Pick(tokenConstructor.AccepetedByAssignStmt), Rhs: []ast.Expr{nc.Expr()}}
}

func (nc *NodeConstructor) BasicLit() *ast.BasicLit {
	return (*utilities.Pick([]func() *ast.BasicLit{nc.basicIntegerLiteral, nc.basicStringLiteral, nc.basicFloatLiteral, nc.basicCharacterLiteral}))()
}

func (nc *NodeConstructor) BinaryExpr() *ast.BinaryExpr {
	return &ast.BinaryExpr{X: nc.Expr(), OpPos: token.NoPos, Op: *utilities.Pick(tokenConstructor.AcceptedByBinaryExpr), Y: nc.Expr()}
}

func (nc *NodeConstructor) BlockStmt() *ast.BlockStmt {
	return &ast.BlockStmt{List: []ast.Stmt{nc.Stmt()}, Lbrace: token.NoPos, Rbrace: token.NoPos}
}

func (nc *NodeConstructor) BranchStmt() *ast.BranchStmt {
	return &ast.BranchStmt{TokPos: token.NoPos, Tok: *utilities.Pick(tokenConstructor.AcceptedByBranchStmt), Label: *utilities.Pick(nc.GeneratedBranchLabels)}
}

func (nc *NodeConstructor) CallExpr() *ast.CallExpr { // TODO: function calls with more than 1 arguments
	return &ast.CallExpr{Fun: nc.Expr(), Lparen: token.NoPos, Args: []ast.Expr{nc.Expr()}, Ellipsis: token.NoPos, Rparen: token.NoPos}
}

func (nc *NodeConstructor) CaseClause() *ast.CaseClause {
	return &ast.CaseClause{Case: token.NoPos, List: []ast.Expr{nc.Expr()}, Colon: token.NoPos, Body: []ast.Stmt{nc.Stmt()}}
}

func (nc *NodeConstructor) ChanType() *ast.ChanType {
	return &ast.ChanType{Begin: token.NoPos, Arrow: token.NoPos, Dir: *utilities.Pick([]ast.ChanDir{ast.SEND, ast.RECV}), Value: nc.Type()}
}

func (nc *NodeConstructor) CommClause() *ast.CommClause {
	return &ast.CommClause{Case: token.NoPos, Colon: token.NoPos, Body: []ast.Stmt{nc.Stmt()}}
}

func (nc *NodeConstructor) CompositeLit() *ast.CompositeLit { // TODO: check Incomplete property
	return &ast.CompositeLit{Type: nc.Type(), Lbrace: token.NoPos, Elts: []ast.Expr{nc.Expr()}, Rbrace: token.NoPos, Incomplete: false}
}

func (nc *NodeConstructor) DeclStmt() *ast.DeclStmt { // either with initial value assignment or declaration only
	return &ast.DeclStmt{Decl: nc.GenDecl()}
}

func (nc *NodeConstructor) DeferStmt() *ast.DeferStmt {
	return &ast.DeferStmt{Defer: token.NoPos, Call: nc.CallExpr()}
}

func (nc *NodeConstructor) Ellipsis() *ast.Ellipsis {
	return &ast.Ellipsis{Ellipsis: token.NoPos, Elt: nc.Expr()}
}

func (nc *NodeConstructor) EmptyStmt() *ast.EmptyStmt {
	return &ast.EmptyStmt{Semicolon: token.NoPos, Implicit: false}
}

func (nc *NodeConstructor) ExprStmt() *ast.ExprStmt {
	return &ast.ExprStmt{X: nc.Expr()}
}

func (nc *NodeConstructor) Field() *ast.Field {
	return &ast.Field{Names: []*ast.Ident{nc.Ident()}, Type: nc.Type(), Tag: nil}
}

func (nc *NodeConstructor) FieldList() *ast.FieldList {
	return &ast.FieldList{Opening: token.NoPos, List: []*ast.Field{nc.Field()}, Closing: token.NoPos}
}

func (nc *NodeConstructor) ForStmt() *ast.ForStmt {
	return &ast.ForStmt{For: token.NoPos, Init: nc.Stmt(), Cond: nc.Expr(), Post: nc.Stmt(), Body: nc.BlockStmt()}
}

func (nc *NodeConstructor) FuncDecl() *ast.FuncDecl { // TODO: Consider adding receiver functions (methods)
	return &ast.FuncDecl{Name: nc.generateFunctionName(), Type: nc.FuncType(), Body: nc.BlockStmt()}
}

func (nc *NodeConstructor) FuncLit() *ast.FuncLit { // TODO:
	return &ast.FuncLit{Type: nc.FuncType(), Body: nc.BlockStmt()}
}

func (nc *NodeConstructor) FuncType() *ast.FuncType { // FIXME:
	return &ast.FuncType{Func: token.NoPos, TypeParams: nc.FieldList(), Params: nc.FieldList(), Results: nc.FieldList()}
}

// Produces only "variable" declarations. "import", "constant", "type" declarations are ignored.
func (nc *NodeConstructor) GenDecl() *ast.GenDecl {
	return &ast.GenDecl{TokPos: token.NoPos, Tok: token.VAR, Lparen: token.NoPos, Rparen: token.NoPos, Specs: []ast.Spec{nc.ValueSpec()}}
}

func (nc *NodeConstructor) GoStmt() *ast.GoStmt {
	return &ast.GoStmt{Go: token.NoPos, Call: nc.CallExpr()}
}

func (nc *NodeConstructor) Ident() *ast.Ident {
	return nc.generateVariableName()
}

func (nc *NodeConstructor) IfStmt() *ast.IfStmt {
	return &ast.IfStmt{If: token.NoPos, Init: nil, Cond: nc.Expr(), Body: &ast.BlockStmt{Lbrace: token.NoPos, List: []ast.Stmt{nc.Stmt()}, Rbrace: token.NoPos}, Else: nil}
}

func (nc *NodeConstructor) ImportSpec() *ast.ImportSpec { // TODO: Store imported packages for later use
	return &ast.ImportSpec{Name: nil, Path: &ast.BasicLit{ValuePos: token.NoPos, Kind: token.STRING, Value: *utilities.Pick(nc.AllowedPackagesToImport)}, EndPos: token.NoPos}
}

func (nc *NodeConstructor) IncDecStmt() *ast.IncDecStmt {
	return &ast.IncDecStmt{X: nc.Expr(), TokPos: token.NoPos, Tok: *utilities.Pick(tokenConstructor.AccepetedByIncDecStmt)}
}

func (nc *NodeConstructor) IndexExpr() *ast.IndexExpr {
	return &ast.IndexExpr{X: nc.Expr(), Lbrack: token.NoPos, Index: nc.Expr(), Rbrack: token.NoPos}
}

func (nc *NodeConstructor) IndexListExpr() *ast.IndexListExpr { // TODO: Multi-dimensional arrays
	return &ast.IndexListExpr{X: nc.Expr(), Lbrack: token.NoPos, Indices: []ast.Expr{nc.Expr()}, Rbrack: token.NoPos}
}

func (nc *NodeConstructor) InterfaceType() *ast.InterfaceType {
	return &ast.InterfaceType{Interface: token.NoPos, Methods: nc.FieldList(), Incomplete: false}
}

func (nc *NodeConstructor) KeyValueExpr() *ast.KeyValueExpr {
	return &ast.KeyValueExpr{Key: nc.Expr(), Colon: token.NoPos, Value: nc.Expr()}
}

func (nc *NodeConstructor) LabeledStmt() *ast.LabeledStmt {
	return &ast.LabeledStmt{Label: nc.generateBranchLabel(), Colon: token.NoPos, Stmt: nc.Stmt()}
}

func (nc *NodeConstructor) MapType() *ast.MapType {
	return &ast.MapType{Map: token.NoPos, Key: nc.Type(), Value: nc.Type()}
}

func (nc *NodeConstructor) ParenExpr() *ast.ParenExpr {
	return &ast.ParenExpr{Lparen: token.NoPos, X: nc.Expr(), Rparen: token.NoPos}
}

func (nc *NodeConstructor) RangeStmt() *ast.RangeStmt {
	return &ast.RangeStmt{For: token.NoPos, Key: nc.Expr(), Value: nc.Expr(), TokPos: token.NoPos, Tok: *utilities.Pick(tokenConstructor.AcceptedByRangeStmt), X: nc.Expr(), Body: nc.BlockStmt()}
}

func (nc *NodeConstructor) ReturnStmt() *ast.ReturnStmt { // TODO: multiple return values
	return &ast.ReturnStmt{Return: token.NoPos, Results: []ast.Expr{nc.Expr()}}
}

func (nc *NodeConstructor) SelectorExpr() *ast.SelectorExpr { // FIXME: randomly produced X and Sel values will never work, maybe choose from imported libraries' exported functions, or previously declared struct instances that has methods
	return &ast.SelectorExpr{X: nc.Expr(), Sel: &ast.Ident{}}
}

func (nc *NodeConstructor) SelectStmt() *ast.SelectStmt {
	return &ast.SelectStmt{
		Select: token.NoPos,
		Body:   &ast.BlockStmt{}, // NOTE: nc.CommClause() ?
	}
}

func (nc *NodeConstructor) SendStmt() *ast.SendStmt {
	return &ast.SendStmt{Chan: nc.Expr(), Arrow: token.NoPos, Value: nc.Expr()}
}

func (nc *NodeConstructor) SliceExpr() *ast.SliceExpr {
	return &ast.SliceExpr{X: nc.Expr(), Lbrack: token.NoPos, Low: nc.Expr(), High: nc.Expr(), Max: nil, Slice3: false, Rbrack: token.NoPos}
}

func (nc *NodeConstructor) StarExpr() *ast.StarExpr {
	return &ast.StarExpr{Star: token.NoPos, X: nc.Expr()}
}

func (nc *NodeConstructor) StructType() *ast.StructType {
	return &ast.StructType{Struct: token.NoPos, Fields: nc.FieldList(), Incomplete: false}
}

func (nc *NodeConstructor) SwitchStmt() *ast.SwitchStmt {
	return &ast.SwitchStmt{Switch: token.NoPos, Init: nc.Stmt(), Tag: nil, Body: nc.BlockStmt()}
}

func (nc *NodeConstructor) TypeAssertExpr() *ast.TypeAssertExpr {
	return &ast.TypeAssertExpr{
		X:      nc.Expr(),
		Lparen: token.NoPos,
		Type:   nc.InterfaceType(),
		Rparen: token.NoPos,
	}
}

func (nc *NodeConstructor) TypeSpec() *ast.TypeSpec {
	return &ast.TypeSpec{
		Doc:        &ast.CommentGroup{},
		Name:       &ast.Ident{},
		TypeParams: &ast.FieldList{},
		Assign:     token.NoPos,
		Type:       nil,
		Comment:    &ast.CommentGroup{},
	}
}

func (nc *NodeConstructor) TypeSwitchStmt() *ast.TypeSwitchStmt {
	return &ast.TypeSwitchStmt{
		Switch: token.NoPos,
		Init:   nil,
		Assign: nil,
		Body:   &ast.BlockStmt{},
	}
}

func (nc *NodeConstructor) UnaryExpr() *ast.UnaryExpr {
	return &ast.UnaryExpr{
		OpPos: token.NoPos,
		Op:    *utilities.Pick(tokenConstructor.AcceptedByUnaryExpr),
		X:     nil,
	}
}

// Returns ValueSpec which is list of pairs of variable names and values to assign
func (nc *NodeConstructor) ValueSpec() *ast.ValueSpec {
	return &ast.ValueSpec{Names: []*ast.Ident{nc.generateVariableName()}, Values: []ast.Expr{nc.Expr()}}
}

var nodeConstructor = NewNodeConstructor()

// MARK: Ideas

// To use in special nodes like SliceExpr.
func (nc *NodeConstructor) ExprInt() {}
