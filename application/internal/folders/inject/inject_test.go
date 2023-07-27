package inject

import (
	"os"
	"testing"

	"github.com/pkg/errors"
)

func Test_Inject(t *testing.T) {
	ti := &TestInfo{
		TargetPackageImportPath: "tde/examples/word-reverse",
		TestFunctionName:        "TDE_WordReverse",
	}

	if err := Inject("../../../examples/word-reverse", ti); err != nil {
		t.Error(errors.Wrap(err, "returned error"))
	}

	if _, err := os.OpenFile("../../../examples/word-reverse/tde/main_tde.go", os.O_RDONLY, os.ModeAppend); err != nil {
		t.Error(errors.Wrap(err, "validation"))
	}

	if err := os.RemoveAll("../../../examples/word-reverse/tde"); err != nil {
		t.Error(errors.Wrap(err, "cleanup"))
	}
}
