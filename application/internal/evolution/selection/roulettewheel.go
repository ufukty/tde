package selection

import (
	"fmt"
	"tde/internal/utilities"
	models "tde/models/program"
)

// O(n) if the key checking is O(1)
func deleteDuplicates(sids []models.Sid) []models.Sid {
	found := map[models.Sid]bool{}
	clean := []models.Sid{}
	for _, sid := range sids {
		if _, ok := found[sid]; !ok {
			clean = append(clean, sid)
		}
	}
	return clean
}

// Successful subjects should have higher fitnesses. Otherwise use reverseFitness.

// FIXME: If there is no subject with better than worst fitness, the input returned without a difference.
func RouletteWheelToEliminate(subjects models.Subjects, layer models.Layer, pick int) (models.Subjects, error) {
	if len(subjects) == 0 || len(subjects) < pick {
		return models.Subjects{}, nil
	}
	var (
		ids, cands   = utilities.MapItems(subjects)
		fitnesses    = reverse(getFitnesses(cands, layer))
		cumulative   = utilities.GetCumulative(fitnesses)
		totalFitness = cumulative[len(cumulative)-1]
		picks        = []models.Sid{}
		choosen      int
		bullet       float64
	)
	if totalFitness == 0.0 {
		return Random(subjects, pick), nil
	}
	for len(picks) < pick {
		for len(picks) < pick { // O(n*logn)
			bullet = utilities.URandFloatForCrypto() * totalFitness
			choosen = utilities.BisectRight(cumulative, bullet)
			picks = append(picks, ids[choosen])
		}
		picks = deleteDuplicates(picks)
	}
	return filterSubjectsByCids(subjects, picks), nil
}

// allows duplicate selections
func RouletteWheelToReproduce(subjects models.Subjects, layer models.Layer, pick int) (models.Subjects, error) {
	if len(subjects) == 0 {
		if pick == 0 {
			return subjects, nil
		} else {
			return subjects, fmt.Errorf("empty population")
		}
	}
	var (
		ids, cands   = utilities.MapItems(subjects)
		fitnesses    = reverse(getFitnesses(cands, layer))
		cumulative   = utilities.GetCumulative(fitnesses)
		totalFitness = cumulative[len(cumulative)-1]
		picks        = []models.Sid{}
		choosen      int
		bullet       float64
	)
	if totalFitness == 0.0 {
		return Random(subjects, pick), nil
	}
	for len(picks) < pick {
		bullet = utilities.URandFloatForCrypto() * totalFitness
		choosen = utilities.BisectRight(cumulative, bullet)
		picks = append(picks, ids[choosen])
	}
	return filterSubjectsByCids(subjects, picks), nil
}
