package nodes

import (
	"fmt"

	"testing"
)

func Test_NodeConstructorSerial(t *testing.T) {
	for i := 0; i < 50; i++ {
		t.Run(fmt.Sprintf("Test_NodeConstructorSerial_Subtest#%d", i), func(t *testing.T) {
			c := &Creator{}
			c.BlockStmt(20)
		})
	}
}
