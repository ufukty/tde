package astwutl

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"

	"github.com/kylelemons/godebug/diff"
	"github.com/pkg/errors"
)

func String(node ast.Node) (string, error) {
	fset := token.NewFileSet()
	buf := bytes.NewBuffer([]byte{})
	err := printer.Fprint(buf, fset, node)
	if err != nil {
		return "", errors.Wrapf(err, "failed print")
	}
	return buf.String(), nil
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
