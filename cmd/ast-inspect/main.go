package main

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"

	"github.com/pkg/errors"
)

func astPrinter(fileset *token.FileSet, astFile *ast.File) error {
	return ast.Print(fileset, astFile)
}

func jsonPrinter(astFile *ast.File) error {
	return json.NewEncoder(os.Stdout).Encode(astFile)
}

func prettyPrint(fileset *token.FileSet, astFile *ast.File) error {
	return printer.Fprint(os.Stdout, fileset, astFile)
}

func main() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not check if stdin has data"))
	}

	if stat.Size() == 0 {
		log.Fatal("Stdin has no data. Pipe the output of cat into this program.")
	}

	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", os.Stdin, parser.ParseComments|parser.AllErrors)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "parser returned an error"))
	}

	if len(os.Args) != 2 {
		log.Fatalln("Not enough args. First argument should be format kind")
	}

	err = nil
	switch os.Args[1] {
	case "json":
		err = jsonPrinter(astFile)
	case "ast":
		err = astPrinter(fset, astFile)
	case "file":
	case "pretty":
		err = prettyPrint(fset, astFile)
	default:
		log.Fatalln("Printer not found:", os.Args[1])
	}

	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed on print"))
	}
}
