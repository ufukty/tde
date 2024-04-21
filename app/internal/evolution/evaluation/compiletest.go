package evaluation

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"tde/internal/evolution/models"
	"tde/internal/utilities/numerics"
	"tde/pkg/testing"
	"time"
)

type CmdStat int

const (
	FailedAtStart = CmdStat(iota)
	Failed
	Success
)

// err is only populated when status=FailedAtStart. use the stderr for errors thrown by the process
func (e *Evaluator) exec(subject *models.Subject, commandName string, arguments ...string) (duration int64, stdout string, stderr string, status CmdStat, err error) {
	stdoutBuffer := bytes.NewBuffer([]byte{})
	stderrBuffer := bytes.NewBuffer([]byte{})

	cmd := exec.Command(commandName, arguments...)
	cmd.Dir = filepath.Join(e.sm.GetPackagePathForSubject(subject.Sid), "tde")
	cmd.Stdout = stdoutBuffer
	cmd.Stderr = stderrBuffer

	start := time.Now()
	err = cmd.Run()
	duration = time.Now().Sub(start).Nanoseconds()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			err = nil
			status = Failed
		} else {
			err = fmt.Errorf("executing %q: %w", cmd.String(), err)
			status = FailedAtStart
		}
	} else {
		status = Success
	}
	return duration, stdoutBuffer.String(), stderrBuffer.String(), status, err
}

// TODO: recover when process fails
func (e *Evaluator) compile(subject *models.Subject) error {
	dur, _, _, status, err := e.exec(subject, "go", "build", "-o=testmain", "-tags=tde", ".")
	switch status {
	case FailedAtStart:
		return fmt.Errorf("compiling the testmain package: %w", err)

	case Failed:
		inv := 1.0 / float64(max(dur, 1))
		subject.Fitness.Code = inv // TODO: use stderr based fitnesses instead duration

	case Success:
		subject.Fitness.Code = 0.0
	}

	return nil
}

func (e *Evaluator) test(subject *models.Subject) error {
	dur, stdout, stderr, status, err := e.exec(subject, "./testmain", "-subject-uuid", string(subject.Sid))
	switch status {
	case FailedAtStart:
		return fmt.Errorf("running the testmain package: %w", err)

	case Failed: // runtime errors lead to early returns
		fmt.Println(stderr)
		inv := 1.0 / float64(max(dur, 1))
		subject.Fitness.Program = inv // TODO: use stderr based fitnesses instead duration

	case Success: // there might be errors on assertions
		fmt.Println(stdout)

		path := filepath.Join(e.sm.GetPackagePathForSubject(subject.Sid), "tde", "results.json")
		t := &testing.T{}
		if err := t.LoadResults(path); err != nil {
			return fmt.Errorf("loading test results: %w", err)
		}
		subject.Fitness.Program = 0 // it's a "program"
		subject.Fitness.Candidate = numerics.Sum(t.AssertionErrorDistance) / float64(max(len(t.AssertionResults), 1))
	}

	return nil
}
