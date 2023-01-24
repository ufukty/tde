package astw

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/pkg/errors"
)

func LoadDir(dirpath string) (*token.FileSet, map[string]*ast.Package, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dirpath, nil, parser.AllErrors)
	if err != nil {
		return nil, nil, errors.Wrap(err, "LoadDir")
	}
	return fset, pkgs, nil
}

func LoadFile(filepath string) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, filepath, nil, parser.AllErrors)
	if err != nil {
		return nil, nil, errors.Wrap(err, "LoadFile")
	}
	return fset, astFile, nil
}

func ParseString(content string) (*token.FileSet, ast.Node, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", content, parser.AllErrors)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ParseString")
	}
	return fset, astFile, nil
}
