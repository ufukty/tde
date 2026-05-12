package astwutl

import (
	"fmt"
	"testing"
)

func Test_String(t *testing.T) {
	_, err := String(TEST_TREE)
	if err != nil {
		t.Error(fmt.Errorf("failed String: %w", err))
	}
}
