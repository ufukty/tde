package nodes

import (
	"go/ast"
	"go/token"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"
)

func FuncDecl(ctx *context.Context, limit int) (*ast.FuncDecl, error) {
	// TODO: Consider adding receiver functions (methods)
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.FuncDecl{
		Name: generateFunctionName(),
		Type: FuncType(ctx, limit-1),
		Body: BlockStmt(ctx, limit-1),
	}, nil
}

// Produces only "variable" declarations. "import", "constant", "type" declarations are ignored.
func GenDecl(ctx *context.Context, limit int) (*ast.GenDecl, error) {
	if limit == 0 {
		return nil, ErrLimitReached
	}
	return &ast.GenDecl{
		// TokPos: token.NoPos,
		// Lparen: token.NoPos,
		// Rparen: token.NoPos,
		Tok: token.VAR,
		Specs: []ast.Spec{
			ValueSpec(ctx, limit-1),
		},
	}, nil
}
