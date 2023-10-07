package search

import (
	"fmt"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/pool"
	models "tde/models/program"
)

// Searches for a AST that compile & run
type ProgramSearch struct {
	Evaluator *evaluation.Evaluator
	Pool      *pool.Pool
	Subpool   *pool.Pool
	Params    *models.Parameters
	Src       models.Sid
	Counter   int

	SubSearch *CodeSearch
	Dests     models.Sid
}

func (s *ProgramSearch) IsEnded() bool {
	return s.Counter == s.Params.Program.Generations
}

func (s *ProgramSearch) Iterate() (models.Subjects, error) {
	// collect findings

	if s.SubSearch == nil {
		s.SubSearch = NewCodeSearch(s.Evaluator, s.Pool)
	}

	products, err := s.SubSearch.Iterate()
	if err != nil {
		return nil, fmt.Errorf("iterating program search: %w", err)
	}

	if len(products) > 0 {
		return products, nil
	}

	if s.SubSearch.IsEnded() {
		s.Counter++
	}

	return models.Subjects{}, nil
}

func (s *ProgramSearch) Terminate() {
	
}

func NewProgramSearch(evaluator *evaluation.Evaluator, pool *pool.Pool, parameters *models.Parameters, sid models.Sid) *ProgramSearch {
	return &ProgramSearch{
		Evaluator: evaluator,
		Pool:      pool,
		Subpool:   pool.Sub(),
		Params:    parameters,
		Src:       sid,
		Counter:   0,
		SubSearch: nil,
		Dests:     "",
	}
}
