package symbols

import (
	"fmt"
	"go/ast"
	"io"
	"os"
	"reflect"
)

func ExampleScopeContent_Markdown() {
	_, _, _, _, pkg, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}

	f, err := os.Create("output.md")
	if err != nil {
		panic(fmt.Errorf("prep, file: %w", err))
	}
	defer f.Close()

	io.Copy(f, PrintPackageAsMarkdown(pkg, 3))
	// Output:
}

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
