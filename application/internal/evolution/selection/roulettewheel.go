package selection

import (
	"tde/internal/utilities"
	models "tde/models/program"

	"golang.org/x/exp/maps"
)

// Successful candidates should have higher fitnesses. Otherwise use reverseFitness.
// One candidate can get selected multiple times.
// If there is no candidate with better than worst fitness, the input returned without a difference.
func RouletteWheel(candidates map[models.CandidateID]*models.Candidate, layer models.Layer) []models.CandidateID {
	if len(candidates) == 0 {
		return []models.CandidateID{}
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
		return maps.Keys(candidates)
	}
	for i := 0; i < len(cands); i++ {
		bullet = utilities.URandFloatForCrypto() * totalFitness
		choosen = utilities.BisectRight(cumulative, bullet)
		picks = append(picks, ids[choosen])
	}
	return picks
}
