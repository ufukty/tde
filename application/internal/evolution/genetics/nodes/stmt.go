package nodes

import (
	"fmt"
	"go/ast"
	"tde/internal/evolution/genetics/nodes/tokens"
	"tde/internal/utilities/pick"
)

func (c *Creator) AssignStmt(l int) (*ast.AssignStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Lhs, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating AssignStmt.Lhs: %w", err)
	}
	Rhs, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating AssignStmt.Rhs: %w", err)
	}
	Tok, err := pick.Pick(tokens.AcceptedByAssignStmt)
	if err != nil {
		return nil, fmt.Errorf("generating AssignStmt.Tok: %w", err)
	}

	return &ast.AssignStmt{
		Lhs: []ast.Expr{Lhs},
		Rhs: []ast.Expr{Rhs},
		Tok: Tok,
	}, nil
}

func (c *Creator) BlockStmt(l int) (*ast.BlockStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Stmt, err := c.Stmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating BlockStmt.List: %w", err)
	}

	return &ast.BlockStmt{
		List: []ast.Stmt{Stmt},
	}, nil
}

// FIXME: branches
func (c *Creator) BranchStmt(l int) (*ast.BranchStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	if (len(GeneratedBranchLabels)) == 0 {
		return nil, ErrNoAvailableValues
	}
	Label, err := pick.Pick(GeneratedBranchLabels)
	if err != nil {
		return nil, fmt.Errorf("generating BranchStmt.Label: %w", err)
	}
	Tok, err := pick.Pick(tokens.AcceptedByBranchStmt)
	if err != nil {
		return nil, fmt.Errorf("generating BranchStmt.Tok: %w", err)
	}

	return &ast.BranchStmt{
		Label: Label,
		Tok:   Tok,
	}, nil
}

func (c *Creator) CaseClause(l int) (*ast.CaseClause, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	List, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating CaseClause.List: %w", err)
	}
	Body, err := c.Stmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating CaseClause.Body: %w", err)
	}

	return &ast.CaseClause{
		List: []ast.Expr{List},
		Body: []ast.Stmt{Body},
	}, nil
}

func (c *Creator) CommClause(l int) (*ast.CommClause, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Stmt, err := c.Stmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating CommClause.Body: %w", err)
	}
	generator, err := pick.Pick([]func(int) (ast.Stmt, error){
		func(l int) (ast.Stmt, error) { return nil, nil },
		func(l int) (ast.Stmt, error) { return c.SendStmt(l) },
		func(l int) (ast.Stmt, error) { return c.ReturnStmt(l) },
	})
	if err != nil {
		return nil, fmt.Errorf("picking generator for CommClause.Comm: %w", err)
	}
	Comm, err := generator(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating CommCaluse.Comm: %w", err)
	}

	return &ast.CommClause{
		Body: []ast.Stmt{Stmt},
		Comm: Comm,
	}, nil
}

// either with initial value assignment or declaration only
func (c *Creator) DeclStmt(l int) (*ast.DeclStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Decl, err := c.GenDecl(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating DeclStmt.Decl: %w", err)
	}

	return &ast.DeclStmt{
		Decl: Decl,
	}, nil
}

func (c *Creator) DeferStmt(l int) (*ast.DeferStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Call, err := c.CallExpr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating DeferStmt.CallExpr: %w", err)
	}

	return &ast.DeferStmt{
		Call: Call,
	}, nil
}

func (c *Creator) ExprStmt(l int) (*ast.ExprStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	X, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating ExprStmt.X: %w", err)
	}

	return &ast.ExprStmt{
		X: X,
	}, nil
}

func (c *Creator) ForStmt(l int) (*ast.ForStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Init, err := c.Stmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating ForStmt.Init: %w", err)
	}
	Cond, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating ForStmt.Cond: %w", err)
	}
	Post, err := c.Stmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating ForStmt.Post: %w", err)
	}
	Body, err := c.BlockStmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating ForStmt.Body: %w", err)
	}

	return &ast.ForStmt{
		Init: Init,
		Cond: Cond,
		Post: Post,
		Body: Body,
	}, nil
}

func (c *Creator) GoStmt(l int) (*ast.GoStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Call, err := c.CallExpr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating GoStmt.Call: %w", err)
	}

	return &ast.GoStmt{
		Call: Call,
	}, nil
}

func (c *Creator) IfStmt(l int) (*ast.IfStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Cond, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating IfStmt.Cond: %w", err)
	}
	Body, err := c.BlockStmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating IfStmt.Body: %w", err)
	}

	return &ast.IfStmt{
		Init: nil,
		Else: nil,
		Cond: Cond,
		Body: Body,
	}, nil
}

func (c *Creator) IncDecStmt(l int) (*ast.IncDecStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	X, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating IncDecStmt.X: %w", err)
	}
	Tok, err := pick.Pick(tokens.AcceptedByIncDecStmt)
	if err != nil {
		return nil, fmt.Errorf("generating IncDecStmt.Tok=pick: %w", err)
	}

	return &ast.IncDecStmt{
		X:   X,
		Tok: Tok,
	}, nil
}

func (c *Creator) LabeledStmt(l int) (*ast.LabeledStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Stmt, err := c.Stmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating LabeledStmt.Stmt: %w", err)
	}

	return &ast.LabeledStmt{
		Stmt:  Stmt,
		Label: generateBranchLabel(),
	}, nil
}

func (c *Creator) RangeStmt(l int) (*ast.RangeStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	X, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating RangeStmt.X: %w", err)
	}
	Key, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating RangeStmt.Key: %w", err)
	}
	Value, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating RangeStmt.Value: %w", err)
	}
	Body, err := c.BlockStmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating RangeStmt.Body: %w", err)
	}
	Tok, err := pick.Pick(tokens.AcceptedByRangeStmt)
	if err != nil {
		return nil, fmt.Errorf("generating RangeStmt.Tok=pick: %w", err)
	}

	return &ast.RangeStmt{
		X:     X,
		Key:   Key,
		Value: Value,
		Body:  Body,
		Tok:   Tok,
	}, nil
}

func (c *Creator) ReturnStmt(l int) (*ast.ReturnStmt, error) {
	// TODO: multiple return values
	if l == 0 {
		return nil, ErrLimitReached
	}
	Results, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating ReturnStmt.Results: %w", err)
	}

	return &ast.ReturnStmt{
		Results: []ast.Expr{Results},
	}, nil
}

func (c *Creator) SelectStmt(l int) (*ast.SelectStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Body, err := c.BlockStmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating SelectStmt.Body: %w", err)
	}

	return &ast.SelectStmt{
		Body: Body,
	}, nil
}

func (c *Creator) SendStmt(l int) (*ast.SendStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Chan, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating SendStmt.Chan: %w", err)
	}
	Value, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating SendStmt.Value: %w", err)
	}

	return &ast.SendStmt{
		Chan:  Chan,
		Value: Value,
	}, nil
}

func (c *Creator) SwitchStmt(l int) (*ast.SwitchStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Init, err := c.Stmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating SwitchStmt.Init: %w", err)
	}
	Body, err := c.BlockStmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating SwitchStmt.Body: %w", err)
	}

	return &ast.SwitchStmt{
		Tag:  nil,
		Init: Init,
		Body: Body,
	}, nil
}

func (c *Creator) TypeSwitchStmt(l int) (*ast.TypeSwitchStmt, error) {
	if l == 0 {
		return nil, ErrLimitReached
	}
	Init, err := c.Stmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating TypeSwitchStmt.Init: %w", err)
	}
	Assign, err := c.Stmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating TypeSwitchStmt.Assign: %w", err)
	}
	Body, err := c.BlockStmt(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating TypeSwitchStmt.Body: %w", err)
	}

	return &ast.TypeSwitchStmt{
		Init:   Init,
		Assign: Assign,
		Body:   Body,
	}, nil
}
