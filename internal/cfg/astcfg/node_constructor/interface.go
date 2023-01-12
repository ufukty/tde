package node_construtor

import (
	"tde/internal/cfg/astcfg/context"
	"tde/internal/utilities"

	"go/ast"
)

// Chooses a node type that confirms ast.Spec interface; initializes and returns
func Spec(ctx context.Context, limit int) ast.Spec {
	if limit == 0 {
		return nil
	}
	return (*utilities.Pick(SpecificationConstructors))(ctx, limit-1)
}

// Chooses a node type that confirms ast.Decl interface; initializes and returns
func Decl(ctx context.Context, limit int) ast.Decl {
	if limit == 0 {
		return nil
	}
	return (*utilities.Pick(DeclarationConstructors))(ctx, limit-1)
}

// Chooses a node type that confirms ast.Expr interface; initializes and returns
func Expr(ctx context.Context, limit int) ast.Expr {
	if limit == 0 {
		return nil
	}
	return (*utilities.Pick(ExpressionConstructors))(ctx, limit-1)
}

// Chooses a node type that confirms ast.Stmt interface; initializes and returns
func Stmt(ctx context.Context, limit int) ast.Stmt {
	if limit == 0 {
		return nil
	}
	return (*utilities.Pick(StatementConstructors))(ctx, limit-1)
}

// Chooses, initialized and returns an instance of random Type Declaration (Expression)
func Type(ctx context.Context, limit int) ast.Expr {
	if limit == 0 {
		return nil
	}
	return (*utilities.Pick(TypeDeclarationConstructors))(ctx, limit-1)
}
