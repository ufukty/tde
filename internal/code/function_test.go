package code

import (
	"fmt"
	"go/parser"
	"testing"

	"github.com/pkg/errors"
)

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
