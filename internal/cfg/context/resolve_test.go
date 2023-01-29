package context

import (
	"tde/internal/astw"
	"testing"

	"github.com/pkg/errors"
)

func Test_Resolve(t *testing.T) {
	_, astPkgs, err := astw.LoadDir("../../astw")
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed on loading package"))
	}
	astPkg := astPkgs["astw"]
	funcDecl, err := astw.FindFuncDecl(astPkg, "WalkWithNils")
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed on preparation"))
	}

	Resolve(astPkg, funcDecl)

}
