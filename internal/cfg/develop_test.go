package cfg

import (
	"bytes"
	"fmt"
	"go/format"
	"go/token"
	ast_utl "tde/internal/astw/utilities"

	"testing"

	"github.com/pkg/errors"
)

func Test_Develop(t *testing.T) {
	_, astPkgs, err := ast_utl.LoadDir("../test_package")
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed test prep"))
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["../test_package/walk.go"]
	funcDecl, _ := ast_utl.FindFuncDecl(astPkg, "WalkWithNils")

	// startCount := traverse.CountNonNilNodes(astFile)

	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprintf("Run: %d", i), func(t *testing.T) {
			Develop(astPkg, astFile, funcDecl)
			err = format.Node(bytes.NewBuffer([]byte{}), token.NewFileSet(), astFile)
			if err != nil {
				t.Error(errors.Wrapf(err, ""))
			}
		})
	}

	// if endCount := traverse.CountNonNilNodes(astFile); endCount == startCount {
	// 	t.Error("\nFailed to add 1 node: startCount=", startCount, " endCount=", endCount)
	// }
}

// func Test_SequentialDevelop(t *testing.T) {
// 	_, astFile, _ := utilities.LoadFile("../astw/walk.go")
// 	funcDecl, _ := utilities.FindFuncDecl(astFile, "WalkWithNils")

// 	var startCount int
// 	for i := 0; i < 2000; i++ {
// 		startCount = traverse.CountNonNilNodes(astFile)
// 		Develop(astFile, funcDecl)
// 		if endCount := traverse.CountNonNilNodes(astFile); endCount == startCount {
// 			t.Error("\nFailed to add 1 node. i=", i, " startCount=", startCount, " endCount=", endCount)
// 		}
// 		if i > 10 {
// 			func() {
// 				defer recover()
// 				err := format.Node(os.Stdout, token.NewFileSet(), astFile)
// 				if err != nil {
// 					t.Error(errors.Wrapf(err, "Failed to format"))
// 				}
// 			}()
// 		}
// 	}

// }
