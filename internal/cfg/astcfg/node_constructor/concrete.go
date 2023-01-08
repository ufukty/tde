package node_construtor

import (
	"go/ast"
	"go/token"
	"tde/internal/utilities"
)

func ArrayType(limit int) *ast.ArrayType {
	// FIXME: // is there any usecase thar is not achievable with a slice but only with a ...T array
	if limit == 0 {
		return nil

	}
	return &ast.ArrayType{
		Lbrack: token.NoPos,
		Len:    nil,
		Elt:    Expr(limit - 1),
	}
}

func AssignStmt(limit int) *ast.AssignStmt {
	if limit == 0 {
		return nil
	}
	return &ast.AssignStmt{
		Lhs:    []ast.Expr{Expr(limit - 1)},
		TokPos: token.NoPos,
		Tok:    *utilities.Pick(tokenConstructor.AccepetedByAssignStmt),
		Rhs:    []ast.Expr{Expr(limit - 1)},
	}
}

func BasicLit(limit int) *ast.BasicLit {
	if limit == 0 {
		return nil
	}
	return (*utilities.Pick([]func() *ast.BasicLit{
		basicIntegerLiteral,
		basicStringLiteral,
		basicFloatLiteral,
		basicCharacterLiteral,
	}))()
}

func BinaryExpr(limit int) *ast.BinaryExpr {
	if limit == 0 {
		return nil
	}
	return &ast.BinaryExpr{
		X:     Expr(limit - 1),
		OpPos: token.NoPos,
		Op:    *utilities.Pick(tokenConstructor.AcceptedByBinaryExpr),
		Y:     Expr(limit - 1),
	}
}

func BlockStmt(limit int) *ast.BlockStmt {
	if limit == 0 {
		return nil
	}
	return &ast.BlockStmt{
		List: []ast.Stmt{
			Stmt(limit - 1),
		},
		Lbrace: token.NoPos,
		Rbrace: token.NoPos,
	}
}

func BranchStmt(limit int) *ast.BranchStmt {
	if limit == 0 {
		return nil
	}
	return &ast.BranchStmt{
		TokPos: token.NoPos,
		Tok:    *utilities.Pick(tokenConstructor.AcceptedByBranchStmt),
		Label:  *utilities.Pick(GeneratedBranchLabels),
	}
}

func CallExpr(limit int) *ast.CallExpr {
	// TODO: function calls with more than 1 arguments
	if limit == 0 {
		return nil
	}
	return &ast.CallExpr{
		Fun:      Expr(limit - 1),
		Lparen:   token.NoPos,
		Args:     []ast.Expr{Expr(limit - 1)},
		Ellipsis: token.NoPos,
		Rparen:   token.NoPos,
	}
}

func CaseClause(limit int) *ast.CaseClause {
	if limit == 0 {
		return nil
	}
	return &ast.CaseClause{
		Case:  token.NoPos,
		List:  []ast.Expr{Expr(limit - 1)},
		Colon: token.NoPos,
		Body: []ast.Stmt{
			Stmt(limit - 1),
		},
	}
}

func ChanType(limit int) *ast.ChanType {
	if limit == 0 {
		return nil
	}
	return &ast.ChanType{
		Begin: token.NoPos,
		Arrow: token.NoPos,
		Dir: *utilities.Pick([]ast.ChanDir{
			ast.SEND,
			ast.RECV,
		}),
		Value: Type(limit - 1),
	}
}

func CommClause(limit int) *ast.CommClause {
	if limit == 0 {
		return nil
	}
	return &ast.CommClause{
		Case:  token.NoPos,
		Colon: token.NoPos,
		Body: []ast.Stmt{
			Stmt(limit - 1),
		},
	}
}

func CompositeLit(limit int) *ast.CompositeLit {
	// TODO: check Incomplete property
	if limit == 0 {
		return nil
	}
	return &ast.CompositeLit{
		Type:       Type(limit - 1),
		Lbrace:     token.NoPos,
		Elts:       []ast.Expr{Expr(limit - 1)},
		Rbrace:     token.NoPos,
		Incomplete: false,
	}
}

func DeclStmt(limit int) *ast.DeclStmt {
	// either with initial value assignment or declaration only
	if limit == 0 {
		return nil
	}
	return &ast.DeclStmt{
		Decl: GenDecl(limit - 1),
	}
}

func DeferStmt(limit int) *ast.DeferStmt {
	if limit == 0 {
		return nil
	}
	return &ast.DeferStmt{
		Defer: token.NoPos,
		Call:  CallExpr(limit - 1),
	}
}

func Ellipsis(limit int) *ast.Ellipsis {
	if limit == 0 {
		return nil
	}
	return &ast.Ellipsis{
		Ellipsis: token.NoPos,
		Elt:      Expr(limit - 1),
	}
}

func EmptyStmt(limit int) *ast.EmptyStmt {
	if limit == 0 {
		return nil
	}
	return &ast.EmptyStmt{
		Semicolon: token.NoPos,
		Implicit:  false,
	}
}

func ExprStmt(limit int) *ast.ExprStmt {
	if limit == 0 {
		return nil
	}
	return &ast.ExprStmt{
		X: Expr(limit - 1),
	}
}

func Field(limit int) *ast.Field {
	if limit == 0 {
		return nil
	}
	return &ast.Field{
		Names: []*ast.Ident{
			Ident(limit - 1),
		},
		Type: Type(limit - 1),
		Tag:  nil,
	}
}

func FieldList(limit int) *ast.FieldList {
	if limit == 0 {
		return nil
	}
	return &ast.FieldList{
		Opening: token.NoPos,
		List: []*ast.Field{
			Field(limit - 1),
		},
		Closing: token.NoPos,
	}
}

func ForStmt(limit int) *ast.ForStmt {
	if limit == 0 {
		return nil
	}
	return &ast.ForStmt{
		For:  token.NoPos,
		Init: Stmt(limit - 1),
		Cond: Expr(limit - 1),
		Post: Stmt(limit - 1),
		Body: BlockStmt(limit - 1),
	}
}

func FuncDecl(limit int) *ast.FuncDecl {
	// TODO: Consider adding receiver functions (methods)
	if limit == 0 {
		return nil
	}
	return &ast.FuncDecl{
		Name: generateFunctionName(),
		Type: FuncType(limit - 1),
		Body: BlockStmt(limit - 1),
	}
}

func FuncLit(limit int) *ast.FuncLit {
	// TODO:
	if limit == 0 {
		return nil
	}
	return &ast.FuncLit{
		Type: FuncType(limit - 1),
		Body: BlockStmt(limit - 1),
	}
}

func FuncType(limit int) *ast.FuncType {
	// FIXME:
	if limit == 0 {
		return nil
	}
	return &ast.FuncType{
		Func:       token.NoPos,
		TypeParams: FieldList(limit - 1),
		Params:     FieldList(limit - 1),
		Results:    FieldList(limit - 1),
	}
}

// Produces only "variable" declarations. "import", "constant", "type" declarations are ignored.
func GenDecl(limit int) *ast.GenDecl {
	if limit == 0 {
		return nil
	}
	return &ast.GenDecl{
		TokPos: token.NoPos,
		Tok:    token.VAR,
		Lparen: token.NoPos,
		Rparen: token.NoPos,
		Specs: []ast.Spec{
			ValueSpec(limit - 1),
		},
	}
}

func GoStmt(limit int) *ast.GoStmt {
	if limit == 0 {
		return nil
	}
	return &ast.GoStmt{
		Go:   token.NoPos,
		Call: CallExpr(limit - 1),
	}
}

func Ident(limit int) *ast.Ident {
	if limit == 0 {
		return nil
	}
	return generateVariableName()
}

// only valid values are types such int, float, string, bool
func IdentType(limit int) *ast.Ident {
	if limit == 0 {
		return nil
	}
	return ast.NewIdent(*utilities.Pick([]string{"int", "float", "string", "bool"}))
}

func IfStmt(limit int) *ast.IfStmt {
	if limit == 0 {
		return nil
	}
	return &ast.IfStmt{
		If:   token.NoPos,
		Init: nil,
		Cond: Expr(limit - 1),
		Body: &ast.BlockStmt{
			Lbrace: token.NoPos,
			List:   []ast.Stmt{Stmt(limit - 1)},
			Rbrace: token.NoPos,
		},
		Else: nil,
	}
}

func ImportSpec(limit int) *ast.ImportSpec {
	// TODO: Store imported packages for later use
	if limit == 0 {
		return nil
	}
	return &ast.ImportSpec{
		Name: nil,
		Path: &ast.BasicLit{
			ValuePos: token.NoPos,
			Kind:     token.STRING,
			Value:    *utilities.Pick(AllowedPackagesToImport),
		},
		EndPos: token.NoPos,
	}
}

func IncDecStmt(limit int) *ast.IncDecStmt {
	if limit == 0 {
		return nil
	}
	return &ast.IncDecStmt{
		X:      Expr(limit - 1),
		TokPos: token.NoPos,
		Tok:    *utilities.Pick(tokenConstructor.AccepetedByIncDecStmt),
	}
}

func IndexExpr(limit int) *ast.IndexExpr {
	if limit == 0 {
		return nil
	}
	return &ast.IndexExpr{
		X:      Expr(limit - 1),
		Lbrack: token.NoPos,
		Index:  Expr(limit - 1),
		Rbrack: token.NoPos,
	}
}

func IndexListExpr(limit int) *ast.IndexListExpr {
	// TODO: Multi-dimensional arrays
	if limit == 0 {
		return nil
	}
	return &ast.IndexListExpr{
		X:       Expr(limit - 1),
		Lbrack:  token.NoPos,
		Indices: []ast.Expr{Expr(limit - 1)},
		Rbrack:  token.NoPos,
	}
}

func InterfaceType(limit int) *ast.InterfaceType {
	if limit == 0 {
		return nil
	}
	return &ast.InterfaceType{
		Interface:  token.NoPos,
		Methods:    FieldList(limit - 1),
		Incomplete: false,
	}
}

func KeyValueExpr(limit int) *ast.KeyValueExpr {
	if limit == 0 {
		return nil
	}
	return &ast.KeyValueExpr{
		Key:   Expr(limit - 1),
		Colon: token.NoPos,
		Value: Expr(limit - 1),
	}
}

func LabeledStmt(limit int) *ast.LabeledStmt {
	if limit == 0 {
		return nil
	}
	return &ast.LabeledStmt{
		Label: generateBranchLabel(),
		Colon: token.NoPos,
		Stmt:  Stmt(limit - 1),
	}
}

func MapType(limit int) *ast.MapType {
	if limit == 0 {
		return nil
	}
	return &ast.MapType{
		Map:   token.NoPos,
		Key:   Type(limit - 1),
		Value: Type(limit - 1),
	}
}

func ParenExpr(limit int) *ast.ParenExpr {
	if limit == 0 {
		return nil
	}
	return &ast.ParenExpr{
		Lparen: token.NoPos,
		X:      Expr(limit - 1),
		Rparen: token.NoPos,
	}
}

func RangeStmt(limit int) *ast.RangeStmt {
	if limit == 0 {
		return nil
	}
	return &ast.RangeStmt{
		For:    token.NoPos,
		Key:    Expr(limit - 1),
		Value:  Expr(limit - 1),
		TokPos: token.NoPos,
		Tok:    *utilities.Pick(tokenConstructor.AcceptedByRangeStmt),
		X:      Expr(limit - 1),
		Body:   BlockStmt(limit - 1),
	}
}

func ReturnStmt(limit int) *ast.ReturnStmt {
	// TODO: multiple return values
	if limit == 0 {
		return nil
	}
	return &ast.ReturnStmt{
		Return:  token.NoPos,
		Results: []ast.Expr{Expr(limit - 1)},
	}
}

func SelectorExpr(limit int) *ast.SelectorExpr {
	// FIXME: randomly produced X and Sel values will never work, maybe choose from imported libraries' exported functions, or previously declared struct instances that has methods
	if limit == 0 {
		return nil
	}
	return &ast.SelectorExpr{
		X:   Expr(limit - 1),
		Sel: &ast.Ident{},
	}
}

func SelectStmt(limit int) *ast.SelectStmt {
	if limit == 0 {
		return nil
	}
	return &ast.SelectStmt{
		Select: token.NoPos,
		Body:   &ast.BlockStmt{}, // NOTE: CommClause() ?
	}
}

func SendStmt(limit int) *ast.SendStmt {
	if limit == 0 {
		return nil
	}
	return &ast.SendStmt{
		Chan:  Expr(limit - 1),
		Arrow: token.NoPos,
		Value: Expr(limit - 1),
	}
}

func SliceExpr(limit int) *ast.SliceExpr {
	if limit == 0 {
		return nil
	}
	return &ast.SliceExpr{
		X:      Expr(limit - 1),
		Lbrack: token.NoPos,
		Low:    Expr(limit - 1),
		High:   Expr(limit - 1),
		Max:    nil,
		Slice3: false,
		Rbrack: token.NoPos,
	}
}

func StarExpr(limit int) *ast.StarExpr {
	if limit == 0 {
		return nil
	}
	return &ast.StarExpr{
		Star: token.NoPos,
		X:    Expr(limit - 1),
	}
}

func StructType(limit int) *ast.StructType {
	if limit == 0 {
		return nil
	}
	return &ast.StructType{
		Struct:     token.NoPos,
		Fields:     FieldList(limit - 1),
		Incomplete: false,
	}
}

func SwitchStmt(limit int) *ast.SwitchStmt {
	if limit == 0 {
		return nil
	}
	return &ast.SwitchStmt{
		Switch: token.NoPos,
		Init:   Stmt(limit - 1),
		Tag:    nil,
		Body:   BlockStmt(limit - 1),
	}
}

func TypeAssertExpr(limit int) *ast.TypeAssertExpr {
	if limit == 0 {
		return nil
	}
	return &ast.TypeAssertExpr{
		X:      Expr(limit - 1),
		Lparen: token.NoPos,
		Type:   InterfaceType(limit - 1),
		Rparen: token.NoPos,
	}
}

// rel. type keyword
func TypeSpec(limit int) *ast.TypeSpec {
	if limit == 0 {
		return nil
	}
	return &ast.TypeSpec{
		Doc:        nil,
		Name:       Ident(limit - 1),
		TypeParams: FieldList(limit - 1),
		Assign:     token.NoPos,
		Type:       nil,
		Comment:    nil,
	}
}

func TypeSwitchStmt(limit int) *ast.TypeSwitchStmt {
	if limit == 0 {
		return nil
	}
	return &ast.TypeSwitchStmt{
		Switch: token.NoPos,
		Init:   Stmt(limit - 1),
		Assign: Stmt(limit - 1),
		Body:   BlockStmt(limit - 1),
	}
}

func UnaryExpr(limit int) *ast.UnaryExpr {
	if limit == 0 {
		return nil
	}
	return &ast.UnaryExpr{
		OpPos: token.NoPos,
		Op:    *utilities.Pick(tokenConstructor.AcceptedByUnaryExpr),
		X:     Expr(limit - 1),
	}
}

// Returns ValueSpec which is list of pairs of variable names and values to assign
func ValueSpec(limit int) *ast.ValueSpec {
	if limit == 0 {
		return nil
	}
	return &ast.ValueSpec{
		Names: []*ast.Ident{
			generateVariableName(),
		},
		Values: []ast.Expr{Expr(limit - 1)},
	}
}
