package copy

import (
	"go/ast"
)

func ObjectMap(src map[string]*ast.Object) map[string]*ast.Object {
	if src == nil {
		return nil
	}
	dst := map[string]*ast.Object{}
	for k, v := range src {
		dst[k] = Object(v)
	}
	return dst
}

func FileMap(src map[string]*ast.File) map[string]*ast.File {
	if src == nil {
		return nil
	}
	dst := map[string]*ast.File{}
	for k, v := range src {
		dst[k] = File(v)
	}
	return dst
}
