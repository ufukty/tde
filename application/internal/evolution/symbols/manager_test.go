package symbols

import (
	"fmt"
	"go/ast"
	"go/types"
	"os"
	"tde/internal/astw/astwutl"
	"tde/internal/utilities/strw"
	"testing"

	"golang.org/x/tools/go/packages"
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
			for typ, symbols := range sm.Context.ByType {
				fmt.Println(typ)
				for _, symbol := range symbols {
					fmt.Printf("  %s", symbol)
				}
			}
		})
	}
}

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

					types.AssignableTo()
				}

				return true
			})
		}

	}
}
