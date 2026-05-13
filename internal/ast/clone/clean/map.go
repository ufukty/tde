package clean

import (
	"go/ast"
)

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
