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
)

var (
	ErrExtensionUnallowed  = errors.New("ErrExtensionUnallowed")
	ErrZipOpen             = errors.New("ErrZipOpen")
	ErrZipFileExceedsLimit = errors.New(fmt.Sprintf("Zip file contains a file that its uncompressed size exceeds the limit for single file (%d)", maxAllowedUncompressedSingleFileSize))
	ErrZipExceedsLimit     = errors.New(fmt.Sprintf("Zip file's uncompressed size exceeds the limit (%d bytes)", maxAllowedUncompressedTotalFileSize))
	ErrRelativePathFound   = errors.New("Relative paths are not allowed in a zip archive")
)

// Regular expression to match relative path segments
var (
	unsafePathFragmentMatcher = regexp.MustCompile(`(^[A-Za-z]\:.*$)|(^\.{0,2}[\\\/].*$)|(^\.{1,2}$)|(^.*[\\\/]{2}.*$)|(^.*([\\\/]\.{1,2}[\\\/]).*$)|(^.*[\\\/]\.{1,2}$)`) // https://regex101.com/r/3wTjZa/1
)

func isPathSafe(name string) bool {
	return !unsafePathFragmentMatcher.MatchString(name)
}

func Unarchive(src string, dest string) error {
	var (
		zipReader     *zip.ReadCloser
		err           error
		totalFileSize uint64
	)

	if !strings.HasSuffix(src, ".zip") {
		return ErrExtensionUnallowed
	}

	zipReader, err = zip.OpenReader(src)
	if err != nil {
		return errors.Wrap(err, "Could not create zip reader")
	}
	defer zipReader.Close()

	for _, containedFileHandler := range zipReader.File {

		if !isPathSafe(containedFileHandler.Name) {
			return ErrRelativePathFound
		}

		if containedFileHandler.UncompressedSize64 > maxAllowedUncompressedSingleFileSize {
			return ErrZipFileExceedsLimit
		}
		totalFileSize += containedFileHandler.UncompressedSize64
		if totalFileSize > maxAllowedUncompressedTotalFileSize {
			return ErrZipExceedsLimit
		}

		rc, err := containedFileHandler.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Create the target file on disk
		path := filepath.Join(dest, containedFileHandler.Name)
		targetFile, err := os.Create(path)
		if err != nil {
			return err
		}
		defer targetFile.Close()

		// Copy the file contents to disk
		_, err = io.Copy(targetFile, rc)
		if err != nil {
			return err
		}
	}

	return nil
}
