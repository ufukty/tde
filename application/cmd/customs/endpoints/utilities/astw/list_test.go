package astw

import (
	"testing"

	"github.com/pkg/errors"
)

func Test_ListPackages(t *testing.T) {
	var m, err = ListPackages(".")
	if err != nil {
		t.Error(errors.Wrapf(err, "act"))
	}
	if _, ok := (*m)["tde/cmd/customs/endpoints/utilities/astw"]; !ok {
		t.Error(errors.Wrapf(err, "assert"))
	}
}
