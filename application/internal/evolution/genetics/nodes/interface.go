package nodes

import (
	"fmt"
	"go/ast"
	"tde/internal/utilities/pick"
)

// About those constructors:
//
// Chooses a node type that confirms ast.Spec, ast.Decl, ast.Expr, ast.Stmt
// interface or is an expression that defines a type; initializes and returns.

func (c *Creator) Spec(l int) (ast.Spec, error) {
	generator, err := pick.Pick([]func(int) (ast.Spec, error){
		func(l int) (ast.Spec, error) { return c.ImportSpec(l) },
		func(l int) (ast.Spec, error) { return c.TypeSpec(l) },
		func(l int) (ast.Spec, error) { return c.ValueSpec(l) },
	})
	if err != nil {
		return nil, fmt.Errorf("picking generator for Spec: %w", err)
	}
	return generator(l)
}

func (c *Creator) Decl(l int) (ast.Decl, error) {
	generator, err := pick.Pick([]func(int) (ast.Decl, error){
		func(l int) (ast.Decl, error) { return c.FuncDecl(l) },
		func(l int) (ast.Decl, error) { return c.GenDecl(l) },
	})
	if err != nil {
		return nil, fmt.Errorf("picking generator for Decl: %w", err)
	}
	return generator(l)
}

func (c *Creator) Expr(l int) (ast.Expr, error) {
	generator, err := pick.Pick([]func(int) (ast.Expr, error){
		func(l int) (ast.Expr, error) { return c.BasicLit(l) },
		func(l int) (ast.Expr, error) { return c.BinaryExpr(l) },
		func(l int) (ast.Expr, error) { return c.CallExpr(l) },
		func(l int) (ast.Expr, error) { return c.CompositeLit(l) },
		func(l int) (ast.Expr, error) { return c.Ellipsis(l) },
		func(l int) (ast.Expr, error) { return c.FuncLit(l) },
		func(l int) (ast.Expr, error) { return c.Ident(l) },
		func(l int) (ast.Expr, error) { return c.IndexExpr(l) },
		func(l int) (ast.Expr, error) { return c.IndexListExpr(l) },
		func(l int) (ast.Expr, error) { return c.KeyValueExpr(l) },
		func(l int) (ast.Expr, error) { return c.ParenExpr(l) },
		func(l int) (ast.Expr, error) { return c.SelectorExpr(l) },
		func(l int) (ast.Expr, error) { return c.SliceExpr(l) },
		func(l int) (ast.Expr, error) { return c.StarExpr(l) },
		func(l int) (ast.Expr, error) { return c.UnaryExpr(l) },
	})
	if err != nil {
		return nil, fmt.Errorf("picking generator for Expr: %w", err)
	}
	return generator(l)
}

func (c *Creator) Stmt(l int) (ast.Stmt, error) {
	generator, err := pick.Pick([]func(int) (ast.Stmt, error){
		func(l int) (ast.Stmt, error) { return c.AssignStmt(l) },
		func(l int) (ast.Stmt, error) { return c.BlockStmt(l) },
		func(l int) (ast.Stmt, error) { return c.BranchStmt(l) },
		func(l int) (ast.Stmt, error) { return c.CaseClause(l) },
		func(l int) (ast.Stmt, error) { return c.CommClause(l) },
		func(l int) (ast.Stmt, error) { return c.DeclStmt(l) },
		func(l int) (ast.Stmt, error) { return c.DeferStmt(l) },
		func(l int) (ast.Stmt, error) { return c.ExprStmt(l) },
		func(l int) (ast.Stmt, error) { return c.ForStmt(l) },
		func(l int) (ast.Stmt, error) { return c.GoStmt(l) },
		func(l int) (ast.Stmt, error) { return c.IfStmt(l) },
		func(l int) (ast.Stmt, error) { return c.IncDecStmt(l) },
		func(l int) (ast.Stmt, error) { return c.LabeledStmt(l) },
		func(l int) (ast.Stmt, error) { return c.RangeStmt(l) },
		func(l int) (ast.Stmt, error) { return c.ReturnStmt(l) },
		func(l int) (ast.Stmt, error) { return c.SelectStmt(l) },
		func(l int) (ast.Stmt, error) { return c.SendStmt(l) },
		func(l int) (ast.Stmt, error) { return c.SwitchStmt(l) },
		func(l int) (ast.Stmt, error) { return c.TypeSwitchStmt(l) },
	})
	if err != nil {
		return nil, fmt.Errorf("picking generator for Stmt: %w", err)
	}
	return generator(l)
}

func (c *Creator) Type(l int) (ast.Expr, error) {
	generator, err := pick.Pick([]func(int) (ast.Expr, error){
		func(l int) (ast.Expr, error) { return c.ArrayType(l) },
		func(l int) (ast.Expr, error) { return c.ChanType(l) },
		func(l int) (ast.Expr, error) { return c.FuncType(l) },
		func(l int) (ast.Expr, error) { return c.InterfaceType(l) },
		func(l int) (ast.Expr, error) { return c.MapType(l) },
		func(l int) (ast.Expr, error) { return c.StructType(l) },
		func(l int) (ast.Expr, error) { return c.IdentType(l) },
		func(l int) (ast.Expr, error) { return c.ParenExprForType(l) },
		func(l int) (ast.Expr, error) { return c.SelectorExprForType(l) },
		func(l int) (ast.Expr, error) { return c.StarExprForType(l) },
	})
	if err != nil {
		return nil, fmt.Errorf("picking generator for Type: %w", err)
	}
	return generator(l)
}
