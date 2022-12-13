package evolution

import (
	models "tde/models/in_program_models"

	"sort"
)

type Evolution struct {
	HallOfFame map[int]*models.Candidate
	Candidates map[models.CandidateID]*models.Candidate
}

func (e *Evolution) InitPopulation(n int) {
	for i := 0; i < n; i++ {
		var candidate = models.NewCandidate()
		candidate.RandomInit()
		e.Candidates[models.CandidateID(candidate.UUID)] = candidate
	}
}

func (e *Evolution) Select() {
	// for _, individual := range e.Individuals {
	// 	individual.Fitness
	// }
}

func (e *Evolution) Measure() {
	for _, candidate := range e.Candidates {
		candidate.Measure()
	}

	// test

	// count failed assert rate

	// penalty for bloat
}

func (e *Evolution) SortedByFitness() []*models.Candidate {
	ordered := []*models.Candidate{}
	for _, ind := range e.Candidates {
		ordered = append(ordered, ind)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Fitness < ordered[j].Fitness
	})
	return ordered

}

func (e *Evolution) IterateOneGeneration() {

}
