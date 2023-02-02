package node_constructor

import (
	"go/ast"
	"go/token"
	"tde/internal/cfg/context"
	utl "tde/internal/utilities"
)

func ImportSpec(ctx *context.Context, limit int) *ast.ImportSpec {
	// TODO: Store imported packages for later use
	if limit == 0 {
		return nil
	}
	return &ast.ImportSpec{
		Name: nil,
		Path: &ast.BasicLit{
			ValuePos: token.NoPos,
			Kind:     token.STRING,
			Value:    *utl.Pick(AllowedPackagesToImport),
		},
		EndPos: token.NoPos,
	}
}

// rel. type keyword
func TypeSpec(ctx *context.Context, limit int) *ast.TypeSpec {
	if limit == 0 {
		return nil
	}
	return &ast.TypeSpec{
		Doc:        nil,
		Name:       Ident(ctx, limit-1),
		TypeParams: FieldList(ctx, limit-1),
		Assign:     token.NoPos,
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
