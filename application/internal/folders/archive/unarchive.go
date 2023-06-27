package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

const (
	maxAllowedUncompressedTotalFileSize  = 4 * 1024 * 1024 // 4 MB
	maxAllowedUncompressedSingleFileSize = 10 * 1024       // 10 KB
	maxAllowedFile                       = 200
	maxAllowedSubfolderDepth             = 10
)

var (
	ErrExtensionUnallowed      = errors.New("ErrExtensionUnallowed")
	ErrZipOpen                 = errors.New("ErrZipOpen")
	ErrZipFileExceedsLimit     = errors.New(fmt.Sprintf("Zip file contains a file that its uncompressed size exceeds the limit for single file (%d)", maxAllowedUncompressedSingleFileSize))
	ErrZipExceedsLimit         = errors.New(fmt.Sprintf("Zip file's uncompressed size exceeds the limit (%d bytes)", maxAllowedUncompressedTotalFileSize))
	ErrRelativePathFound       = errors.New("Relative paths are not allowed in a zip archive")
	ErrTooManyFiles            = errors.New("A module upload can not have more than 200 files")
	ErrSubfolderExceedingDepth = errors.New("More than 10 nested subfolders are unallowed.")
)

// Regular expression to match relative path segments
var (
	unsafePathFragmentMatcher = regexp.MustCompile(`(^[A-Za-z]\:.*$)|(^\.{0,2}[\\\/].*$)|(^\.{1,2}$)|(^.*[\\\/]{2}.*$)|(^.*([\\\/]\.{1,2}[\\\/]).*$)|(^.*[\\\/]\.{1,2}$)`) // https://regex101.com/r/3wTjZa/1
)

func isPathSafe(name string) bool {
	return !unsafePathFragmentMatcher.MatchString(name)
}

func isDirInsideDest(dir string, dest string) bool {
	return strings.HasPrefix(dir, dest)
}

func checkFileDepth(path string, destDepth int) error {
	if len(strings.Split(path, "/")) > destDepth+maxAllowedFile {
		return ErrSubfolderExceedingDepth
	}
	return nil
}

func createParentFoldersForFile(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return errors.Wrap(err, "os.MkdirAll")
	}
	return nil
}

func Unarchive(src string, dest string) error {
	var (
		zipReader     *zip.ReadCloser
		err           error
		totalFileSize uint64
		totalFile     int
	)

	if !strings.HasSuffix(src, ".zip") {
		return ErrExtensionUnallowed
	}

	zipReader, err = zip.OpenReader(src)
	if err != nil {
		return errors.Wrap(err, "could not create zip reader")
	}
	defer zipReader.Close()

	var moduleDepth = len(strings.Split(dest, "/"))

	for _, containedFileHandler := range zipReader.File {
		var filename = containedFileHandler.Name

		if totalFile++; totalFile == maxAllowedFile {
			return ErrTooManyFiles
		}
		if !isPathSafe(filename) {
			return errors.Wrap(ErrRelativePathFound, filename)
		}
		if containedFileHandler.UncompressedSize64 > maxAllowedUncompressedSingleFileSize {
			return errors.Wrap(ErrZipFileExceedsLimit, fmt.Sprintf("%s => %d", filename, containedFileHandler.UncompressedSize64))
		}
		if totalFileSize += containedFileHandler.UncompressedSize64; totalFileSize > maxAllowedUncompressedTotalFileSize {
			return ErrZipExceedsLimit
		}

		rc, err := containedFileHandler.Open()
		if err != nil {
			return errors.Wrapf(err, "could not open the file inside zip '%s'", filename)
		}
		defer rc.Close()

		path := filepath.Clean(filepath.Join(dest, filename))

		if !isDirInsideDest(path, dest) {
			return errors.Wrap(ErrRelativePathFound, filename)
		}
		if err = checkFileDepth(path, moduleDepth); err != nil {
			return errors.Wrap(err, filename)
		}
		if err = createParentFoldersForFile(path); err != nil {
			return errors.Wrapf(err, "creating parent dirs for file %s", filename)
		}

		targetFile, err := os.Create(path)
		if err != nil {
			return errors.Wrapf(err, "could not create file to write zip file '%s'", filename)
		}
		defer targetFile.Close()

		_, err = io.Copy(targetFile, rc)
		if err != nil {
			return errors.Wrapf(err, "could not extract the file '%s'", filename)
		}
	}

	return nil
}
