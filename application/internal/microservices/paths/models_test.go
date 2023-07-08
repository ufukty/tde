package paths

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Order(t *testing.T) {

	eps := Sort([]Endpoint{
		CustomsModuleList,
		CustomsModuleContext,
		CustomsModuleUpload,
		CustomsModuleDownload,
		CustomsModuleAstFile,
		CustomsModuleAstPackage,
		CustomsModuleAstFuncDecl,
	})

	for i := 1; i < len(eps); i++ {
		a := eps[i-1].String()
		b := eps[i].String()
		if strings.HasPrefix(b, a) {
			t.Errorf("should be before:\n\t%s\n\t%s", a, b)
		}
	}

	for _, ep := range eps {
		fmt.Println(ep)
	}
}
