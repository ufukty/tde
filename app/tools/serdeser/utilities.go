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

func discoverFileForStructDefinitions(file *ast.File) (reqModels, resModels []*ast.Ident) {
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
					reqModels = append(reqModels, n.Name)
				} else if checkSuffix(n.Name.Name, "Response") {
					resModels = append(resModels, n.Name)
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

func stripSuffix(str, suffix string) string {
	return str[0 : len(str)-len(suffix)]
}

type ReqResPair struct {
	Request  *ast.Ident
	Response *ast.Ident
}

func FindReqResPairs(reqModels, resModels []*ast.Ident) (pairs []ReqResPair) {
	for _, reqModel := range reqModels {
		if len(reqModel.Name) < len("Request") {
			continue
		}
		reqName := stripSuffix(reqModel.Name, "Request")
		for _, resModel := range resModels {
			if len(resModel.Name) < len("Request") {
				continue
			}
			resName := stripSuffix(resModel.Name, "Response")
			if reqName == resName {
				pairs = append(pairs, ReqResPair{reqModel, resModel})
			}
		}
	}
	return
}
