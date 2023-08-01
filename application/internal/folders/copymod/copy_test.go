package copymod

import (
	"fmt"
	"os"
	"testing"
)

func TestModule(t *testing.T) {
	var dst, err = os.MkdirTemp(os.TempDir(), "tde_test_folders_copy_copy_module")
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	fmt.Println("created temp dir:", dst)

	if err = CopyModule(dst, "../../../", true, []string{}, []string{}, []string{}, true); err != nil {
		t.Fatal(fmt.Errorf("action: %w", err))
	}
}
