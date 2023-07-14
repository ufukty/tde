package utilities

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"tde/i18n"

	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
)

func LoadDir(dirpath string) (*token.FileSet, map[string]*ast.Package, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dirpath, nil, parser.AllErrors)
	if err != nil {
		return nil, nil, errors.Wrap(err, "LoadDir")
	}
	return fset, pkgs, nil
}

func LoadFile(filepath string) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, filepath, nil, parser.AllErrors)
	if err != nil {
		return nil, nil, errors.Wrap(err, "LoadFile")
	}
	return fset, astFile, nil
}

func ParseString(content string) (*token.FileSet, ast.Node, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", content, parser.AllErrors)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ParseString")
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
		return nil, errors.Wrap(err, "Can't get the working directory")
	}
	err = os.Chdir(path)
	if err != nil {
		return nil, errors.Wrap(err, "Can't switch to the directory of requested package")
	}
	defer os.Chdir(wd)

	_, pkgs, err = LoadDir(".")
	if err != nil {
		return nil, errors.Wrap(i18n.ErrAstConversionFailed, err.Error())
	}
	pkgList = maps.Keys(pkgs)
	if l := len(pkgList); l == 0 {
		return nil, i18n.ErrNoPackagesFound
	} else if l > 1 {
		return nil, i18n.ErrMultiplePackagesFound
	}
	return pkgs[pkgList[0]], nil
}
