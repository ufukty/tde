package genetics

import (
	"fmt"
	"go/parser"
	"go/token"
	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone/clean"
	"tde/internal/evolution/genetics/nodes"
	"testing"
)

func TestGrow(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "testdata/words", nil, parser.AllErrors)
	if err != nil {
		t.Fatal(fmt.Errorf("prep 1: %w", err))
	}
	fd, err := astwutl.FindFuncDecl(f, "WordReverse")
	if err != nil {
		t.Fatal(fmt.Errorf("prep 2: %w", err))
	}
	mfd := clean.FuncDecl(fd)

	nc := nodes.NewCreator()
	Grow(nc, mfd)

	if astwutl.CompareRecursively(fd, mfd) {
		t.Fatal("assert: change is expected")
	}
}
