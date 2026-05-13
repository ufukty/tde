package export

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"

	"github.com/kylelemons/godebug/diff"
)

func String(n ast.Node) (string, error) {
	var b bytes.Buffer
	if err := printer.Fprint(&b, token.NewFileSet(), n); err != nil {
		return "", fmt.Errorf("printing: %w", err)
	}
	return b.String(), nil
}

func Diff(current, next ast.Node) (string, error) {
	c, err := String(current)
	if err != nil {
		return "", fmt.Errorf("printing the current version: %w", err)
	}
	n, err := String(next)
	if err != nil {
		return "", fmt.Errorf("printing the new version: %w", err)
	}
	var b strings.Builder
	for line := range strings.Lines(diff.Diff(c, n)) {
		if !strings.HasPrefix(line, " ") {
			b.WriteString(line + "\n")
		}
	}
	return b.String(), nil
}
