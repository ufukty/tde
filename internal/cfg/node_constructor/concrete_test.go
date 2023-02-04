package node_constructor

import (
	"fmt"
	"tde/internal/cfg/context_resolution/context"

	"testing"
)

func Test_NodeConstructorSerial(t *testing.T) {
	for i := 0; i < 50; i++ {
		t.Run(fmt.Sprintf("Test_NodeConstructorSerial_Subtest#%d", i), func(t *testing.T) {
			BlockStmt(context.NewContext(), 20)
		})
	}
}
