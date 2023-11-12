package examiner

import (
	"go/ast"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"
)

func examineFuncDecl(ctx *context.Context, decl *ast.FuncDecl) {
	if decl.Recv == nil {
		ctx.AddFuncDeclaration(decl)
	} else {
		ctx.AddMethodDeclaration(decl)
	}
}

// Only adds declarations for functions, imports, types, variables and constants. Won't examine function bodies.
func examineFile(ctx *context.Context, file *ast.File) {
	for _, decl := range file.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			examineGenDecl(ctx, decl)
		case *ast.FuncDecl:
			examineFuncDecl(ctx, decl)
		}
	}
}

func Pkg(ctx *context.Context, pkg *ast.Package) {
	for _, file := range pkg.Files {
		examineFile(ctx, file)
	}
}
