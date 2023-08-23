package examiner

import (
	"tde/internal/cfg/ctxres/context"

	"go/ast"
	"go/token"
)

func examineSingularAssignment(ctx *context.Context, lhs, rhs ast.Expr) {
	if lhs, ok := lhs.(*ast.Ident); ok {
		ctx.AddVariable(lhs)
	}
}

func examineAssignStmt(ctx *context.Context, stmt *ast.AssignStmt) {
	if stmt.Tok == token.DEFINE {
		for i := 0; i < len(stmt.Lhs); i++ {
			examineSingularAssignment(ctx, stmt.Lhs[i], stmt.Rhs[i])
		}
	}
}

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
