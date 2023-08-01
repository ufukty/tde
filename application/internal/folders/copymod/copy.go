package copymod

import (
	"fmt"
	"log"
	"strings"

	"os"
	"path/filepath"

	"golang.org/x/exp/slices"
)

var DefaultInclExt = []string{"go", "mod", "sum"}
var DefaultSkipDirs = []string{".git", "build", "docs", ".vscode", "vendor"}

func CopyModule(dst string, src string, incSubdirs bool, skipDirs, skipSubdirs, includeExt []string, enableLogging bool) error {
	skipSubdirs = append(skipSubdirs, DefaultSkipDirs...)
	includeExt = append(includeExt, DefaultInclExt...)

	return filepath.Walk(src, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk dir %q: %w", file, err)
		}
		srcRel, err := filepath.Rel(src, file)
		if err != nil {
			return fmt.Errorf("could not get the entry's in-module path: %w", err)
		}
		dstAbs := filepath.Join(dst, srcRel)
		switch info.Mode() & os.ModeType {
		case os.ModeDir:
			var (
				isInSkipDirs    = slices.Index(skipDirs, srcRel) != -1
				isInSkipSubdirs = slices.Index(skipSubdirs, filepath.Base(srcRel)) != -1
			)
			if !incSubdirs || isInSkipDirs || isInSkipSubdirs {
				if enableLogging {
					log.Println(srcRel, "(skip dir)")
				}
				return filepath.SkipDir
			}
			if enableLogging {
				log.Println(srcRel)
			}
			if err := createDirIfDoesNotExists(dstAbs, 0755); err != nil {
				return fmt.Errorf("could not create dir %q: %w", srcRel, err)
			}
		case os.ModeSymlink:
			if enableLogging {
				log.Println(srcRel, "(skip symlink)")
			}
			break
		default:
			ext := strings.TrimPrefix(filepath.Ext(filepath.Base(srcRel)), ".")
			if slices.Index(includeExt, ext) == -1 {
				if enableLogging {
					log.Println(srcRel, "(skip file)")
				}
				return nil
			}
			if enableLogging {
				log.Println(srcRel)
			}
			if err := CopyFile(file, dstAbs); err != nil {
				return fmt.Errorf("copying %q: %w", srcRel, err)
			}
		}
		return nil // keep walk
	})
}
