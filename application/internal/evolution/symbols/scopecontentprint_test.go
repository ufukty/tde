package symbols

import (
	"fmt"
	"go/types"
	"os"
)

func ExampleScopeContent_Markdown_Universe() {
	astpkg, _, _, info, _, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}

	f, err := os.Create("dump.universe.md")
	if err != nil {
		panic(fmt.Errorf("prep, file: %w", err))
	}
	defer f.Close()
	PrintPackageAsMarkdown(f, info, astpkg, types.Universe, 10)
	// Output:
}

func ExampleScopeContent_Markdown_Package() {
	astpkg, _, _, info, pkg, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}

	f, err := os.Create("dump.package.md")
	if err != nil {
		panic(fmt.Errorf("prep, file: %w", err))
	}
	defer f.Close()
	PrintPackageAsMarkdown(f, info, astpkg, pkg.Scope(), 10)
	// Output:
}
