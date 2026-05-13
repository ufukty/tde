package models

import (
	"fmt"
	"go/ast"

	"tde/internal/ast/clone/clean"
	"tde/internal/ast/parse"

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

func funcInFile(pkg *ast.Package, name string) (*ast.File, *ast.FuncDecl, int, error) {
	if pkg != nil && pkg.Files != nil {
		for _, f := range pkg.Files {
			for i, d := range f.Decls {
				if fd, ok := d.(*ast.FuncDecl); ok {
					if fd.Name != nil && fd.Name.Name == name {
						return f, fd, i, nil
					}
				}
			}
		}
	}
	return nil, nil, -1, fmt.Errorf("could not find: %s", name)
}

func LoadContext(module, pkgpath, funcname string) (*Context, error) {
	p, err := parse.Package(pkgpath)
	if err != nil {
		return nil, fmt.Errorf("parsing the package: %w", err)
	}
	f, fd, fdi, err := funcInFile(p, funcname)
	if err != nil {
		return nil, fmt.Errorf("locating the function in AST: %w", err)
	}
	return &Context{
		Module:        map[string]*ast.Package{}, // TODO: context for module
		Package:       p,
		File:          f,
		FuncDecl:      fd,
		funcDeclIndex: fdi,
		orgFuncDecl:   fd,
	}, nil
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
