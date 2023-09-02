package list

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func Test_AssertThisPackage(t *testing.T) {
	var m, err = ListAllPackages(".")
	if err != nil {
		t.Error(errors.Wrapf(err, "act"))
	}
	if _, ok := m["tde/internal/folders/list"]; !ok {
		t.Error(errors.Wrapf(err, "assert"))
	}
}

func Test_ListPackages(t *testing.T) {
	var m, err = ListAllPackages("../../../")
	if err != nil {
		t.Error(errors.Wrapf(err, "act"))
	}
	if len(m) == 0 {
		t.Fatal("assert, expecting more results")
	}
	for name, pkg := range m {
		fmt.Println("*", name, " ", pkg.ImportPath, " ", pkg.Dir)
	}
}

func Test_FindImportPath(t *testing.T) {
	var testCases = map[string]string{
		".":                       "tde/internal/folders/list",
		"../../../examples/words": "tde/examples/words",
	}

	for input, want := range testCases {
		var pkgs, err = ListPackagesInDir(input)
		if err != nil {
			t.Fatal(fmt.Errorf("action: %w", err))
		}
		got := pkgs.First().ImportPath
		if got != want {
			t.Fatal(fmt.Errorf("assert. want %q got %q", want, got))
		}
	}
}
