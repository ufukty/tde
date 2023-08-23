package nodes

import (
	"tde/internal/evolution/genetics/mutation/cfg/ctxres/context"
	utl "tde/internal/utilities"

	"go/ast"
	"go/token"
)

func ImportSpec(ctx *context.Context, limit int) *ast.ImportSpec {
	// TODO: Store imported packages for later use
	if limit == 0 {
		return nil
	}
	return &ast.ImportSpec{
		// EndPos: token.NoPos,
		Name: nil,
		Path: &ast.BasicLit{
			// ValuePos: token.NoPos,
			Kind:  token.STRING,
			Value: *utl.Pick(AllowedPackagesToImport),
		},
	}
}

// rel. type keyword
func TypeSpec(ctx *context.Context, limit int) *ast.TypeSpec {
	if limit == 0 {
		return nil
	}
	return &ast.TypeSpec{
		// Assign:     token.NoPos,
		Doc:        nil,
		Name:       Ident(ctx, limit-1),
		TypeParams: FieldList(ctx, limit-1),
		Type:       nil,
		Comment:    nil,
	}
}

// Returns ValueSpec which is list of pairs of variable names and values to assign
func ValueSpec(ctx *context.Context, limit int) *ast.ValueSpec {
	if limit == 0 {
		return nil
	}
	return &ast.ValueSpec{
		Names: []*ast.Ident{
			generateNewIdent(),
		},
		Values: []ast.Expr{Expr(ctx, limit-1)},
	}
}
