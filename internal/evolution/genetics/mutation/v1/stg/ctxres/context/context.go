package context

import "go/ast"

type Scope struct {
	Imports   []*ast.ImportSpec // package scope
	Variables []*ast.Ident      // from package scope to inner scopes such as bodies of flow control blocks
	Types     []*ast.TypeSpec   // package scope
	Methods   []*ast.FuncDecl   // package scope, functions that has receivers
	Functions []*ast.FuncDecl   // package scope functions that has no receiver
}

func NewScope() Scope {
	return Scope{
		Functions: []*ast.FuncDecl{},
		Imports:   []*ast.ImportSpec{},
		Types:     []*ast.TypeSpec{},
		Variables: []*ast.Ident{},
	}
}

type Context struct {
	Scopes []Scope
}

func NewContext() *Context {
	return &Context{Scopes: []Scope{NewScope()}}
}

func (ctx *Context) ScopeIn() {
	ctx.Scopes = append(ctx.Scopes, NewScope())
}

func (ctx *Context) ScopeOut() {
	ctx.Scopes = ctx.Scopes[:len(ctx.Scopes)-1]
}

func (ctx *Context) GetCurrentScope() *Scope {
	return &ctx.Scopes[len(ctx.Scopes)-1]
}

func (ctx *Context) AddFuncDeclaration(item *ast.FuncDecl) {
	ctx.GetCurrentScope().Functions = append(ctx.GetCurrentScope().Functions, item)
}

func (ctx *Context) AddMethodDeclaration(item *ast.FuncDecl) {
	ctx.GetCurrentScope().Methods = append(ctx.GetCurrentScope().Methods, item)
}

func (ctx *Context) AddImport(item *ast.ImportSpec) {
	ctx.GetCurrentScope().Imports = append(ctx.GetCurrentScope().Imports, item)
}

func (ctx *Context) AddType(item *ast.TypeSpec) {
	ctx.GetCurrentScope().Types = append(ctx.GetCurrentScope().Types, item)
}

func (ctx *Context) AddVariable(item *ast.Ident) {
	ctx.GetCurrentScope().Variables = append(ctx.GetCurrentScope().Variables, item)
}

func (ctx *Context) GetFunctions() []*ast.FuncDecl {
	ret := []*ast.FuncDecl{}
	for _, scope := range ctx.Scopes {
		ret = append(ret, scope.Functions...)
	}
	return ret
}

func (ctx *Context) GetImports() []*ast.ImportSpec {
	ret := []*ast.ImportSpec{}
	for _, scope := range ctx.Scopes {
		ret = append(ret, scope.Imports...)
	}
	return ret
}

func (ctx *Context) GetTypes() []*ast.TypeSpec {
	ret := []*ast.TypeSpec{}
	for _, scope := range ctx.Scopes {
		ret = append(ret, scope.Types...)
	}
	return ret
}

func (ctx *Context) GetVariables() []*ast.Ident {
	ret := []*ast.Ident{}
	for _, scope := range ctx.Scopes {
		ret = append(ret, scope.Variables...)
	}
	return ret
}
