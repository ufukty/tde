package context

import "go/ast"

type Scope struct {
	Functions []ast.Ident
	Libraries []ast.Ident
	Types     []ast.Expr
	Variables []ast.Ident
}

func NewScope() Scope {
	return Scope{
		Functions: []ast.Ident{},
		Libraries: []ast.Ident{},
		Types:     []ast.Expr{},
		Variables: []ast.Ident{},
	}
}

type Context struct {
	Scopes []Scope
}

func NewContext() Context {
	return Context{Scopes: []Scope{NewScope()}}
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

func (ctx *Context) AddFunction(funcDecl ast.FuncDecl) {
	ctx.GetCurrentScope().Functions = append(ctx.GetCurrentScope().Functions, *funcDecl.Name)
}

func (ctx *Context) AddLibrary(funcDecl ast.FuncDecl) {
	ctx.GetCurrentScope().Functions = append(ctx.GetCurrentScope().Functions, *funcDecl.Name)
}

func (ctx *Context) AddType(funcDecl ast.FuncDecl) {
	ctx.GetCurrentScope().Functions = append(ctx.GetCurrentScope().Functions, *funcDecl.Name)
}

func (ctx *Context) AddVariable(funcDecl ast.FuncDecl) {
	ctx.GetCurrentScope().Functions = append(ctx.GetCurrentScope().Functions, *funcDecl.Name)
}

func (ctx *Context) GetFunctions() []ast.Ident {
	ret := []ast.Ident{}
	for _, scope := range ctx.Scopes {
		ret = append(ret, scope.Functions...)
	}
	return ret
}

func (ctx *Context) GetLibraries() []ast.Ident {
	ret := []ast.Ident{}
	for _, scope := range ctx.Scopes {
		ret = append(ret, scope.Libraries...)
	}
	return ret
}

func (ctx *Context) GetTypes() []ast.Expr {
	ret := []ast.Expr{}
	for _, scope := range ctx.Scopes {
		ret = append(ret, scope.Types...)
	}
	return ret
}

func (ctx *Context) GetVariables() []ast.Ident {
	ret := []ast.Ident{}
	for _, scope := range ctx.Scopes {
		ret = append(ret, scope.Variables...)
	}
	return ret
}
