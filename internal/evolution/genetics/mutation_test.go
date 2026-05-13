package genetics

import (
	"fmt"
	"go/parser"
	"go/token"
	"testing"

	"tde/internal/ast/clone/clean"
	"tde/internal/ast/compare"
	"tde/internal/ast/find"
	"tde/internal/evolution/genetics/nodes"
)

func TestGrow(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "testdata/words", nil, parser.AllErrors)
	if err != nil {
		t.Fatal(fmt.Errorf("prep 1: %w", err))
	}
	fd, err := find.Function(f, "WordReverse")
	if err != nil {
		t.Fatal(fmt.Errorf("prep 2: %w", err))
	}
	mfd := clean.FuncDecl(fd)

	nc := nodes.NewCreator()
	Grow(nc, mfd)

	if compare.Recursively(fd, mfd) {
		t.Fatal("assert: change is expected")
	}
}
