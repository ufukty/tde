package imports

import (
	"fmt"
	"go/ast"
	"go/token"
	"tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/utilities/pick"
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
func ImportPackage(ctx *models.MutationParameters) error {
	if ctx.File == nil {
		return fmt.Errorf("value for the parameter f is nil")
	}

	importDecls := listImportDecls(ctx.File)
	p, err := pick.Pick(ctx.AllowedPackages)
	if err != nil {
		return fmt.Errorf("picking one out of allowed packages %w", err)
	}
	importSpec := createImportSpecFromImportPath(p)

	if len(importDecls) == 0 {
		newGenDecl := &ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: []ast.Spec{importSpec},
		}
		ctx.File.Decls = append(ctx.File.Decls, newGenDecl)
	} else {
		p, err := pick.Pick(ctx.AllowedPackages)
		if err != nil {
			return fmt.Errorf("picking one out of allowed packages %w", err)
		}
		if checkExitingImports(importDecls, p) {
			return models.ErrNoChangeNeeded
		}
		importDecls[0].Specs = append(importDecls[0].Specs, importSpec)
	}

	return nil
}
