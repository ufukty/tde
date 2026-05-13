package ctxres

import (
	"fmt"
	"slices"
	"testing"

	"tde/internal/astw/astwutl"
	"tde/internal/astw/parse"
	"tde/internal/astw/traverse"
)

func Test_GetContextForSpot(t *testing.T) {
	_, astPkgs, err := parse.Dir("testdata")
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	astPkg := astPkgs["test_package"]
	funcDecl, err := astwutl.FindFuncDecl(astPkg, "WalkWithNils")
	if err != nil {
		t.Error(fmt.Errorf("Failed on preparation: %w", err))
	}

	tFuncDecl := traverse.GetTraversableNodeForASTNode(funcDecl)
	funcBody := traverse.GetTraversableNodeForASTNode(funcDecl.Body).GetTraversableSubnodes()
	choosenSpot := funcBody[len(funcBody)-1]

	ctx, err := GetContextForSpot(
		astPkg,
		tFuncDecl,
		choosenSpot,
	)
	if err != nil {
		t.Error(fmt.Errorf("getting context for spot: %w", err))
	}

	vars := []string{}
	for _, id := range ctx.Scopes[1].Variables {
		vars = append(vars, id.Name)
	}
	if !slices.Contains(vars, "root") {
		t.Error("validation: variable 'root' not found in context")
	}
	if !slices.Contains(vars, "callback") {
		t.Error("validation: variable 'callback' not found in context")
	}

	for i, scope := range ctx.Scopes {
		fmt.Println("scope:", i)
		for _, variable := range scope.Variables {
			fmt.Println("variable:", variable.Name)
		}
		for _, function := range scope.Functions {
			fmt.Println("function:", function.Name)
		}
		for _, import_ := range scope.Imports {
			fmt.Println("import:", import_.Path.Value)
		}
		for _, type_ := range scope.Types {
			fmt.Println("types:", type_.Name)
		}
		for _, method := range scope.Methods {
			fmt.Println("method:", method.Recv.List[0].Type, method.Name)
		}
		fmt.Println()
	}

	// fmt.Println("Context:\n", ctx)
}
