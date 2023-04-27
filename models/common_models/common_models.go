package common_models

import "go/ast"

type OnPremisesRunnerConfig struct {
	Address string
	Port    string
	Token   string
}

type EvolutionTarget struct {
	// TODO: Module          map[string]*ast.Package // import path -> package
	Package  *ast.Package
	File     *ast.File
	FuncDecl *ast.FuncDecl
}

type EvolutionConfig struct {
	Timeout    int      // in seconds
	Runner     string   // ip address
	Continue   string   // session
	Model      string   //
	Ratios     string   //
	Population int      //
	Iterate    int      //
	Size       int      //
	Package    []string // packages allowed to import
	TestName   string
}
