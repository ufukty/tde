package cfg

import (
	"tde/internal/astw/traverse"
	ast_utl "tde/internal/astw/utilities"

	"testing"
)

func Test_Develop(t *testing.T) {
	_, astPkgs, _ := ast_utl.LoadDir("../astw")
	astPkg := astPkgs["astw"]
	astFile := astPkg.Files["walk"]
	funcDecl, _ := ast_utl.FindFuncDecl(astPkg, "WalkWithNils")

	startCount := traverse.CountNonNilNodes(astFile)
	Develop(astPkg, astFile, funcDecl)
	if endCount := traverse.CountNonNilNodes(astFile); endCount == startCount {
		t.Error("\nFailed to add 1 node: startCount=", startCount, " endCount=", endCount)
	}
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
