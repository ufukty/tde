package symbols

import (
	"bytes"
	"fmt"
	"go/types"
	"io"
	"strings"
)

func markdownPackageRecHelper(scope *types.Scope, name string, d int, limit int) *bytes.Buffer {
	h := strings.Repeat("#", d+2)
	f := bytes.NewBuffer([]byte{})
	fmt.Fprintf(f, "%s %s\n\n", h, name)
	fmt.Fprintln(f, NewScopeContent(scope).Markdown())
	if limit != d {
		for i := 0; i < scope.NumChildren(); i++ {
			firstline := strings.Split(scope.String(), "\n")[0]
			title := strings.Join(strings.Split(firstline, " ")[0:2], " ")
			io.Copy(f, markdownPackageRecHelper(scope.Child(i), title, d+1, limit))
		}
	}
	return f
}

// use limit to restrain what depth/level of nested-scopes will be visited
func PrintPackageAsMarkdown(pkg *types.Package, limit int) *bytes.Buffer {
	f := bytes.NewBuffer([]byte{})
	fmt.Fprintf(f, "\n# Universe\n\n")
	fmt.Fprintln(f, NewScopeContent(types.Universe).Markdown())
	io.Copy(f, markdownPackageRecHelper(pkg.Scope(), "The Package", 0, limit))
	return f
}
