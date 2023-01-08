package node_construtor

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"testing"
)

func Test_NodeConstructorSerial(t *testing.T) {
	for i := 0; i < 100000; i++ {
		BlockStmt(100)
	}
}

func Test_PrintAST(t *testing.T) {
	n := BlockStmt(100)
	fset := token.NewFileSet()
	ast.Print(fset, n)
	printer.Fprint(os.Stdout, fset, n)
}
