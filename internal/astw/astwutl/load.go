package astwutl

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"maps"
	"slices"
)

func LoadDir(dirpath string) (*token.FileSet, map[string]*ast.Package, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dirpath, nil, parser.AllErrors|parser.SkipObjectResolution)
	if err != nil {
		return nil, nil, fmt.Errorf("parser: %w", err)
	}
	return fset, pkgs, nil
}

func LoadFile(filepath string) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, filepath, nil, parser.AllErrors|parser.SkipObjectResolution)
	if err != nil {
		return nil, nil, fmt.Errorf("parser: %w", err)
	}
	return fset, astFile, nil
}

func ParseString(content string) (*token.FileSet, ast.Node, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", content, parser.AllErrors|parser.SkipObjectResolution)
	if err != nil {
		return nil, nil, fmt.Errorf("parser: %w", err)
	}
	return fset, astFile, nil
}

func LoadPackageFromDir(path string) (*ast.Package, error) {
	var (
		pkgs    map[string]*ast.Package
		pkgList []string
		err     error
	)
	_, pkgs, err = LoadDir(path)
	if err != nil {
		return nil, fmt.Errorf("root: %w", err)
	}
	pkgList = slices.Collect(maps.Keys(pkgs))
	if l := len(pkgList); l != 1 {
		return nil, fmt.Errorf("unexpected number of packages: %d", l)
	}
	return pkgs[pkgList[0]], nil
}
