package evolution

import (
	"fmt"
	"go/ast"
	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone/clean"
	"testing"
)

// use this test with memory profiling. go test -memprofile=mem.out .
func Test_Scale(t *testing.T) {
	const POPULATION = 10000

	_, file, err := astwutl.LoadFile("testdata/walk.go")
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	candidates := []*ast.File{}
	for i := 0; i < POPULATION; i++ {
		candidates = append(candidates, clean.File(file))
	}
	if len(candidates) != POPULATION {
		t.Fatalf("assert: got: %d", len(candidates))
	}
}
