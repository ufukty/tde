package astcfg

import (
	"testing"
)

func Test_DictInitialization(t *testing.T) {
	if len(Dict_NodeTypeClassToNodeType[Statement]) == 0 {
		t.Error("No statement found.")
	}
	if len(Dict_NodeTypeClassToNodeType[Expression]) == 0 {
		t.Error("No expression found.")
	}
	if len(Dict_NodeTypeClassToNodeType[Declaration]) == 0 {
		t.Error("No declaration found.")
	}
	for nodeType := range nc.Dictionary {
		if _, ok := Dict_NodeTypeToNodeTypeClass[nodeType]; !ok {
			t.Error("Node type class is not found for node type:", nodeType)
		}
	}
}
