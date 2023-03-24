package archive

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

var DefaultSkipDirs = []string{".git", "build", "docs", ".vscode"}

func Directory(relativePath string, includeSubfolders bool, skipDirs []string) (path string, err error) {
	target, err := os.CreateTemp(os.TempDir(), "tde.CodeArchive.*.zip")
	if err != nil {
		return "", errors.Wrap(err, "failed to create temporary zip file")
	}
	defer target.Close()

	zipWriter := zip.NewWriter(target)
	defer zipWriter.Close()

	err = filepath.Walk(relativePath, func(subPath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "failed to walk directory")
		}

		inZipSubPath, err := filepath.Rel(relativePath, subPath)
		if err != nil {
			return errors.Wrap(err, "failed to clean path for a file")
		}

		if fileInfo.IsDir() {
			if !includeSubfolders || slices.Index(skipDirs, inZipSubPath) != -1 {
				log.Println("skip dir:", inZipSubPath)
				return filepath.SkipDir
			}
			return nil // keep walk
		}

		log.Println("archiving:", inZipSubPath)

		subFile, err := os.Open(subPath)
		if err != nil {
			return errors.Wrap(err, "failed to open file")
		}
		defer subFile.Close()

		zipSubPathWriter, err := zipWriter.Create(inZipSubPath)
		if err != nil {
			return errors.Wrap(err, "failed to get a writer for a file to write it into the archive")
		}

		_, err = io.Copy(zipSubPathWriter, subFile)
		if err != nil {
			return errors.Wrap(err, "failed to add a file into archive")
		}

		return nil
	})
	if err != nil {
		return target.Name(), errors.Wrap(err, "failed on archiving the directory")
	}

	return target.Name(), nil
}
