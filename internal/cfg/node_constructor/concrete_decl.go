package node_constructor

import (
	"go/ast"
	"go/token"
	"tde/internal/cfg/context"
)

func FuncDecl(ctx *context.Context, limit int) *ast.FuncDecl {
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

// Produces only "variable" declarations. "import", "constant", "type" declarations are ignored.
func GenDecl(ctx *context.Context, limit int) *ast.GenDecl {
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
