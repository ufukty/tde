package list

import (
	"fmt"
	"testing"
)

func Test_AssertThisPackage(t *testing.T) {
	m, err := ListAllPackages(".")
	if err != nil {
		t.Error(fmt.Errorf("act: %w", err))
	}
	if _, ok := m["tde/internal/evolution/evaluation/list"]; !ok {
		t.Error(fmt.Errorf("assert: %w", err))
	}
}

func Test_ListPackages(t *testing.T) {
	m, err := ListAllPackages("../../../")
	if err != nil {
		t.Error(fmt.Errorf("act: %w", err))
	}
	if len(m) == 0 {
		t.Fatal("assert, expecting more results")
	}
	for name, pkg := range m {
		fmt.Println("*", name, " ", pkg.ImportPath, " ", pkg.Dir)
	}
}

func Test_FindImportPath(t *testing.T) {
	testCases := map[string]string{
		".":        "tde/internal/evolution/evaluation/list",
		"testdata": "tde/internal/evolution/evaluation/list/testdata",
	}

	for input, want := range testCases {
		pkgs, err := ListPackagesInDir(input)
		if err != nil {
			t.Fatal(fmt.Errorf("action: %w", err))
		}
		got := pkgs.First().ImportPath
		if got != want {
			t.Fatal(fmt.Errorf("assert. want %q got %q", want, got))
		}
	}
}
