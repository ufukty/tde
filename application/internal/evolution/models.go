package evolution

import "go/ast"

type OnPremisesRunnerConfig struct {
	Address string
	Port    string
	Token   string
}

type Target struct {
	// TODO: Module          map[string]*ast.Package // import path -> package
	Package  *ast.Package
	File     *ast.File
	FuncDecl *ast.FuncDecl
}
