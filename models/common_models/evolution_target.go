package common_models

import (
	astw_utl "tde/internal/astw/utilities"
	"tde/internal/folders/types"

	"fmt"
	"go/ast"
	"path/filepath"

	"github.com/pkg/errors"
)

type RunnerType int

const (
	RUNNER_CLOUD = RunnerType(iota)
	RUNNER_ON_PREM
)

type EvolverParameters struct {
	GenerationsToIterate int
	RunnerType           RunnerType
}

type OnPremisesRunnerConfig struct {
	Address string
	Port    string
	Token   string
}

type EvolutionTarget struct {
	// TODO: Module          map[string]*ast.Package // import path -> package
	Package         *ast.Package
	File            *ast.File
	FuncDecl        *ast.FuncDecl
	AllowedPackages []string
}

func NewEvolutionTarget(modulePath types.AbsolutePath, packagePath types.InModulePath, importPath string, funcName string) (*EvolutionTarget, error) {
	pkgName := filepath.Base(importPath)

	_, pkgs, err := astw_utl.LoadDir(string(modulePath.Join(packagePath)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to create ast representation for the directory: \"%s\"", packagePath))
	}

	var pkg *ast.Package
	var ok bool
	if pkg, ok = pkgs[pkgName]; !ok {
		return nil, errors.Wrap(err, fmt.Sprintf("directory doesn't contain named package: \"%s\"", pkgName))
	}

	file, funcDecl, err := astw_utl.FindFuncDeclInPkg(pkg, funcName)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("directory doesn't contain named function: \"%s\"", funcName))
	}

	return &EvolutionTarget{
		Package:         pkg,
		File:            file,
		FuncDecl:        funcDecl,
		AllowedPackages: []string{},
	}, nil
}
