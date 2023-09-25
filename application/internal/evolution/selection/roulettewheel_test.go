package selection

import (
	"fmt"
	"sort"
	"tde/internal/utilities"
	models "tde/models/program"
	"testing"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func filterBestAndWorstIds(candidates map[models.CandidateID]*models.Candidate) (sortedIds []models.CandidateID, bests []models.CandidateID, worsts []models.CandidateID) {
	_, cands := utilities.MapItems(candidates)
	sort.Slice(cands, func(i, j int) bool {
		return cands[i].Fitness.AST < cands[j].Fitness.AST
	})
	for _, cand := range cands {
		sortedIds = append(sortedIds, cand.UUID)
	}
	for i := 0; i < int(float64(len(candidates))*0.5); i++ {
		bests = append(bests, cands[i].UUID)
		worsts = append(worsts, cands[len(candidates)-1-i].UUID)
	}
	return
}

type freqCounter struct {
	candidates                   map[models.CandidateID]*models.Candidate
	idFreqs                      map[models.CandidateID]int
	sortedIds, bestIds, worstIds []models.CandidateID
}

func newFreqCounter(candidates map[models.CandidateID]*models.Candidate) *freqCounter {
	ids := maps.Keys(candidates)
	freqs := map[models.CandidateID]int{}
	for _, id := range ids {
		freqs[id] = 0
	}
	sortedIds, bestIds, worstIds := filterBestAndWorstIds(candidates)
	return &freqCounter{
		candidates: candidates,
		idFreqs:    freqs,
		sortedIds:  sortedIds,
		bestIds:    bestIds,
		worstIds:   worstIds,
	}
}

func (fc *freqCounter) count(candidates []models.CandidateID) (survivingBest int, survivingWorst int) {
	for _, id := range candidates {
		fc.idFreqs[id]++
		if slices.Contains(fc.bestIds, id) {
			survivingBest++
		}
		if slices.Contains(fc.worstIds, id) {
			survivingWorst++
		}
	}
	return
}

func (fc freqCounter) PrintHistogram() {
	if len(fc.candidates) == 0 {
		return
	}
	fmt.Println("Histogram of frequencies:")
	maxFreq := slices.Max(maps.Values(fc.idFreqs))
	for _, id := range fc.sortedIds {
		freq := fc.idFreqs[id]
		fmt.Printf("    %2s %.2f %3d%% %s\n", string(id), fc.candidates[id].Fitness.AST, int(float64(freq)/float64(maxFreq)*100), utilities.StringFill("*", int(float64(freq)/float64(maxFreq)*40)))
	}
}

func prepare(candidates map[models.CandidateID]*models.Candidate, picks []models.CandidateID) map[models.CandidateID]*models.Candidate {
	ret := map[models.CandidateID]*models.Candidate{}
	for _, id := range picks {
		delete(ret, id)
	}
	return ret
}

func candidatesForDataset(dataset []float64) map[models.CandidateID]*models.Candidate {
	var candidates = map[models.CandidateID]*models.Candidate{}
	for i, f := range dataset {
		id := models.CandidateID(fmt.Sprintf("%d", i))
		candidates[id] = &models.Candidate{UUID: id, Fitness: models.Fitness{AST: f}}
	}
	return candidates
}

func Test_RouletteWheelDistributionWithDatasets(t *testing.T) {
	const (
		ndatasets      = 10
		runsPerDataset = 1000
		bulletsPerRun  = 10
	)

	var datasets = [][]float64{
		{},
		{0.0},
		{1.0},
		{1.0, 0.22, 0.20, 0.28, 0.18, 0.35, 0.93, 0.21, 0.12, 0.39, 0.33, 0.0, 0.34, 0.26, 0.28, 0.30, 0.34, 0.22, 0.0, 0.29, 0.21, 0.22, 0.39, 0.39, 0.32, 0.32, 0.15, 0.24, 0.92, 0.28, 0.19, 0.0, 0.74, 0.10, 0.22, 0.30, 0.16, 0.35, 1.0, 0.31, 1.0, 0.22, 0.20, 0.28, 0.18, 0.35, 0.93, 0.21, 0.12, 0.39, 0.33, 0.0, 0.34, 0.26, 0.28, 0.30, 0.34, 0.22, 0.0, 0.29, 0.21, 0.22, 0.39, 0.39, 0.32, 0.32, 0.15, 0.24, 0.92, 0.28, 0.19, 0.0, 0.74, 0.10, 0.22, 0.30, 0.16, 0.35, 1.0, 0.31},
		{0.000, 0.025, 0.050, 0.075, 0.100, 0.125, 0.200, 0.225, 0.250, 0.275, 0.300, 0.325, 0.350, 0.375, 0.400, 0.425, 0.450, 0.475, 0.500, 0.525, 0.550, 0.575, 0.600, 0.625, 0.650, 0.675, 0.700, 0.725, 0.750, 0.775, 0.800, 0.825, 0.850, 0.875, 0.900, 0.925, 0.950, 0.975, 1.000},
	}

	for i, dataset := range datasets {
		fmt.Println("\nRunning the dataset:", i)

		var (
			candidates  = candidatesForDataset(dataset)
			freqCounter = newFreqCounter(candidates)
		)

		var imbalancedRuns = 0
		for j := 0; j < runsPerDataset; j++ {
			picks := RouletteWheel(candidates, models.AST, true)
			if len(candidates) != len(picks) {
				t.Fatal(fmt.Errorf("assert: len(picks) = %d", len(picks)))
			}
			if survivingBest, survivingWorst := freqCounter.count(picks); survivingWorst > survivingBest {
				fmt.Printf("Run %d: Imbalanced: %d, %d\n", j, survivingWorst, survivingBest)
				imbalancedRuns++
			}
		}

		if imbalancedRuns > runsPerDataset*0.33 {
			t.Fatal(fmt.Errorf("assert %d of the runs suffer imbalanced elimination", imbalancedRuns))
		}
		fmt.Printf("imbalanced runs: %d\n", imbalancedRuns)
		freqCounter.PrintHistogram()
	}
}
