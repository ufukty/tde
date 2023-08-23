package nodes

import (
	"fmt"
	"tde/internal/genetics/mutation/cfg/ctxres/context"

	"testing"
)

func Test_NodeConstructorSerial(t *testing.T) {
	for i := 0; i < 50; i++ {
		t.Run(fmt.Sprintf("Test_NodeConstructorSerial_Subtest#%d", i), func(t *testing.T) {
			BlockStmt(context.NewContext(), 20)
		})
	}
}
