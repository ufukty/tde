package ast_wrapper

import (
	"go/ast"
	"go/printer"
	"go/token"
	"tde/internal/utilities"

	"github.com/pkg/errors"
)

func String(node ast.Node) (string, error) {
	fset := token.NewFileSet()
	sw := utilities.NewStringWriter()
	err := printer.Fprint(sw, fset, node)
	if err != nil {
		return "", errors.Wrapf(err, "failed print")
	}
	return sw.String(), nil
}
