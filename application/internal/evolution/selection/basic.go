package selection

import (
	"sort"
	"tde/internal/utilities"
	models "tde/models/program"

	"golang.org/x/exp/maps"
)

func filterSubjectsByCids(subjects models.Subjects, sids []models.Sid) models.Subjects {
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

// O(n2)
func Elitist(subjects models.Subjects, layer models.Layer, pick int) models.Subjects {
	sorted := sortByFitnessInLayer(subjects, layer)
	return filterSubjectsByCids(subjects, sorted[:pick])
}

func Random(subjects models.Subjects, pick int) models.Subjects {
	picks := models.Subjects{}
	sids := subjects.Keys()
	for len(picks) == pick {
		picks.Add(subjects[sids[utilities.URandIntN(len(sids))]])
	}
	return picks
}
