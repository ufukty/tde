package discovery

import (
	"fmt"
	"testing"
)

func Test_CombinedForDir(t *testing.T) {
	c, err := CombinedDetailsForTest("testdata", "TDE_WordReverse")
	if err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}
	if c.Package.Name != "words" {
		t.Fatal(fmt.Errorf("assert. got %q. want %q", c.Package.Name, "words"))
	}
	if c.Package.ImportPath != "tde/internal/evolution/evaluation/discovery/testdata" {
		t.Fatal(fmt.Errorf("assert. got %q. want %q", c.Package.ImportPath, "tde/examples/words"))
	}
	if c.Target.Name != "WordReverse" {
		t.Fatal(fmt.Errorf("assert. got %q. want %q", c.Target.Name, "WordReverse"))
	}
	if c.Test.Name != "TDE_WordReverse" {
		t.Fatal(fmt.Errorf("assert. got %q. want %q", c.Target.Name, "TDE_WordReverse"))
	}
	fmt.Println(c)
}
