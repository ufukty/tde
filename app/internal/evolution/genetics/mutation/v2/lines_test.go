package mutation

import (
	"testing"
)

func Test_RemoveLine(t *testing.T) {
	// _, _, originalFuncDecl, err := loadTestPackage()
	// if err != nil {
	// 	t.Error(errors.Wrapf(err, "prep"))
	// }

	// modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)
	// if err := RemoveALine(modifiedFuncDecl.Body); err != nil {
	// 	t.Error(fmt.Errorf("act: %w", err))
	// }

	// codeForOriginal, err := astwutl.String(originalFuncDecl)
	// if err != nil {
	// 	t.Error("validation prep")
	// }
	// codeForModified, err := astwutl.String(modifiedFuncDecl)
	// if err != nil {
	// 	t.Error("validation prep")
	// }

	// fmt.Println("Differences in code:\n", diff.Diff(codeForOriginal, codeForModified))

	// if astwutl.CompareRecursively(originalFuncDecl, modifiedFuncDecl) {
	// 	t.Error("validation")
	// }
}

func Test_RemoveLineMany(t *testing.T) {
	// _, _, originalFuncDecl, err := loadTestPackage()
	// if err != nil {
	// 	t.Error(errors.Wrapf(err, "prep"))
	// }

	// for i := 0; i < 1000; i++ {
	// 	modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)

	// 	if err := RemoveALine(modifiedFuncDecl.Body); err != nil {
	// 		t.Error(fmt.Errorf("act: %w", err))
	// 	}

	// 	if astwutl.CompareRecursively(originalFuncDecl, modifiedFuncDecl) {
	// 		t.Error("validation", i)
	// 	}
	// }
}

func Test_SiblingSwap(t *testing.T) {
	// _, _, originalFuncDecl, err := loadTestPackage()
	// if err != nil {
	// 	t.Error(errors.Wrapf(err, "prep"))
	// }

	// modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)
	// if err := SwapTwoLines(modifiedFuncDecl.Body); err != nil {
	// 	t.Error(fmt.Errorf("act: %w", err))
	// }

	// codeForOriginal, err := astwutl.String(originalFuncDecl)
	// if err != nil {
	// 	t.Error("validation prep")
	// }
	// codeForModified, err := astwutl.String(modifiedFuncDecl)
	// if err != nil {
	// 	t.Error("validation prep")
	// }

	// fmt.Println("Differences in code:\n", diff.Diff(codeForOriginal, codeForModified))

	// if astwutl.CompareRecursively(originalFuncDecl, modifiedFuncDecl) {
	// 	t.Error("validation")
	// }
}

func Test_SiblingSwapMany(t *testing.T) {
	// _, _, originalFuncDecl, err := loadTestPackage()
	// if err != nil {
	// 	t.Error(errors.Wrapf(err, "prep"))
	// }

	// for i := 0; i < 1000; i++ {
	// 	modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)

	// 	if err := SwapTwoLines(modifiedFuncDecl.Body); err != nil {
	// 		t.Error(fmt.Errorf("act: %w", err))
	// 	}

	// 	if astwutl.CompareRecursively(originalFuncDecl, modifiedFuncDecl) {
	// 		t.Error("validation", i)
	// 	}
	// }
}
