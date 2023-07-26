package types

import (
	"go/ast"
	"tde/internal/folders/list"
)

type ImplFuncDetails struct {
	Name string
	Path string
	Line int
}

type TestFuncDetails struct {
	Name  string
	Path  string
	Line  int
	Calls []*ast.CallExpr
}

type DiscoveryResponse struct {
	ModuleAbsolutePath      string
	TargetPackageImportPath string
}

type TestDetails struct {
	PackagePath string // eg. .../examples/word-reverse

	Package      *list.Package // eg. examples/word-reverse/word_reverse
	ImplFuncFile string        // eg. .../examples/word-reverse/word_reverse.go
	ImplFuncName string        // eg. WordReverse
	ImplFuncLine int

	TestFuncFile string // eg. .../examples/word-reverse/word_reverse_tde.go
	TestFuncName string // eg. TDE_WordReverse
	TestFuncLine int
}
