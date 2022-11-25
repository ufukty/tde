package evolution

import (
	"sort"

	"github.com/google/uuid"
)

type Evolution struct {
	Individuals map[uuid.UUID]*Individual
}

func (e *Evolution) InitPopulation(n int) {
	for i := 0; i < n; i++ {
		var individual = NewIndividual()
		individual.RandomInit()
		e.Individuals[individual.ID] = individual
	}
}

func (e *Evolution) Select() {
	// for _, individual := range e.Individuals {
	// 	individual.Fitness
	// }
}

func (e *Evolution) Measure() {
	for _, individual := range e.Individuals {
		individual.Measure()
	}

	// test

	// count failed assert rate

	// penalty for bloat
}

func (e *Evolution) SortedByFitness() []*Individual {
	ordered := []*Individual{}
	for _, ind := range e.Individuals {
		ordered = append(ordered, ind)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Fitness < ordered[j].Fitness
	})
	return ordered

}

func (e *Evolution) IterateOneGeneration() {

}
