package astcfg

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"testing"
)

func Test_NodeConstructorSerial(t *testing.T) {
	for i := 0; i < 10; i++ {
		nodeConstructor.BlockStmt()
	}
}

func Test_PrintAST(t *testing.T) {
	n := nodeConstructor.BlockStmt()
	fset := token.NewFileSet()
	ast.Print(fset, n)
	printer.Fprint(os.Stdout, fset, n)
}
