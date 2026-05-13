package symbols

import (
	"fmt"
	"go/ast"
	"go/types"
	"maps"
	"os"
	"reflect"
	"slices"

	"tde/internal/ast/export"
)

func ExampleAstToScope() {
	f, _, _, info, _, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}

	ast.Inspect(f, func(n ast.Node) bool {
		if n == f || n == nil {
			return true
		}

		if i, ok := n.(*ast.Ident); ok {
			if def, ok := info.Defs[i]; ok && def != nil {
				fmt.Printf("%-15s => (%s)\n", reflect.TypeOf(n), def.Name())
			}
		}
		if scope, ok := info.Scopes[n]; ok {
			fmt.Printf("%-15s => %s\n", reflect.TypeOf(n), NewScopeContent(scope).Brief())
		}

		return true
	})
	// Output:
}

func ExamplePackageScopeWriteTo() {
	_, _, _, _, pkg, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}

	pkg.Scope().WriteTo(os.Stdout, 0, true)
	// Output:
}

func findIdent(objects map[*ast.Ident]types.Object, obj types.Object) (*ast.Ident, bool) {
	for k, v1 := range objects {
		if v1 == obj {
			return k, true
		}
	}
	return nil, false
}

func pop(slice []*types.Scope) ([]*types.Scope, *types.Scope) {
	return slice[:len(slice)-1], slice[len(slice)-1]
}

func ExampleScopeLookup() {
	p, _, _, info, pkg, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}

	fmt.Println("len", len(info.Scopes), len(info.Defs), len(info.Types))

	ss := []*types.Scope{pkg.Scope()}
	var s *types.Scope

	for len(ss) > 0 {
		ss, s = pop(ss)
		fmt.Println(findMeaningfulPathToScope(info, p, s))
		fmt.Println("    Names:", s.Names())
		for _, name := range s.Names() {
			obj := s.Lookup(name)
			def, ok := findIdent(info.Defs, obj)
			if ok {
				fmt.Printf("    %-20s => (%T) %s\n", obj, def, def)
			}
		}
		for i := 0; i < s.NumChildren(); i++ {
			ss = append(ss, s.Child(i))
		}

	}

	ns, ss := slices.Collect(maps.Keys(info.Scopes)), slices.Collect(maps.Values(info.Scopes))
	for i := 0; i < 20; i++ {
		fmt.Printf(">>>%T %s\nDefines the scope: %s\n\n", ns[i], ns[i], ss[i])
	}

	// Output:
}

func ExampleExprTypes() {
	p, _, _, info, _, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}

	ast.Inspect(p, func(n ast.Node) bool {
		if e, ok := n.(ast.Expr); ok {
			if t, ok := info.Types[e]; ok {
				pr, err := export.String(e)
				if err != nil {
					panic(fmt.Errorf("printing the expression: %w", err))
				}
				if ft, ok := t.Type.(*types.Signature); ok {
					fmt.Printf("%-20T %-30s %-50s", e, pr, t.Type.String())
					os := ft.Results()
					for i := 0; i < os.Len(); i++ {
						v := os.At(i)
						fmt.Printf(" (%s) %s", v.Type(), v.Name())
					}
					fmt.Println()
				} else {
					// fmt.Printf("%-20T %-50s %s\n", e, t.Type.String(), t.Value)
				}
			}
		}

		return true
	})

	// Output:
}
