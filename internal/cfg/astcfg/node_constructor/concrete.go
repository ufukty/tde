package node_construtor

import (
	"go/ast"
	"go/token"
	"tde/internal/utilities"
)

func ArrayType(counter int) *ast.ArrayType {
	// FIXME: // is there any usecase thar is not achievable with a slice but only with a ...T array
	if counter == 0 {
		return nil

	}
	return &ast.ArrayType{
		Lbrack: token.NoPos,
		Len:    nil,
		Elt:    Expr(counter - 1),
	}
}

func AssignStmt(counter int) *ast.AssignStmt {
	if counter == 0 {
		return nil
	}
	return &ast.AssignStmt{
		Lhs:    []ast.Expr{Expr(counter - 1)},
		TokPos: token.NoPos,
		Tok:    *utilities.Pick(tokenConstructor.AccepetedByAssignStmt),
		Rhs:    []ast.Expr{Expr(counter - 1)},
	}
}

func BasicLit(counter int) *ast.BasicLit {
	if counter == 0 {
		return nil
	}
	return (*utilities.Pick([]func() *ast.BasicLit{
		basicIntegerLiteral,
		basicStringLiteral,
		basicFloatLiteral,
		basicCharacterLiteral,
	}))()
}

func BinaryExpr(counter int) *ast.BinaryExpr {
	if counter == 0 {
		return nil
	}
	return &ast.BinaryExpr{
		X:     Expr(counter - 1),
		OpPos: token.NoPos,
		Op:    *utilities.Pick(tokenConstructor.AcceptedByBinaryExpr),
		Y:     Expr(counter - 1),
	}
}

func BlockStmt(counter int) *ast.BlockStmt {
	if counter == 0 {
		return nil
	}
	return &ast.BlockStmt{
		List: []ast.Stmt{
			Stmt(counter - 1),
		},
		Lbrace: token.NoPos,
		Rbrace: token.NoPos,
	}
}

func BranchStmt(counter int) *ast.BranchStmt {
	if counter == 0 {
		return nil
	}
	return &ast.BranchStmt{
		TokPos: token.NoPos,
		Tok:    *utilities.Pick(tokenConstructor.AcceptedByBranchStmt),
		Label:  *utilities.Pick(GeneratedBranchLabels),
	}
}

func CallExpr(counter int) *ast.CallExpr {
	// TODO: function calls with more than 1 arguments
	if counter == 0 {
		return nil
	}
	return &ast.CallExpr{
		Fun:      Expr(counter - 1),
		Lparen:   token.NoPos,
		Args:     []ast.Expr{Expr(counter - 1)},
		Ellipsis: token.NoPos,
		Rparen:   token.NoPos,
	}
}

func CaseClause(counter int) *ast.CaseClause {
	if counter == 0 {
		return nil
	}
	return &ast.CaseClause{
		Case:  token.NoPos,
		List:  []ast.Expr{Expr(counter - 1)},
		Colon: token.NoPos,
		Body: []ast.Stmt{
			Stmt(counter - 1),
		},
	}
}

func ChanType(counter int) *ast.ChanType {
	if counter == 0 {
		return nil
	}
	return &ast.ChanType{
		Begin: token.NoPos,
		Arrow: token.NoPos,
		Dir: *utilities.Pick([]ast.ChanDir{
			ast.SEND,
			ast.RECV,
		}),
		Value: Type(counter - 1),
	}
}

func CommClause(counter int) *ast.CommClause {
	if counter == 0 {
		return nil
	}
	return &ast.CommClause{
		Case:  token.NoPos,
		Colon: token.NoPos,
		Body: []ast.Stmt{
			Stmt(counter - 1),
		},
	}
}

func CompositeLit(counter int) *ast.CompositeLit {
	// TODO: check Incomplete property
	if counter == 0 {
		return nil
	}
	return &ast.CompositeLit{
		Type:       Type(counter - 1),
		Lbrace:     token.NoPos,
		Elts:       []ast.Expr{Expr(counter - 1)},
		Rbrace:     token.NoPos,
		Incomplete: false,
	}
}

func DeclStmt(counter int) *ast.DeclStmt {
	// either with initial value assignment or declaration only
	if counter == 0 {
		return nil
	}
	return &ast.DeclStmt{
		Decl: GenDecl(counter - 1),
	}
}

func DeferStmt(counter int) *ast.DeferStmt {
	if counter == 0 {
		return nil
	}
	return &ast.DeferStmt{
		Defer: token.NoPos,
		Call:  CallExpr(counter - 1),
	}
}

func Ellipsis(counter int) *ast.Ellipsis {
	if counter == 0 {
		return nil
	}
	return &ast.Ellipsis{
		Ellipsis: token.NoPos,
		Elt:      Expr(counter - 1),
	}
}

func EmptyStmt(counter int) *ast.EmptyStmt {
	if counter == 0 {
		return nil
	}
	return &ast.EmptyStmt{
		Semicolon: token.NoPos,
		Implicit:  false,
	}
}

func ExprStmt(counter int) *ast.ExprStmt {
	if counter == 0 {
		return nil
	}
	return &ast.ExprStmt{
		X: Expr(counter - 1),
	}
}

func Field(counter int) *ast.Field {
	if counter == 0 {
		return nil
	}
	return &ast.Field{
		Names: []*ast.Ident{
			Ident(counter - 1),
		},
		Type: Type(counter - 1),
		Tag:  nil,
	}
}

func FieldList(counter int) *ast.FieldList {
	if counter == 0 {
		return nil
	}
	return &ast.FieldList{
		Opening: token.NoPos,
		List: []*ast.Field{
			Field(counter - 1),
		},
		Closing: token.NoPos,
	}
}

func ForStmt(counter int) *ast.ForStmt {
	if counter == 0 {
		return nil
	}
	return &ast.ForStmt{
		For:  token.NoPos,
		Init: Stmt(counter - 1),
		Cond: Expr(counter - 1),
		Post: Stmt(counter - 1),
		Body: BlockStmt(counter - 1),
	}
}

func FuncDecl(counter int) *ast.FuncDecl {
	// TODO: Consider adding receiver functions (methods)
	if counter == 0 {
		return nil
	}
	return &ast.FuncDecl{
		Name: generateFunctionName(),
		Type: FuncType(counter - 1),
		Body: BlockStmt(counter - 1),
	}
}

func FuncLit(counter int) *ast.FuncLit {
	// TODO:
	if counter == 0 {
		return nil
	}
	return &ast.FuncLit{
		Type: FuncType(counter - 1),
		Body: BlockStmt(counter - 1),
	}
}

func FuncType(counter int) *ast.FuncType {
	// FIXME:
	if counter == 0 {
		return nil
	}
	return &ast.FuncType{
		Func:       token.NoPos,
		TypeParams: FieldList(counter - 1),
		Params:     FieldList(counter - 1),
		Results:    FieldList(counter - 1),
	}
}

// Produces only "variable" declarations. "import", "constant", "type" declarations are ignored.
func GenDecl(counter int) *ast.GenDecl {
	if counter == 0 {
		return nil
	}
	return &ast.GenDecl{
		TokPos: token.NoPos,
		Tok:    token.VAR,
		Lparen: token.NoPos,
		Rparen: token.NoPos,
		Specs: []ast.Spec{
			ValueSpec(counter - 1),
		},
	}
}

func GoStmt(counter int) *ast.GoStmt {
	if counter == 0 {
		return nil
	}
	return &ast.GoStmt{
		Go:   token.NoPos,
		Call: CallExpr(counter - 1),
	}
}

func Ident(counter int) *ast.Ident {
	if counter == 0 {
		return nil
	}
	return generateVariableName()
}

// only valid values are types such int, float, string, bool
func IdentType(counter int) *ast.Ident {
	if counter == 0 {
		return nil
	}
	return ast.NewIdent(*utilities.Pick([]string{"int", "float", "string", "bool"}))
}

func IfStmt(counter int) *ast.IfStmt {
	if counter == 0 {
		return nil
	}
	return &ast.IfStmt{
		If:   token.NoPos,
		Init: nil,
		Cond: Expr(counter - 1),
		Body: &ast.BlockStmt{
			Lbrace: token.NoPos,
			List:   []ast.Stmt{Stmt(counter - 1)},
			Rbrace: token.NoPos,
		},
		Else: nil,
	}
}

func ImportSpec(counter int) *ast.ImportSpec {
	// TODO: Store imported packages for later use
	if counter == 0 {
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

func IncDecStmt(counter int) *ast.IncDecStmt {
	if counter == 0 {
		return nil
	}
	return &ast.IncDecStmt{
		X:      Expr(counter - 1),
		TokPos: token.NoPos,
		Tok:    *utilities.Pick(tokenConstructor.AccepetedByIncDecStmt),
	}
}

func IndexExpr(counter int) *ast.IndexExpr {
	if counter == 0 {
		return nil
	}
	return &ast.IndexExpr{
		X:      Expr(counter - 1),
		Lbrack: token.NoPos,
		Index:  Expr(counter - 1),
		Rbrack: token.NoPos,
	}
}

func IndexListExpr(counter int) *ast.IndexListExpr {
	// TODO: Multi-dimensional arrays
	if counter == 0 {
		return nil
	}
	return &ast.IndexListExpr{
		X:       Expr(counter - 1),
		Lbrack:  token.NoPos,
		Indices: []ast.Expr{Expr(counter - 1)},
		Rbrack:  token.NoPos,
	}
}

func InterfaceType(counter int) *ast.InterfaceType {
	if counter == 0 {
		return nil
	}
	return &ast.InterfaceType{
		Interface:  token.NoPos,
		Methods:    FieldList(counter - 1),
		Incomplete: false,
	}
}

func KeyValueExpr(counter int) *ast.KeyValueExpr {
	if counter == 0 {
		return nil
	}
	return &ast.KeyValueExpr{
		Key:   Expr(counter - 1),
		Colon: token.NoPos,
		Value: Expr(counter - 1),
	}
}

func LabeledStmt(counter int) *ast.LabeledStmt {
	if counter == 0 {
		return nil
	}
	return &ast.LabeledStmt{
		Label: generateBranchLabel(),
		Colon: token.NoPos,
		Stmt:  Stmt(counter - 1),
	}
}

func MapType(counter int) *ast.MapType {
	if counter == 0 {
		return nil
	}
	return &ast.MapType{
		Map:   token.NoPos,
		Key:   Type(counter - 1),
		Value: Type(counter - 1),
	}
}

func ParenExpr(counter int) *ast.ParenExpr {
	if counter == 0 {
		return nil
	}
	return &ast.ParenExpr{
		Lparen: token.NoPos,
		X:      Expr(counter - 1),
		Rparen: token.NoPos,
	}
}

func RangeStmt(counter int) *ast.RangeStmt {
	if counter == 0 {
		return nil
	}
	return &ast.RangeStmt{
		For:    token.NoPos,
		Key:    Expr(counter - 1),
		Value:  Expr(counter - 1),
		TokPos: token.NoPos,
		Tok:    *utilities.Pick(tokenConstructor.AcceptedByRangeStmt),
		X:      Expr(counter - 1),
		Body:   BlockStmt(counter - 1),
	}
}

func ReturnStmt(counter int) *ast.ReturnStmt {
	// TODO: multiple return values
	if counter == 0 {
		return nil
	}
	return &ast.ReturnStmt{
		Return:  token.NoPos,
		Results: []ast.Expr{Expr(counter - 1)},
	}
}

func SelectorExpr(counter int) *ast.SelectorExpr {
	// FIXME: randomly produced X and Sel values will never work, maybe choose from imported libraries' exported functions, or previously declared struct instances that has methods
	if counter == 0 {
		return nil
	}
	return &ast.SelectorExpr{
		X:   Expr(counter - 1),
		Sel: &ast.Ident{},
	}
}

func SelectStmt(counter int) *ast.SelectStmt {
	if counter == 0 {
		return nil
	}
	return &ast.SelectStmt{
		Select: token.NoPos,
		Body:   &ast.BlockStmt{}, // NOTE: CommClause() ?
	}
}

func SendStmt(counter int) *ast.SendStmt {
	if counter == 0 {
		return nil
	}
	return &ast.SendStmt{
		Chan:  Expr(counter - 1),
		Arrow: token.NoPos,
		Value: Expr(counter - 1),
	}
}

func SliceExpr(counter int) *ast.SliceExpr {
	if counter == 0 {
		return nil
	}
	return &ast.SliceExpr{
		X:      Expr(counter - 1),
		Lbrack: token.NoPos,
		Low:    Expr(counter - 1),
		High:   Expr(counter - 1),
		Max:    nil,
		Slice3: false,
		Rbrack: token.NoPos,
	}
}

func StarExpr(counter int) *ast.StarExpr {
	if counter == 0 {
		return nil
	}
	return &ast.StarExpr{
		Star: token.NoPos,
		X:    Expr(counter - 1),
	}
}

func StructType(counter int) *ast.StructType {
	if counter == 0 {
		return nil
	}
	return &ast.StructType{
		Struct:     token.NoPos,
		Fields:     FieldList(counter - 1),
		Incomplete: false,
	}
}

func SwitchStmt(counter int) *ast.SwitchStmt {
	if counter == 0 {
		return nil
	}
	return &ast.SwitchStmt{
		Switch: token.NoPos,
		Init:   Stmt(counter - 1),
		Tag:    nil,
		Body:   BlockStmt(counter - 1),
	}
}

func TypeAssertExpr(counter int) *ast.TypeAssertExpr {
	if counter == 0 {
		return nil
	}
	return &ast.TypeAssertExpr{
		X:      Expr(counter - 1),
		Lparen: token.NoPos,
		Type:   InterfaceType(counter - 1),
		Rparen: token.NoPos,
	}
}

// rel. type keyword
func TypeSpec(counter int) *ast.TypeSpec {
	if counter == 0 {
		return nil
	}
	return &ast.TypeSpec{
		Doc:        nil,
		Name:       Ident(counter - 1),
		TypeParams: FieldList(counter - 1),
		Assign:     token.NoPos,
		Type:       nil,
		Comment:    nil,
	}
}

func TypeSwitchStmt(counter int) *ast.TypeSwitchStmt {
	if counter == 0 {
		return nil
	}
	return &ast.TypeSwitchStmt{
		Switch: token.NoPos,
		Init:   Stmt(counter - 1),
		Assign: Stmt(counter - 1),
		Body:   BlockStmt(counter - 1),
	}
}

func UnaryExpr(counter int) *ast.UnaryExpr {
	if counter == 0 {
		return nil
	}
	return &ast.UnaryExpr{
		OpPos: token.NoPos,
		Op:    *utilities.Pick(tokenConstructor.AcceptedByUnaryExpr),
		X:     Expr(counter - 1),
	}
}

// Returns ValueSpec which is list of pairs of variable names and values to assign
func ValueSpec(counter int) *ast.ValueSpec {
	if counter == 0 {
		return nil
	}
	return &ast.ValueSpec{
		Names: []*ast.Ident{
			generateVariableName(),
		},
		Values: []ast.Expr{Expr(counter - 1)},
	}
}
