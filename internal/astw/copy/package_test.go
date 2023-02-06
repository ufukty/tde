package copy

import (
	"go/ast"
	"tde/internal/astw/utilities"

	"testing"

	"github.com/pkg/errors"
)

func BenchmarkCopyPackage(b *testing.B) {
	_, pkgs, err := utilities.LoadDir("../../test_package")
	if err != nil {
		b.Error(errors.Wrap(err, "Failed on prep"))
	}
	var (
		pkg *ast.Package
		ok  bool
	)
	if pkg, ok = pkgs["test_package"]; !ok {
		b.Error("Failed on prep. Can't find 'test_package' in imported packages.")
	}
	for i := 0; i < b.N; i++ {
		Package(pkg)
	}
}
