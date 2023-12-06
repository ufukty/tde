package discovery

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"tde/internal/astw/astwutl"
	"tde/internal/utilities/functional"
	"tde/internal/utilities/osw"

	"golang.org/x/exp/maps"
)

// all paths returned will be relative to <moduleRoot>
func TestFunctionsInFile(path string) ([]TestFunction, error) {
	tests := []TestFunction{}

	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	ast.Inspect(astFile, func(node ast.Node) bool {
		if funcDecl, ok := node.(*ast.FuncDecl); ok {
			name := funcDecl.Name.Name
			if strings.Index(name, "TDE_") == 0 {
				tests = append(tests, TestFunction{
					Name:      name,
					Path:      path,
					LineStart: astwutl.LineNumberOfPosition(fset, node.Pos()),
					LineEnd:   astwutl.LineNumberOfPosition(fset, node.End()),
					Calls:     FunctionCalls(funcDecl),
				})
			}
		}
		return node == astFile
	})

	return tests, nil
}

func TestFunctionInDir(path string, funcname string) (*TestFunction, error) {
	tests, _, err := TestFunctionsInDir(path)
	if err != nil {
		return nil, fmt.Errorf("listing all test functions in dir %q: %w", path, err)
	}
	matches := functional.Mapf(tests, func(i int, v TestFunction) (TestFunction, bool) { return v, v.Name == funcname })
	switch len(matches) {
	case 0:
		tests := functional.Map(tests, func(i int, v TestFunction) string { return v.Name })
		return nil, fmt.Errorf("%q not found. Found tests are %s", funcname, strings.Join(tests, ", "))
	case 1:
		return &matches[0], nil
	default:
		return nil, fmt.Errorf("multiple tests have found with name %q", funcname)
	}
}

// skipped is for filename:syntax-error
func TestFunctionsInDir(path string) (tests []TestFunction, skipped map[string]error, err error) {
	tests = []TestFunction{}
	skipped = map[string]error{}
	files, err := osw.Files(path)
	if err != nil {
		return nil, nil, fmt.Errorf("listing files in dir %q: %w", path, err)
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), "_tde.go") {
			continue
		}
		testsInFile, err := TestFunctionsInFile(filepath.Join(path, file.Name()))
		if err != nil {
			skipped[filepath.Join(path, file.Name())] = err
		} else {
			tests = append(tests, testsInFile...)
		}
	}
	return
}

// skipped is for filename:syntax-error
func TestFunctionsInSubdirs(path string) (tests []TestFunction, skipped map[string]error, err error) {
	// excludeDirs := copymod.DefaultSkipDirs
	// excludePaths := functional.Map(excludeDirs, func(i int, v string) string {
	// 	return filepath.Join(path, v)
	// })
	tests, skipped, err = TestFunctionsInDir(path)
	if err != nil {
		return nil, nil, fmt.Errorf(": %w", err)
	}
	dirs, err := osw.Dirs(path)
	if err != nil {
		return nil, nil, fmt.Errorf("listing dirs in %q: %w", path, err)
	}
	for _, dir := range dirs {
		testsInSubdirs, skippedInSubdirs, err := TestFunctionsInSubdirs(filepath.Join(path, dir.Name()))
		if err != nil {
			return nil, nil, fmt.Errorf("(in recursion) listing test functions in subdirs of %q: %w", dir.Name(), err)
		}
		tests = append(tests, testsInSubdirs...)
		maps.Copy(skipped, skippedInSubdirs)
	}
	return tests, skipped, nil
}
