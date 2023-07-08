package paths_test

import (
	"math/rand"
	"tde/config"
	"tde/internal/microservices/paths"

	"fmt"
	"strings"
	"testing"
)

func Test_Order(t *testing.T) {

	eps := []paths.Endpoint{
		config.CustomsModuleList,
		config.CustomsModuleContext,
		config.CustomsModuleUpload,
		config.CustomsModuleDownload,
		config.CustomsModuleAstFile,
		config.CustomsModuleAstPackage,
		config.CustomsModuleAstFuncDecl,
	}

	for take := 0; take < 10; take++ {
		eps = paths.Sort(eps)

		for i := 0; i < len(eps); i++ {
			for j := 0; i < len(eps); i++ {
				if !(i < j) {
					continue
				}
				a := eps[i].Url()
				b := eps[j].Url()
				if strings.HasPrefix(b, a) {
					t.Errorf("should've be otherway:\n\t%s\n\t%s", a, b)
				}
			}
		}
		for _, ep := range eps {
			fmt.Println(ep.Url())
		}
		fmt.Println()

		var rn = rand.Intn(len(eps) - 1)
		var tmp = eps[rn]
		eps = append(eps[:rn], eps[rn+1:]...)
		eps = append(eps, tmp)
	}

}
