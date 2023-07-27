package evaluation

import (
	tester_package "tde/internal/folders/inject"

	"fmt"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func prepare(t *testing.T) {
	ti := &tester_package.TestInfo{
		TargetPackageImportPath: "tde/examples/word-reverse",
		TestFunctionName:        "TDE_WordReverse",
	}

	if err := tester_package.Inject("../../examples/word-reverse", ti); err != nil {
		t.Error(errors.Wrap(err, "returned error"))
	}

	if _, err := os.OpenFile("../../examples/word-reverse/tde/main_tde.go", os.O_RDONLY, os.ModeAppend); err != nil {
		t.Error(errors.Wrap(err, "validation"))
	}
}

func Test_Compile(t *testing.T) {
	prepare(t)

	runner := Runner{}
	output, err := runner.compile("../../examples/word-reverse/tde", "00000000-0000-0000-0000-000000000000")
	if err != nil {
		t.Error(errors.Wrapf(err, "in testing"))
	}

	fmt.Println(output)

	if err := os.RemoveAll("../../examples/word-reverse/tde"); err != nil {
		t.Error(errors.Wrap(err, "cleanup"))
	}
}
