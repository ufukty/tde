package upload

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

func findModuleRoot() (string, error) {
	cmd := exec.Command("go", "env")
	bytes, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "failed on running \"go env\" command")
	}

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		seperatorIndex := strings.Index(line, "=")
		if line[:seperatorIndex] == "GOMOD" {
			valuePart := line[seperatorIndex+1:]
			strippedQuotes := valuePart[1 : len(valuePart)-1]
			dir := filepath.Dir(strippedQuotes)
			return dir, nil
		}
	}

	return "", errors.New("could not find GOMOD in output of \"go env\" command")
}

func ArchiveDirectory(relativePath string, includeSubfolders bool, skipDirs []string) (string, error) {
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
