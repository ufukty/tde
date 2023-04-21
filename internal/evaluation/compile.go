package evaluation

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"tde/internal/folders/slot_manager"
	models "tde/models/program"

	"github.com/pkg/errors"
)

type Runner struct {
	slotManagerSession *slot_manager.Session
}

func NewRunner(slotManagerSession *slot_manager.Session) *Runner {
	return &Runner{
		slotManagerSession: slotManagerSession,
	}
}

func (r *Runner) compile(pkgPath string, candidateID models.CandidateID) (string, error) {
	cmd := exec.Command("go", "run", "-tags=tde", ".", "-candidate-uuid", string(candidateID))
	cmd.Dir = pkgPath
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "", errors.Wrap(err, "cmd.Output() returned error for command \"go run .\"")
	}
	return string(bytes), nil
}

func (r *Runner) run(candidate *models.Candidate) error {
	pkgPath := r.slotManagerSession.GetPackagePathForCandidate(candidate.UUID)
	output, err := r.compile(filepath.Join(pkgPath, "tde"), candidate.UUID)
	if err != nil {
		return errors.Wrapf(err, "failed to run testing package for candidate=\"%s\"", candidate.UUID)
	}
	fmt.Println(output)
	return nil
}

func (r *Runner) Pipeline(candidates []*models.Candidate) {
	for _, candidate := range candidates {
		r.run(candidate)
	}
}
