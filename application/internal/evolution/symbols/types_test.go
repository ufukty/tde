package symbols

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"tde/internal/astw/astwutl"
	"tde/internal/astw/traced"
	"testing"
)

var tcases = map[string]string{
	"basic":     "testdata/words",
	"populated": "testdata/evolution",
}

func Test_SymbolsManager(t *testing.T) {
	for tname, tcase := range tcases {
		t.Run(tname, func(t *testing.T) {
			sm, err := NewSymbolsManager(tcase)
			if err != nil {
				t.Fatal(fmt.Errorf("prep: %w", err))
			}

			fmt.Println("package:")
			for idt, typ := range sm.Context.Package {
				fmt.Printf("  %s: %s\n", idt.Name, typ.String())
			}
		})
	}
}

var testdatafolders = [][]string{
	{"testdata/evolution/walk.go", "WalkWithNils"},
	{"testdata/words/words.go", "Reverse"},
}

var testcase int64 = 0

func prepare() (*ast.File, *ast.FuncDecl, ast.Node, *types.Info, *types.Package, error) {
	path, funcname := testdatafolders[testcase][0], testdatafolders[testcase][1]
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("parser: %w", err)
	}
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
	pkg, err := conf.Check("main", fset, []*ast.File{file}, info)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("check: %w", err)
	}
	funcdecl, err := astwutl.FindFuncDecl(file, funcname)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("find func decl: %w", err)
	}
	spot := funcdecl.Body.List[0]
	return file, funcdecl, spot, info, pkg, nil
}

func Example_FindFuncScope() {
	_, funcdecl, _, info, _, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}
	funcscope := info.Defs[funcdecl.Name].Parent()
	fmt.Println("funcscope", funcscope)
}

func Test_ConvertScopesMap(t *testing.T) {
	_, _, _, info, _, err := prepare()
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
}

func Test_Trace(t *testing.T) {
	file, _, spot, info, _, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	parents := traced.Parents(file, spot)
	fmt.Println("parents:", parents)
	for _, parent := range parents {
		fmt.Println("scope for a parent:", info.Scopes[parent])
	}
}

func Test_PkgScope(t *testing.T) {
	_, _, _, _, pkg, err := prepare()
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

}

func ExampleFindingIdentsDefinedInImports() {
	_, _, _, _, pkg, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}

	for _, ipkg := range pkg.Imports() {
		// fmt.Println(ipkg.Name(), ipkg.Scope().Names())
		qual := types.RelativeTo(pkg)
		for _, symbolname := range ipkg.Scope().Names() {
			symbol := ipkg.Scope().Lookup(symbolname)
			str := types.ObjectString(symbol, qual)
			fmt.Println(str)
		}
	}
	// Output: [Append Appendf Appendln Errorf FormatString Formatter Fprint Fprintf Fprintln Fscan Fscanf Fscanln GoStringer Print Printf Println Scan ScanState Scanf Scanln Scanner Sprint Sprintf Sprintln Sscan Sscanf Sscanln State Stringer stringReader]
}

func FilterCompatibleTypes(target types.Type, set []types.Type) (comptbl []types.Type) {
	for _, typ := range set {
		if types.Identical(target, typ) {
			comptbl = append(comptbl, typ)
		}
	}
	return
}

func ExampleFindExpressionOfCommons() {
	var content = `package main
	
	var boolean = true
	var integer = 0
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", content, parser.AllErrors)
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("main", fset, []*ast.File{file}, nil)
	if err != nil {
		panic(fmt.Errorf("check: %w", err))
	}

	fmt.Println(pkg.Scope().Names())
	for _, name := range pkg.Scope().Names() {
		obj := pkg.Scope().Lookup(name)
		typ := obj.Type()
		fmt.Println(typ)

	}
	// Output: [boolean integer]
	// bool
	// int
}

func ExampleFindingPkgScopeInInfo() {
	_, _, _, _, pkg, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
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
}

func ExampleIdenticalTypes() {
	_, _, _, _, pkg, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
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
}

func ExampleFindingFuncScopeInInfo() {
	_, funcdecl, _, info, _, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}

	funcscope := info.Scopes[funcdecl.Type]
	fmt.Println(funcscope.Names()) // Output: [in]
}

func ExampleFuncSignature() {
	_, funcdecl, _, info, _, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}

	returntype := info.Defs[funcdecl.Name].Type().(*types.Signature).Results()

	fmt.Println(returntype)
	// Output:
}

// func Example() {
// 	file, funcdecl, spot, info, pkg, err := prepare()
// 	if err != nil {
// 		panic(fmt.Errorf("prep: %w", err))
// 	}
// }
