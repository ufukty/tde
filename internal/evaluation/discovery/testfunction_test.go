package discovery

import (
	"fmt"
	"testing"
)

func Test_TestFunctionsInFile(t *testing.T) {
	fns, err := TestFunctionsInFile("testdata/words_tde.go")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to detect positions and names of test functions that is in the user-provided test file: %q", err))
	} else if len(fns) != 1 {
		t.Fatal("Got wrong number of results:", len(fns))
	} else if fns[0].Name != "TDE_WordReverse" {
		t.Fatalf("Want 'TDE_Word_Reverse' got %q", fns[0].Name)
	}
	fmt.Println(fns)
}
