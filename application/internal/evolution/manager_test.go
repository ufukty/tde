package evolution

import (
	"fmt"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/evaluation/discovery"
	"tde/internal/evolution/evaluation/inject"
	"tde/internal/evolution/evaluation/list"
	"tde/internal/evolution/evaluation/slotmgr"
	models "tde/models/program"
)

func prepare(parameters *models.Parameters) (*Manager, *discovery.CombinedDetails, error) {
	pkgs, err := list.ListPackagesInDir("testdata/words")
	if err != nil {
		return nil, nil, fmt.Errorf("lising packages in the package arranged for this test: %w", err)
	}
	sample, err := inject.WithCreatingSample(pkgs.First().Module.Dir, pkgs.First(), "TDE_WordReverse")
	if err != nil {
		return nil, nil, fmt.Errorf("creating the sample (injected one): %w", err)
	}
	cd, err := discovery.CombinedDetailsForTest("testdata/words", "TDE_WordReverse")
	if err != nil {
		return nil, nil, fmt.Errorf("getting combined details: %w", err)
	}
	sm := slotmgr.New(sample, cd)
	ev := evaluation.NewEvaluator(sm)
	em := NewManager(parameters, ev)
	return em, cd, nil
}

// func Test_Manager(t *testing.T) {
// 	parameters := &models.Parameters{
// 		Population:  0,
// 		Generations: 0,
// 		Size:        0,
// 		Packages:    []string{},
// 		Code: models.SearchParameters{
// 			Cap:         20,
// 			Generations: 3,
// 			Depth:       3,
// 			Evaluations: 2,
// 		},
// 		Program: models.SearchParameters{
// 			Cap:         4,
// 			Generations: 2,
// 			Depth:       2,
// 			Evaluations: 2,
// 		},
// 		Candidate: models.SearchParameters{
// 			Cap:         0,
// 			Generations: 2,
// 			Depth:       0,
// 			Evaluations: 0,
// 		},
// 	}

// 	em, cd, err := prepare(parameters)
// 	if err != nil {
// 		t.Fatalf("prep: %w", err)
// 	}
// 	em.Init()
// }
