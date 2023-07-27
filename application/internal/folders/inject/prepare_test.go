package inject

import (
	"fmt"
	"path/filepath"
	"tde/internal/folders/list"
	"testing"

	"github.com/pkg/errors"
)

func Test_Preparation(t *testing.T) {
	abs, err := filepath.Abs("../../../")
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	var pkgInfo = &list.Package{
		ImportPath: "tde/examples/word-reverse",
	}
	dupl, err := WithCreatingSample(abs, "examples/word-reverse", pkgInfo, "TDE_WordReverse")
	if err != nil {
		t.Error(errors.Wrapf(err, ""))
	}
	fmt.Println("dst:", dupl)
}
