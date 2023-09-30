package selection

import (
	"fmt"
	"tde/internal/utilities"
	models "tde/models/program"
)

// O(n) if the key checking is O(1)
func deleteDuplicates(cids []models.CandidateID) []models.CandidateID {
	found := map[models.CandidateID]bool{}
	clean := []models.CandidateID{}
	for _, cid := range cids {
		if _, ok := found[cid]; !ok {
			clean = append(clean, cid)
		}
	}
	return clean
}

// Successful candidates should have higher fitnesses. Otherwise use reverseFitness.
// One candidate can get selected multiple times.
// If there is no candidate with better than worst fitness, the input returned without a difference.
func RouletteWheel(candidates models.Candidates, layer models.Layer, pick int) (models.Candidates, error) {
	if len(candidates) == 0 {
		return candidates, fmt.Errorf("empty population")
	}
	if len(candidates) < pick {
		return candidates, fmt.Errorf("there are less candidates then needs to get pick")
	}
	var (
		ids, cands   = utilities.MapItems(candidates)
		fitnesses    = reverse(getFitnesses(cands, layer))
		cumulative   = utilities.GetCumulative(fitnesses)
		totalFitness = cumulative[len(cumulative)-1]
		picks        = []models.CandidateID{}
		choosen      int
		bullet       float64
	)
	if totalFitness == 0.0 {
		return candidates, fmt.Errorf("all candidates are worst so it's not possible to select better ones.")
	}
	for len(picks) < pick {
		for len(picks) < pick { // O(n*logn)
			bullet = utilities.URandFloatForCrypto() * totalFitness
			choosen = utilities.BisectRight(cumulative, bullet)
			picks = append(picks, ids[choosen])
		}
		picks = deleteDuplicates(picks)
	}
	return filterCandidatesByCids(candidates, picks), nil
}
