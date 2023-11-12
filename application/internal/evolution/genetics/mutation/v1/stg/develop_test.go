package stg

import (
	"bytes"
	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone"
	"tde/internal/utilities"

	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"reflect"
	"testing"

	"github.com/kr/pretty"
	"github.com/pkg/errors"
)

func loadTestPackage() (*ast.Package, *ast.File, *ast.FuncDecl, error) {
	_, astPkgs, err := astwutl.LoadDir("testdata")
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "could not load test package")
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["testdata/walk.go"]
	funcDecl, err := astwutl.FindFuncDecl(astPkg, "WalkWithNils")
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "could not find test function")
	}
	return astPkg, astFile, funcDecl, nil
}

func fprintSafe(fdecl *ast.FuncDecl) (code []byte, err error) {
	code = []byte{}
	buffer := bytes.NewBuffer([]byte{})
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %q", r)
		}
	}()
	if printer.Fprint(buffer, token.NewFileSet(), fdecl) != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Test_Develop(t *testing.T) {
	astPkg, astFile, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on prep"))
	}

	subjectFuncDecl := clone.FuncDecl(originalFuncDecl)
	newNode, err := Develop(astPkg, astFile, subjectFuncDecl, 1)
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed on Develop"))
	}
	fmt.Println("typeOf: ", reflect.TypeOf(newNode))
	if astwutl.CompareRecursivelyWithAddresses(subjectFuncDecl, originalFuncDecl) == true {
		pretty.Println(newNode)
		pretty.Println(subjectFuncDecl.Body)
		t.Error("Failed to see change on subject")
	}
}

func Benchmark_Develop(b *testing.B) {
	astPkg, astFile, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		b.Error(errors.Wrapf(err, "failed on prep"))
	}

	for i := 0; i < b.N; i++ {
		subjectFuncDecl := clone.FuncDecl(originalFuncDecl)
		newNode, err := Develop(astPkg, astFile, subjectFuncDecl, 1)
		if err != nil {
			b.Error(errors.Wrapf(err, "Failed on Develop"))
		}
		if astwutl.CompareRecursivelyWithAddresses(subjectFuncDecl, originalFuncDecl) == true {
			if _, ok := newNode.(*ast.BranchStmt); ok { // empty branch statement always leads fail in ast->code convertion
				continue
			}
			b.Errorf("Failed to see change on subject #%d\n", i)
		}
	}
}

func Test_DevelopProgressively(t *testing.T) {
	astPkg, astFile, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on prep"))
	}
	var best = originalFuncDecl
	for i := 0; i < 2000; i++ {
		subject := clone.FuncDecl(best)
		newNode, err := Develop(astPkg, astFile, subject, 20)
		if err != nil {
			t.Error(errors.Wrapf(err, "Failed on Develop i = %d, typeOf = %v", i, reflect.TypeOf(newNode)))
		}
		if code, err := fprintSafe(best); err == nil {
			fmt.Printf("Code:\n%s\n", utilities.IndentLines(string(code), 4))
			best = subject
		}
	}
}

func Test_DevelopFindUnbreakingChange(t *testing.T) {
	astPkg, astFile, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on prep"))
	}

	nonBreakingChangeFound := false
	for i := 0; i < 200; i++ {
		subject := clone.FuncDecl(originalFuncDecl)
		Develop(astPkg, astFile, subject, 20)

		if code, err := fprintSafe(subject); err == nil {
			fmt.Printf("Code:\n%s\n", utilities.IndentLines(string(code), 4))
			nonBreakingChangeFound = true
		}
	}

	if !nonBreakingChangeFound {
		t.Error("No non-breaking subjects found.")
	}
}
