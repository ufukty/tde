package utilities

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/pkg/errors"
)

func WorkingDir() (string, error) {
	stdOut, _, err := RunCommandForOutput("pwd", "-P")
	if err != nil {
		return "", errors.Wrap(err, "Failed to run command pwd")
	}
	return stdOut, nil
}

func Dirs(path string) ([]fs.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("listing dir entries: %w", err)
	}
	return FilteredMap(entries, func(i int, entry fs.DirEntry) (fs.DirEntry, bool) { return entry, entry.IsDir() }), nil
}

func Files(path string) ([]fs.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("listing dir entries: %w", err)
	}
	return FilteredMap(entries, func(i int, entry fs.DirEntry) (fs.DirEntry, bool) { return entry, !entry.IsDir() }), nil
}
