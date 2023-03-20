package utilities

import (
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

func Diff(current, new ast.Node) (string, error) {
	printCurrent, err := String(current)
	if err != nil {
		return "", errors.Wrap(err, "printing the current version")
	}
	printNew, err := String(new)
	if err != nil {
		return "", errors.Wrap(err, "printing the new version")
	}

	diffStr := diff.Diff(printCurrent, printNew)
	changedLines := ""
	for _, str := range strings.Split(diffStr, "\n") {
		if strings.Index(str, " ") != 0 {
			changedLines = changedLines + str + "\n"
		}
	}
	return changedLines, nil
}
