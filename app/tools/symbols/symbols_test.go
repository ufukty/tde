package symbols

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"tde/internal/astw/astwutl"

	"golang.org/x/exp/maps"
)

var testdatafolders = [][]string{
	{"/usr/local/go/src/fmt", "writePadding"},
	{"/usr/local/go/src/reflect", "DeepEqual"},
	{"testdata/evolution", "WalkWithNils"},
	{"testdata/words", "Reverse"},
}

var testcase int64 = 0

func prepare() (*ast.Package, *ast.FuncDecl, ast.Node, *types.Info, *types.Package, error) {
	path, funcname := testdatafolders[testcase][0], testdatafolders[testcase][1]
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, 0)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("parser: %w", err)
	}
	astpkg := maps.Values(pkgs)[0] // there should be exactly 1 package at tested dir
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs:       map[*ast.Ident]types.Object{},
		Implicits:  map[ast.Node]types.Object{},
		InitOrder:  []*types.Initializer{},
		Instances:  map[*ast.Ident]types.Instance{},
		Scopes:     map[ast.Node]*types.Scope{},
		Selections: map[*ast.SelectorExpr]*types.Selection{},
		Types:      map[ast.Expr]types.TypeAndValue{},
		Uses:       map[*ast.Ident]types.Object{},
	}
	pkg, err := conf.Check("main", fset, maps.Values(astpkg.Files), info)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("check: %w", err)
	}
	funcdecl, err := astwutl.FindFuncDecl(astpkg, funcname)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("find func decl: %w", err)
	}
	spot := funcdecl.Body.List[0]
	return astpkg, funcdecl, spot, info, pkg, nil
}
