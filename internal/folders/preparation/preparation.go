package preparation

import (
	"os"
	"tde/internal/folders/copy_module"
	"tde/internal/folders/tester_package"
	"tde/internal/folders/types"

	"github.com/pkg/errors"
)

func duplicateInTmp(src types.AbsolutePath) (types.AbsolutePath, error) {
	dst, err := os.MkdirTemp(os.TempDir(), "tde-*")
	if err != nil {
		return "", errors.Wrap(err, "MkdirTemp")
	}

	err = copy_module.Module(string(src), dst, true, copy_module.DefaultSkipDirs)
	if err != nil {
		return "", errors.Wrap(err, "copy_module.Module")
	}

	return types.AbsolutePath(dst), nil
}

func mountTesterPackage(pkg types.AbsolutePath, packageImportPath, testName string) error {
	testInfo := &tester_package.TestInfo{
		TargetPackageImportPath: packageImportPath,
		TestFunctionName:        testName,
	}
	err := tester_package.Create(string(pkg), testInfo)
	if err != nil {
		return errors.Wrap(err, "mountTesterPackage")
	}
	return nil
}

// returns the path of duplication of module. Duplicated module is free of directories unrelated,
func Prepare(modulePath types.AbsolutePath, pkg types.InModulePath, packageImportPath string, testName string) (types.AbsolutePath, error) {
	dupl, err := duplicateInTmp(modulePath)
	if err != nil {
		return "", errors.Wrap(err, "duplicateInTmp")
	}

	err = mountTesterPackage(dupl.Join(pkg), packageImportPath, testName)
	if err != nil {
		return "", errors.Wrap(err, "mountTesterPackage")
	}

	return dupl, nil
}
