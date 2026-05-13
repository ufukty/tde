package parse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"maps"
	"slices"
)

func Dir(path string) (*token.FileSet, map[string]*ast.Package, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.AllErrors|parser.SkipObjectResolution)
	if err != nil {
		return nil, nil, fmt.Errorf("parser: %w", err)
	}
	return fset, pkgs, nil
}

func File(path string) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, path, nil, parser.AllErrors|parser.SkipObjectResolution)
	if err != nil {
		return nil, nil, fmt.Errorf("parser: %w", err)
	}
	return fset, astFile, nil
}

func String(content string) (*token.FileSet, ast.Node, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", content, parser.AllErrors|parser.SkipObjectResolution)
	if err != nil {
		return nil, nil, fmt.Errorf("parser: %w", err)
	}
	return fset, astFile, nil
}

func Package(path string) (*ast.Package, error) {
	_, pkgs, err := Dir(path)
	if err != nil {
		return nil, fmt.Errorf("root: %w", err)
	}
	list := slices.Collect(maps.Keys(pkgs))
	if l := len(list); l != 1 {
		return nil, fmt.Errorf("unexpected number of packages: %d", l)
	}
	return pkgs[list[0]], nil
}
