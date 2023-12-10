package symbols

import (
	"fmt"
	"go/ast"
	"os"
	"tde/internal/astw/astwutl"
	"tde/internal/utilities/strw"
	"testing"

	"golang.org/x/tools/go/packages"
)

func Test_Packages(t *testing.T) {
	cfg := &packages.Config{
		Mode: packages.NeedDeps |
			packages.NeedImports |
			packages.NeedSyntax |
			packages.NeedTypes |
			packages.NeedTypesInfo |
			packages.NeedFiles,
		BuildFlags: []string{"-tags", "tde"},
	}
	pkgs, err := packages.Load(cfg, "./testdata")
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			ast.Inspect(file, func(n ast.Node) bool {
				// if e, ok := n.(ast.Expr); ok {
				// 	ty := pkg.TypesInfo.TypeOf(e)
				// 	p, err := astwutl.String(e)
				// 	if err != nil {
				// 		t.Fatal(fmt.Errorf("printing: %w", err))
				// 	}
				// 	fmt.Printf("(%T) %s\n%s\n\n", ty, ty, strw.IndentLines(p, 3))
				// }

				switch n := n.(type) {
				case *ast.IfStmt:
					e := n.Cond
					ty := pkg.TypesInfo.TypeOf(e)
					p, err := astwutl.String(e)
					if err != nil {
						t.Fatal(fmt.Errorf("printing: %w", err))
					}
					fmt.Printf("(%T) %s\n%s\n\n", ty, ty, strw.IndentLines(p, 3))
				}

				return true
			})
		}

	}
}
