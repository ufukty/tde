package symbols

import (
	"fmt"
	"go/ast"
)

func ExampleScopeIdent() {
	f, _, _, info, pkg, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}
	chd := pkg.Scope().Child(0).Child(1)
	fmt.Println(chd)
	for n, scp := range info.Scopes {
		if chd == scp {

			if ft, ok := n.(*ast.FuncType); ok {
				fd := findFuncTypeParent(f, ft)

				if fd != nil {
					fmt.Println(fd.Name)
				} else {
					fmt.Println(ft)
				}

			} else {
				fmt.Println(ft)
			}
		}
	}
	// Output:
}
