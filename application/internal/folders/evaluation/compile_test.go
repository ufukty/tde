package evaluation

import (
	"os"
	"path/filepath"
	"tde/internal/folders/inject"
	"tde/internal/folders/list"

	"fmt"
	"testing"
)

func Test_Compile(t *testing.T) {
	pkgs, err := list.ListPackagesInDir("../../../examples/word-reverse")
	if err != nil {
		t.Fatal(fmt.Errorf("arrange, listing packages in target dir: %w", err))
	}
	sample, err := inject.WithCreatingSample("../../../", pkgs.First(), "TDE_WordReverse")
	if err != nil {
		t.Fatal(fmt.Errorf("arrange, creating sample: %w", err))
	}
	fmt.Println("sample dir:", sample)
	path := filepath.Join(sample, "examples/word-reverse/tde")
	fmt.Println("test-runner package path:", path)
	runner := Runner{}
	output, err := runner.compile(path, "00000000-0000-0000-0000-000000000000")
	if err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}
	fmt.Println("Combined output:", output)
	byteContents, err := os.ReadFile(filepath.Join(path, "results.json"))
	if err != nil {
		t.Fatal(fmt.Errorf("assert prep: %w", err))
	}
	contents := string(byteContents)
	if contents == "" {
		t.Fatal(fmt.Errorf("assert. results.json is empty"))
	}
	fmt.Println(contents)
}
