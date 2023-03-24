package copy_module

import (
	ucopy "tde/internal/utilities/copy"

	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

var DefaultSkipDirs = []string{".git", "build", "docs", ".vscode"}

func Module(srcMod string, dstMod string, includeSubfolders bool, skipDirs []string) error {
	return filepath.Walk(srcMod, func(srcAbs string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "failed to walk directory")
		}

		srcRel, err := filepath.Rel(srcMod, srcAbs)
		if err != nil {
			return errors.Wrap(err, "getting relative path from absolute")
		}

		dstAbs := filepath.Join(dstMod, srcRel)

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if !includeSubfolders || slices.Index(skipDirs, srcRel) != -1 {
				// log.Println("skip dir:", srcRel)
				return filepath.SkipDir
			}
			// log.Println("copy dir:", srcRel)
			if err := ucopy.CreateIfNotExists(dstAbs, 0755); err != nil {
				return errors.Wrap(err, "CreateIfNotExists")
			}

		case os.ModeSymlink:
			break

		default:
			// log.Println("copy file:", srcRel)
			if err := ucopy.File(srcAbs, dstAbs); err != nil {
				return errors.Wrap(err, "File")
			}
		}

		return nil // keep walk
	})
}
