package nodes

import (
	"go/ast"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"
	"tde/internal/evolution/genetics/mutation/v1/tokens"
	"tde/internal/utilities/pick"
)

func AssignStmt(ctx *context.Context, limit int) (*ast.AssignStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.AssignStmt{
		// TokPos: token.NoPos,
		Lhs: []ast.Expr{Expr(ctx, limit-1)},
		Rhs: []ast.Expr{Expr(ctx, limit-1)},
		Tok: *pick.Pick(tokens.AcceptedByAssignStmt),
	}, nil
}

func BlockStmt(ctx *context.Context, limit int) (*ast.BlockStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.BlockStmt{
		// Lbrace: token.NoPos,
		// Rbrace: token.NoPos,
		List: []ast.Stmt{Stmt(ctx, limit-1)},
	}, nil
}

func BranchStmt(ctx *context.Context, limit int) (*ast.BranchStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	if (len(GeneratedBranchLabels)) == 0 {
		return nil, ErrNoAvailableValues
	}
	return &ast.BranchStmt{
		// TokPos: token.NoPos,
		Label: *pick.Pick(GeneratedBranchLabels),       // FIXME:
		Tok:   *pick.Pick(tokens.AcceptedByBranchStmt), // FIXME:
	}, nil
}

func CaseClause(ctx *context.Context, limit int) (*ast.CaseClause, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.CaseClause{
		// Case:  token.NoPos,
		// Colon: token.NoPos,
		List: []ast.Expr{Expr(ctx, limit-1)},
		Body: []ast.Stmt{Stmt(ctx, limit-1)},
	}, nil
}

var commClauseCommGenerators = []func(*context.Context, int) (ast.Stmt, error){
	func(ctx *context.Context, limit int) ast.Stmt { return nil, error },
	func(ctx *context.Context, limit int) ast.Stmt { return SendStmt(ctx, limit), error },
	func(ctx *context.Context, limit int) ast.Stmt { return ReturnStmt(ctx, limit), error },
}

func CommClause(ctx *context.Context, limit int) (*ast.CommClause, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.CommClause{
		// Case:  token.NoPos,
		// Colon: token.NoPos,
		Body: []ast.Stmt{Stmt(ctx, limit-1)},
		Comm: (*pick.Pick(commClauseCommGenerators))(ctx, limit-1),
	}, nil
}

func DeclStmt(ctx *context.Context, limit int) (*ast.DeclStmt, error) {
	// either with initial value assignment or declaration only
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.DeclStmt{
		Decl: GenDecl(ctx, limit-1),
	}, nil
}

func DeferStmt(ctx *context.Context, limit int) (*ast.DeferStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.DeferStmt{
		// Defer: token.NoPos,
		Call: CallExpr(ctx, limit-1),
	}, nil
}

func EmptyStmt(ctx *context.Context, limit int) (*ast.EmptyStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.EmptyStmt{
		// Semicolon: token.NoPos,
		Implicit: false,
	}, nil
}

func ExprStmt(ctx *context.Context, limit int) (*ast.ExprStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.ExprStmt{
		X: Expr(ctx, limit-1),
	}, nil
}

func ForStmt(ctx *context.Context, limit int) (*ast.ForStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.ForStmt{
		// For:  token.NoPos,
		Init: Stmt(ctx, limit-1),
		Cond: Expr(ctx, limit-1),
		Post: Stmt(ctx, limit-1),
		Body: BlockStmt(ctx, limit-1),
	}, nil
}

func GoStmt(ctx *context.Context, limit int) (*ast.GoStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.GoStmt{
		// Go:   token.NoPos,
		Call: CallExpr(ctx, limit-1),
	}, nil
}

func IfStmt(ctx *context.Context, limit int) (*ast.IfStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.IfStmt{
		// If:   token.NoPos,
		Init: nil,
		Else: nil,
		Cond: Expr(ctx, limit-1),
		Body: BlockStmt(ctx, limit-1),
	}, nil
}

func IncDecStmt(ctx *context.Context, limit int) (*ast.IncDecStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.IncDecStmt{
		// TokPos: token.NoPos,
		X:   Expr(ctx, limit-1),
		Tok: *pick.Pick(tokens.AcceptedByIncDecStmt),
	}, nil
}

func LabeledStmt(ctx *context.Context, limit int) (*ast.LabeledStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.LabeledStmt{
		// Colon: token.NoPos,
		Stmt:  Stmt(ctx, limit-1),
		Label: generateBranchLabel(),
	}, nil
}

func RangeStmt(ctx *context.Context, limit int) (*ast.RangeStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.RangeStmt{
		// For:    token.NoPos,
		// TokPos: token.NoPos,
		X:     Expr(ctx, limit-1),
		Key:   Expr(ctx, limit-1),
		Value: Expr(ctx, limit-1),
		Body:  BlockStmt(ctx, limit-1),
		Tok:   *pick.Pick(tokens.AcceptedByRangeStmt),
	}, nil
}

func ReturnStmt(ctx *context.Context, limit int) (*ast.ReturnStmt, error) {
	// TODO: multiple return values
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.ReturnStmt{
		// Return:  token.NoPos,
		Results: []ast.Expr{Expr(ctx, limit-1)},
	}, nil
}

func SelectStmt(ctx *context.Context, limit int) (*ast.SelectStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.SelectStmt{
		// Select: token.NoPos,
		Body: BlockStmt(ctx, limit-1),
	}, nil
}

func SendStmt(ctx *context.Context, limit int) (*ast.SendStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.SendStmt{
		// Arrow: token.NoPos,
		Chan:  Expr(ctx, limit-1),
		Value: Expr(ctx, limit-1),
	}, nil
}

func SwitchStmt(ctx *context.Context, limit int) (*ast.SwitchStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.SwitchStmt{
		// Switch: token.NoPos,
		Tag:  nil,
		Init: Stmt(ctx, limit-1),
		Body: BlockStmt(ctx, limit-1),
	}, nil
}

func TypeSwitchStmt(ctx *context.Context, limit int) (*ast.TypeSwitchStmt, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.TypeSwitchStmt{
		// Switch: token.NoPos,
		Init:   Stmt(ctx, limit-1),
		Assign: Stmt(ctx, limit-1),
		Body:   BlockStmt(ctx, limit-1),
	}, nil
}
