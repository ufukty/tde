package context

import (
	"go/ast"
	"go/token"
)
func ExamineImportDeclaration(ctx *Context, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		if spec, ok := spec.(*ast.ImportSpec); ok {
			ctx.AddImport(spec)
		}
	}
}

func ExamineVariableDeclaration(ctx *Context, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		if spec, ok := spec.(*ast.ImportSpec); ok {
			ctx.AddImport(spec)
		}
	}
}

func ExamineTypeDeclaration(ctx *Context, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		if spec, ok := spec.(*ast.TypeSpec); ok {
			ctx.AddType(spec)
		}
	}
}

func ExamineGenDecl(ctx *Context, decl *ast.GenDecl) {
	switch decl.Tok {
	case token.IMPORT:
		ExamineImportDeclaration(ctx, decl)
	case token.VAR, token.CONST:
		ExamineVariableDeclaration(ctx, decl)
	case token.TYPE:
		ExamineTypeDeclaration(ctx, decl)
	}
}

func ExamineFuncDecl(ctx *Context, decl *ast.FuncDecl) {
	if decl.Recv == nil {
		ctx.AddFuncDeclaration(decl)
	} else {
		ctx.AddMethodDeclaration(decl)
	}
}

// Only adds declarations for functions, imports, types, variables and constants. Won't examine function bodies.
func ExamineFile(ctx *Context, file *ast.File) {
	for _, decl := range file.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			ExamineGenDecl(ctx, decl)
		case *ast.FuncDecl:
			ExamineFuncDecl(ctx, decl)
		}
	}
}
