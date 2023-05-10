package cp_printer

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/kylelemons/godebug/diff"
	"github.com/pkg/errors"
)

const testFilePath = "../../../internal/test-package/walk.go"

func Test_Print(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), testFilePath, nil, parser.AllErrors)
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}
	Println(f)
}

// this function tests if the Fprint can produce code in valid Go syntax by putting it into a Go file, compile and run
func Test_ReverseCheck(t *testing.T) {
	content, err := os.ReadFile(testFilePath)
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	f, err := parser.ParseFile(token.NewFileSet(), testFilePath, nil, parser.AllErrors)
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	module, err := os.MkdirTemp(os.TempDir(), "Test_ReverseCheck")
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	file, err := os.Create(filepath.Join(module, "main.go"))
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	fmt.Fprintf(file, `package main

import (
	"os"
	"go/printer"
	"go/ast"
	"go/token"
)

var f = `)
	Fprint(file, f)
	fmt.Fprintf(file, `


func main() {
	printer.Fprint(os.Stdout, token.NewFileSet(), f)
}`)

	// fmt.Println(module)

	cmd := exec.Command("go", "mod", "init", "na/na")
	cmd.Dir = module
	err = cmd.Run()
	if err != nil {
		t.Error(errors.Wrapf(err, "performing module creation"))
	}

	stdout := bytes.NewBuffer([]byte{})

	cmd = exec.Command("go", "run", ".")
	cmd.Dir = module
	cmd.Stdout = stdout
	cmd.Run()
	if err != nil {
		t.Error(errors.Wrapf(err, "performing running test program"))
	}

	fmt.Println(diff.Diff(string(content), stdout.String()))
}
