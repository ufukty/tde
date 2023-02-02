package node_constructor

import (
	"tde/internal/cfg/context"

	"testing"
)

func Test_NodeConstructorSerial(t *testing.T) {
	for i := 0; i < 1000; i++ {
		BlockStmt(context.NewContext(), 10)
	}
}

func Test_PrintAST(t *testing.T) {
	BlockStmt(context.NewContext(), 50)
}
