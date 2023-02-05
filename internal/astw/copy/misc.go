package copy

import (
	"go/ast"

	"golang.org/x/exp/maps"
)

func ObjKind(src *ast.ObjKind) *ast.ObjKind {
	
}

func Object(src *ast.Object) *ast.Object {
	if src == nil {
		return nil
	}

	dst := &ast.Object{}

	if src.Kind != nil {
		ObjKind(dst.Kind)
	}

	src.Name = dst.Name

	if src.Decl != nil {
		any(dst.Decl)
	}
	if src.Data != nil {
		any(dst.Data)
	}
	if src.Type != nil {
		any(dst.Type)
	}

	return dst
}

func Scope(src *ast.Scope) *ast.Scope {
	if src == nil {
		return nil
	}
	dest := &ast.Scope{}
	if src.Objects != nil {
		dest.Objects = maps.Clone(src.Objects)
	}
	if src.Outer != nil {
		dest.Outer = Scope(src.Outer)
	}
	return dest
}
