package evolution

import (
	"tde/internal/utilities"
	models "tde/models/in_program_models"

	"math/rand"
)

type Number64 interface {
	~float64 | ~int64
}

func CumulativeArray[N Number64](input []N) []N {
	var output = []N{}
	var total = N(0)
	for _, v := range input {
		total += v
		output = append(output, total)
	}
	return output
}

func GetFitnessArray(individuals []models.Candidate) []float64 {
	fitnesses := []float64{}
	for _, individual := range individuals {
		fitnesses = append(fitnesses, individual.Fitness)
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
		cumulativeFitnesses   = CumulativeArray(reversedFitnesses)
		upperBoundLastFitness = cumulativeFitnesses[len(cumulativeFitnesses)-1]
		choosedIndividuals    = []int{selectionSize}
		choosen               int
		rouletteBullet        float64
	)

	for i := 0; i < selectionSize; i++ {
		rouletteBullet = rand.Float64() * upperBoundLastFitness
		choosen = utilities.BinaryRangeSearch(cumulativeFitnesses, rouletteBullet)
		choosedIndividuals = append(choosedIndividuals, choosen)
	}

}
