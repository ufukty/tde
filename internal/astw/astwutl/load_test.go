package astwutl

import (
	"fmt"
	"testing"
)

func Test_ParseString(t *testing.T) {
	_, _, err := ParseString(TEST_FILE)
	if err != nil {
		t.Error(fmt.Errorf("failed ParseString: %w", err))
	}
}

func Test_LoadDir(t *testing.T) {
	_, _, err := LoadDir("testdata")
	if err != nil {
		t.Error(fmt.Errorf("Failed on loading dir: %w", err))
	}
}
