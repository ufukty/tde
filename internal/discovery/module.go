package discovery

import (
	"fmt"
	"go/ast"
	"strings"
	astw_utl "tde/internal/astw/utilities"
	"tde/internal/utilities"

	"io/fs"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

type TargetFunction struct {
	Name string
	Path string
	Line int
}

type TDEFunction struct {
	Name  string
	Path  string
	Line  int
	Calls []*ast.CallExpr
}

func FindTargetFunction(targetFilePath, targetFunctionName string) (*TargetFunction, error) {

	fset, astFile, err := astw_utl.LoadFile(targetFilePath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("could not create AST for the file \"%s\"", targetFilePath))
	}

	var targetFunction *TargetFunction
	ast.Inspect(astFile, func(node ast.Node) bool {
		if node, ok := node.(*ast.FuncDecl); ok {
			if node.Name.Name == targetFunctionName {
				targetFunction = &TargetFunction{
					Name: targetFunctionName,
					Path: targetFilePath,
					Line: astw_utl.LineNumberOfPosition(fset, node.Pos()),
				}
			}
		}
		return node == astFile
	})
	if targetFunction == nil {
		return nil, errors.New(fmt.Sprintf("Could not find function %s in the file %s", targetFunctionName, targetFilePath))
	}

	return targetFunction, nil
}

func DiscoverTDEFunction(funcDecl *ast.FuncDecl) []*ast.CallExpr {
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

func GetTargetFuncForTDEFunc(tdeFile, tdeFuncName string) (targetFile, targetFuncName string) {
	targetFile = strings.TrimSuffix(tdeFile, "_tde.go") + ".go"
	targetFuncName = strings.TrimPrefix(tdeFuncName, "TDE_")
	return
}

// TODO: Call it from language-server
func DiscoverFile(path string) ([]TDEFunction, error) {
	testFunctions := []TDEFunction{}

	fset, astFile, err := astw_utl.LoadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load file and parse into AST tree")
	}

	ast.Inspect(astFile, func(node ast.Node) bool {
		if funcDecl, ok := node.(*ast.FuncDecl); ok {
			functionName := funcDecl.Name.Name
			if strings.Index(functionName, "TDE_") == 0 {
				testFunctions = append(testFunctions, TDEFunction{
					Name:  functionName,
					Path:  path,
					Line:  astw_utl.LineNumberOfPosition(fset, node.Pos()),
					Calls: DiscoverTDEFunction(funcDecl),
				})
			}
		}
		return node == astFile
	})

	return testFunctions, nil
}

func DiscoverSubdirs(path string) ([]TDEFunction, error) {
	excludeDirs := []string{".git", "build", "artifacts"}
	excludePaths := utilities.Map(excludeDirs, func(i int, in string) string { return filepath.Join(path, in) })

	testFunctions := []TDEFunction{}
	err := filepath.Walk(path, func(filePath string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "failed to walk directory: "+filePath)
		}

		if fileInfo.IsDir() {
			if slices.Contains(excludePaths, filePath) {
				return filepath.SkipDir
			} else {
				return nil
			}
		}

		if strings.HasSuffix(filePath, "_tde.go") {
			testFunctionsInFile, err := DiscoverFile(filePath)
			if err != nil {
				// log.Println(errors.Wrap(err, "Could not discover file: "+filePath))
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
