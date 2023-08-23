package examiner

import (
	"tde/internal/evolution/genetics/mutation/cfg/ctxres/context"

	"go/ast"
	"go/token"
)

func examineImportDeclaration(ctx *context.Context, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		if spec, ok := spec.(*ast.ImportSpec); ok {
			ctx.AddImport(spec)
		}
	}
}

func examineVariableDeclaration(ctx *context.Context, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		if spec, ok := spec.(*ast.ValueSpec); ok {
			for _, name := range spec.Names {
				ctx.AddVariable(name)
			}
		}
	}
}

func examineTypeDeclaration(ctx *context.Context, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		if spec, ok := spec.(*ast.TypeSpec); ok {
			ctx.AddType(spec)
		}
	}
}

func examineGenDecl(ctx *context.Context, decl *ast.GenDecl) {
	switch decl.Tok {
	case token.IMPORT:
		examineImportDeclaration(ctx, decl)
	case token.VAR, token.CONST:
		examineVariableDeclaration(ctx, decl)
	case token.TYPE:
		examineTypeDeclaration(ctx, decl)
	}
}
