package discovery

import (
	"fmt"
	"tde/internal/folders/list"
)

func CombinedDetailsForTest(path string, testname string) (*CombinedDetails, error) {
	pkgs, err := list.ListAllPackages(path)
	if err != nil {
		return nil, fmt.Errorf("locating the Go mo0dule and package of current dir: %w", err)
	}
	pkg := pkgs.First()
	test, err := TestFunctionInDir(path, testname)
	if err != nil {
		return nil, fmt.Errorf("locating test functions in the package: %w", err)
	}
	targetFile, targetFunctionName := ExpectedTargetFileAndFuncNameFor(test.Path, test.Name)
	target, err := TargetFunctionInFile(targetFile, targetFunctionName)
	if err != nil {
		return nil, fmt.Errorf("finding the target file %q and function %q: %w", targetFile, targetFunctionName, err)
	}
	combined := &CombinedDetails{
		Package: pkg,
		Target:  target,
		Test:    test,
	}
	return combined, nil
}
