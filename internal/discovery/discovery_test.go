package discovery

import (
	"fmt"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func Test_FindImportPath(t *testing.T) {
	var (
		path string
		err  error
	)

	path, err = FindImportPathOfThePackage()
	if err != nil {
		t.Error(errors.Wrap(err, "failed on finding import path"))
	} else if path != "tde/internal/discovery" {
		t.Error(errors.Wrapf(err, "got the wrong import path '%s'", path))
	}

	fmt.Println("Import path for package:", path)

	err = os.Chdir("../../examples/word-reverse/word_reverse")
	if err != nil {
		t.Error(errors.Wrap(err, "chdir is failed"))
	}

	path, err = FindImportPathOfThePackage()
	if err != nil {
		t.Error(errors.Wrap(err, "failed on finding import path"))
	} else if path != "tde/examples/word-reverse/word_reverse" {
		t.Error(errors.Wrapf(err, "got the wrong import path '%s'", path))
	}

	fmt.Println("Import path for package:", path)
}

func Test_FindModulePath(t *testing.T) {

	fmt.Println(FindModulePath())
}
