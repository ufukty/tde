package code

import (
	"fmt"
	"go/format"
	"go/parser"
	"tde/internal/code_fragment"
	"testing"

	"github.com/pkg/errors"
)

func TestInsertGeneretedTextIntoFunctionBody(t *testing.T) {

	input := `package main
	
	import "fmt"
	
	func Main() {
		fmt.Println(foo(), " world ")
	}

	func foo() string {
		sfsdfdsf
	}

	func bar() {
		panic("bar")
	}
	`

	c := Code{}
	c.LoadFromString(input)
	f, err := c.FindFunction("foo")
	if err != nil {
		t.Error(errors.Wrap(err, "Could not get function object"))
	}

	// fmt.Printf("%+v", f.Body)

	// fmt.Println(input[:f.Body.Lbrace] + "\nreturn \"Goodbye\"\n" + input[f.Body.Rbrace-1:])

	st := input[:f.Body.Lbrace] + " " + string(code_fragment.ProduceRandomFragment()) + " " + input[f.Body.Rbrace-1:]

	// "return \"Goodbye\""

	// str := utilities.NewStringWriter()
	// format.Node(str, c.fset, c.astFile)
	str, err := format.Source([]byte(st))
	if err != nil {
		t.Error(errors.Wrap(err, "Could not format the produced code"))
	}
	fmt.Println(string(str))
}

func TestParseBlockStatement(t *testing.T) {
	in := `var a = 5
	a = a * b
	return a
	`
	// fset := token.NewFileSet()
	astExpr, err := parser.ParseExpr(in)
	if err != nil {
		t.Error(errors.Wrap(err, "Failed to load code from string"))
	}

	fmt.Printf("%+v\n", astExpr)

	// c := Code{}
	// err := c.LoadFromString(in)
	// if err != nil {
	// 	t.Error(errors.Wrap(err, "Failed to load code from string"))
	// }

}
