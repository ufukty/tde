package evolution

import (
	"context"
	"tde/internal/evaluation"
	"tde/internal/folders/slot_manager"
	models "tde/models/in_program_models"

	"sort"

	"golang.org/x/exp/maps"
)

type Evolution struct {
	Evaluation *evaluation.Evaluator
	HallOfFame map[int]*models.Candidate
	Candidates map[models.CandidateID]*models.Candidate
}

func NewEvolution(slotManagerSession *slot_manager.Session) *Evolution {
	return &Evolution{
		Evaluation: evaluation.NewEvaluator(slotManagerSession),
		HallOfFame: map[int]*models.Candidate{},
		Candidates: map[models.CandidateID]*models.Candidate{},
	}
}

func (e *Evolution) InitPopulation(n int) {
	for i := 0; i < n; i++ {
		var candidate = models.NewCandidate()
		// init candidate body randomly
		e.Candidates[models.CandidateID(candidate.UUID)] = candidate
	}
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
