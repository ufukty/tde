package upload

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func ArchiveDirectory(relativePath string) (string, error) {
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

		fmt.Println("archiving:", inZipSubPath)

		if fileInfo.IsDir() {
			if subPath == relativePath {
				return nil
			} else {
				return filepath.SkipDir
			}
		}

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
