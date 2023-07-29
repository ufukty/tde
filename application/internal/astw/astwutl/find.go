package astwutl

import (
	"fmt"
	"go/ast"
	"go/token"
)

func FindFuncDecl(toppest ast.Node, name string) (*ast.FuncDecl, error) {
	found := []*ast.FuncDecl{}

	ast.Inspect(toppest, func(n ast.Node) bool {
		if n != nil {
			if functionDeclaration, ok := n.(*ast.FuncDecl); ok {
				if functionDeclaration.Name.Name == name {
					found = append(found, functionDeclaration)
					return false
				}
			}
		}
		return true
	})

	if len(found) == 0 {
		return nil, fmt.Errorf("could not find function '%s'", name)
	} else if len(found) > 1 {
		return nil, fmt.Errorf("more than one function definition with the name  '%s'", name)
	} else {
		return found[0], nil
	}
}

func FindFuncDeclInPkg(pkg *ast.Package, name string) (*ast.File, *ast.FuncDecl, error) {
	funcDecl := []*ast.FuncDecl{}
	file := []*ast.File{}

	var fileVisited *ast.File

	ast.Inspect(pkg, func(n ast.Node) bool {
		if n != nil {
			if file, ok := n.(*ast.File); ok {
				fileVisited = file
			}
			if functionDeclaration, ok := n.(*ast.FuncDecl); ok {
				if functionDeclaration.Name.Name == name {
					funcDecl = append(funcDecl, functionDeclaration)
					file = append(file, fileVisited)
					return false
				}
			}
		}
		return true
	})

	if len(funcDecl) == 0 {
		return nil, nil, fmt.Errorf("could not find function '%s'", name)
	} else if len(funcDecl) > 1 {
		return nil, nil, fmt.Errorf("more than one function definition with the name  '%s'", name)
	} else {
		return file[0], funcDecl[0], nil
	}
}

func FindFunctionInFile(path string, name string) (*ast.FuncDecl, *token.FileSet, error) {
	fset, astFile, err := LoadFile(path)
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
