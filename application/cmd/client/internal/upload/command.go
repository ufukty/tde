package upload

import (
	module "tde/cmd/customs/endpoints/module/post"
	"tde/internal/command"
	"tde/internal/folders/archive"
	"tde/internal/folders/discovery"

	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

type Command struct {
	OnlyArchive bool                `long:"only-archive"`
	ExcludeDirs command.MultiString `short:"e" long:"exclude-dir"`
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

	c.ExcludeDirs = append(c.ExcludeDirs, archive.DefaultSkipDirs...)
	zipPath, err := archive.Directory(modulePath, true, c.ExcludeDirs, c.ExcludeDirs)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not create archive for module"))
	}

	fmt.Println("Archive path:", zipPath)
	if c.OnlyArchive {
		os.Exit(0)
	}

	fileHandler, err = os.Open(zipPath)
	if err != nil {
		panic(errors.Wrap(err, "Could not open file to upload"))
	}
	defer fileHandler.Close()

	req, err = module.NewRequest(fileHandler)
	if err != nil {
		panic(errors.Wrap(err, "Could not create request"))
	}

	resp, err = req.Send("POST", "http://127.0.0.1:8087/api/v1.0.0/customs/module")
	if err != nil {
		panic(errors.Wrap(err, "Failed"))
	}

	fmt.Println("Archive ID:", resp.ArchiveID)
}
