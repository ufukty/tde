package selection

import (
	"tde/internal/utilities"
	models "tde/models/program"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func getFitnesses(candidates []*models.Candidate, layer models.Layer) []float64 {
	fitnesses := []float64{}
	for _, individual := range candidates {
		fitnesses = append(fitnesses, individual.Fitness.InLayer(layer))
	}
	return fitnesses
}

// reduces all items of the array into [0, 1] range.
// or [1, 0] if reverse is true
func normalizeFitnesses(fitnesses []float64, reverse bool) []float64 {
	if len(fitnesses) == 0 {
		return fitnesses
	}
	var (
		normalized = []float64{}
		maxFitness = slices.Max(fitnesses)
		minFitness = slices.Min(fitnesses)
		deltaF     = maxFitness - minFitness
	)
	if deltaF == 0 {
		return fitnesses
	}
	for _, f := range fitnesses {
		lerped := (f - minFitness) / deltaF
		if reverse {
			lerped = 1.0 - lerped
		}
		normalized = append(normalized, lerped)
	}
	return normalized
}

// Chooses random candidates by chances that are relative to their fitnesses
// Candidates with higher numerical value for their fitnesses have higher chance to get selected by default.
// When reverseFitness is True, behaviour gets reversed.
// One candidate can get selected multiple times.
func RouletteWheel(candidates map[models.CandidateID]*models.Candidate, layer models.Layer, reverseFitness bool) []models.CandidateID {
	if len(candidates) == 0 {
		return []models.CandidateID{}
	}
	if len(candidates) == 1 {
		return maps.Keys(candidates)
	}
	var (
		ids, cands   = utilities.MapItems(candidates)
		fitnesses    = normalizeFitnesses(getFitnesses(cands, layer), reverseFitness)
		cumulative   = utilities.GetCumulative(fitnesses)
		totalFitness = cumulative[len(cumulative)-1]
		picks        = []models.CandidateID{}
		choosen      int
		bullet       float64
	)
	for i := 0; i < len(cands); i++ {
		bullet = utilities.URandFloatForCrypto() * totalFitness
		choosen = utilities.BisectRight(cumulative, bullet)
		picks = append(picks, ids[choosen])
	}
	return picks
}
