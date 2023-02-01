package node_constructor

import (
	"tde/internal/cfg/context"
	utl "tde/internal/utilities"

	"go/ast"
)

// Chooses a node type that confirms ast.Spec interface; initializes and returns
func Spec(ctx *context.Context, limit int) ast.Spec {
	return (*utl.Pick(SpecificationConstructors))(ctx, limit)
}

// Chooses a node type that confirms ast.Decl interface; initializes and returns
func Decl(ctx *context.Context, limit int) ast.Decl {
	return (*utl.Pick(DeclarationConstructors))(ctx, limit)
}

// Chooses a node type that confirms ast.Expr interface; initializes and returns
func Expr(ctx *context.Context, limit int) ast.Expr {
	return (*utl.Pick(ExpressionConstructors))(ctx, limit)
}

// Chooses a node type that confirms ast.Stmt interface; initializes and returns
func Stmt(ctx *context.Context, limit int) ast.Stmt {
	return (*utl.Pick(StatementConstructors))(ctx, limit)
}

// Chooses, initialized and returns an instance of random Type Declaration (Expression)
func Type(ctx *context.Context, limit int) ast.Expr {
	return (*utl.Pick(TypeDeclarationConstructors))(ctx, limit)
}
