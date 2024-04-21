package canonicalize

import (
	"path/filepath"
	"slices"
	"tde/internal/evolution/evaluation/list"
)

func isInStdLib(pkg string) bool {
	return slices.Index(stdlibpkgs, pkg) != -1
}

func isInModule(modpkgdirs map[string]string, pkg string) bool {
	if _, ok := modpkgdirs[pkg]; ok {
		return true
	}
	return false
}

// items of "pkgs" to be like either of fmt or mymodule/cmd/mypackage
func CanonicalizePaths(goroot, modroot string, modpkgs list.Packages, pkgs []string) (canonicalized []string) {
	modpkgdirs := map[string]string{}
	for _, pkg := range modpkgs {
		modpkgdirs[pkg.ImportPath] = pkg.Dir
	}

	for _, pkg := range pkgs {
		if isInStdLib(pkg) {
			canonicalized = append(canonicalized, filepath.Join(goroot, "src", pkg))
		} else if isInModule(modpkgdirs, pkg) {
			canonicalized = append(canonicalized, modpkgdirs[pkg])
		} else {
			canonicalized = append(canonicalized, pkg)
		}
	}
	return
}
