package preps

import (
	copymod "tde/internal/folders/copymod"
	"tde/internal/folders/list"
	testinject "tde/internal/folders/testinject"

	"fmt"
	"os"
)

func duplicateInTmp(srcModule string) (string, error) {
	dst, err := os.MkdirTemp(os.TempDir(), "tde-*")
	if err != nil {
		return "", fmt.Errorf("MkdirTemp: %w", err)
	}
	err = copymod.Copy(srcModule, dst, true, copymod.DefaultSkipDirs)
	if err != nil {
		return "", fmt.Errorf("copymod.Module: %w", err)
	}
	return dst, nil
}

func mountTesterPackage(pkg string, packageImportPath string, testName string) error {
	testInfo := &testinject.TestInfo{
		TargetPackageImportPath: packageImportPath,
		TestFunctionName:        testName,
	}
	err := testinject.Inject(string(pkg), testInfo)
	if err != nil {
		return fmt.Errorf("mountTesterPackage: %w", err)
	}
	return nil
}

// returns the path of duplication of module. Duplicated module is free of directories unrelated,
func Prepare(srcModule string, targetPackage string, pkg *list.Package, testName string) (string, error) {
	dupl, err := duplicateInTmp(srcModule)
	if err != nil {
		return "", fmt.Errorf("duplicateInTmp: %w", err)
	}
	err = mountTesterPackage(targetPackage, pkg.ImportPath, testName)
	if err != nil {
		return "", fmt.Errorf("mountTesterPackage: %w", err)
	}
	return dupl, nil
}
