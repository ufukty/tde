package astcfg

import (
	astw "tde/internal/ast_wrapper"
	"testing"
)

func Test_Develop(t *testing.T) {
	_, astFile, _ := astw.LoadFile("../../ast_wrapper/walk.go")
	funcDecl, _ := astw.FindFuncDecl(astFile, "WalkWithNils")

	var startCount int
	for i := 0; i < 200; i++ {
		startCount = astw.CountNonNilNodes(astFile)
		Develop(astFile, funcDecl)
		if endCount := astw.CountNonNilNodes(astFile); endCount == startCount {
			t.Error("\nFailed to add 1 node. i=", i, " startCount=", startCount, " endCount=", endCount)
		}
		// err := format.Node(os.Stdout, token.NewFileSet(), astFile)
		// if err != nil {
		// 	t.Error(errors.Wrapf(err, "Failed to format"))
		// }
	}

}
