package paths

import (
	"fmt"
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
	for _, ep := range eps {
		fmt.Println(ep)
	}
}
