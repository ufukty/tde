package discovery

import (
	"fmt"
	"go/ast"
	"strings"
	astw_utl "tde/internal/astw/utilities"
	"tde/internal/folders/types"
	"tde/internal/utilities"

	"io/fs"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

func DiscoverImplFileForImplFuncDetails(moduleRoot, implFilePathInMod, implFuncName string) (*types.ImplFuncDetails, error) {
	fset, astFile, err := astw_utl.LoadFile(filepath.Join(moduleRoot, implFilePathInMod))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("could not create AST for the file \"%s\"", implFilePathInMod))
	}

	var implFuncDetails *types.ImplFuncDetails
	ast.Inspect(astFile, func(node ast.Node) bool {
		if node, ok := node.(*ast.FuncDecl); ok {
			if node.Name.Name == implFuncName {
				implFuncDetails = &types.ImplFuncDetails{
					Name: implFuncName,
					Path: implFilePathInMod,
					Line: astw_utl.LineNumberOfPosition(fset, node.Pos()),
				}
			}
		}
		return node == astFile
	})
	if implFuncDetails == nil {
		return nil, errors.New(fmt.Sprintf("Could not find function %s in the file %s", implFuncName, implFilePathInMod))
	}

	return implFuncDetails, nil
}

func ListOfCallsInTestFunc(funcDecl *ast.FuncDecl) []*ast.CallExpr {
	list := []*ast.CallExpr{}
	ast.Inspect(funcDecl, func(node ast.Node) bool {
		if node, ok := node.(*ast.CallExpr); ok {
			if !slices.Contains(list, node) {
				list = append(list, node)
			}
		}
		return true
	})
	return list
}

func GetExpectedImplFileAndImplFuncName(testFile, testFuncName string) (implFile, implFuncName string) {
	implFile = strings.TrimSuffix(testFile, "_tde.go") + ".go"
	implFuncName = strings.TrimPrefix(testFuncName, "TDE_")
	return
}

// all paths returned will be relative to <moduleRoot>
func DiscoverTestFileForTestFuncDetails(moduleRoot, testFilePathAbs string) ([]types.TestFuncDetails, error) {
	testFunctions := []types.TestFuncDetails{}

	fset, astFile, err := astw_utl.LoadFile(testFilePathAbs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load file and parse into AST tree")
	}

	relPath, err := filepath.Rel(moduleRoot, testFilePathAbs)
	if err != nil {
		return nil, errors.Wrap(err, "path is not in module root")
	}

	ast.Inspect(astFile, func(node ast.Node) bool {
		if funcDecl, ok := node.(*ast.FuncDecl); ok {
			functionName := funcDecl.Name.Name
			if strings.Index(functionName, "TDE_") == 0 {
				testFunctions = append(testFunctions, types.TestFuncDetails{
					Name:  functionName,
					Path:  relPath,
					Line:  astw_utl.LineNumberOfPosition(fset, node.Pos()),
					Calls: ListOfCallsInTestFunc(funcDecl),
				})
			}
		}
		return node == astFile
	})

	return testFunctions, nil
}

// will return any test function in and under <inModPath>
// all paths returned will be relative to <moduleRoot>
func DiscoverSubdirsForTestFuncDetails(moduleRoot, dirPathAbs string) ([]types.TestFuncDetails, error) {
	excludeDirs := []string{".git", "build", "artifacts"}
	excludePaths := utilities.Map(excludeDirs, func(i int, v string) string { return filepath.Join(dirPathAbs, v) })

	testFunctions := []types.TestFuncDetails{}
	err := filepath.Walk(dirPathAbs, func(filePathAbs string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "failed to walk directory: "+filePathAbs)
		}

		if fileInfo.IsDir() {
			if slices.Contains(excludePaths, filePathAbs) {
				return filepath.SkipDir
			} else {
				return nil
			}
		}

		if strings.HasSuffix(filePathAbs, "_tde.go") {
			testFunctionsInFile, err := DiscoverTestFileForTestFuncDetails(moduleRoot, filePathAbs)
			if err != nil {
				return nil
			}
			testFunctions = append(testFunctions, testFunctionsInFile...)
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "Could not walk directories")
	}

	return testFunctions, nil
}

func WhereAmI() (modulePath, pkgPathInMod, importPath string, err error) {
	modulePath, err = GetModulePath()
	if err != nil {
		return "", "", "", errors.Wrap(err, "Failed to get current module path")
	}

	pkgPathInMod, err = GetPackagePathInModule(modulePath)
	if err != nil {
		return "", "", "", errors.Wrap(err, "Failed to get current package path")
	}

	importPath, err = GetImportPathOfPackage()
	if err != nil {
		return "", "", "", errors.Wrap(err, "Failed to get import path of package in current directory. Maybe you are not in a directory that contains a Go package.")
	}

	return
}
