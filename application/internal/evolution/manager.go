package evolution

import (
	"fmt"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/pool"
	"tde/internal/evolution/search"
	"tde/internal/evolution/selection"
	models "tde/models/program"

	"golang.org/x/exp/maps"
)

type Manager struct {
	Evaluator *evaluation.Evaluator
	Params    *models.Parameters
	Pool      *pool.Pool
	Searches  search.CandidateSearches
	Balances  BalanceBook
}

func NewManager(parameters *models.Parameters, evaluator *evaluation.Evaluator) *Manager {
	return &Manager{
		Evaluator: evaluator,
		Params:    parameters,
	}
}

// Pick parents for genetic operations
func PickParents(params models.SearchParameters, s models.Subjects) (co models.Subjects, mu models.Subjects, err error) {
	parents, err := selection.RouletteWheel(s, models.Candidate, params.Evaluations)
	return
}

func (m *Manager) Init(start *models.Subject) error {
	m.Pool = pool.New(start.Sid)

	for i := 0; i < m.Params.Generations; i++ {
		m.Balances.NewGen()
		candidates := m.Pool.FilterValidIn(models.Candidate)

		for _, search := range m.Searches {
			products, err := search.Iterate()
			if err != nil {
				return fmt.Errorf("iterating a candidate search (sid: %s): %w", search.Src, err)
			}
			if products != nil && len(products) > 0 {
				candidates.Join(products)
			}
		}

		if err := m.Searches.Prune(); err != nil {
			return fmt.Errorf("pruning failed searches: %w", err)
		}

		offsprings := models.Subjects{}

		if err != nil {
			return fmt.Errorf("picking parents with RouletteWheel: %w", err)
		}

		// TODO: mutations

		// TODO: crossovers

		err := m.Evaluator.Pipeline(maps.Values(offsprings))
		if err != nil {
			return fmt.Errorf("evaluating point-0: %w", err)
		}

		// span candidate searches for subjects that doesn't finish the unit test
		for _, o := range offsprings {
			if !o.IsValidIn(models.Candidate) {
				cs := search.NewCandidateSearch(m.Evaluator, m.Pool, m.Params, start.Sid)
			}
		}
	}

	return nil
}
