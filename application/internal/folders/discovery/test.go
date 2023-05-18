package discovery

import (
	"path/filepath"
	"tde/internal/folders/types"

	"github.com/pkg/errors"
)

func ResolveTestDetailsInCurrentDir(testName string) (*types.TestDetails, error) {
	modPath, pkgPathInMod, importPath, err := WhereAmI()
	if err != nil {
		return nil, errors.Wrap(err, "WhereAmI")
	}

	funcDetailsInSubdirs, err := DiscoverSubdirsForTestFuncDetails(modPath, filepath.Join(modPath, pkgPathInMod))
	if err != nil {
		return nil, errors.Wrap(err, "DiscoverSubdirsForTestFuncDetails")
	}

	var foundFuncDetails *types.TestFuncDetails
	for _, funcDetails := range funcDetailsInSubdirs {
		if funcDetails.Name == testName {
			foundFuncDetails = &funcDetails
			break
		}
	}
	if foundFuncDetails == nil {
		return nil, errors.New("Given test name is not found in subdirs")
	}

	implFile, implFuncName := GetExpectedImplFileAndImplFuncName(foundFuncDetails.Path, foundFuncDetails.Name)

	implFuncDetails, err := DiscoverImplFileForImplFuncDetails(modPath, implFile, implFuncName)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not find implementation file \"%s\" and function \"%s\"", implFile, implFuncName)
	}

	testDetails := types.TestDetails{
		PackagePath:   types.InModulePath(pkgPathInMod),
		PackageImport: importPath,
		ImplFuncFile:  types.InModulePath(implFile),
		ImplFuncName:  implFuncName,
		ImplFuncLine:  implFuncDetails.Line,
		TestFuncFile:  types.InModulePath(foundFuncDetails.Path),
		TestFuncName:  testName,
		TestFuncLine:  foundFuncDetails.Line,
	}

	return &testDetails, nil
}
