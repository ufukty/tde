package discovery

import (
	"tde/internal/folders/list"
	"tde/internal/utilities"

	"fmt"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func Test_FindImportPath(t *testing.T) {
	var (
		pkgs list.Packages
		path string
		err  error
	)

	pkgs, err = list.ListPackagesInDir(".")
	if err != nil {
		t.Fatal(errors.Wrap(err, "failed on finding import path"))
	}
	if pkgs.First().ImportPath != "tde/internal/folders/discovery" {
		t.Fatal(fmt.Errorf("got the wrong import path %q: %w", path, err))
	}

	fmt.Println("Import path for package:", path)

	initialDir, _ := utilities.CurrentDir()
	err = os.Chdir("../../../examples/word-reverse")
	if err != nil {
		t.Fatal(errors.Wrap(err, "chdir is failed"))
	}
	defer func() {
		err := os.Chdir(initialDir)
		if err != nil {
			t.Fatal(fmt.Errorf("failed to revert working directory: %q: %w", initialDir, err))
		}
	}()

	path, err = GetImportPathOfPackage()
	if err != nil {
		t.Fatal(errors.Wrap(err, "failed on finding import path"))
	} else if path != "tde/examples/word-reverse" {
		t.Fatal(fmt.Errorf("got the wrong import path %q: %w", path, err))
	}

	fmt.Println("Import path for package:", path)
}

func Test_DetectTestFunctions(t *testing.T) {
	fns, err := DiscoverTestFileForTestFuncDetails("../../../", "../../../examples/word-reverse/word_reverse_tde.go")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to detect positions and names of test functions that is in the user-provided test file", err))
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
		t.Fatal(fmt.Errorf("prep", err))
	}

	pkgRelMod, err := GetPackagePathInModule(mod)
	if err != nil {
		t.Fatal(fmt.Errorf("return", err))
	}

	fmt.Println(pkgRelMod)
}
