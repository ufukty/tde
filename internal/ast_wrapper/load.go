package ast_wrapper

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/pkg/errors"
)

func LoadFile(filepath string) (*token.FileSet, ast.Node, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, filepath, nil, parser.AllErrors)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not parse file")
	}
	return fset, astFile, nil
}

func ParseString(content string) (*token.FileSet, ast.Node, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", content, parser.AllErrors)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not parse string")
	}
	return fset, astFile, nil
}
