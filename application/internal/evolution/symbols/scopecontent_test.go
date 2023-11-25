package symbols

import (
	"fmt"
	"go/ast"
	"os"
	"reflect"
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
