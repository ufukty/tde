package copymod

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
)

func TestModule(t *testing.T) {
	dst := filepath.Join(os.TempDir(), "tde_test_folders_copy_copy_module")
	fmt.Println("destination:", dst)

	err := Copy("../../../", dst, true, DefaultSkipDirs)
	if err != nil {
		t.Error(errors.Wrapf(err, "return value"))
	}
}
