package upload

import (
	"fmt"
	"log"
	"os"
	"tde/internal/command"
	"tde/internal/folders/archive"
	"tde/internal/folders/discovery"

	"github.com/pkg/errors"
)

type Command struct {
	OnlyArchive bool                `long:"only-archive"`
	ExcludeDirs command.MultiString `short:"e" long:"exclude-dir"`
}

func (c *Command) Run() {
	modulePath, err := discovery.GetModulePath()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not find the path of Go module root"))
	}

	c.ExcludeDirs = append(c.ExcludeDirs, archive.DefaultSkipDirs...)
	zipPath, err := archive.Directory(modulePath, true, c.ExcludeDirs, c.ExcludeDirs)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not create archive for module"))
	}

	fmt.Println("Archive path:", zipPath)
	if c.OnlyArchive {
		os.Exit(0)
	} else {
		fmt.Println("cont")
	}
}
