package evaluation

import (
	"tde/internal/evolution/evaluation/slotmgr"
	models "tde/models/program"
	"tde/pkg/testing"

	"fmt"
	"os/exec"
	"path/filepath"
)

type Evaluator struct {
	sm *slotmgr.SlotManager
}

func NewEvaluator(sm *slotmgr.SlotManager) *Evaluator {
	return &Evaluator{
		sm: sm,
	}
}

// FIXME: count errors on code creation
// TODO: ...and populate fitness component for it
func syntaxCheckAndProduceCode(subjects []*models.Subject) {
	for _, subject := range subjects {
		buffer, ok, err := ProduceCodeFromASTSafe(subject.AST.File) // produce code from ast.File to capture changes in import list too
		if err != nil || !ok {
			subject.Fitness.AST = 1.0
			continue
		}
		subject.File = buffer.Bytes()
	}
}

// FIXME: recover when process fails
func (e *Evaluator) run(sid models.Sid) error {
	cmd := exec.Command("go", "run", "-tags=tde", ".", "-subject-uuid", string(sid))
	cmd.Dir = filepath.Join(e.sm.GetPackagePathForCandidate(sid), "tde")
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command %q in dir: %q returned %q: %w", cmd.String(), cmd.Dir, string(bytes), err)
	}
	return nil
}

func (e *Evaluator) test(sid models.Sid) error {
	if err := e.run(sid); err != nil {
		return fmt.Errorf("running the injected package in subject: %w", err)
	}
	return nil
}

func (e *Evaluator) collectResult(sid models.Sid) (*testing.T, error) {
	path := filepath.Join(e.sm.GetPackagePathForCandidate(sid), "tde", "results.json")
	results := &testing.T{}
	if err := results.LoadResults(path); err != nil {
		return nil, fmt.Errorf("parsing: %w", err)
	}
	return results, nil
}

// TODO:
func (e *Evaluator) populateFitnessWithResults(subject *models.Subject, results *testing.T) error {
	// subject.Fitness.Program
	// results.AssertionResults
	return nil
}

func (e *Evaluator) runSubjects(subjects []*models.Subject) error {
	for _, subject := range subjects {
		if err := e.test(subject.Sid); err != nil {
			return fmt.Errorf("testing the subject %q: %w", subject.Sid, err)
		}
		results, err := e.collectResult(subject.Sid)
		if err != nil {
			return fmt.Errorf("collecting test results for the subject %q: %w", subject.Sid, err)
		}
		if err := e.populateFitnessWithResults(subject, results); err != nil {
			return fmt.Errorf("populating fitness score with results for th subject %q: %w", subject.Sid, err)
		}
	}
	return nil
}

// TODO: Syntax Check
// TODO: Print code (full package vs changed func body ??)
// TODO: Send whole generation into sandboxed environment
// TODO: Get test results
// TODO: Return test results
// TODO: skip subjects its fitness already set
func (e *Evaluator) Pipeline(subjects []*models.Subject) error {
	syntaxCheckAndProduceCode(subjects)
	if err := e.sm.PlaceSubjectsIntoSlots(subjects); err != nil {
		return fmt.Errorf("placing subjects into slots: %w", err)
	}
	if err := e.runSubjects(subjects); err != nil {
		return fmt.Errorf("running subjects: %w", err)
	}
	if err := e.sm.FreeAllSlots(); err != nil {
		return fmt.Errorf("restoring slots for next generation: %w", err)
	}
	return nil
}
