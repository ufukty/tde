package discovery

import (
	"tde/internal/astw/astwutl"

	"fmt"
)

func TargetFunctionInFile(path string, funcname string) (*TargetFunction, error) {
	funcDecl, fset, err := astwutl.FindFunctionInFile(path, funcname)
	if err != nil {
		return nil, fmt.Errorf("searching ast of file %q for the function %q: %w", path, funcname, err)
	}
	return &TargetFunction{
		Name:      funcname,
		Path:      path,
		LineStart: astwutl.LineNumberOfPosition(fset, funcDecl.Pos()),
		LineEnd:   astwutl.LineNumberOfPosition(fset, funcDecl.End()),
	}, nil
}
