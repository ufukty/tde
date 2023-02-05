package utilities

import (
	"go/ast"
	"tde/internal/astw/copy"

	"testing"
)

func TestCompare(t *testing.T) {

	if Compare(TEST_TREE, TEST_TREE) != true {
		t.Error("Failed for same inputs")
	}

	TEST_TREE_NEW := copy.File(TEST_TREE)
	if Compare(TEST_TREE, TEST_TREE) != true {
		t.Error("Failed for same inputs")
	}

	TEST_TREE_NEW.Decls = append(TEST_TREE_NEW.Decls, &ast.GenDecl{})
	if Compare(TEST_TREE, TEST_TREE_NEW) != false {
		t.Error("Failed for same inputs")
	}
}
