package tester_package

import (
	"os"
	"testing"

	"github.com/pkg/errors"
)

func Test_Inject(t *testing.T) {
	ti := &TestInfo{
		TargetPackageImportPath: "tde/examples/word_reverse",
		TestFunctionName:        "TDE_WordReverse",
	}

	if err := Create("../../../examples/word_reverse", ti); err != nil {
		t.Error(errors.Wrap(err, "returned error"))
	}

	if _, err := os.OpenFile("../../../examples/word_reverse/tde/main_tde.go", os.O_RDONLY, os.ModeAppend); err != nil {
		t.Error(errors.Wrap(err, "validation"))
	}

	if err := os.RemoveAll("../../../examples/word_reverse/tde"); err != nil {
		t.Error(errors.Wrap(err, "cleanup"))
	}
}
