package astcfg

import (
	astw "tde/internal/ast_wrapper"
	"testing"

	"github.com/pkg/errors"
)

func Test_Develop(t *testing.T) {
	_, astFile, err := astw.LoadFile("../../ast_wrapper/walk.go")
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed on prep"))
	}
	funcDecl, err := astw.FindFuncDecl(astFile, "WalkWithNils")
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed on prep"))
	}

	Develop(astFile, funcDecl)
}
