package context

import (
	"go/ast"
	"go/token"
)

func examineImportDeclaration(ctx *Context, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		if spec, ok := spec.(*ast.ImportSpec); ok {
			ctx.AddImport(spec)
		}
	}
}

func examineVariableDeclaration(ctx *Context, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		if spec, ok := spec.(*ast.ValueSpec); ok {
			for _, name := range spec.Names {
				ctx.AddVariable(name)
			}
		}
	}
}

func examineTypeDeclaration(ctx *Context, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		if spec, ok := spec.(*ast.TypeSpec); ok {
			ctx.AddType(spec)
		}
	}
}

func examineGenDecl(ctx *Context, decl *ast.GenDecl) {
	switch decl.Tok {
	case token.IMPORT:
		examineImportDeclaration(ctx, decl)
	case token.VAR, token.CONST:
		examineVariableDeclaration(ctx, decl)
	case token.TYPE:
		examineTypeDeclaration(ctx, decl)
	}
}

func examineFuncDecl(ctx *Context, decl *ast.FuncDecl) {
	if decl.Recv == nil {
		ctx.AddFuncDeclaration(decl)
	} else {
		ctx.AddMethodDeclaration(decl)
	}
}

// Only adds declarations for functions, imports, types, variables and constants. Won't examine function bodies.
func examineFile(ctx *Context, file *ast.File) {
	for _, decl := range file.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			examineGenDecl(ctx, decl)
		case *ast.FuncDecl:
			examineFuncDecl(ctx, decl)
		}
	}
}

func Package(ctx *Context, pkg *ast.Package) {
	for _, file := range pkg.Files {
		examineFile(ctx, file)
	}
}
