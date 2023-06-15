package customs_proxy

import (
	"go/ast"
	"go/parser"
	"go/token"
	"tde/internal/astw/utilities"
	"testing"

	"github.com/pkg/errors"
)

func Test_JSON(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "../../../../internal/test-package/walk.go", nil, parser.AllErrors)
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	var fd *ast.FuncDecl
	ast.Inspect(f, func(n ast.Node) bool {
		if n, ok := n.(*ast.FuncDecl); ok {
			fd = n
		}
		return fd != nil
	})

	buf, err := encodeAsJson(fd)
	if err != nil {
		t.Error(errors.Wrapf(err, "encoding"))
	}

	fdg, err := decodeFromJson(buf)
	if err != nil {
		t.Error(errors.Wrapf(err, "decoding"))
	}

	if same := utilities.CompareNonNodeFields(fd, fdg); !same {
		t.Error("validate non-node fields only")
	}

	if same := utilities.CompareRecursively(fd, fdg); !same {
		t.Error("validate all")
	}
}
