package discovery

import (
	"tde/internal/astw/astwutl"
	"tde/internal/folders/list"
	"tde/internal/utilities"

	"fmt"
	"go/ast"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

func DiscoverImplFileForImplFuncDetails(moduleRoot, implFilePathInMod, implFuncName string) (*ImplFuncDetails, error) {
	fset, astFile, err := astwutl.LoadFile(filepath.Join(moduleRoot, implFilePathInMod))
	if err != nil {
		return nil, fmt.Errorf("could not create AST for the file %q: %w", implFilePathInMod, err)
	}

	var implFuncDetails *ImplFuncDetails
	ast.Inspect(astFile, func(node ast.Node) bool {
		if node, ok := node.(*ast.FuncDecl); ok {
			if node.Name.Name == implFuncName {
				implFuncDetails = &ImplFuncDetails{
					Name: implFuncName,
					Path: implFilePathInMod,
					Line: astwutl.LineNumberOfPosition(fset, node.Pos()),
				}
			}
		}
		return node == astFile
	})
	if implFuncDetails == nil {
		return nil, fmt.Errorf("Could not find function %s in the file %s", implFuncName, implFilePathInMod)
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
func DiscoverTestFileForTestFuncDetails(moduleRoot, testFilePathAbs string) ([]TestFuncDetails, error) {
	testFunctions := []TestFuncDetails{}

	fset, astFile, err := astwutl.LoadFile(testFilePathAbs)
	if err != nil {
		return nil, fmt.Errorf("failed to load file and parse into AST tree: %w", err)
	}

	relPath, err := filepath.Rel(moduleRoot, testFilePathAbs)
	if err != nil {
		return nil, fmt.Errorf("path is not in module root: %w", err)
	}

	ast.Inspect(astFile, func(node ast.Node) bool {
		if funcDecl, ok := node.(*ast.FuncDecl); ok {
			functionName := funcDecl.Name.Name
			if strings.Index(functionName, "TDE_") == 0 {
				testFunctions = append(testFunctions, TestFuncDetails{
					Name:  functionName,
					Path:  relPath,
					Line:  astwutl.LineNumberOfPosition(fset, node.Pos()),
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
func DiscoverSubdirsForTestFuncDetails(moduleRoot, dirPathAbs string) ([]TestFuncDetails, error) {
	excludeDirs := []string{".git", "build", "artifacts"}
	excludePaths := utilities.Map(excludeDirs, func(i int, v string) string { return filepath.Join(dirPathAbs, v) })

	testFunctions := []TestFuncDetails{}
	err := filepath.Walk(dirPathAbs, func(filePathAbs string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk directory %q: %w", filePathAbs, err)
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
		return nil, fmt.Errorf("Could not walk directories: %w", err)
	}

	return testFunctions, nil
}

func WhereAmI() (modulePath, pkgPathInMod string, pkg *list.Package, err error) {
	modulePath, err = GetModulePath()
	if err != nil {
		return "", "", nil, fmt.Errorf("Failed to get current module path: %w", err)
	}
	pkgPathInMod, err = GetPackagePathInModule(modulePath)
	if err != nil {
		return "", "", nil, fmt.Errorf("Failed to get current package path: %w", err)
	}
	var pkgs list.Packages
	pkgs, err = list.ListAllPackages(".")
	if err != nil {
		return "", "", nil, fmt.Errorf("Failed to get import path of package in current directory. Maybe you are not in a directory that contains a Go package: %w", err)
	}
	pkg = pkgs.First()
	return
}

func FindTest(testName string) (*TestDetails, error) {
	modPath, pkgPathInMod, pkg, err := WhereAmI()
	if err != nil {
		return nil, errors.Wrap(err, "WhereAmI")
	}

	funcDetailsInSubdirs, err := DiscoverSubdirsForTestFuncDetails(modPath, filepath.Join(modPath, pkgPathInMod))
	if err != nil {
		return nil, errors.Wrap(err, "DiscoverSubdirsForTestFuncDetails")
	}

	var foundFuncDetails *TestFuncDetails
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

	testDetails := TestDetails{
		PackagePath:  pkgPathInMod,
		Package:      pkg,
		ImplFuncFile: implFile,
		ImplFuncName: implFuncName,
		ImplFuncLine: implFuncDetails.Line,
		TestFuncFile: foundFuncDetails.Path,
		TestFuncName: testName,
		TestFuncLine: foundFuncDetails.Line,
	}

	return &testDetails, nil
}
