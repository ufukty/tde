package test_package

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
