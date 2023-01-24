package context

import "go/ast"

type Scope struct {
	// Functions []ast.Ident
	Libraries        []ast.Ident
	TypeDeclarations []ast.Expr
	Variables        []ast.Ident
}

func NewScope() Scope {
	return Scope{
		// Functions: []ast.Ident{},
		Libraries:        []ast.Ident{},
		TypeDeclarations: []ast.Expr{},
		Variables:        []ast.Ident{},
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

func (ctx *Context) AddLibrary(item ast.Ident) {
	ctx.GetCurrentScope().Libraries = append(ctx.GetCurrentScope().Libraries, item)
}

func (ctx *Context) AddTypeDeclaration(item ast.Expr) {
	ctx.GetCurrentScope().TypeDeclarations = append(ctx.GetCurrentScope().TypeDeclarations, item)
}

func (ctx *Context) AddVariable(item ast.Ident) {
	ctx.GetCurrentScope().Variables = append(ctx.GetCurrentScope().Variables, item)
}

func (ctx *Context) GetLibraries() []ast.Ident {
	ret := []ast.Ident{}
	for _, scope := range ctx.Scopes {
		ret = append(ret, scope.Libraries...)
	}
	return ret
}

func (ctx *Context) GetTypeDeclarations() []ast.Expr {
	ret := []ast.Expr{}
	for _, scope := range ctx.Scopes {
		ret = append(ret, scope.TypeDeclarations...)
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
