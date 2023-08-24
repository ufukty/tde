package astwutl

import (
	"testing"

	"github.com/pkg/errors"
)

func Test_ParseString(t *testing.T) {
	_, _, err := ParseString(TEST_FILE)
	if err != nil {
		t.Error(errors.Wrapf(err, "failed ParseString"))
	}
}

func Test_LoadDir(t *testing.T) {
	_, _, err := LoadDir("testdata")
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed on loading dir"))
	}
}
