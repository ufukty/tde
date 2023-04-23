package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func createNewFileName(inputPath string) string {
	filename := filepath.Base(inputPath)
	fileExt := filepath.Ext(filename)
	filenameWithoutExt := filename[0 : len(filename)-len(fileExt)]
	return filepath.Join(filepath.Dir(inputPath), filenameWithoutExt+".sd.go")
}
func checkSuffix(str, suffix string) bool {
	return strings.LastIndex(str, suffix) == len(str)-len(suffix)
}

func discoverFileForStructDefinitions(file *ast.File) (reqStructs, resStructs []*ast.Ident) {
	ast.Inspect(file, func(n ast.Node) bool {
		if n == nil || n == file {
			return true
		}
		switch n := n.(type) {
		case *ast.File, *ast.GenDecl:
			return true
		case *ast.TypeSpec:
			if _, ok := n.Type.(*ast.StructType); ok {
				if checkSuffix(n.Name.Name, "Request") {
					reqStructs = append(reqStructs, n.Name)
				} else if checkSuffix(n.Name.Name, "Response") {
					resStructs = append(resStructs, n.Name)
				}
			}
			return false
		}
		return false
	})
	return
}

func parseFile(path string) (*ast.File, error) {
	file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.AllErrors)
	if err != nil {
		return nil, errors.Wrap(err, "Outside error")
	}
	return file, nil
}
