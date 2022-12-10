package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"

	"github.com/pkg/errors"
)

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
	ast.Print(fset, astFile)
}
