package imports

import (
	"fmt"
	"go/ast"
	"go/token"
	"tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/utilities"
)

func listImportDecls(f *ast.File) (importDecls []*ast.GenDecl) {
	for _, decl := range f.Decls {
		if decl, ok := decl.(*ast.GenDecl); ok {
			if decl.Tok == token.IMPORT {
				importDecls = append(importDecls, decl)
			}
		}
	}
	return
}

func createImportSpecFromImportPath(importPath string) *ast.ImportSpec {
	return &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf(`"%s"`, importPath),
		},
	}
}

// return early if the package already imported
func checkExitingImports(importDecls []*ast.GenDecl, importPath string) (found bool) {
	for _, decl := range importDecls {
		for _, spec := range decl.Specs {
			if importSpec, ok := spec.(*ast.ImportSpec); ok {
				if importSpec.Path.Value == importPath {
					return true
				}
			}
		}
	}
	return false
}

// Always returns with desired result. But don't modify, if the package is already imported
func ImportPackage(f *ast.File, importPath string) (modified bool) {
	if f == nil {
		panic("value for the parameter f is nil")
	}

	importDecls := listImportDecls(f)
	importSpec := createImportSpecFromImportPath(importPath)

	if len(importDecls) == 0 {
		newGenDecl := ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: []ast.Spec{importSpec},
		}
		f.Decls = append([]ast.Decl{&newGenDecl}, f.Decls...)
	} else {
		if checkExitingImports(importDecls, importPath) {
			return false
		}
		importDecls[0].Specs = append(importDecls[0].Specs, importSpec)
	}

	return true
}

func GeneticOperation(ctx *models.MutationParameters) bool {
	return ImportPackage(ctx.File, *utilities.Pick(ctx.AllowedPackages))
}
