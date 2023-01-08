package node_construtor

import (
	"go/ast"
	"tde/internal/utilities"
)

// Chooses a node type that confirms ast.Spec interface; initializes and returns
func Spec(limit int) ast.Spec {
	if limit == 0 {
		return nil
	}
	return (*utilities.Pick(SpecificationConstructors))(limit - 1)
}

// Chooses a node type that confirms ast.Decl interface; initializes and returns
func Decl(limit int) ast.Decl {
	if limit == 0 {
		return nil
	}
	return (*utilities.Pick(DeclarationConstructors))(limit - 1)
}

// Chooses a node type that confirms ast.Expr interface; initializes and returns
func Expr(limit int) ast.Expr {
	if limit == 0 {
		return nil
	}
	return (*utilities.Pick(ExpressionConstructors))(limit - 1)
}

// Chooses a node type that confirms ast.Stmt interface; initializes and returns
func Stmt(limit int) ast.Stmt {
	if limit == 0 {
		return nil
	}
	return (*utilities.Pick(StatementConstructors))(limit - 1)
}

// Chooses, initialized and returns an instance of random Type Declaration (Expression)
func Type(limit int) ast.Expr {
	if limit == 0 {
		return nil
	}
	return (*utilities.Pick(TypeDeclarationConstructors))(limit - 1)
}
