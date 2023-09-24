package selection

import (
	"tde/internal/utilities"
	models "tde/models/program"

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
	var (
		reversedFitnesses = []float64{}
		maxFitness        = slices.Max(fitnesses)
		minFitness        = slices.Min(fitnesses)
		deltaF            = maxFitness - minFitness
	)
	for _, f := range fitnesses {
		lerped := (f - minFitness) / deltaF
		if reverse {
			lerped = 1.0 - lerped
		}
		reversedFitnesses = append(reversedFitnesses, lerped)
	}
	return reversedFitnesses
}

// Chooses random candidates by chances that are relative to their fitnesses
// Candidates with higher numerical value for their fitnesses have higher chance to get selected by default.
// When reverse is True, behaviour gets reversed.
func RouletteWheel(candidates map[models.CandidateID]*models.Candidate, layer models.Layer, bullets int, reverse bool) []models.CandidateID {
	var (
		ids, cands            = utilities.MapItems(candidates)
		fitnesses             = normalizeFitnesses(getFitnesses(cands, layer), reverse)
		cumulativeFitnesses   = utilities.GetCumulative(fitnesses)
		upperBoundLastFitness = cumulativeFitnesses[len(cumulativeFitnesses)-1]
		choosenIndividuals    = []models.CandidateID{}
		choosen               int
		bullet                float64
		survivors             = len(fitnesses) - bullets
	)
	for i := 0; i < survivors; i++ {
		bullet = utilities.URandFloatForCrypto() * upperBoundLastFitness
		choosen = utilities.BisectRight(cumulativeFitnesses, bullet)
		choosenIndividuals = append(choosenIndividuals, ids[choosen])
	}
	return choosenIndividuals
}
