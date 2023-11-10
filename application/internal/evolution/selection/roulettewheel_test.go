package selection

import (
	"fmt"
	"sort"
	"tde/internal/evolution/models"
	"tde/internal/utilities"
	"testing"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Test_normalize(t *testing.T) {
	type Case struct {
		input []float64
		want  []float64
	}
	cases := []Case{
		{
			input: []float64{},
			want:  []float64{},
		},
		{
			input: []float64{0.0},
			want:  []float64{1.0},
		},
		{
			input: []float64{1.0},
			want:  []float64{0.0},
		},
		{
			input: []float64{0.5},
			want:  []float64{0.5},
		},
		{
			input: []float64{0.2, 0.3, 0.4},
			want:  []float64{0.8, 0.7, 0.6},
		},
	}
	for i, tcase := range cases {
		got := reverse(tcase.input)
		if d := slices.Compare(tcase.want, got); d != 0 {
			t.Errorf("Run %d: %d differences:\n    passed  : %v\n    expected: %v\n    got     : %v", i, d, tcase.input, tcase.want, got)
		}
	}
}

func Test_RouletteWheelFrequencyDistribution(t *testing.T) {
	const runsPerDataset = 1000

	var datasets = [][]float64{
		{0.0},
		{0.0, 1.0},
		{1.0, 0.0},
		{0.0, 0.0, 1.0},
		{0.0, 1.0, 0.0},
		{0.0, 1.0, 1.0},
		{1.0, 0.0, 0.0},
		{1.0, 0.0, 1.0},
		{1.0, 1.0, 0.0},
		{1.0, 0.22, 0.20, 0.28, 0.18, 0.35, 0.93, 0.21, 0.12, 0.39, 0.33, 0.0, 0.34, 0.26, 0.28, 0.30, 0.34, 0.22, 0.0, 0.29, 0.21, 0.22, 0.39, 0.39, 0.32, 0.32, 0.15, 0.24, 0.92, 0.28, 0.19, 0.0, 0.74, 0.10, 0.22, 0.30, 0.16, 0.35, 1.0, 0.31, 1.0, 0.22, 0.20, 0.28, 0.18, 0.35, 0.93, 0.21, 0.12, 0.39, 0.33, 0.0, 0.34, 0.26, 0.28, 0.30, 0.34, 0.22, 0.0, 0.29, 0.21, 0.22, 0.39, 0.39, 0.32, 0.32, 0.15, 0.24, 0.92, 0.28, 0.19, 0.0, 0.74, 0.10, 0.22, 0.30, 0.16, 0.35, 1.0, 0.31},
		{0.000, 0.025, 0.050, 0.075, 0.100, 0.125, 0.200, 0.225, 0.250, 0.275, 0.300, 0.325, 0.350, 0.375, 0.400, 0.425, 0.450, 0.475, 0.500, 0.525, 0.550, 0.575, 0.600, 0.625, 0.650, 0.675, 0.700, 0.725, 0.750, 0.775, 0.800, 0.825, 0.850, 0.875, 0.900, 0.925, 0.950, 0.975, 1.000},
	}

	for _, dataset := range datasets {
		t.Run(fmt.Sprintf("%v", dataset), func(t *testing.T) {

			var subjects = subjectsForDataset(dataset)
			var freqCounter = newFreqCounter(subjects)
			var imbalancedRuns = 0

			for j := 0; j < runsPerDataset; j++ {
				picks := RouletteWheelToEliminate(subjects, models.AST, int(len(subjects)/2))
				survivingBest, survivingWorst := freqCounter.count(maps.Keys(picks))
				if survivingWorst > survivingBest {
					fmt.Printf("Run %3d: Imbalanced: B:%d / W:%d\n", j, survivingBest, survivingWorst)
					imbalancedRuns++
				}
			}

			if imbalancedRuns > runsPerDataset*0.50 {
				t.Fatal(fmt.Errorf("assert %d runs suffer imbalanced elimination", imbalancedRuns))
			}

			fmt.Printf("imbalanced runs: %d\n", imbalancedRuns)
			freqCounter.PrintHistogram()
		})
	}
}

func Test_RouletteWheelAllFailingSubjects(t *testing.T) {
	var datasets = [][]float64{
		{1.0},
		{1.0, 1.0},
		{1.0, 1.0, 1.0},
	}

	for _, dataset := range datasets {
		subjects := subjectsForDataset(dataset)
		for pick := 0; pick <= len(dataset); pick++ {
			t.Run(fmt.Sprintf("%v>%d", dataset, pick), func(t *testing.T) {
				fmt.Printf("%v>%d\n", dataset, pick)
				selection := RouletteWheelToEliminate(subjects, models.AST, pick)
				if len(selection) != pick {
					t.Fatal(fmt.Errorf("assert, selection length: expected %d, got %d items", pick, len(selection)))
				}
			})
		}
	}
}

// MARK: test utilities

func filterBestAndWorstIds(subjects map[models.Sid]*models.Subject) (sortedIds []models.Sid, bests []models.Sid, worsts []models.Sid) {
	_, cands := utilities.MapItems(subjects)
	sort.Slice(cands, func(i, j int) bool {
		return cands[i].Fitness.AST < cands[j].Fitness.AST
	})
	for _, cand := range cands {
		sortedIds = append(sortedIds, cand.Sid)
	}
	half := int(float64(len(subjects)) * 0.5)
	for i := 0; i < half; i++ {
		bests = append(bests, cands[i].Sid)
		worsts = append(worsts, cands[len(subjects)-1-i].Sid)
	}
	return
}

type freqCounter struct {
	subjects                     map[models.Sid]*models.Subject
	idFreqs                      map[models.Sid]int
	sortedIds, bestIds, worstIds []models.Sid
}

func newFreqCounter(subjects map[models.Sid]*models.Subject) *freqCounter {
	ids := maps.Keys(subjects)
	freqs := map[models.Sid]int{}
	for _, id := range ids {
		freqs[id] = 0
	}
	sortedIds, bestIds, worstIds := filterBestAndWorstIds(subjects)
	return &freqCounter{
		subjects:  subjects,
		idFreqs:   freqs,
		sortedIds: sortedIds,
		bestIds:   bestIds,
		worstIds:  worstIds,
	}
}

func (fc *freqCounter) count(subjects []models.Sid) (survivingBest int, survivingWorst int) {
	for _, id := range subjects {
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
	if len(fc.subjects) == 0 {
		return
	}
	fmt.Println("Histogram of frequencies:")
	maxFreq := slices.Max(maps.Values(fc.idFreqs))
	for _, id := range fc.sortedIds {
		freq := fc.idFreqs[id]
		fmt.Printf("    %2s %.2f %3d%% %s\n", string(id), fc.subjects[id].Fitness.AST, int(float64(freq)/float64(maxFreq)*100), utilities.StringFill("*", int(float64(freq)/float64(maxFreq)*40)))
	}
}

func prepare(subjects map[models.Sid]*models.Subject, picks []models.Sid) map[models.Sid]*models.Subject {
	ret := map[models.Sid]*models.Subject{}
	for _, id := range picks {
		delete(ret, id)
	}
	return ret
}

func subjectsForDataset(dataset []float64) map[models.Sid]*models.Subject {
	var subjects = map[models.Sid]*models.Subject{}
	for i, f := range dataset {
		id := models.Sid(fmt.Sprintf("%d", i))
		subjects[id] = &models.Subject{Sid: id, Fitness: models.Fitness{AST: f}}
	}
	return subjects
}
