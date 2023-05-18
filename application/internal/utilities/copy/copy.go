package copy

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func Directory(scrDir, dest string) error {
	entries, err := os.ReadDir(scrDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateIfNotExists(destPath, 0755); err != nil {
				return err
			}
			if err := Directory(sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := SymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := File(sourcePath, destPath); err != nil {
				return err
			}
		}

		// if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
		// 	return err
		// }

		fInfo, err := entry.Info()
		if err != nil {
			return err
		}

		isSymlink := fInfo.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, fInfo.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

func File(srcFile, dstFile string) error {
	dst, err := os.Create(dstFile)
	if err != nil {
		return errors.Wrap(err, "os.Create(dstFile)")
	}
	defer dst.Close()

	src, err := os.Open(srcFile)
	if err != nil {
		return errors.Wrap(err, "os.Open(srcFile)")
	}
	defer src.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return errors.Wrap(err, "io.Copy(dst, src)")
	}

	return nil
}

func exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateIfNotExists(dir string, perm os.FileMode) error {
	if exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func SymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}
