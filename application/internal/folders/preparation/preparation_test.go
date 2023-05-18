package preparation

import (
	"fmt"
	"path/filepath"
	"tde/internal/folders/types"
	"testing"

	"github.com/pkg/errors"
)

func Test_Preparation(t *testing.T) {
	abs, err := filepath.Abs("../../../")
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	dupl, err := Prepare(
		types.AbsolutePath(abs),
		types.InModulePath("examples/word-reverse"),
		"tde/examples/word-reverse",
		"TDE_WordReverse",
	)

	if err != nil {
		t.Error(errors.Wrapf(err, ""))
	}
	fmt.Println("dst:", dupl)
}
