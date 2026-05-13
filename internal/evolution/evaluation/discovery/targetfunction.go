package discovery

import (
	"fmt"
	"go/ast"
	"go/token"

	"tde/internal/astw/astwutl"
	"tde/internal/astw/parse"
)

func find(path string, name string) (*ast.FuncDecl, *token.FileSet, error) {
	fset, astFile, err := parse.File(path)
	if err != nil {
		return nil, nil, fmt.Errorf("parsing file %q: %w", path, err)
	}
	var found *ast.FuncDecl
	ast.Inspect(astFile, func(node ast.Node) bool {
		if node, ok := node.(*ast.FuncDecl); ok {
			if node.Name.Name == name {
				found = node
			}
		}
		return found == nil
	})
	if found == nil {
		return nil, nil, fmt.Errorf("not found")
	}
	return found, fset, err
}

func TargetFunctionInFile(path string, funcname string) (*TargetFunction, error) {
	funcDecl, fset, err := find(path, funcname)
	if err != nil {
		return nil, fmt.Errorf("searching ast of file %q for the function %q: %w", path, funcname, err)
	}
	return &TargetFunction{
		Name:      funcname,
		Path:      path,
		LineStart: astwutl.LineNumberOfPosition(fset, funcDecl.Pos()),
		LineEnd:   astwutl.LineNumberOfPosition(fset, funcDecl.End()),
	}, nil
}
