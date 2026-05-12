package osw

import (
	"fmt"
	"io/fs"
	"os"

	"tde/internal/utilities/functional"
)

func WorkingDir() (string, error) {
	stdOut, _, err := RunCommandForOutput("pwd", "-P")
	if err != nil {
		return "", fmt.Errorf("Failed to run command pwd: %w", err)
	}
	return stdOut, nil
}

func Dirs(path string) ([]fs.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("listing dir entries: %w", err)
	}
	return functional.Mapf(entries, func(i int, entry fs.DirEntry) (fs.DirEntry, bool) { return entry, entry.IsDir() }), nil
}

func Files(path string) ([]fs.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("listing dir entries: %w", err)
	}
	return functional.Mapf(entries, func(i int, entry fs.DirEntry) (fs.DirEntry, bool) { return entry, !entry.IsDir() }), nil
}
