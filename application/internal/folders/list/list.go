package list

import (
	"bytes"
	"encoding/json"
	"io"
	"os/exec"

	"github.com/pkg/errors"
)

func ListAllPackages(path string) (Packages, error) {
	var (
		cmd      *exec.Cmd
		stdbuf   = bytes.NewBuffer([]byte{})
		err      error
		packages = Packages{}
		decoder  *json.Decoder
	)
	cmd = exec.Command("go", "list", "-e", "-json", "./...")
	cmd.Dir = path
	cmd.Stdout = stdbuf
	cmd.Stderr = stdbuf
	err = cmd.Run()
	if err != nil {
		return nil, errors.Wrap(err, stdbuf.String())
	}

	decoder = json.NewDecoder(stdbuf)
	for {
		var pkg = &Package{}
		err = decoder.Decode(pkg)
		if err == io.EOF {
			return packages, nil
		} else if err != nil {
			return nil, errors.Wrap(err, "JSON decoding")
		}
		(packages)[pkg.ImportPath] = pkg
	}
}

func ListPackagesInDir(path string) (Packages, error) {
	var (
		cmd      *exec.Cmd
		stdbuf   = bytes.NewBuffer([]byte{})
		err      error
		packages = Packages{}
		decoder  *json.Decoder
	)
	cmd = exec.Command("go", "list", "-e", "-json", ".")
	cmd.Dir = path
	cmd.Stdout = stdbuf
	cmd.Stderr = stdbuf
	err = cmd.Run()
	if err != nil {
		return nil, errors.Wrap(err, stdbuf.String())
	}

	decoder = json.NewDecoder(stdbuf)
	for {
		var pkg = &Package{}
		err = decoder.Decode(pkg)
		if err == io.EOF {
			return packages, nil
		} else if err != nil {
			return nil, errors.Wrap(err, "JSON decoding")
		}
		(packages)[pkg.ImportPath] = pkg
	}
}
