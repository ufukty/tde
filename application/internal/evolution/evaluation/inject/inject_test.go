package inject

import (
	"fmt"
	"os"
	"testing"
)

func Test_Inject(t *testing.T) {
	defer os.RemoveAll("testdata/tde") // clean up

	if err := Inject("testdata", &TestInfo{
		TargetPackageImportPath: "tde/internal/evolution/evaluation/inject/testdata",
		TestFunctionName:        "TDE_WordReverse",
	}); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}

	_, err := os.Stat("testdata/tde/main_tde.go")
	if err != nil {
		t.Fatal(fmt.Errorf("assert: %w", err))
	}
}
