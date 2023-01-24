package code

import (
	"os"
	"testing"

	"github.com/pkg/errors"
)

func TestPrint(t *testing.T) {
	var err error
	code := Code{}

	err = code.LoadFromFile("code_test.go")
	if err != nil {
		t.Error(errors.Wrap(err, "could not load the code from file"))
	}

	code.Print(os.Stdout)
}

func TestPrintListOfTokens(t *testing.T) {
	var err error
	code := Code{}

	err = code.LoadFromFile("code_test.go")
	if err != nil {
		t.Error(errors.Wrap(err, "could not load the code from file"))
	}

	code.PrintListOfTokens()
}

func TestRenameFunction(t *testing.T) {
	var err error
	code := Code{}

	err = code.LoadFromFile("code_test.go")
	if err != nil {
		t.Error(errors.Wrap(err, "could not load the code from file"))
	}

	err = code.RenameFunction("HelloWorld", "TestyTest")
	if err == nil {
		t.Error(errors.Wrap(err, "should fail when inexisting function renamed"))
	}

	err = code.RenameFunction("TestRenameFunction", "TestyTest")
	if err != nil {
		t.Error(errors.Wrap(err, "should pass"))
	}
}

func TestCAASTCFG(t *testing.T) {
	txt := `package main
	
	import "fmt"

	func Main() {
		return
	}
	`

	c := Code{}
	c.LoadFromString(txt)
	// f, _ := c.FindFunction("Main")

	// caast := caast.CAAST{}
	// caast.Develop()

	// c.Print(os.Stdout)
}
