package astwutl

import (
	"go/ast"
	"testing"

	"tde/internal/astw/clone"
)

func TestComparRecursively(t *testing.T) {
	if CompareRecursively(TEST_TREE, TEST_TREE) != true {
		t.Error("Failed for same inputs")
	}
	if CompareRecursivelyWithAddresses(TEST_TREE, TEST_TREE) != true {
		t.Error("Failed for same inputs")
	}

	TEST_TREE_NEW := clone.File(TEST_TREE)
	if CompareRecursively(TEST_TREE, TEST_TREE_NEW) != true {
		t.Error("Failed for same inputs")
	}
	if CompareRecursivelyWithAddresses(TEST_TREE, TEST_TREE_NEW) != false {
		t.Error("Failed for same inputs")
	}

	TEST_TREE_NEW.Decls = append(TEST_TREE_NEW.Decls, &ast.GenDecl{})
	if CompareRecursivelyWithAddresses(TEST_TREE, TEST_TREE_NEW) != false {
		t.Error("Failed for changed inputs")
	}
}
