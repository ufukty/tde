package node_constructor

import (
	"tde/internal/cfg/astcfg/context"
	"tde/internal/utilities"

	"go/ast"
	"go/token"
)

func ArrayType(ctx context.Context, limit int) *ast.ArrayType {
	// FIXME: // is there any usecase thar is not achievable with a slice but only with a ...T array
	if limit == 0 {
		return nil

	}
	return &ast.ArrayType{
		Lbrack: token.NoPos,
		Len:    nil,
		Elt:    Expr(ctx, limit-1),
	}
}

func AssignStmt(ctx context.Context, limit int) *ast.AssignStmt {
	if limit == 0 {
		return nil
	}
	return &ast.AssignStmt{
		Lhs:    []ast.Expr{Expr(ctx, limit-1)},
		TokPos: token.NoPos,
		Tok:    *utilities.Pick(tokenConstructor.AccepetedByAssignStmt),
		Rhs:    []ast.Expr{Expr(ctx, limit-1)},
	}
}

func BasicLit(ctx context.Context, limit int) *ast.BasicLit {
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

func BinaryExpr(ctx context.Context, limit int) *ast.BinaryExpr {
	if limit == 0 {
		return nil
	}
	return &ast.BinaryExpr{
		X:     Expr(ctx, limit-1),
		OpPos: token.NoPos,
		Op:    *utilities.Pick(tokenConstructor.AcceptedByBinaryExpr),
		Y:     Expr(ctx, limit-1),
	}
}

func BlockStmt(ctx context.Context, limit int) *ast.BlockStmt {
	if limit == 0 {
		return nil
	}
	return &ast.BlockStmt{
		List: []ast.Stmt{
			Stmt(ctx, limit-1),
		},
		Lbrace: token.NoPos,
		Rbrace: token.NoPos,
	}
}

func BranchStmt(ctx context.Context, limit int) *ast.BranchStmt {
	if limit == 0 {
		return nil
	}
	return &ast.BranchStmt{
		TokPos: token.NoPos,
		Tok:    *utilities.Pick(tokenConstructor.AcceptedByBranchStmt), // FIXME:
		Label:  *utilities.Pick(GeneratedBranchLabels),                 // FIXME:
	}
}

func CallExpr(ctx context.Context, limit int) *ast.CallExpr {
	// TODO: function calls with more than 1 arguments
	if limit == 0 {
		return nil
	}
	return &ast.CallExpr{
		Fun:      Expr(ctx, limit-1),
		Lparen:   token.NoPos,
		Args:     []ast.Expr{Expr(ctx, limit-1)},
		Ellipsis: token.NoPos,
		Rparen:   token.NoPos,
	}
}

func CaseClause(ctx context.Context, limit int) *ast.CaseClause {
	if limit == 0 {
		return nil
	}
	return &ast.CaseClause{
		Case:  token.NoPos,
		List:  []ast.Expr{Expr(ctx, limit-1)},
		Colon: token.NoPos,
		Body: []ast.Stmt{
			Stmt(ctx, limit-1),
		},
	}
}

func ChanType(ctx context.Context, limit int) *ast.ChanType {
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
		Value: Type(ctx, limit-1),
	}
}

func CommClause(ctx context.Context, limit int) *ast.CommClause {
	if limit == 0 {
		return nil
	}
	return &ast.CommClause{
		Case:  token.NoPos,
		Colon: token.NoPos,
		Body: []ast.Stmt{
			Stmt(ctx, limit-1),
		},
	}
}

func CompositeLit(ctx context.Context, limit int) *ast.CompositeLit {
	// TODO: check Incomplete property
	if limit == 0 {
		return nil
	}
	return &ast.CompositeLit{
		Type:       Type(ctx, limit-1),
		Lbrace:     token.NoPos,
		Elts:       []ast.Expr{Expr(ctx, limit-1)},
		Rbrace:     token.NoPos,
		Incomplete: false,
	}
}

func DeclStmt(ctx context.Context, limit int) *ast.DeclStmt {
	// either with initial value assignment or declaration only
	if limit == 0 {
		return nil
	}
	return &ast.DeclStmt{
		Decl: GenDecl(ctx, limit-1),
	}
}

func DeferStmt(ctx context.Context, limit int) *ast.DeferStmt {
	if limit == 0 {
		return nil
	}
	return &ast.DeferStmt{
		Defer: token.NoPos,
		Call:  CallExpr(ctx, limit-1),
	}
}

func Ellipsis(ctx context.Context, limit int) *ast.Ellipsis {
	if limit == 0 {
		return nil
	}
	return &ast.Ellipsis{
		Ellipsis: token.NoPos,
		Elt:      Expr(ctx, limit-1),
	}
}

func EmptyStmt(ctx context.Context, limit int) *ast.EmptyStmt {
	if limit == 0 {
		return nil
	}
	return &ast.EmptyStmt{
		Semicolon: token.NoPos,
		Implicit:  false,
	}
}

func ExprStmt(ctx context.Context, limit int) *ast.ExprStmt {
	if limit == 0 {
		return nil
	}
	return &ast.ExprStmt{
		X: Expr(ctx, limit-1),
	}
}

func Field(ctx context.Context, limit int) *ast.Field {
	if limit == 0 {
		return nil
	}
	return &ast.Field{
		Names: []*ast.Ident{
			Ident(ctx, limit-1),
		},
		Type: Type(ctx, limit-1),
		Tag:  nil,
	}
}

func FieldList(ctx context.Context, limit int) *ast.FieldList {
	if limit == 0 {
		return nil
	}
	return &ast.FieldList{
		Opening: token.NoPos,
		List: []*ast.Field{
			Field(ctx, limit-1),
		},
		Closing: token.NoPos,
	}
}

func ForStmt(ctx context.Context, limit int) *ast.ForStmt {
	if limit == 0 {
		return nil
	}
	return &ast.ForStmt{
		For:  token.NoPos,
		Init: Stmt(ctx, limit-1),
		Cond: Expr(ctx, limit-1),
		Post: Stmt(ctx, limit-1),
		Body: BlockStmt(ctx, limit-1),
	}
}

func FuncDecl(ctx context.Context, limit int) *ast.FuncDecl {
	// TODO: Consider adding receiver functions (methods)
	if limit == 0 {
		return nil
	}
	return &ast.FuncDecl{
		Name: generateFunctionName(),
		Type: FuncType(ctx, limit-1),
		Body: BlockStmt(ctx, limit-1),
	}
}

func FuncLit(ctx context.Context, limit int) *ast.FuncLit {
	// TODO:
	if limit == 0 {
		return nil
	}
	return &ast.FuncLit{
		Type: FuncType(ctx, limit-1),
		Body: BlockStmt(ctx, limit-1),
	}
}

func FuncType(ctx context.Context, limit int) *ast.FuncType {
	// FIXME:
	if limit == 0 {
		return nil
	}
	return &ast.FuncType{
		Func:       token.NoPos,
		TypeParams: FieldList(ctx, limit-1),
		Params:     FieldList(ctx, limit-1),
		Results:    FieldList(ctx, limit-1),
	}
}

// Produces only "variable" declarations. "import", "constant", "type" declarations are ignored.
func GenDecl(ctx context.Context, limit int) *ast.GenDecl {
	if limit == 0 {
		return nil
	}
	return &ast.GenDecl{
		TokPos: token.NoPos,
		Tok:    token.VAR,
		Lparen: token.NoPos,
		Rparen: token.NoPos,
		Specs: []ast.Spec{
			ValueSpec(ctx, limit-1),
		},
	}
}

func GoStmt(ctx context.Context, limit int) *ast.GoStmt {
	if limit == 0 {
		return nil
	}
	return &ast.GoStmt{
		Go:   token.NoPos,
		Call: CallExpr(ctx, limit-1),
	}
}

func Ident(ctx context.Context, limit int) *ast.Ident {
	if limit == 0 {
		return nil
	}
	return generateVariableName()
}

// only valid values are types such int, float, string, bool
func IdentType(ctx context.Context, limit int) *ast.Ident {
	if limit == 0 {
		return nil
	}
	return ast.NewIdent(*utilities.Pick([]string{"int", "float", "string", "bool"}))
}

func IfStmt(ctx context.Context, limit int) *ast.IfStmt {
	if limit == 0 {
		return nil
	}
	return &ast.IfStmt{
		If:   token.NoPos,
		Init: nil,
		Cond: Expr(ctx, limit-1),
		Body: &ast.BlockStmt{
			Lbrace: token.NoPos,
			List:   []ast.Stmt{Stmt(ctx, limit-1)},
			Rbrace: token.NoPos,
		},
		Else: nil,
	}
}

func ImportSpec(ctx context.Context, limit int) *ast.ImportSpec {
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

func IncDecStmt(ctx context.Context, limit int) *ast.IncDecStmt {
	if limit == 0 {
		return nil
	}
	return &ast.IncDecStmt{
		X:      Expr(ctx, limit-1),
		TokPos: token.NoPos,
		Tok:    *utilities.Pick(tokenConstructor.AccepetedByIncDecStmt),
	}
}

func IndexExpr(ctx context.Context, limit int) *ast.IndexExpr {
	if limit == 0 {
		return nil
	}
	return &ast.IndexExpr{
		X:      Expr(ctx, limit-1),
		Lbrack: token.NoPos,
		Index:  Expr(ctx, limit-1),
		Rbrack: token.NoPos,
	}
}

func IndexListExpr(ctx context.Context, limit int) *ast.IndexListExpr {
	// TODO: Multi-dimensional arrays
	if limit == 0 {
		return nil
	}
	return &ast.IndexListExpr{
		X:       Expr(ctx, limit-1),
		Lbrack:  token.NoPos,
		Indices: []ast.Expr{Expr(ctx, limit-1)},
		Rbrack:  token.NoPos,
	}
}

func InterfaceType(ctx context.Context, limit int) *ast.InterfaceType {
	if limit == 0 {
		return nil
	}
	return &ast.InterfaceType{
		Interface:  token.NoPos,
		Methods:    FieldList(ctx, limit-1),
		Incomplete: false,
	}
}

func KeyValueExpr(ctx context.Context, limit int) *ast.KeyValueExpr {
	if limit == 0 {
		return nil
	}
	return &ast.KeyValueExpr{
		Key:   Expr(ctx, limit-1),
		Colon: token.NoPos,
		Value: Expr(ctx, limit-1),
	}
}

func LabeledStmt(ctx context.Context, limit int) *ast.LabeledStmt {
	if limit == 0 {
		return nil
	}
	return &ast.LabeledStmt{
		Label: generateBranchLabel(),
		Colon: token.NoPos,
		Stmt:  Stmt(ctx, limit-1),
	}
}

func MapType(ctx context.Context, limit int) *ast.MapType {
	if limit == 0 {
		return nil
	}
	return &ast.MapType{
		Map:   token.NoPos,
		Key:   Type(ctx, limit-1),
		Value: Type(ctx, limit-1),
	}
}

func ParenExpr(ctx context.Context, limit int) *ast.ParenExpr {
	if limit == 0 {
		return nil
	}
	return &ast.ParenExpr{
		Lparen: token.NoPos,
		X:      Expr(ctx, limit-1),
		Rparen: token.NoPos,
	}
}

func RangeStmt(ctx context.Context, limit int) *ast.RangeStmt {
	if limit == 0 {
		return nil
	}
	return &ast.RangeStmt{
		For:    token.NoPos,
		Key:    Expr(ctx, limit-1),
		Value:  Expr(ctx, limit-1),
		TokPos: token.NoPos,
		Tok:    *utilities.Pick(tokenConstructor.AcceptedByRangeStmt),
		X:      Expr(ctx, limit-1),
		Body:   BlockStmt(ctx, limit-1),
	}
}

func ReturnStmt(ctx context.Context, limit int) *ast.ReturnStmt {
	// TODO: multiple return values
	if limit == 0 {
		return nil
	}
	return &ast.ReturnStmt{
		Return:  token.NoPos,
		Results: []ast.Expr{Expr(ctx, limit-1)},
	}
}

func SelectorExpr(ctx context.Context, limit int) *ast.SelectorExpr {
	// FIXME: randomly produced X and Sel values will never work, maybe choose from imported libraries' exported functions, or previously declared struct instances that has methods
	if limit == 0 {
		return nil
	}
	return &ast.SelectorExpr{
		X:   Expr(ctx, limit-1),
		Sel: &ast.Ident{},
	}
}

func SelectStmt(ctx context.Context, limit int) *ast.SelectStmt {
	if limit == 0 {
		return nil
	}
	return &ast.SelectStmt{
		Select: token.NoPos,
		Body:   &ast.BlockStmt{}, // NOTE: CommClause() ?
	}
}

func SendStmt(ctx context.Context, limit int) *ast.SendStmt {
	if limit == 0 {
		return nil
	}
	return &ast.SendStmt{
		Chan:  Expr(ctx, limit-1),
		Arrow: token.NoPos,
		Value: Expr(ctx, limit-1),
	}
}

func SliceExpr(ctx context.Context, limit int) *ast.SliceExpr {
	if limit == 0 {
		return nil
	}
	return &ast.SliceExpr{
		X:      Expr(ctx, limit-1),
		Lbrack: token.NoPos,
		Low:    Expr(ctx, limit-1),
		High:   Expr(ctx, limit-1),
		Max:    nil,
		Slice3: false,
		Rbrack: token.NoPos,
	}
}

func StarExpr(ctx context.Context, limit int) *ast.StarExpr {
	if limit == 0 {
		return nil
	}
	return &ast.StarExpr{
		Star: token.NoPos,
		X:    Expr(ctx, limit-1),
	}
}

func StructType(ctx context.Context, limit int) *ast.StructType {
	if limit == 0 {
		return nil
	}
	return &ast.StructType{
		Struct:     token.NoPos,
		Fields:     FieldList(ctx, limit-1),
		Incomplete: false,
	}
}

func SwitchStmt(ctx context.Context, limit int) *ast.SwitchStmt {
	if limit == 0 {
		return nil
	}
	return &ast.SwitchStmt{
		Switch: token.NoPos,
		Init:   Stmt(ctx, limit-1),
		Tag:    nil,
		Body:   BlockStmt(ctx, limit-1),
	}
}

func TypeAssertExpr(ctx context.Context, limit int) *ast.TypeAssertExpr {
	if limit == 0 {
		return nil
	}
	return &ast.TypeAssertExpr{
		X:      Expr(ctx, limit-1),
		Lparen: token.NoPos,
		Type:   InterfaceType(ctx, limit-1),
		Rparen: token.NoPos,
	}
}

// rel. type keyword
func TypeSpec(ctx context.Context, limit int) *ast.TypeSpec {
	if limit == 0 {
		return nil
	}
	return &ast.TypeSpec{
		Doc:        nil,
		Name:       Ident(ctx, limit-1),
		TypeParams: FieldList(ctx, limit-1),
		Assign:     token.NoPos,
		Type:       nil,
		Comment:    nil,
	}
}

func TypeSwitchStmt(ctx context.Context, limit int) *ast.TypeSwitchStmt {
	if limit == 0 {
		return nil
	}
	return &ast.TypeSwitchStmt{
		Switch: token.NoPos,
		Init:   Stmt(ctx, limit-1),
		Assign: Stmt(ctx, limit-1),
		Body:   BlockStmt(ctx, limit-1),
	}
}

func UnaryExpr(ctx context.Context, limit int) *ast.UnaryExpr {
	if limit == 0 {
		return nil
	}
	return &ast.UnaryExpr{
		OpPos: token.NoPos,
		Op:    *utilities.Pick(tokenConstructor.AcceptedByUnaryExpr),
		X:     Expr(ctx, limit-1),
	}
}

// Returns ValueSpec which is list of pairs of variable names and values to assign
func ValueSpec(ctx context.Context, limit int) *ast.ValueSpec {
	if limit == 0 {
		return nil
	}
	return &ast.ValueSpec{
		Names: []*ast.Ident{
			generateVariableName(),
		},
		Values: []ast.Expr{Expr(ctx, limit-1)},
	}
}
