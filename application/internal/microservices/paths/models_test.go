package paths_test

import (
	"tde/config"
	"tde/internal/microservices/paths"

	"fmt"
	"strings"
	"testing"
)

func Test_Order(t *testing.T) {

	eps := paths.Sort([]paths.Endpoint{
		config.CustomsModuleList,
		config.CustomsModuleContext,
		config.CustomsModuleUpload,
		config.CustomsModuleDownload,
		config.CustomsModuleAstFile,
		config.CustomsModuleAstPackage,
		config.CustomsModuleAstFuncDecl,
	})

	for i := 1; i < len(eps); i++ {
		a := eps[i-1].Url()
		b := eps[i].Url()
		if strings.HasPrefix(b, a) {
			t.Errorf("should be before:\n\t%s\n\t%s", a, b)
		}
	}

	for _, ep := range eps {
		fmt.Println(ep)
	}
}
