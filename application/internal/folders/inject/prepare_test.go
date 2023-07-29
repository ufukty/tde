package inject

import (
	"fmt"
	"tde/internal/folders/list"
	"testing"
)

func Test_Preparation(t *testing.T) {
	var (
		pkgs list.Packages
		dupl string
		err  error
	)
	if pkgs, err = list.ListPackagesInDir("../../../examples/word-reverse"); err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	if dupl, err = WithCreatingSample(pkgs.First().Module.Dir, pkgs.First(), "TDE_WordReverse"); err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
	}
	fmt.Println("dst:", dupl)
}
