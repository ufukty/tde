package selection

import (
	models "tde/models/program"
)

func getFitnesses(subjects []*models.Subject, layer models.Layer) []float64 {
	fitnesses := []float64{}
	for _, individual := range subjects {
		fitnesses = append(fitnesses, individual.Fitness.InLayer(layer))
	}
	return fitnesses
}

// reverses fitness from [0,1] to [1,0] range
func reverse(fitnesses []float64) []float64 {
	normalized := []float64{}
	for _, f := range fitnesses {
		normalized = append(normalized, 1.0-f)
	}
	return normalized
}
