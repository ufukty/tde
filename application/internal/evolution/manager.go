package evolution

import (
	"tde/internal/evolution/evaluation"
	models "tde/models/program"

	"sort"

	"github.com/pkg/errors"
)

type Manager struct {
	Evaluator  *evaluation.Evaluator
	Target     *Target
	HallOfFame map[int]*models.Candidate
	Candidates map[models.CandidateID]*models.Candidate
}

func NewManager(target *Target) *Manager {
	return &Manager{
		Target:     target,
		HallOfFame: map[int]*models.Candidate{},
		Candidates: map[models.CandidateID]*models.Candidate{},
	}
}

func (e *Manager) Init() error {
	return nil
}

func (e *Manager) Select() {
	// for _, individual := range e.Individuals {
	// 	individual.Fitness
	// }
}

func (e *Manager) SortedByFitness() []*models.Candidate {
	ordered := []*models.Candidate{}
	for _, ind := range e.Candidates {
		ordered = append(ordered, ind)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Fitness.Flat() < ordered[j].Fitness.Flat()
	})
	return ordered
}

// This won't perform evaluation and will expect the fitnesses are already set
func (e *Manager) IterateLoop() {
	// TODO: selection
	// TODO: reproduction
}
