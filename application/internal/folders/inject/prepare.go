package inject

import (
	"strings"
	"tde/internal/folders/copymod"
	"tde/internal/folders/list"

	"fmt"
	"os"
	"path/filepath"
)

func duplicateInTmp(src string) (string, error) {
	dst, err := os.MkdirTemp(os.TempDir(), "deepthinker-sample-*")
	if err != nil {
		return "", fmt.Errorf("creating empty dir in tmp for future writes: %w", err)
	}
	err = copymod.CopyModule(dst, src, true, []string{}, []string{}, []string{}, false)
	if err != nil {
		return "", fmt.Errorf("copying contents of original module to sample module dir: %w", err)
	}
	return dst, nil
}

func inject(sample string, pkg *list.Package, testname string) error {
	var pkgRel = strings.TrimPrefix(pkg.Dir, pkg.Module.Dir)
	var pkgAbs = filepath.Join(sample, pkgRel)
	var testInfo = &TestInfo{
		TargetPackageImportPath: pkg.ImportPath,
		TestFunctionName:        testname,
	}
	if err := Inject(pkgAbs, testInfo); err != nil {
		return fmt.Errorf("callling inject function: %w", err)
	}
	return nil
}

// creates a dir in temp folder for the module. returns path for the copy
//   - mod is the path to module root
//   - pkg is details about target package that contains target and test functions
func WithCreatingSample(mod string, pkg *list.Package, testname string) (string, error) {
	sample, err := duplicateInTmp(mod)
	if err != nil {
		return "", fmt.Errorf("creating sample: %w", err)
	}
	if err := inject(sample, pkg, testname); err != nil {
		return "", fmt.Errorf("injecting test-runner package: %w", err)
	}
	return sample, nil
}
