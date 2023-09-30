package selection

import (
	"sort"
	models "tde/models/program"

	"golang.org/x/exp/maps"
)

func filterCandidatesByCids(candidates models.Candidates, cids []models.CandidateID) models.Candidates {
	cands := models.Candidates{}
	for _, pick := range cids {
		cand := candidates[pick]
		cands[cand.UUID] = cand
	}
	return cands
}

func sortByFitnessInLayer(candidates models.Candidates, layer models.Layer) []models.CandidateID {
	sorted := maps.Keys(candidates)
	for _, ind := range candidates {
		sorted = append(sorted, ind.UUID)
	}
	sort.Slice(sorted, func(i, j int) bool {
		return candidates[sorted[i]].Fitness.InLayer(layer) < candidates[sorted[j]].Fitness.InLayer(layer)
	})
	return sorted
}

// O(n2)
func Elitist(candidates models.Candidates, layer models.Layer, pick int) models.Candidates {
	sorted := sortByFitnessInLayer(candidates, layer)
	return filterCandidatesByCids(candidates, sorted[:pick])
}
