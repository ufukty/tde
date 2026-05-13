package evolution

import (
	"fmt"
	"go/ast"
	"testing"

	"tde/internal/ast/clone/clean"
	"tde/internal/ast/parse"
)

// use this test with memory profiling. go test -memprofile=mem.out .
func Test_Scale(t *testing.T) {
	const POPULATION = 10000

	_, file, err := parse.File("testdata/evolution/walk.go")
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	subjects := []*ast.File{}
	for i := 0; i < POPULATION; i++ {
		subjects = append(subjects, clean.File(file))
	}
	if len(subjects) != POPULATION {
		t.Fatalf("assert: got: %d", len(subjects))
	}
}
