package discovery

import (
	"tde/internal/folders/list"

	"fmt"
	"testing"
)

func Test_FindImportPath(t *testing.T) {
	var testCases = map[string]string{
		".":                              "tde/internal/folders/discovery",
		"../../../examples/word-reverse": "tde/examples/word-reverse",
	}

	for input, want := range testCases {
		var pkgs, err = list.ListPackagesInDir(input)
		if err != nil {
			t.Fatal(fmt.Errorf("action: %w", err))
		}
		got := pkgs.First().ImportPath
		if got != want {
			t.Fatal(fmt.Errorf("assert. want %q got %q", want, got))
		}
	}
}

func Test_DetectTestFunctions(t *testing.T) {
	fns, err := DiscoverTestFileForTestFuncDetails("../../../", "../../../examples/word-reverse/word_reverse_tde.go")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to detect positions and names of test functions that is in the user-provided test file: %q", err))
	} else if len(fns) != 1 {
		t.Fatal("Got wrong number of results:", len(fns))
	} else if fns[0].Name != "TDE_WordReverse" {
		t.Fatalf("Want 'TDE_Word_Reverse' got %q", fns[0].Name)
	}
	fmt.Println(fns)
}

func Test_GetPackagePathInModule(t *testing.T) {
	mod, err := GetModulePath()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	pkgRelMod, err := GetPackagePathInModule(mod)
	if err != nil {
		t.Fatal(fmt.Errorf("return: %w", err))
	}

	fmt.Println(pkgRelMod)
}
