package astwutl

import (
	"fmt"
	"testing"
)

func Test_FindFuncDecl(t *testing.T) {
	_, astNode, err := ParseString(TEST_FILE)
	if err != nil {
		t.Error(fmt.Errorf("failed ParseString: %w", err))
	}

	funcDecl, err := FindFuncDecl(astNode, "Addition")
	if err != nil {
		t.Error(fmt.Errorf("failed FindFuncDecl: %w", err))
	}

	if funcDecl.Name.Name != "Addition" {
		t.Error(fmt.Errorf("failed name check: %w", err))
	}
}
