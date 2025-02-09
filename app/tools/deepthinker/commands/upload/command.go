package upload

import (
	"flag"
	"fmt"
	"log"
	"os"
	"tde/cmd/customs/endpoints"
	"tde/internal/evolution/evaluation/archive"
	"tde/internal/evolution/evaluation/discovery"
	"tde/internal/utilities/valuable"
)

type Args struct {
	OnlyArchive bool
	Output      string
	ExcludeDirs valuable.Strings
	IncludeExts valuable.Strings
	Verbose     bool
}

func Run() error {
	args := &Args{}
	flag.BoolVar(&args.OnlyArchive, "only-archive", false, "")
	flag.StringVar(&args.Output, "output", "", "")
	flag.Var(&args.ExcludeDirs, "exclude-dir", "")
	flag.Var(&args.IncludeExts, "include-ext", "")
	flag.BoolVar(&args.Verbose, "v", false, "")
	flag.Parse()

	modulePath, err := discovery.ModuleRoot()
	if err != nil {
		return fmt.Errorf("Could not find the path of Go module root: %w", err)
	}

	args.IncludeExts = append(args.IncludeExts, archive.DefaultInclExt...)
	args.ExcludeDirs = append(args.ExcludeDirs, archive.DefaultSkipDirs...)

	zipPath := ""
	if args.OnlyArchive && args.Output != "" {
		err = archive.DirectoryToFile(args.Output, modulePath, true, args.ExcludeDirs, args.ExcludeDirs, args.IncludeExts, args.Verbose)
	} else {
		zipPath, err = archive.Directory(modulePath, true, args.ExcludeDirs, args.ExcludeDirs, args.IncludeExts, args.Verbose)
	}
	if err != nil {
		return fmt.Errorf("Could not create archive for module: %w", err)
	}

	if args.Verbose || (args.OnlyArchive && args.Output == "") {
		fmt.Println("Archived into:", zipPath)
	}
	if args.OnlyArchive {
		fmt.Println("Done.")
		os.Exit(0)
	}

	fh, err := os.Open(zipPath)
	if err != nil {
		return fmt.Errorf("Could not open file to upload: %w", err)
	}
	defer fh.Close()

	up, err := endpoints.NewUploadRequest(fh)
	if err != nil {
		return fmt.Errorf("Could not create request: %w", err)
	}

	if args.Verbose {
		log.Println("Uploading...")
	}
	resp, err := up.Send()
	if err != nil {
		return fmt.Errorf("Failed: %w", err)
	}

	fmt.Println("Done. Archive ID:", resp.ArchiveID)
	return nil
}
