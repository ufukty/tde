package selection

import (
	"tde/internal/evolution/models"
	"tde/internal/utilities"
)

// O(n2)
func Elitist(subjects models.Subjects, layer models.Layer, pick int) models.Subjects {
	sorted := sortByFitnessInLayer(subjects, layer)
	return filterSubjectsBySids(subjects, sorted[:pick])
}

func Random(subjects models.Subjects, pick int) models.Subjects {
	picks := models.Subjects{}
	sids := subjects.Keys()
	for len(picks) < pick {
		picks.Add(subjects[sids[utilities.URandIntN(len(sids))]])
	}
	return picks
}
