package evolution

import (
	"math"
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

func GetFitnessArray(individuals []Individual) []float64 {
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
	return fitnesses
}

func BinarySearchSmallestOfGreaters(values []float64, key float64) int {
	var (
		lo  = 0
		mid int
		hi  = len(values) - 1
	)

	for hi-lo > 1 {
		mid = int(math.Floor(float64(lo+hi) / 2))
		if values[mid] <= key {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return hi
}

func SelectionRouletteWheel(individuals []Individual, selectionSize int) {
	var (
		fitnesses             = GetFitnessArray(individuals)
		reversedFitnesses     = ReversedFitnesses(fitnesses)
		cumulativeFitnesses   = CumulativeArray[float64](reversedFitnesses)
		upperBoundLastFitness = cumulativeFitnesses[len(cumulativeFitnesses)-1]
		choosedIndividuals    = []int{selectionSize}
		choosen               int
		rouletteBullet        float64
	)

	for i := 0; i < selectionSize; i++ {
		rouletteBullet = rand.Float64() * upperBoundLastFitness
		choosen = BinarySearchSmallestOfGreaters(cumulativeFitnesses, rouletteBullet)
		choosedIndividuals = append(choosedIndividuals, choosen)
	}

}
