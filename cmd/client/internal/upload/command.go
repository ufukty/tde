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

var DefaultExcludeDirs = []string{".git", "build", "docs", ".vscode"}

type Command struct {
	OnlyArchive bool                `long:"only-archive"`
	ExcludeDirs command.MultiString `short:"e" long:"exclude-dir"`
}

func (c *Command) Run() {
	c.ExcludeDirs = append(c.ExcludeDirs, DefaultExcludeDirs...)

	modulePath, err := discovery.FindModulePath()
	if err != nil {
		utl.Terminate(errors.Wrap(err, "Could not find the path of Go module root"))
	}

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
