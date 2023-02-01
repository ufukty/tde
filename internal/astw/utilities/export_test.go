package utilities

import (
	"testing"

	"github.com/pkg/errors"
)

func Test_String(t *testing.T) {
	_, err := String(TEST_TREE)
	if err != nil {
		t.Error(errors.Wrapf(err, "failed String"))
	}
}
