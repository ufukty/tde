package node_constructor

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"tde/internal/cfg/astcfg/context"
	"testing"
)

func Test_NodeConstructorSerial(t *testing.T) {
	for i := 0; i < 1000; i++ {
		BlockStmt(context.NewContext(), 10)
	}
}

func Test_PrintAST(t *testing.T) {
	n := BlockStmt(context.NewContext(), 100)
	fset := token.NewFileSet()
	ast.Print(fset, n)
	printer.Fprint(os.Stdout, fset, n)
}
