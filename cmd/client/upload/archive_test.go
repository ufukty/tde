package upload

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func Test_ArchiveDirectory(t *testing.T) {
	path, err := ArchiveDirectory("../../../")
	if err != nil {
		t.Error(errors.Wrapf(err, "failed to archive directory"))
	}
	fmt.Println("Output zip:", path)
}
