package discovery

import (
	"go/ast"
	"path/filepath"
	"strings"
	"tde/internal/code"
	"tde/internal/utilities"
	"tde/models/in_program_models"

	"github.com/pkg/errors"
)

var (
	ModuleNotFound = errors.New("this directory is not part of a Go module")
)

// Returns the import path for the package inside working directory
func FindImportPathOfThePackage() (string, error) {
	out, err := utilities.RunCommandForOutput("go", "list")
	if err != nil {
		return "", errors.Wrap(err, "running 'go list' is failed on the working directory")
	}
	return utilities.StripOnlyLineFromCommandOuput(out)
}

// Returns the absolute path of the module that working directory is in it
func FindModulePath() (string, error) {
	path, err := utilities.RunCommandForOutput("go", "env", "GOMOD")
	if err != nil {
		return "", errors.Wrap(err, "failed to run 'go env GOMOD'")
	}
	path, err = utilities.StripOnlyLineFromCommandOuput(path)
	if err != nil {
		return "", errors.Wrap(err, "could not strip GOMOD path from the output of 'go env GOMOD'")
	}
	if path == "/dev/null" {
		return "", ModuleNotFound
	}
	return filepath.Dir(path), nil
}

// func FindTestFunctionName(filepath string, lineStart int, lineEnd int) (string, error) {
// 	c := code.Code{}
// 	err := c.LoadFromFile(filepath)
// 	if err != nil {
// 		return "", errors.Wrap(err, "failed on load test file")
// 	}
// }

type TestFunctionDetails struct {
	Name     string // function name
	Position int    // character position
	Line     int    // starts with 1
}

// TODO: Use it by language-server
func DetectTestFunctions(filepath string) ([]TestFunctionDetails, error) {
	c := code.NewCode()
	err := c.LoadFromFile(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load file and parse into AST tree")
	}

	testFunctions := []TestFunctionDetails{}

	c.InspectWithTrace(func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		depth := len(parentTrace)
		if depth == 2 {
			if n, ok := n.(*ast.FuncDecl); ok {
				functionName := n.Name.Name
				if strings.Index(functionName, "TDE") == 0 {
					testFunctions = append(testFunctions, TestFunctionDetails{
						Name:     functionName,
						Line:     c.LineNumberOfPosition(n.Pos()),
						Position: int(n.Pos()),
					})
				}
			}
		}
		return depth < 2
	})

	return testFunctions, nil
}

func Discover() (*in_program_models.DiscoveryResponse, error) {
	modulePath, err := FindModulePath()
	if err != nil {
		return nil, errors.Wrap(err, "failed on finding module absoute path")
	}
	packageImportPath, err := FindImportPathOfThePackage()
	if err != nil {
		return nil, errors.Wrap(err, "failed on target package's import path")
	}
	return &in_program_models.DiscoveryResponse{
		ModuleAbsolutePath:      modulePath,
		TargetPackageImportPath: packageImportPath,
	}, nil
}
