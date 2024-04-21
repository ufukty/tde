package main

import (
	"go/parser"
	"go/token"
	"log"
	"testing"

	"github.com/pkg/errors"
)

func TestJSONPrinter(t *testing.T) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "testdata/walk.go", nil, parser.ParseComments|parser.AllErrors)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "prepare"))
	}

	err = jsonPrinter(astFile)
	if err != nil {
		t.Error(errors.Wrapf(err, "evaluation"))
	}
}
