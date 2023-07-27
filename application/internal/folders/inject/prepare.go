package inject

import (
	"path/filepath"
	copymod "tde/internal/folders/copymod"
	"tde/internal/folders/list"

	"fmt"
	"os"
)

func duplicateInTmp(path string) (string, error) {
	dst, err := os.MkdirTemp(os.TempDir(), "tde-*")
	if err != nil {
		return "", fmt.Errorf("MkdirTemp: %w", err)
	}
	err = copymod.Copy(dst, path, true, []string{}, []string{}, []string{}, false)
	if err != nil {
		return "", fmt.Errorf("copymod.Module: %w", err)
	}
	return dst, nil
}

func mountTesterPackage(pkg string, pkgImportPath string, testName string) error {
	testInfo := &TestInfo{
		TargetPackageImportPath: pkgImportPath,
		TestFunctionName:        testName,
	}
	err := Inject(pkg, testInfo)
	if err != nil {
		return fmt.Errorf("mountTesterPackage: %w", err)
	}
	return nil
}

// creates a dir in temp folder for the module. returns path for the copy
func WithCreatingSample(path string, pkgPath string, pkgInfo *list.Package, testName string) (string, error) {
	dupl, err := duplicateInTmp(path)
	if err != nil {
		return "", fmt.Errorf("duplicateInTmp: %w", err)
	}
	err = mountTesterPackage(filepath.Join(path, pkgPath), pkgInfo.ImportPath, testName)
	if err != nil {
		return "", fmt.Errorf("mountTesterPackage: %w", err)
	}
	return dupl, nil
}
