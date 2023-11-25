package symbols

import (
	"fmt"
	"io"
	"os"
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
