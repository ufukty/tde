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
// One subject can get selected multiple times.
// If there is no subject with better than worst fitness, the input returned without a difference.
func RouletteWheel(subjects models.Subjects, layer models.Layer, pick int) (models.Subjects, error) {
	if len(subjects) == 0 {
		return subjects, fmt.Errorf("empty population")
	}
	if len(subjects) < pick {
		return subjects, fmt.Errorf("there are less subjects then needs to get pick")
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
		return subjects, fmt.Errorf("all subjects are worst so it's not possible to select better ones.")
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
