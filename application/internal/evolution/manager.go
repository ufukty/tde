package evolution

import (
	"fmt"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/pool"
	"tde/internal/evolution/search"
	models "tde/models/program"
)

type Manager struct {
	Evaluator *evaluation.Evaluator
	Params    *models.Parameters
	Pool      *pool.Pool
	Branches  []*search.CandidateSearch
	Balances  BalanceBook
	// Target    *Target
	// HallOfFame map[int]*models.Candidate
	// Subjects models.Subjects
}

func NewManager(parameters *models.Parameters, evaluator *evaluation.Evaluator) *Manager {
	return &Manager{
		Evaluator: evaluator,
		Params:    parameters,
		Pool:      pool.New(),
	}
}

func (m *Manager) iterateBranches() {
	for _, ss := range m.Branches {
		ss.Iterate()
	}
}

func (m *Manager) collectResultsFromBranches() {
	// for _, ss := range m.Branches {
	// 	ss.Collect()
	// }
}

func (m *Manager) endSearches() {
	// for _, ss := range m.Branches {
	// 	ss.Collect()
	// }
}

func (m *Manager) startSearches() {
	// for _, ss := range m.Branches {
	// 	ss.Collect()
	// }
}

func (m *Manager) Init(target *Target) error {
	var startingPoint, err = newCandidate(target.Package, target.File, target.FuncDecl)
	if err != nil {
		return fmt.Errorf("creating a subject instance representing the target: %w", err)
	}

	// var ss = search.NewSolutionSearch(m.Pool, m.Params, startingPoint.UUID)
	// e.Searches = append(e.Searches, ss)

	for i := 0; i < m.Params.Generations; i++ {
		m.Balances.NewGen()

		m.iterateBranches()
		m.collectResultsFromBranches()
		m.endSearches()
		m.startSearches()

		// sols := e.Pool.Filter(models.Solution)

		if err = m.Evaluator.Pipeline([]*models.Subject{startingPoint}); err != nil {
			return fmt.Errorf("evaluating point-0: %w", err)
		}
	}

	return nil
}
