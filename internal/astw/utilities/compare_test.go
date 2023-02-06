package utilities

import (
	"go/ast"
	"tde/internal/astw/copy"

	"testing"
)

func TestComparRecursively(t *testing.T) {
	if CompareRecursively(TEST_TREE, TEST_TREE) != true {
		t.Error("Failed for same inputs")
	}

	TEST_TREE_NEW := copy.File(TEST_TREE)
	if CompareRecursively(TEST_TREE, TEST_TREE_NEW) != true {
		t.Error("Failed for same inputs")
	}

	TEST_TREE_NEW.Decls = append(TEST_TREE_NEW.Decls, &ast.GenDecl{})
	if CompareRecursively(TEST_TREE, TEST_TREE_NEW) != false {
		t.Error("Failed for changed inputs")
	}
}
