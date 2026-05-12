package osw

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

var (
	MoreThanOneLineFound = fmt.Errorf("given command output has more than one line")
	NoLinesFound         = fmt.Errorf("given command output has no output that is terminated with '\\n' character")
)

func RunCommandForOutput(commandName string, arguments ...string) (stdout string, stderr string, err error) {
	cmd := exec.Command(commandName, arguments...)
	var (
		outputStream bytes.Buffer
		errStream    bytes.Buffer
	)
	cmd.Stdout = &outputStream
	cmd.Stderr = &errStream
	err_ := cmd.Run()
	if err_ != nil {
		return outputStream.String(), errStream.String(),
			fmt.Errorf("exec.Command is failed for command %s: %w", commandName, err_)
	}
	return outputStream.String(), errStream.String(), nil
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
	dir, _, err := RunCommandForOutput("pwd", "-P")
	if err != nil {
		return "", fmt.Errorf("failed to run 'pwd': %w", err)
	}
	dir, err = StripOnlyLineFromCommandOuput(dir)
	if err != nil {
		return "", fmt.Errorf("failed to get current dir from output of 'pwd': %w", err)
	}
	return dir, nil
}
