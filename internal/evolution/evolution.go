package evolution

import (
	"context"
	"go/ast"
	"tde/internal/evaluation"
	models "tde/models/program"

	"sort"

	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
)

type Evolution struct {
	Evaluation *evaluation.Evaluator
	Target     *models.EvolutionTarget
	HallOfFame map[int]*models.Candidate
	Candidates map[models.CandidateID]*models.Candidate
}

func NewEvolution(evaluator *evaluation.Evaluator, target *models.EvolutionTarget) *Evolution {
	return &Evolution{
		Evaluation: evaluator,
		Target:     target,
		HallOfFame: map[int]*models.Candidate{},
		Candidates: map[models.CandidateID]*models.Candidate{},
	}
}

type EvolutionTarget struct {
	Package  *ast.Package
	File     *ast.File
	FuncDecl *ast.FuncDecl
}

func (e *Evolution) InitPopulation(n int) error {
	for i := 0; i < n; i++ {
		var candidate, err = models.NewCandidate(e.Target.Package, e.Target.File, e.Target.FuncDecl)
		if err != nil {
			return errors.Wrap(err, "failed to create new Candidate instance")
		}
		e.Candidates[models.CandidateID(candidate.UUID)] = candidate
	}
	return nil
}

func (e *Evolution) Select() {
	// for _, individual := range e.Individuals {
	// 	individual.Fitness
	// }
}

func (e *Evolution) SortedByFitness() []*models.Candidate {
	ordered := []*models.Candidate{}
	for _, ind := range e.Candidates {
		ordered = append(ordered, ind)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Fitness.Flat() < ordered[j].Fitness.Flat()
	})
	return ordered
}

func (e *Evolution) IterateLoop(ctx context.Context) {
	e.Evaluation.Pipeline(maps.Values(e.Candidates))

	// TODO: selection

	// TODO: reproduction

}
