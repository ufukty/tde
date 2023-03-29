package upload

import (
	"fmt"
	"os"
	"tde/internal/command"
	"tde/internal/folders/archive"
	"tde/internal/folders/discovery"
	utl "tde/internal/utilities"

	"github.com/pkg/errors"
)

type Command struct {
	OnlyArchive bool                `long:"only-archive"`
	ExcludeDirs command.MultiString `short:"e" long:"exclude-dir"`
}

func (c *Command) Run() {
	modulePath, err := discovery.GetModulePath()
	if err != nil {
		utl.Terminate(errors.Wrap(err, "Could not find the path of Go module root"))
	}

	c.ExcludeDirs = append(c.ExcludeDirs, archive.DefaultSkipDirs...)
	zipPath, err := archive.Directory(modulePath, true, c.ExcludeDirs)
	if err != nil {
		utl.Terminate(errors.Wrap(err, "Could not create archive for module"))
	}

	fmt.Println("Archive path:", zipPath)
	if c.OnlyArchive {
		os.Exit(0)
	} else {
		fmt.Println("cont")
	}
}
