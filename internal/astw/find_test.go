package astw

import (
	"testing"

	"github.com/pkg/errors"
)

func Test_FindFuncDecl(t *testing.T) {
	_, astNode, err := ParseString(TEST_FILE)
	if err != nil {
		t.Error(errors.Wrapf(err, "failed ParseString"))
	}

	funcDecl, err := FindFuncDecl(astNode, "Addition")
	if err != nil {
		t.Error(errors.Wrapf(err, "failed FindFuncDecl"))
	}

	if funcDecl.Name.Name != "Addition" {
		t.Error(errors.Wrapf(err, "failed name check"))
	}
}
