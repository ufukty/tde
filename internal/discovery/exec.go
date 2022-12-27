package discovery

import (
	"bytes"
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
	return lines[0], nil
}
