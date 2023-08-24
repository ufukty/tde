package customs_proxy

import (
	"tde/internal/astw/astwutl"
	
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/pkg/errors"
)

func Test_JSON(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "testdata/walk.go", nil, parser.AllErrors)
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

	if same := astwutl.CompareNonNodeFields(fd, fdg); !same {
		t.Error("validate non-node fields only")
	}

	if same := astwutl.CompareRecursively(fd, fdg); !same {
		t.Error("validate all")
	}
}
