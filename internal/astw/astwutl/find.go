package astwutl

import (
	"fmt"
	"go/ast"
)

func FindFuncInsideFile(f *ast.File, name string) (*ast.FuncDecl, error) {
	if f != nil && f.Decls != nil {
		for _, d := range f.Decls {
			if fd, ok := d.(*ast.FuncDecl); ok && fd.Name != nil && fd.Name.Name == name {
				return fd, nil
			}
		}
	}
	return nil, fmt.Errorf("not found")
}

func FindFuncInsidePackage(p *ast.Package, name string) (*ast.FuncDecl, error) {
	if p != nil && p.Files != nil {
		for _, f := range p.Files {
			if fd, err := FindFuncInsideFile(f, name); err == nil {
				return fd, nil
			}
		}
	}
	return nil, fmt.Errorf("not found")
}

func FindFuncDecl(root ast.Node, name string) (*ast.FuncDecl, error) {
	switch root := (root).(type) {
	case *ast.Package:
		return FindFuncInsidePackage(root, name)
	case *ast.File:
		return FindFuncInsideFile(root, name)
	}
	return nil, fmt.Errorf("expected ast.Package or ast.File")
}
