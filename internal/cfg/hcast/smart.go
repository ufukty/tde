package hcast

import "go/ast"

var variablesInUse = []string{}

func GetVaribleInUseInSameTypeWith() {}

func AddVariable() error { return nil }

func AddVariableAtSameType(n *ast.Ident) error { return nil }

func MathOperation() error { return nil }

func ExternalLibraryCall() error { return nil }

var allowedPackagesToImport = []string{
	"math", "fmt", "tde/pkg/httpmitm",
}
