package archive

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

var DefaultInclExt = []string{"go", "mod", "sum"}
var DefaultSkipDirs = []string{".git", "build", "docs", ".vscode", "vendor"}

func directory(dst io.Writer, src string, incSubdirs bool, skipDirs, skipSubdirs, includeExt []string, enableLogging bool) (err error) {
	zipWriter := zip.NewWriter(dst)
	defer zipWriter.Close()

	err = filepath.Walk(src, func(subPath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "failed to walk directory")
		}

		inZipSubPath, err := filepath.Rel(src, subPath)
		if err != nil {
			return errors.Wrap(err, "failed to clean path for a file")
		}

		if fileInfo.IsDir() {
			var (
				isInSkipDirs    = slices.Index(skipDirs, inZipSubPath) != -1
				isInSkipSubdirs = slices.Index(skipSubdirs, filepath.Base(inZipSubPath)) != -1
			)
			if !incSubdirs || isInSkipDirs || isInSkipSubdirs {
				if enableLogging {
					log.Println("skip dir:", inZipSubPath)
				}
				return filepath.SkipDir
			}
			return nil // keep walk
		} else {
			ext := strings.TrimPrefix(filepath.Ext(filepath.Base(inZipSubPath)), ".")
			if slices.Index(includeExt, ext) == -1 {
				if enableLogging {
					log.Println("skip ext:", inZipSubPath)
				}
				return nil
			}
		}

		if enableLogging {
			log.Println("archiving:", inZipSubPath)
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
		return errors.Wrap(err, "failed on archiving the directory")
	}

	return nil
}

func Directory(relativePath string, includeSubfolders bool, skipDirs, skipSubdirs, includeExt []string, enableLogging bool) (path string, err error) {
	target, err := os.CreateTemp(os.TempDir(), "tde.CodeArchive.*.zip")
	if err != nil {
		return "", errors.Wrap(err, "failed to create temporary zip file")
	}
	defer target.Close()
	return target.Name(), directory(target, relativePath, includeSubfolders, skipDirs, skipSubdirs, includeExt, enableLogging)
}

func DirectoryToFile(target string, relativePath string, includeSubfolders bool, skipDirs, skipSubdirs, includeExt []string, enableLogging bool) error {
	fh, err := os.Create(target)
	if err != nil {
		return errors.Wrap(err, "failed to create temporary zip file")
	}
	defer fh.Close()
	return directory(fh, relativePath, includeSubfolders, skipDirs, skipSubdirs, includeExt, enableLogging)
}
