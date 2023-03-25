package types

import "path/filepath"

type (
	TempPath     string // absolute path. result of the MkdirTemp
	AbsolutePath string
	ModulePath   string
	InModulePath string // doesn't contain module's path since it can differ from original to the one we are working on
	SlotPath     string // doesn't contain temp dir's path at beginning and in-module paths at ending
)

func (base ModulePath) Rel(target ModulePath) (string, error) {
	return filepath.Rel(string(base), string(target))
}

func (ap AbsolutePath) Join(imp InModulePath) AbsolutePath {
	return AbsolutePath(filepath.Join(string(ap), string(imp)))
}

func (tmp TempPath) Join(slot SlotPath) AbsolutePath {
	return AbsolutePath(filepath.Join(string(tmp), string(slot)))
}

func (tmp TempPath) FindInModulePath(slot SlotPath, inModule InModulePath) string {
	return filepath.Join(
		filepath.Join(string(tmp), string(slot)),
		string(inModule),
	)
}
