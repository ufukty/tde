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
func syntaxCheckAndProduceCode(candidates []*models.Candidate) {
	for _, candidate := range candidates {
		buffer, ok, err := ProduceCodeFromASTSafe(candidate.AST.File) // produce code from ast.File to capture changes in import list too
		if err != nil || !ok {
			candidate.Fitness.AST = 1.0
			continue
		}
		candidate.File = buffer.Bytes()
	}
}

// FIXME: recover when process fails
func (e *Evaluator) run(cid models.CandidateID) error {
	cmd := exec.Command("go", "run", "-tags=tde", ".", "-candidate-uuid", string(cid))
	cmd.Dir = filepath.Join(e.sm.GetPackagePathForCandidate(cid), "tde")
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command %q in dir: %q returned %q: %w", cmd.String(), cmd.Dir, string(bytes), err)
	}
	return nil
}

func (e *Evaluator) test(cid models.CandidateID) error {
	if err := e.run(cid); err != nil {
		return fmt.Errorf("running the injected package in candidate: %w", err)
	}
	return nil
}

func (e *Evaluator) collectResult(cid models.CandidateID) (*testing.T, error) {
	path := filepath.Join(e.sm.GetPackagePathForCandidate(cid), "tde", "results.json")
	results := &testing.T{}
	if err := results.LoadResults(path); err != nil {
		return nil, fmt.Errorf("parsing: %w", err)
	}
	return results, nil
}

// TODO:
func (e *Evaluator) populateFitnessWithResults(candidate *models.Candidate, results *testing.T) error {
	// candidate.Fitness.Program
	// results.AssertionResults
	return nil
}

func (e *Evaluator) runCandidates(candidates []*models.Candidate) error {
	for _, candidate := range candidates {
		if err := e.test(candidate.UUID); err != nil {
			return fmt.Errorf("testing the candidate %q: %w", candidate.UUID, err)
		}
		results, err := e.collectResult(candidate.UUID)
		if err != nil {
			return fmt.Errorf("collecting test results for the candidate %q: %w", candidate.UUID, err)
		}
		if err := e.populateFitnessWithResults(candidate, results); err != nil {
			return fmt.Errorf("populating fitness score with results for th candidate %q: %w", candidate.UUID, err)
		}
	}
	return nil
}

// TODO: Syntax Check
// TODO: Print code (full package vs changed func body ??)
// TODO: Send whole generation into sandboxed environment
// TODO: Get test results
// TODO: Return test results
func (e *Evaluator) Pipeline(candidates []*models.Candidate) error {
	syntaxCheckAndProduceCode(candidates)
	if err := e.sm.PlaceCandidatesIntoSlots(candidates); err != nil {
		return fmt.Errorf("placing candidates into slots: %w", err)
	}
	if err := e.runCandidates(candidates); err != nil {
		return fmt.Errorf("running candidates: %w", err)
	}
	if err := e.sm.FreeAllSlots(); err != nil {
		return fmt.Errorf("restoring slots for next generation: %w", err)
	}
	return nil
}
