package astwutl

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"tde/i18n"

	"golang.org/x/exp/maps"
)

func LoadDir(dirpath string) (*token.FileSet, map[string]*ast.Package, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dirpath, nil, parser.AllErrors|parser.SkipObjectResolution)
	if err != nil {
		return nil, nil, fmt.Errorf("LoadDir: %w", err)
	}
	return fset, pkgs, nil
}

func LoadFile(filepath string) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, filepath, nil, parser.AllErrors|parser.SkipObjectResolution)
	if err != nil {
		return nil, nil, fmt.Errorf("LoadFile: %w", err)
	}
	return fset, astFile, nil
}

func ParseString(content string) (*token.FileSet, ast.Node, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", content, parser.AllErrors|parser.SkipObjectResolution)
	if err != nil {
		return nil, nil, fmt.Errorf("ParseString: %w", err)
	}
	return fset, astFile, nil
}

func LoadPackageFromDir(path string) (*ast.Package, error) {
	var (
		pkgs    map[string]*ast.Package
		pkgList []string
		wd      string
		err     error
	)

	wd, err = os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Can't get the working directory: %w", err)
	}
	err = os.Chdir(path)
	if err != nil {
		return nil, fmt.Errorf("Can't switch to the directory of requested package: %w", err)
	}
	defer os.Chdir(wd)

	_, pkgs, err = LoadDir(".")
	if err != nil {
		return nil, fmt.Errorf(err.Error()+": %w", i18n.ErrAstConversionFailed)
	}
	pkgList = maps.Keys(pkgs)
	if l := len(pkgList); l == 0 {
		return nil, i18n.ErrNoPackagesFound
	} else if l > 1 {
		return nil, i18n.ErrMultiplePackagesFound
	}
	return pkgs[pkgList[0]], nil
}
