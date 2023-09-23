package selection

import (
	"tde/internal/utilities"
	models "tde/models/program"

	"golang.org/x/exp/slices"
)

func getFitnessArray(individuals []*models.Candidate) []float64 {
	fitnesses := []float64{}
	for _, individual := range individuals {
		fitnesses = append(fitnesses, individual.Fitness.Flat())
	}
	return fitnesses
}

func reversedFitnesses(fitnesses []float64) []float64 {
	var reversedFitnesses = []float64{}
	var maxFitness = slices.Max(fitnesses)
	for _, v := range fitnesses {
		reversedFitnesses = append(reversedFitnesses, maxFitness-v)
	}
	return reversedFitnesses
}

func RouletteWheel(candidates []*models.Candidate, bullets int) []*models.Candidate {
	var (
		fitnesses             = getFitnessArray(candidates)
		cumulativeFitnesses   = utilities.GetCumulative(fitnesses)
		upperBoundLastFitness = cumulativeFitnesses[len(cumulativeFitnesses)-1]
		choosenIndividuals    = []*models.Candidate{}
		choosen               int
		rouletteBullet        float64
		survivors             = len(fitnesses) - bullets
	)
	for i := 0; i < survivors; i++ {
		rouletteBullet = utilities.URandFloatForCrypto() * upperBoundLastFitness
		choosen = utilities.BisectRight(cumulativeFitnesses, rouletteBullet)
		choosenIndividuals = append(choosenIndividuals, candidates[choosen])
	}
	return choosenIndividuals
}

func RouletteWheelReverse(candidates []*models.Candidate, survivors int) []*models.Candidate {
	var (
		fitnesses             = getFitnessArray(candidates)
		reversedFitnesses     = reversedFitnesses(fitnesses)
		cumulativeFitnesses   = utilities.GetCumulative(reversedFitnesses)
		upperBoundLastFitness = cumulativeFitnesses[len(cumulativeFitnesses)-1]
		choosenIndividuals    = []*models.Candidate{}
		choosen               int
		rouletteBullet        float64
	)
	for i := 0; i < survivors; i++ {
		rouletteBullet = utilities.URandFloatForCrypto() * upperBoundLastFitness
		choosen = utilities.BisectRight(cumulativeFitnesses, rouletteBullet)
		choosenIndividuals = append(choosenIndividuals, candidates[choosen])
	}
	return choosenIndividuals
}
