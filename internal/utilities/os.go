package utilities

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func Terminate(msg ...any) {
	fmt.Println(msg...)
	os.Exit(1)
}

func WorkingDir() (string, error) {
	stdOut, _, err := RunCommandForOutput("pwd", "-P")
	if err != nil {
		return "", errors.Wrap(err, "Failed to run command pwd")
	}
	return stdOut, nil
}
