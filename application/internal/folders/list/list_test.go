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
	for name, pkg := range m {
		fmt.Println("*", name, " ", pkg.ImportPath, " ", pkg.Dir)
	}
}
