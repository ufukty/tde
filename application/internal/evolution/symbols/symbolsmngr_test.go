package symbols

import (
	"flag"
	"fmt"
	"os"
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
		Mode:       packages.NeedFiles | packages.NeedSyntax,
		BuildFlags: []string{"-tags", "tde"},
	}
	pkgs, err := packages.Load(cfg, flag.Args()...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	// Print the names of the source files
	// for each package listed on the command line.
	for _, pkg := range pkgs {
		fmt.Println(pkg.ID, pkg.GoFiles)
		if pkg.TypesInfo != nil {
			for def := range pkg.TypesInfo.Defs {
				fmt.Println(def)
			}
		}
	}
}
