package evolution

import (
	"fmt"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/genetics/mutation"
	"tde/internal/evolution/genetics/mutation/common"
	"tde/internal/evolution/pool"
	"tde/internal/evolution/selection"
	models "tde/models/program"

	"golang.org/x/exp/maps"
)

// 1. merkezlerin sonuca ortalama uzaklığı derinlikle azalır.
// 2. merkezden uzaklaşmak gereklidir. hedef merkezden uzaktadır.

// doğru yönü bulmak için başlangıçtan çok sayıda deneme yapmak gerekir
//

// Searches for a AST that compile
type CodeSearch struct {
	Evaluator *evaluation.Evaluator
	Pool      *pool.Pool
	Subpool   *pool.Pool
	Params    *models.Parameters
	Src       models.Sid
	Counter   int
}

func (s *CodeSearch) CollectProducts() models.Subjects {
	return s.Subpool.FilterValidIn(models.Code)
}

func (s *CodeSearch) IsEnded() bool {
	return s.Counter == s.Params.Code.Generations
}

func (s *CodeSearch) Iterate() (models.Subjects, error) {
	var (
		subjects            = s.Subpool.All()
		pop                 = len(subjects)
		noEliminations      = min(int(float64(pop)/2.0), s.Params.Code.Evaluations) // no need to eliminate more than we can generate
		popAfterElimination = pop - noEliminations
		availableCap        = s.Params.Code.Cap - popAfterElimination
		noGenerations       = min(availableCap, s.Params.Code.Evaluations) // can't reproduce more than allowed anyway
	)

	// selection
	if noEliminations > 0 {
		next := selection.Elitist(subjects, models.Code, popAfterElimination)
		diff := subjects.Diff(next)
		for sid := range diff {
			s.Subpool.Delete(sid)
		}
	}

	products := models.Subjects{}
	offsprings := models.Subjects{}

	// reproduction
	if noGenerations > 0 {
		parents, err := selection.RouletteWheel(s.Pool.FilterByDepth(s.Params.Code.Depth), models.AST, noGenerations)
		if err != nil {
			return models.Subjects{}, fmt.Errorf("picking parents with roulette wheel: %w", err)
		}

		for _, subj := range parents {
			offspring := subj.Clone()
			offsprings.Add(offspring)
			op := mutation.Pick()
			opctx := &common.GeneticOperationContext{
				Package:         offspring.AST.Package,
				File:            offspring.AST.File,
				FuncDecl:        offspring.AST.FuncDecl,
				AllowedPackages: s.Params.Packages,
			}
			if ok := op(opctx); !ok {
				return models.Subjects{}, fmt.Errorf("applying mutation on offspring")
			}
		}
	}

	// evaluate
	s.Evaluator.Pipeline(maps.Values(offsprings))

	// check results

	// limit reached
	// early finding

	s.Counter++
	return products, nil
}

func NewCodeSearch(evaluator *evaluation.Evaluator, pool *pool.Pool, parameters *models.Parameters, sid models.Sid) *CodeSearch {
	return &CodeSearch{
		Evaluator: evaluator,
		Pool:      pool,
		Subpool:   pool.Sub(),
		Params:    parameters,
		Src:       sid,
		Counter:   0,
	}
}
