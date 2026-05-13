package astwutl

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"

	"github.com/kylelemons/godebug/diff"
)

func String(node ast.Node) (string, error) {
	fset := token.NewFileSet()
	buf := bytes.NewBuffer([]byte{})
	err := printer.Fprint(buf, fset, node)
	if err != nil {
		return "", fmt.Errorf("failed print: %w", err)
	}
	return buf.String(), nil
}

func Diff(current, new ast.Node) (string, error) {
	printCurrent, err := String(current)
	if err != nil {
		return "", fmt.Errorf("printing the current version: %w", err)
	}
	printNew, err := String(new)
	if err != nil {
		return "", fmt.Errorf("printing the new version: %w", err)
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
