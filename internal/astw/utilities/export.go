package utilities

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
	"tde/internal/utilities"

	"github.com/kylelemons/godebug/diff"
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

func Diff(current, new ast.Node) string {
	bufferCurrent := bytes.NewBuffer([]byte{})
	ast.Fprint(bufferCurrent, token.NewFileSet(), current, nil)

	bufferNew := bytes.NewBuffer([]byte{})
	ast.Fprint(bufferNew, token.NewFileSet(), current, nil)

	diffStr := diff.Diff(bufferCurrent.String(), bufferNew.String())
	changedLines := ""
	for _, str := range strings.Split(diffStr, "\n") {
		if strings.Index(str, " ") != 0 {
			changedLines = changedLines + str
		}
	}
	return changedLines
}
