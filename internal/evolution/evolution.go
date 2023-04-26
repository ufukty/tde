package evolution

import (
	"context"
	"go/ast"
	models "tde/models/program"

	"sort"

	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
)

type Evaluator interface {
	Pipeline(candidates []*models.Candidate)
}

type EvolutionManager struct {
	Evaluation Evaluator
	Target     *models.EvolutionTarget
	HallOfFame map[int]*models.Candidate
	Candidates map[models.CandidateID]*models.Candidate
}

func NewEvolutionManager(evaluator Evaluator, target *models.EvolutionTarget) *EvolutionManager {
	return &EvolutionManager{
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

func (e *EvolutionManager) InitPopulation(n int) error {
	for i := 0; i < n; i++ {
		var candidate, err = models.NewCandidate(e.Target.Package, e.Target.File, e.Target.FuncDecl)
		if err != nil {
			return errors.Wrap(err, "failed to create new Candidate instance")
		}
		e.Candidates[models.CandidateID(candidate.UUID)] = candidate
	}
	return nil
}

func (e *EvolutionManager) Select() {
	// for _, individual := range e.Individuals {
	// 	individual.Fitness
	// }
}

func (e *EvolutionManager) SortedByFitness() []*models.Candidate {
	ordered := []*models.Candidate{}
	for _, ind := range e.Candidates {
		ordered = append(ordered, ind)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Fitness.Flat() < ordered[j].Fitness.Flat()
	})
	return ordered
}

func (e *EvolutionManager) IterateLoop(ctx context.Context) {
	e.Evaluation.Pipeline(maps.Values(e.Candidates))

	// TODO: selection

	// TODO: reproduction

}
