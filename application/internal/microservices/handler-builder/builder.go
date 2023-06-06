package handler_builder

import (
	"fmt"
	"go/ast"
	"strings"
)

func FindRequestSpec(file *ast.File) (requestSpec *ast.TypeSpec) {
	ast.Inspect(file, func(n ast.Node) bool {
		var t, ok = n.(*ast.TypeSpec)
		if !ok || !strings.HasSuffix(t.Name.Name, "Request") {
			return requestSpec == nil
		}
		requestSpec = t
		return false
	})
	return
}

func FindResponseSpec(file *ast.File) (responseSpec *ast.TypeSpec) {
	ast.Inspect(file, func(n ast.Node) bool {
		var t, ok = n.(*ast.TypeSpec)
		if !ok || !strings.HasSuffix(t.Name.Name, "Response") {
			return responseSpec == nil
		}
		responseSpec = t
		return false
	})
	return
}

type Mime string

const (
	NA                  = Mime("")
	APPLICATION_JSON    = Mime("application/json")
	APPLICATION_YAML    = Mime("application/yaml")
	APPLICATION_ZIP     = Mime("application/zip")
	APPLICATION_JPEG    = Mime("application/jpeg")
	APPLICATION_PNG     = Mime("application/png")
	MULTIPART_FORM_DATA = Mime("multipart/form-data")
)

func FindTagContentType(tag string) (Mime, bool) {
	var pairs = ParseTag(tag)
	var contentType, ok = pairs["Content-Type"]
	if !ok {
		return NA, false
	}
	switch contentType {
	case "application/json":
		return APPLICATION_JSON, true
	case "application/yaml":
		return APPLICATION_YAML, true
	case "application/zip":
		return APPLICATION_ZIP, true
	case "application/jpeg":
		return APPLICATION_JPEG, true
	case "application/png":
		return APPLICATION_PNG, true
	case "multipart/form-data":
		return MULTIPART_FORM_DATA, true
	}
	return NA, false
}

func FindRequestContentType(reqSpec *ast.TypeSpec) Mime {
	var t, ok = reqSpec.Type.(*ast.StructType)
	if !ok {
		panic("Type specification is not a struct type specification")
	}

	for _, field := range t.Fields.List {
		if field.Names[0].Name == "Body" {
			var contentType, ok = FindTagContentType(field.Tag.Value)
			if !ok {
				continue
			}

			fmt.Println(contentType)
		}
	}
	return ""
}

func Boilerer(src *ast.File) *ast.File {
	var dst = &ast.File{
		Name:  src.Name,
		Decls: []ast.Decl{},
	}
	var requestSpec = FindRequestSpec(src)
	// var responseSpec = FindResponseSpec(src)

	FindRequestContentType(requestSpec)
	return dst
}
