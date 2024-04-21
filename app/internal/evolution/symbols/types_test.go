package symbols

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"tde/internal/astw/astwutl"
	"testing"

	"golang.org/x/exp/maps"
)

type tcase1 struct {
	name        string
	funcname    string
	pkgpath     string
	pkgid       string
	allowedpkgs []string
}

var tcs1 = []tcase1{
	{
		name:        "fmt",
		funcname:    "writePadding",
		pkgpath:     "/usr/local/go/src/fmt",
		pkgid:       "fmt",
		allowedpkgs: []string{},
	},
	{
		name:        "reflect",
		funcname:    "DeepEqual",
		pkgpath:     "/usr/local/go/src/reflect",
		pkgid:       "reflect",
		allowedpkgs: []string{},
	},
	{
		name:        "evolution",
		funcname:    "WalkWithNils",
		pkgpath:     "./testdata/evolution",
		pkgid:       "tde/internal/evolution/symbols/testdata/evolution",
		allowedpkgs: []string{},
	},
	{
		name:        "words",
		funcname:    "Reverse",
		pkgpath:     "./testdata/words",
		pkgid:       "tde/internal/evolution/symbols/testdata/words",
		allowedpkgs: []string{"/usr/local/go/src/fmt", "/usr/local/go/src/math"},
	},
}

func prepare(tc tcase1) (*ast.Package, *ast.FuncDecl, ast.Node, *types.Info, *types.Package, error) {
	path, funcname := tc.pkgpath, tc.funcname
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

func TestFindFuncScope(t *testing.T) {
	for _, tc := range tcs1 {
		t.Run(tc.name, func(t *testing.T) {
			_, funcdecl, _, info, _, err := prepare(tc)
			if err != nil {
				t.Fatal(fmt.Errorf("prep: %w", err))
			}
			funcscope := info.Defs[funcdecl.Name].Parent()
			fmt.Println(funcscope.Names())
			// Output: [WalkCallbackFunction WalkWithNils increaseLastChildIndex isNodeNil walkAstTypeFieldsIfSet walkHelper]
		})
	}
}

func Test_ConvertScopesMap(t *testing.T) {
	for _, tc := range tcs1 {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, info, _, err := prepare(tc)
			if err != nil {
				t.Fatal(fmt.Errorf("prep: %w", err))
			}

			scopeNodes := map[*types.Scope][]ast.Node{}
			for node, scope := range info.Scopes {
				if _, ok := scopeNodes[scope]; !ok {
					scopeNodes[scope] = []ast.Node{}
				}
				scopeNodes[scope] = append(scopeNodes[scope], node)
			}
			fmt.Println(scopeNodes)

			// if nodes, ok := scopeNodes[pkg.Scope()]; ok {
			// 	for _, node := range nodes {
			// 		fmt.Println(reflect.TypeOf(node), node)
			// 	}
			// }

			// if funcscope, ok := info.Scopes[funcdecl]; ok {
			// 	if funcscopeobjects, ok := scopeNodes[funcscope]; ok {
			// 		fmt.Println(len(funcscopeobjects))
			// 		fmt.Println("")
			// 	}
			// }
		})
	}
}

func Test_PkgScope(t *testing.T) {
	for _, tc := range tcs1 {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, _, pkg, err := prepare(tc)
			if err != nil {
				t.Fatal(fmt.Errorf("prep: %w", err))
			}

			children := []*types.Scope{}
			for i := 0; i < pkg.Scope().NumChildren(); i++ {
				scope := pkg.Scope().Child(i)
				children = append(children, scope)
				fmt.Println("scope==", scope)
				children2 := []*types.Scope{}
				for j := 0; j < scope.NumChildren(); j++ {
					scope := scope.Child(i)
					children2 = append(children2, scope)
				}
				fmt.Println(children2)
			}

		})
	}
}

func TestFindingIdentsDefinedInImports(t *testing.T) {
	for _, tc := range tcs1 {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, _, pkg, err := prepare(tc)
			if err != nil {
				t.Error(fmt.Errorf("prep: %w", err))
			}

			for _, ipkg := range pkg.Imports() {
				// fmt.Println(ipkg.Name(), ipkg.Scope().Names())
				qual := types.RelativeTo(pkg)
				if ipkg.Name() == "fmt" {
					for _, symbolname := range ipkg.Scope().Names() {
						symbol := ipkg.Scope().Lookup(symbolname)
						str := types.ObjectString(symbol, qual)
						fmt.Println(str)
					}
				}
			}
			// Output:
			// func fmt.Append(b []byte, a ...any) []byte
			// func fmt.Appendf(b []byte, format string, a ...any) []byte
			// func fmt.Appendln(b []byte, a ...any) []byte
			// func fmt.Errorf(format string, a ...any) error
			// func fmt.FormatString(state fmt.State, verb rune) string
			// type fmt.Formatter interface{Format(f fmt.State, verb rune)}
			// func fmt.Fprint(w io.Writer, a ...any) (n int, err error)
			// func fmt.Fprintf(w io.Writer, format string, a ...any) (n int, err error)
			// func fmt.Fprintln(w io.Writer, a ...any) (n int, err error)
			// func fmt.Fscan(r io.Reader, a ...any) (n int, err error)
			// func fmt.Fscanf(r io.Reader, format string, a ...any) (n int, err error)
			// func fmt.Fscanln(r io.Reader, a ...any) (n int, err error)
			// type fmt.GoStringer interface{GoString() string}
			// func fmt.Print(a ...any) (n int, err error)
			// func fmt.Printf(format string, a ...any) (n int, err error)
			// func fmt.Println(a ...any) (n int, err error)
			// func fmt.Scan(a ...any) (n int, err error)
			// type fmt.ScanState interface{Read(buf []byte) (n int, err error); ReadRune() (r rune, size int, err error); SkipSpace(); Token(skipSpace bool, f func(rune) bool) (token []byte, err error); UnreadRune() error; Width() (wid int, ok bool)}
			// func fmt.Scanf(format string, a ...any) (n int, err error)
			// func fmt.Scanln(a ...any) (n int, err error)
			// type fmt.Scanner interface{Scan(state fmt.ScanState, verb rune) error}
			// func fmt.Sprint(a ...any) string
			// func fmt.Sprintf(format string, a ...any) string
			// func fmt.Sprintln(a ...any) string
			// func fmt.Sscan(str string, a ...any) (n int, err error)
			// func fmt.Sscanf(str string, format string, a ...any) (n int, err error)
			// func fmt.Sscanln(str string, a ...any) (n int, err error)
			// type fmt.State interface{Flag(c int) bool; Precision() (prec int, ok bool); Width() (wid int, ok bool); Write(b []byte) (n int, err error)}
			// type fmt.Stringer interface{String() string}
			// type fmt.stringReader string
		})
	}
}

func FilterCompatibleTypes(target types.Type, set []types.Type) (comptbl []types.Type) {
	for _, typ := range set {
		if types.Identical(target, typ) {
			comptbl = append(comptbl, typ)
		}
	}
	return
}

func TestFindExpressionOfCommons(t *testing.T) {
	var content = `package main
	
var boolean = true
var integer = 0
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", content, parser.AllErrors)
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("main", fset, []*ast.File{file}, nil)
	if err != nil {
		t.Fatal(fmt.Errorf("check: %w", err))
	}

	for _, name := range pkg.Scope().Names() {
		obj := pkg.Scope().Lookup(name)
		typ := obj.Type()
		fmt.Printf("%s => (%T) %s\n", name, typ, typ)
	}
	// Output:
	// boolean => (*types.Basic) bool
	// integer => (*types.Basic) int
}

func TestFindingPkgScopeInInfo(t *testing.T) {
	for _, tc := range tcs1 {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, _, pkg, err := prepare(tc)
			if err != nil {
				t.Fatal(fmt.Errorf("prep: %w", err))
			}
			fmt.Println(pkg.Scope().Names())
			for _, name := range pkg.Scope().Names() {
				obj := pkg.Scope().Lookup(name)
				typ := obj.Type()
				fmt.Println(typ)
			}
			// Output: [Indent Reverse]
			// func(in string) string
			// func(in string) string
		})
	}
}

func TestIdenticalTypes(t *testing.T) {
	for _, tc := range tcs1 {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, _, pkg, err := prepare(tc)
			if err != nil {
				t.Fatal(fmt.Errorf("prep: %w", err))
			}
			fmt.Println(pkg.Scope().Names())
			for _, name := range pkg.Scope().Names() {
				obj := pkg.Scope().Lookup(name)
				typ := obj.Type()
				fmt.Println(typ)
			}

			id := types.Identical(
				pkg.Scope().Lookup(pkg.Scope().Names()[0]).Type(),
				pkg.Scope().Lookup(pkg.Scope().Names()[1]).Type(),
			)
			fmt.Println(id)

			// Output: [WordReverse]
		})
	}
}

func TestFindingFuncScopeInInfo(t *testing.T) {
	for _, tc := range tcs1 {
		t.Run(tc.name, func(t *testing.T) {
			_, funcdecl, _, info, _, err := prepare(tc)
			if err != nil {
				t.Fatal(fmt.Errorf("prep: %w", err))
			}

			funcscope := info.Scopes[funcdecl.Type]
			fmt.Println(funcscope.Names()) // Output: [in]
		})
	}
}

func TestFuncSignature(t *testing.T) {
	for _, tc := range tcs1 {
		t.Run(tc.name, func(t *testing.T) {
			_, funcdecl, _, info, _, err := prepare(tc)
			if err != nil {
				t.Fatal(fmt.Errorf("prep: %w", err))
			}

			returntype := info.Defs[funcdecl.Name].Type().(*types.Signature).Results()

			fmt.Println(returntype)
			// Output:
		})
	}
}

// func Test(t *testing.T) {
// 	file, funcdecl, spot, info, pkg, err := prepare(tc)
// 	if err != nil {
// 		t.Fatal(fmt.Errorf("prep: %w", err))
// 	}
// }
