package in_program_models

import (
	"fmt"
	"path/filepath"
	astw_utl "tde/internal/astw/utilities"
	"tde/internal/folders/types"

	"go/ast"

	"github.com/pkg/errors"
)

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
