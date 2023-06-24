package astw

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func Test_ListPackages(t *testing.T) {
	var m, err = ListPackages(".")
	if err != nil {
		t.Error(errors.Wrapf(err, "act"))
	}
	for pkgName, pkg := range *m {
		fmt.Println(pkgName, pkg.Dir)
	}
}
