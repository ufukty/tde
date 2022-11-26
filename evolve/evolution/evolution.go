package evolution

import (
	"sort"

	"github.com/google/uuid"
)

type Program []byte

type Evolution struct {
	HallOfFame map[int]*Candidate
	Candidates map[uuid.UUID]*Candidate
}

func (e *Evolution) InitPopulation(n int) {
	for i := 0; i < n; i++ {
		var candidate = NewCandidate()
		candidate.RandomInit()
		e.Candidates[candidate.ID] = candidate
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

func (e *Evolution) SortedByFitness() []*Candidate {
	ordered := []*Candidate{}
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
