package evolution

import (
	"tde/internal/embedding"
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
		// init candidate body randomly
		e.Candidates[models.CandidateID(candidate.UUID)] = candidate
	}
}

func (e *Evolution) Measure(embed embedding.EmbeddingConfig) {
	validCandidates := []*models.Candidate{}

	for _, candidate := range e.Candidates {
		if CheckSyntax(candidate) {
			validCandidates = append(validCandidates, candidate)
		}
	}

	embed.WriteCandidatesIntoFile(validCandidates)

	// test

	// count failed assert rate

	// penalty for bloat
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
		return ordered[i].Fitness < ordered[j].Fitness
	})
	return ordered

}

func (e *Evolution) IterateOneGeneration() {

}
