package inject

import (
	"os"
	"testing"

	"github.com/pkg/errors"
)

func Test_Inject(t *testing.T) {
	ti := &TestInfo{
		TargetPackageImportPath: "tde/examples/words",
		TestFunctionName:        "TDE_WordReverse",
	}

	if err := Inject("../../../examples/words", ti); err != nil {
		t.Error(errors.Wrap(err, "returned error"))
	}

	if _, err := os.OpenFile("../../../examples/words/tde/main_tde.go", os.O_RDONLY, os.ModeAppend); err != nil {
		t.Error(errors.Wrap(err, "validation"))
	}

	if err := os.RemoveAll("../../../examples/words/tde"); err != nil {
		t.Error(errors.Wrap(err, "cleanup"))
	}
}
