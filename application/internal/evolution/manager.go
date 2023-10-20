package evolution

import (
	"fmt"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/genetics/crossover/subtreeswitch"
	"tde/internal/evolution/genetics/mutation"
	"tde/internal/evolution/pool"
	"tde/internal/evolution/search"
	"tde/internal/evolution/selection"
	"tde/internal/utilities"
	models "tde/models/program"

	"golang.org/x/exp/maps"
)

type SolutionSearch struct {
	Evaluator *evaluation.Evaluator
	Params    *models.Parameters
	Pool      *pool.Pool
	Searches  search.CandidateSearches
	Context   *models.Context
}

func NewSolutionSearch(e *evaluation.Evaluator, params *models.Parameters, context *models.Context) *SolutionSearch {
	return &SolutionSearch{
		Evaluator: e,
		Params:    params,
		Context:   context,
	}
}

// Pick parents for genetic operations
func (m *SolutionSearch) PickParents(candidates models.Subjects) (co []*[2]*models.Subject, mu models.Subjects, err error) {
	n := m.Params.Solution.Evaluations
	parents, err := selection.RouletteWheel(candidates, models.Candidate, n)
	if err != nil {
		return nil, nil, fmt.Errorf("running RouletteWheel: %w", err)
	}
	nCo := int((float64(n)) / 20)
	coA, err := selection.RouletteWheel(parents, models.Candidate, nCo)
	if err != nil {
		return nil, nil, fmt.Errorf("picking crossover parents from picked parents: %w", err)
	}
	coB, err := selection.RouletteWheel(parents, models.Candidate, nCo)
	if err != nil {
		return nil, nil, fmt.Errorf("picking crossover parents from picked parents: %w", err)
	}
	co = utilities.SliceZipToSlice(maps.Values(coA), maps.Values(coB))
	nMu := n - 2*nCo
	mu, err = selection.RouletteWheel(parents, models.Candidate, nMu)
	if err != nil {
		return nil, nil, fmt.Errorf("picking mutation parents from picked parents: %w", err)
	}
	return
}

func (m *SolutionSearch) diversify() error {
	candidates := m.Pool.FilterValidIn(models.Candidate)

	co, mu, err := m.PickParents(candidates)
	if err != nil {
		return fmt.Errorf("picking parents: %w", err)
	}

	offsprings := models.Subjects{}

	// cross overs
	if len(co) > 0 {
		for _, pair := range co {
			oA, oB := pair[0].Clone(), pair[1].Clone()
			if ok := subtreeswitch.SubtreeSwitch(oA.AST, oB.AST); !ok {
				return fmt.Errorf("failed at subtree switch")
			}
			offsprings.Add(oA)
			offsprings.Add(oB)
		}
	}

	// mutations
	if len(mu) > 0 {
		for _, subj := range mu {
			offspring := subj.Clone()
			if err := mutation.Mutate(*m.Context, offspring, m.Params.Packages); err != nil {
				return fmt.Errorf("failed at mutation")
			}
			offsprings.Add(offspring)
		}
	}

	if len(offsprings) > 0 {

		if err := m.Evaluator.Pipeline(offsprings); err != nil {
			return fmt.Errorf("evaluating offsprings: %w", err)
		}

		// span candidate searches for subjects that doesn't finish the unit test
		for _, o := range offsprings {
			if !o.IsValidIn(models.Candidate) {
				m.Searches.Append(
					search.NewCandidateSearch(m.Evaluator, m.Pool, m.Params, m.Context, o.Sid),
				)
			}
		}
	}

	return nil
}

func (m *SolutionSearch) Init() error {
	m.Pool = pool.New(m.Context.NewSubject())

	for i := 0; i < m.Params.Generations; i++ {
		candidates := m.Pool.FilterValidIn(models.Candidate)

		if products, err := m.Searches.Iterate(); err != nil {
			return fmt.Errorf("pruning failed searches: %w", err)
		} else {
			candidates.Join(products)
		}

		if err := m.Searches.Prune(); err != nil {
			return fmt.Errorf("pruning failed searches: %w", err)
		}
		if err := m.diversify(); err != nil {
			return fmt.Errorf("diversifying: %w", err)
		}
	}

	return nil
}
