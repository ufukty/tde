package upload

import (
	"fmt"
	module "tde/cmd/customs/endpoints/module/post"
	"tde/internal/command"
	"tde/internal/folders/archive"
	"tde/internal/folders/discovery"

	"log"
	"os"

	"github.com/pkg/errors"
)

type Command struct {
	OnlyArchive bool                `long:"only-archive"`
	ExcludeDirs command.MultiString `short:"e" long:"exclude-dir"`
	IncludeExts command.MultiString `short:"i" long:"include-ext"`
	Verbose     bool                `short:"v"`
}

func (c *Command) Run() {
	var (
		err         error
		modulePath  string
		req         *module.Request
		resp        *module.Response
		fileHandler *os.File
	)
	modulePath, err = discovery.GetModulePath()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not find the path of Go module root"))
	}

	c.IncludeExts = append(c.IncludeExts, archive.DefaultInclExt...)
	c.ExcludeDirs = append(c.ExcludeDirs, archive.DefaultSkipDirs...)
	zipPath, err := archive.Directory(modulePath, true, c.ExcludeDirs, c.ExcludeDirs, c.IncludeExts, c.Verbose)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not create archive for module"))
	}

	if c.Verbose || c.OnlyArchive {
		fmt.Println("Archived into:", zipPath)
	}
	if c.OnlyArchive {
		os.Exit(0)
	}

	fileHandler, err = os.Open(zipPath)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not open file to upload"))
	}
	defer fileHandler.Close()

	req, err = module.NewRequest(fileHandler)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not create request"))
	}

	if c.Verbose {
		log.Println("Uploading...")
	}
	resp, err = req.Send("POST", "http://127.0.0.1:8087/api/v1.0.0/customs/module")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Failed"))
	}

	fmt.Println("Archive ID:", resp.ArchiveID)
}
