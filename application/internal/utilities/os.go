package utilities

import (
	"github.com/pkg/errors"
)

func WorkingDir() (string, error) {
	stdOut, _, err := RunCommandForOutput("pwd", "-P")
	if err != nil {
		return "", errors.Wrap(err, "Failed to run command pwd")
	}
	return stdOut, nil
}
