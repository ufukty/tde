package evolution

import (
	"tde/internal/utilities"
	models "tde/models/program"

	"go/format"
)

func GetFitnessArray(individuals []models.Candidate) []float64 {
	fitnesses := []float64{}
	for _, individual := range individuals {
		fitnesses = append(fitnesses, individual.Fitness.Flat())
	}
	return fitnesses
}

func ReversedFitnesses(fitnesses []float64) []float64 {
	var reversedFitnesses = []float64{}
	for _, v := range fitnesses {
		reversedFitnesses = append(reversedFitnesses, 1.0-v)
	}
	return reversedFitnesses
}

func SelectionRouletteWheel(individuals []models.Candidate, selectionSize int) {
	var (
		fitnesses             = GetFitnessArray(individuals)
		reversedFitnesses     = ReversedFitnesses(fitnesses)
		cumulativeFitnesses   = utilities.GetCumulative(reversedFitnesses)
		upperBoundLastFitness = cumulativeFitnesses[len(cumulativeFitnesses)-1]
		choosenIndividuals    = []int{selectionSize}
		choosen               int
		rouletteBullet        float64
	)

	for i := 0; i < selectionSize; i++ {
		rouletteBullet = utilities.URandFloatForCrypto() * upperBoundLastFitness
		choosen = utilities.BisectRight(cumulativeFitnesses, rouletteBullet)
		choosenIndividuals = append(choosenIndividuals, choosen)
	}

}

func CheckSyntax(c *models.Candidate) bool {
	_, err := format.Source(c.File)
	return err == nil
}
