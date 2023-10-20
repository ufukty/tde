package models

import (
	"fmt"
	"go/ast"
	"tde/internal/astw/astwutl"
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
		Module:   map[string]*ast.Package{}, // TODO:
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
