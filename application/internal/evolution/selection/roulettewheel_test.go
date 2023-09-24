package selection

import (
	"fmt"
	"tde/internal/utilities"
	models "tde/models/program"
	"testing"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Test_RouletteWheelDistribution(t *testing.T) {
	const (
		runs    = 1000
		bullets = 10
	)
	var (
		candidates = map[models.CandidateID]*models.Candidate{
			"1":  {UUID: "1", Fitness: models.Fitness{AST: 1.0}},
			"2":  {UUID: "2", Fitness: models.Fitness{AST: 0.22}},
			"3":  {UUID: "3", Fitness: models.Fitness{AST: 0.20}},
			"4":  {UUID: "4", Fitness: models.Fitness{AST: 0.28}},
			"5":  {UUID: "5", Fitness: models.Fitness{AST: 0.18}},
			"6":  {UUID: "6", Fitness: models.Fitness{AST: 0.35}},
			"7":  {UUID: "7", Fitness: models.Fitness{AST: 0.93}},
			"8":  {UUID: "8", Fitness: models.Fitness{AST: 0.21}},
			"9":  {UUID: "9", Fitness: models.Fitness{AST: 0.12}},
			"10": {UUID: "10", Fitness: models.Fitness{AST: 0.39}},
			"11": {UUID: "11", Fitness: models.Fitness{AST: 0.33}},
			"12": {UUID: "12", Fitness: models.Fitness{AST: 0.0}},
			"13": {UUID: "13", Fitness: models.Fitness{AST: 0.34}},
			"14": {UUID: "14", Fitness: models.Fitness{AST: 0.26}},
			"15": {UUID: "15", Fitness: models.Fitness{AST: 0.28}},
			"16": {UUID: "16", Fitness: models.Fitness{AST: 0.30}},
			"17": {UUID: "17", Fitness: models.Fitness{AST: 0.34}},
			"18": {UUID: "18", Fitness: models.Fitness{AST: 0.22}},
			"19": {UUID: "19", Fitness: models.Fitness{AST: 0.0}},
			"20": {UUID: "20", Fitness: models.Fitness{AST: 0.29}},
			"21": {UUID: "21", Fitness: models.Fitness{AST: 0.21}},
			"22": {UUID: "22", Fitness: models.Fitness{AST: 0.22}},
			"23": {UUID: "23", Fitness: models.Fitness{AST: 0.39}},
			"24": {UUID: "24", Fitness: models.Fitness{AST: 0.39}},
			"25": {UUID: "25", Fitness: models.Fitness{AST: 0.32}},
			"26": {UUID: "26", Fitness: models.Fitness{AST: 0.32}},
			"27": {UUID: "27", Fitness: models.Fitness{AST: 0.15}},
			"28": {UUID: "28", Fitness: models.Fitness{AST: 0.24}},
			"29": {UUID: "29", Fitness: models.Fitness{AST: 0.92}},
			"30": {UUID: "30", Fitness: models.Fitness{AST: 0.28}},
			"31": {UUID: "31", Fitness: models.Fitness{AST: 0.19}},
			"32": {UUID: "32", Fitness: models.Fitness{AST: 0.0}},
			"33": {UUID: "33", Fitness: models.Fitness{AST: 0.74}},
			"34": {UUID: "34", Fitness: models.Fitness{AST: 0.10}},
			"35": {UUID: "35", Fitness: models.Fitness{AST: 0.22}},
			"36": {UUID: "36", Fitness: models.Fitness{AST: 0.30}},
			"37": {UUID: "37", Fitness: models.Fitness{AST: 0.16}},
			"38": {UUID: "38", Fitness: models.Fitness{AST: 0.35}},
			"39": {UUID: "39", Fitness: models.Fitness{AST: 1.0}},
			"40": {UUID: "40", Fitness: models.Fitness{AST: 0.31}},
		}
		bestIds   = []models.CandidateID{"12", "19", "32", "34", "9", "27", "37", "5", "31", "3"}
		worstIds  = []models.CandidateID{"38", "6", "10", "23", "24", "33", "29", "7", "1", "39"}
		sortedIds = []models.CandidateID{"39", "1", "7", "29", "33", "24", "23", "10", "6", "38", "17", "13", "11", "26", "25", "40", "36", "16", "20", "4", "30", "15", "14", "28", "35", "22", "2", "18", "8", "21", "3", "31", "5", "37", "27", "9", "34", "32", "19", "12"}
		idFreqs   = map[models.CandidateID]int{"1": 0, "2": 0, "3": 0, "4": 0, "5": 0, "6": 0, "7": 0, "8": 0, "9": 0, "10": 0, "11": 0, "12": 0, "13": 0, "14": 0, "15": 0, "16": 0, "17": 0, "18": 0, "19": 0, "20": 0, "21": 0, "22": 0, "23": 0, "24": 0, "25": 0, "26": 0, "27": 0, "28": 0, "29": 0, "30": 0, "31": 0, "32": 0, "33": 0, "34": 0, "35": 0, "36": 0, "37": 0, "38": 0, "39": 0, "40": 0}
	)

	var imbalancedRuns = 0
	for i := 0; i < runs; i++ {
		pickedToEliminate := RouletteWheel(candidates, models.AST, bullets, false)
		if len(candidates)-bullets != len(pickedToEliminate) {
			t.Fatal(fmt.Errorf("assert: len(candidates) = %d", len(pickedToEliminate)))
		}

		cleaned := map[models.CandidateID]*models.Candidate{}
		maps.Copy(cleaned, candidates)
		for _, id := range pickedToEliminate {
			delete(cleaned, id)
		}

		var (
			survivingBest  = 0
			survivingWorst = 0
		)
		for id := range cleaned {
			idFreqs[id]++
			if slices.Contains(bestIds, id) {
				survivingBest++
			}
			if slices.Contains(worstIds, id) {
				survivingWorst++
			}
		}
		if survivingWorst > survivingBest {
			fmt.Printf("Run %d: Imbalanced: %d, %d\n", i, survivingWorst, survivingBest)
			imbalancedRuns++
		} else {
			// fmt.Printf("Run %d: Balanced: %d, %d\n", i, survivingWorst, survivingBest)
		}
	}

	if imbalancedRuns > runs*0.33 {
		t.Fatal(fmt.Errorf("assert %d of the runs suffer imbalanced elimination", imbalancedRuns))
	}
	fmt.Printf("\nimbalanced runs: %d\n", imbalancedRuns)

	fmt.Println("\nHistogram of frequencies")
	maxFreq := slices.Max(maps.Values(idFreqs))
	for _, id := range sortedIds {
		freq := idFreqs[id]
		fmt.Printf("%2s %.2f %s\n", string(id), candidates[id].Fitness.AST, utilities.StringFill("*", int(float64(freq)/float64(maxFreq)*40)))
	}
}
