package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	maxAllowedUncompressedTotalFileSize  = 10 * 1024 * 1024 // 10 MB
	maxAllowedUncompressedSingleFileSize = 100 * 1024       // 100 KB
	maxAllowedFile                       = 500
	maxAllowedSubfolderDepth             = 10
)

var (
	ErrExtensionUnallowed      = fmt.Errorf("ErrExtensionUnallowed")
	ErrZipOpen                 = fmt.Errorf("ErrZipOpen")
	ErrZipFileExceedsLimit     = fmt.Errorf(fmt.Sprintf("Zip file contains a file that its uncompressed size exceeds the limit for single file (%d)", maxAllowedUncompressedSingleFileSize))
	ErrZipExceedsLimit         = fmt.Errorf(fmt.Sprintf("Zip file's uncompressed size exceeds the limit (%d bytes)", maxAllowedUncompressedTotalFileSize))
	ErrRelativePathFound       = fmt.Errorf("Relative paths are not allowed in a zip archive")
	ErrTooManyFiles            = fmt.Errorf(fmt.Sprintf("A module upload can not have more than %d files", maxAllowedFile))
	ErrSubfolderExceedingDepth = fmt.Errorf(fmt.Sprintf("More than %d nested subfolders are unallowed.", maxAllowedSubfolderDepth))
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

func isPathTooDeep(path string, destDepth int) bool {
	return len(strings.Split(path, "/")) > destDepth+maxAllowedFile
}

func createParentFoldersForFile(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("os.MkdirAll: %w", err)
	}
	return nil
}

func Unarchive(src string, dest string) (err error, errFile string) {
	var (
		zipReader     *zip.ReadCloser
		totalFileSize uint64
		totalFile     int
	)

	if !strings.HasSuffix(src, ".zip") {
		return ErrExtensionUnallowed, ""
	}

	zipReader, err = zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("could not create zip reader: %w", err), ""
	}
	defer zipReader.Close()

	moduleDepth := len(strings.Split(dest, "/"))

	for _, containedFileHandler := range zipReader.File {
		filename := containedFileHandler.Name

		if totalFile++; totalFile == maxAllowedFile {
			return ErrTooManyFiles, ""
		}
		if !isPathSafe(filename) {
			return ErrRelativePathFound, filename
		}
		if containedFileHandler.UncompressedSize64 > maxAllowedUncompressedSingleFileSize {
			return ErrZipFileExceedsLimit, filename
		}
		if totalFileSize += containedFileHandler.UncompressedSize64; totalFileSize > maxAllowedUncompressedTotalFileSize {
			return ErrZipExceedsLimit, ""
		}

		rc, err := containedFileHandler.Open()
		if err != nil {
			return err, filename
		}
		defer rc.Close()

		path := filepath.Clean(filepath.Join(dest, filename))

		if !isDirInsideDest(path, dest) {
			return ErrRelativePathFound, filename
		}
		if isPathTooDeep(path, moduleDepth) {
			return ErrSubfolderExceedingDepth, filename
		}
		if err = createParentFoldersForFile(path); err != nil {
			return fmt.Errorf("creating parent dirs for file: %w", err), ""
		}

		targetFile, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("could not create file to write zip file': %w", err), filename
		}
		defer targetFile.Close()

		_, err = io.Copy(targetFile, rc)
		if err != nil {
			return fmt.Errorf("could not extract the file: %w", err), filename
		}
	}

	return nil, ""
}
