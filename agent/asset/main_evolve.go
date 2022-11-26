//go:build evolve
// +build evolve

package main

// usage from commandline, after this file places into user's project root
//   go build -tags evolve

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func checkFolder(path string) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Could not list the directory '%s'", path))
	}
	foldersToCheck := []string{}
	filesToCheck := []string{}
	for _, dirEntry := range dirEntries {
		path := filepath.Join(path, dirEntry.Name())
		if dirEntry.IsDir() {
			foldersToCheck = append(foldersToCheck, path)
		} else {
			filesToCheck = append(filesToCheck, path)
		}
	}

	for _, filePath := range filesToCheck {
		fmt.Println(filePath)
	}

	for _, folderPath := range foldersToCheck {
		fmt.Println(folderPath)
	}
}

func main() {
	targetFunc := os.Args[1]

	// search all files that ends with xxx_evolve.go for a function starts with "Evolve<Target>"
	fmt.Println(targetFunc)
}
