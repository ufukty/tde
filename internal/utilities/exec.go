package utilities

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

var (
	MoreThanOneLineFound = errors.New("given command output has more than one line")
	NoLinesFound         = errors.New("given command output has no output that is terminated with '\\n' character")
)

func RunCommandForOutput(commandName string, arguments ...string) (string, error) {
	cmd := exec.Command(commandName, arguments...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", errors.Wrapf(err, "exec.Command is failed for command %s", commandName)
	}
	return out.String(), nil
}

func StripOnlyLineFromCommandOuput(output string) (string, error) {
	lines := strings.Split(output, "\n")
	if len(lines) < 2 {
		return "", NoLinesFound
	} else if len(lines) > 2 {
		return "", MoreThanOneLineFound
	}
	lastLine := lines[0]

	if strings.LastIndex(lastLine, "\r\n") != -1 {
		fmt.Println("========1")
		return strings.TrimSuffix(lastLine, "\r\n"), nil
	} else if strings.LastIndex(lastLine, "\n") != -1 {
		fmt.Println("========2")
		return strings.TrimSuffix(lastLine, "\n"), nil
	} else {
		return lastLine, nil
	}
}

func CurrentDir() (string, error) {
	dir, err := RunCommandForOutput("pwd", "-P")
	if err != nil {
		return "", errors.Wrap(err, "failed to run 'pwd'")
	}
	dir, err = StripOnlyLineFromCommandOuput(dir)
	if err != nil {
		return "", errors.Wrap(err, "failed to get current dir from output of 'pwd'")
	}
	return dir, nil
}
