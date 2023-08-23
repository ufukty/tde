package nodes

import (
	"tde/internal/genetics/mutation/cfg/ctxres/context"
	"tde/internal/tokenw"
	utl "tde/internal/utilities"

	"go/ast"
)

func AssignStmt(ctx *context.Context, limit int) *ast.AssignStmt {
	if limit == 0 {
		return nil
	}
	return &ast.AssignStmt{
		// TokPos: token.NoPos,
		Lhs: []ast.Expr{Expr(ctx, limit-1)},
		Rhs: []ast.Expr{Expr(ctx, limit-1)},
		Tok: *utl.Pick(tokenw.AcceptedByAssignStmt),
	}
}

func BlockStmt(ctx *context.Context, limit int) *ast.BlockStmt {
	if limit == 0 {
		return nil
	}
	return &ast.BlockStmt{
		// Lbrace: token.NoPos,
		// Rbrace: token.NoPos,
		List: []ast.Stmt{Stmt(ctx, limit-1)},
	}
}

func BranchStmt(ctx *context.Context, limit int) *ast.BranchStmt {
	if limit == 0 {
		return nil
	}
	if (len(GeneratedBranchLabels)) == 0 {
		return nil
	}
	return &ast.BranchStmt{
		// TokPos: token.NoPos,
		Label: *utl.Pick(GeneratedBranchLabels),       // FIXME:
		Tok:   *utl.Pick(tokenw.AcceptedByBranchStmt), // FIXME:
	}
}

func CaseClause(ctx *context.Context, limit int) *ast.CaseClause {
	if limit == 0 {
		return nil
	}
	return &ast.CaseClause{
		// Case:  token.NoPos,
		// Colon: token.NoPos,
		List: []ast.Expr{Expr(ctx, limit-1)},
		Body: []ast.Stmt{Stmt(ctx, limit-1)},
	}
}

var commClauseCommGenerators = []func(*context.Context, int) ast.Stmt{
	func(ctx *context.Context, limit int) ast.Stmt { return nil },
	func(ctx *context.Context, limit int) ast.Stmt { return SendStmt(ctx, limit) },
	func(ctx *context.Context, limit int) ast.Stmt { return ReturnStmt(ctx, limit) },
}

func CommClause(ctx *context.Context, limit int) *ast.CommClause {
	if limit == 0 {
		return nil
	}
	return &ast.CommClause{
		// Case:  token.NoPos,
		// Colon: token.NoPos,
		Body: []ast.Stmt{Stmt(ctx, limit-1)},
		Comm: (*utl.Pick(commClauseCommGenerators))(ctx, limit-1),
	}
}

func DeclStmt(ctx *context.Context, limit int) *ast.DeclStmt {
	// either with initial value assignment or declaration only
	if limit == 0 {
		return nil
	}
	return &ast.DeclStmt{
		Decl: GenDecl(ctx, limit-1),
	}
}

func DeferStmt(ctx *context.Context, limit int) *ast.DeferStmt {
	if limit == 0 {
		return nil
	}
	return &ast.DeferStmt{
		// Defer: token.NoPos,
		Call: CallExpr(ctx, limit-1),
	}
}

func EmptyStmt(ctx *context.Context, limit int) *ast.EmptyStmt {
	if limit == 0 {
		return nil
	}
	return &ast.EmptyStmt{
		// Semicolon: token.NoPos,
		Implicit: false,
	}
}

func ExprStmt(ctx *context.Context, limit int) *ast.ExprStmt {
	if limit == 0 {
		return nil
	}
	return &ast.ExprStmt{
		X: Expr(ctx, limit-1),
	}
}

func ForStmt(ctx *context.Context, limit int) *ast.ForStmt {
	if limit == 0 {
		return nil
	}
	return &ast.ForStmt{
		// For:  token.NoPos,
		Init: Stmt(ctx, limit-1),
		Cond: Expr(ctx, limit-1),
		Post: Stmt(ctx, limit-1),
		Body: BlockStmt(ctx, limit-1),
	}
}

func GoStmt(ctx *context.Context, limit int) *ast.GoStmt {
	if limit == 0 {
		return nil
	}
	return &ast.GoStmt{
		// Go:   token.NoPos,
		Call: CallExpr(ctx, limit-1),
	}
}

func IfStmt(ctx *context.Context, limit int) *ast.IfStmt {
	if limit == 0 {
		return nil
	}
	return &ast.IfStmt{
		// If:   token.NoPos,
		Init: nil,
		Else: nil,
		Cond: Expr(ctx, limit-1),
		Body: BlockStmt(ctx, limit-1),
	}
}

func IncDecStmt(ctx *context.Context, limit int) *ast.IncDecStmt {
	if limit == 0 {
		return nil
	}
	return &ast.IncDecStmt{
		// TokPos: token.NoPos,
		X:   Expr(ctx, limit-1),
		Tok: *utl.Pick(tokenw.AcceptedByIncDecStmt),
	}
}

func LabeledStmt(ctx *context.Context, limit int) *ast.LabeledStmt {
	if limit == 0 {
		return nil
	}
	return &ast.LabeledStmt{
		// Colon: token.NoPos,
		Stmt:  Stmt(ctx, limit-1),
		Label: generateBranchLabel(),
	}
}

func RangeStmt(ctx *context.Context, limit int) *ast.RangeStmt {
	if limit == 0 {
		return nil
	}
	return &ast.RangeStmt{
		// For:    token.NoPos,
		// TokPos: token.NoPos,
		X:     Expr(ctx, limit-1),
		Key:   Expr(ctx, limit-1),
		Value: Expr(ctx, limit-1),
		Body:  BlockStmt(ctx, limit-1),
		Tok:   *utl.Pick(tokenw.AcceptedByRangeStmt),
	}
}

func ReturnStmt(ctx *context.Context, limit int) *ast.ReturnStmt {
	// TODO: multiple return values
	if limit == 0 {
		return nil
	}
	return &ast.ReturnStmt{
		// Return:  token.NoPos,
		Results: []ast.Expr{Expr(ctx, limit-1)},
	}
}

func SelectStmt(ctx *context.Context, limit int) *ast.SelectStmt {
	if limit == 0 {
		return nil
	}
	return &ast.SelectStmt{
		// Select: token.NoPos,
		Body: BlockStmt(ctx, limit-1),
	}
}

func SendStmt(ctx *context.Context, limit int) *ast.SendStmt {
	if limit == 0 {
		return nil
	}
	return &ast.SendStmt{
		// Arrow: token.NoPos,
		Chan:  Expr(ctx, limit-1),
		Value: Expr(ctx, limit-1),
	}
}

func SwitchStmt(ctx *context.Context, limit int) *ast.SwitchStmt {
	if limit == 0 {
		return nil
	}
	return &ast.SwitchStmt{
		// Switch: token.NoPos,
		Tag:  nil,
		Init: Stmt(ctx, limit-1),
		Body: BlockStmt(ctx, limit-1),
	}
}

func TypeSwitchStmt(ctx *context.Context, limit int) *ast.TypeSwitchStmt {
	if limit == 0 {
		return nil
	}
	return &ast.TypeSwitchStmt{
		// Switch: token.NoPos,
		Init:   Stmt(ctx, limit-1),
		Assign: Stmt(ctx, limit-1),
		Body:   BlockStmt(ctx, limit-1),
	}
}
