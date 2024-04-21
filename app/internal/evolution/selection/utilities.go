package selection

import (
	"sort"
	"tde/internal/evolution/models"

	"golang.org/x/exp/maps"
)

func getFitnesses(subjects []*models.Subject, layer models.Layer) []float64 {
	fitnesses := []float64{}
	for _, individual := range subjects {
		fitnesses = append(fitnesses, individual.Fitness.InLayer(layer))
	}
	return fitnesses
}

// reverses fitness from [0,1] to [1,0] range
func reverse(fitnesses []float64) []float64 {
	normalized := []float64{}
	for _, f := range fitnesses {
		normalized = append(normalized, 1.0-f)
	}
	return normalized
}

func filterSubjectsBySids(subjects models.Subjects, sids []models.Sid) models.Subjects {
	cands := models.Subjects{}
	for _, pick := range sids {
		cand := subjects[pick]
		cands[cand.Sid] = cand
	}
	return cands
}

func sortByFitnessInLayer(subjects models.Subjects, layer models.Layer) []models.Sid {
	sorted := maps.Keys(subjects)
	for _, ind := range subjects {
		sorted = append(sorted, ind.Sid)
	}
	sort.Slice(sorted, func(i, j int) bool {
		return subjects[sorted[i]].Fitness.InLayer(layer) < subjects[sorted[j]].Fitness.InLayer(layer)
	})
	return sorted
}
