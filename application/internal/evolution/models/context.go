package models

import (
	"fmt"
	"go/ast"
	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone/clean"

	"github.com/google/uuid"
)

type Context struct {
	Module   map[string]*ast.Package
	Package  *ast.Package
	File     *ast.File
	FuncDecl *ast.FuncDecl

	funcDeclIndex int
	orgFuncDecl   *ast.FuncDecl
}

func LoadContext(module, pkgpath, funcname string) (*Context, error) {
	pkg, err := astwutl.LoadPackageFromDir(pkgpath)
	if err != nil {
		return nil, fmt.Errorf("loading the package in %q: %w", pkgpath, err)
	}
	file, funcdecl, err := astwutl.FindFuncDeclInPkg(pkg, funcname)
	if err != nil {
		return nil, fmt.Errorf("searching %q in the package: %w", funcname, err)
	}
	ctx := &Context{
		Module:   map[string]*ast.Package{}, // TODO: context for module
		Package:  pkg,
		File:     file,
		FuncDecl: funcdecl,

		funcDeclIndex: -1,
		orgFuncDecl:   funcdecl,
	}
	for i, decl := range file.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok && funcDecl.Name.Name == funcname {
			ctx.funcDeclIndex = i
		}
	}
	return ctx, nil
}

// needed before printing file
func (ctx *Context) Swap(funcDecl *ast.FuncDecl) {
	ctx.File.Decls[ctx.funcDeclIndex] = funcDecl
}

// needed after printing
func (ctx *Context) Restore() {
	ctx.File.Decls[ctx.funcDeclIndex] = ctx.orgFuncDecl
}

func (ctx *Context) NewSubject() *Subject {
	return &Subject{
		Sid:     Sid(uuid.New().String()), // UUID v4,
		Parent:  "-1",
		AST:     clean.FuncDecl(ctx.orgFuncDecl),
		Imports: []*ast.ImportSpec{},
		Code:    []byte{},
		Fitness: Fitness{
			AST:       1.0,
			Code:      1.0,
			Program:   1.0,
			Candidate: 1.0,
		},
		ExecTimeInMs: 0,
	}
}
