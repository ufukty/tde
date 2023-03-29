package discovery

import (
	"fmt"
	"os"
	"tde/internal/utilities"
	"testing"

	"github.com/pkg/errors"
)

func Test_FindImportPath(t *testing.T) {
	var (
		path string
		err  error
	)

	path, err = GetImportPathOfPackage()
	if err != nil {
		t.Error(errors.Wrap(err, "failed on finding import path"))
	} else if path != "tde/internal/folders/discovery" {
		t.Error(errors.Wrapf(err, "got the wrong import path '%s'", path))
	}

	fmt.Println("Import path for package:", path)

	initialDir, _ := utilities.CurrentDir()
	err = os.Chdir("../../../examples/word_reverse")
	if err != nil {
		t.Error(errors.Wrap(err, "chdir is failed"))
	}
	defer func() {
		err := os.Chdir(initialDir)
		if err != nil {
			t.Error(errors.Wrapf(err, "failed to revert working directory: %q", initialDir))
		}
	}()

	path, err = GetImportPathOfPackage()
	if err != nil {
		t.Error(errors.Wrap(err, "failed on finding import path"))
	} else if path != "tde/examples/word_reverse" {
		t.Error(errors.Wrapf(err, "got the wrong import path '%s'", path))
	}

	fmt.Println("Import path for package:", path)
}

func Test_DetectTestFunctions(t *testing.T) {
	fns, err := DiscoverTestFileForTestFuncDetails("../../../", "../../../examples/word_reverse/word_reverse_tde.go")
	if err != nil {
		t.Error(errors.Wrapf(err, "failed to detect positions and names of test functions that is in the user-provided test file"))
	} else if len(fns) != 1 {
		t.Error("Got wrong number of results:", len(fns))
	} else if fns[0].Name != "TDE_WordReverse" {
		t.Errorf("Want 'TDE_Word_Reverse' got '%s'", fns[0].Name)
	}
	fmt.Println(fns)
}

func Test_GetPackagePathInModule(t *testing.T) {
	mod, err := GetModulePath()
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	pkgRelMod, err := GetPackagePathInModule(mod)
	if err != nil {
		t.Error(errors.Wrapf(err, "return"))
	}

	fmt.Println(pkgRelMod)
}
