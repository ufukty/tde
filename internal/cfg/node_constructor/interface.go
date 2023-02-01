package node_constructor

import (
	"tde/internal/cfg/context"
	"tde/internal/utilities"

	"go/ast"
)

// Chooses a node type that confirms ast.Spec interface; initializes and returns
func Spec(ctx *context.Context, limit int) ast.Spec {
	return (*utilities.Pick(SpecificationConstructors))(ctx, limit)
}

// Chooses a node type that confirms ast.Decl interface; initializes and returns
func Decl(ctx *context.Context, limit int) ast.Decl {
	return (*utilities.Pick(DeclarationConstructors))(ctx, limit)
}

// Chooses a node type that confirms ast.Expr interface; initializes and returns
func Expr(ctx *context.Context, limit int) ast.Expr {
	return (*utilities.Pick(ExpressionConstructors))(ctx, limit)
}

// Chooses a node type that confirms ast.Stmt interface; initializes and returns
func Stmt(ctx *context.Context, limit int) ast.Stmt {
	return (*utilities.Pick(StatementConstructors))(ctx, limit)
}

// Chooses, initialized and returns an instance of random Type Declaration (Expression)
func Type(ctx *context.Context, limit int) ast.Expr {
	return (*utilities.Pick(TypeDeclarationConstructors))(ctx, limit)
}
