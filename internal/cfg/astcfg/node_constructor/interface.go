package node_construtor

import (
	"go/ast"
	"tde/internal/utilities"
)

// Chooses a node type that confirms ast.Spec interface; initializes and returns
func Spec(credits int) ast.Spec {
	if credits == 0 {
		return nil
	}
	return (*utilities.Pick(SpecificationConstructors))(credits - 1)
}

// Chooses a node type that confirms ast.Decl interface; initializes and returns
func Decl(credits int) ast.Decl {
	if credits == 0 {
		return nil
	}
	return (*utilities.Pick(DeclarationConstructors))(credits - 1)
}

// Chooses a node type that confirms ast.Expr interface; initializes and returns
func Expr(credits int) ast.Expr {
	if credits == 0 {
		return nil
	}
	return (*utilities.Pick(ExpressionConstructors))(credits - 1)
}

// Chooses a node type that confirms ast.Stmt interface; initializes and returns
func Stmt(credits int) ast.Stmt {
	if credits == 0 {
		return nil
	}
	return (*utilities.Pick(StatementConstructors))(credits - 1)
}

// Chooses, initialized and returns an instance of random Type Declaration (Expression)
func Type(credits int) ast.Expr {
	if credits == 0 {
		return nil
	}
	return (*utilities.Pick(TypeDeclarationConstructors))(credits - 1)
}
