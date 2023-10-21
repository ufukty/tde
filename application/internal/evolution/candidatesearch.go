package evolution

import (
	"fmt"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/pool"
	models "tde/models/program"
)

// Searches for a AST that compile, run & solve some user problems
type CandidateSearch struct {
	Evaluator *evaluation.Evaluator
	Pool      *pool.Pool
	Subpool   *pool.Pool
	Params    *models.Parameters
	Src       models.Sid
	Counter   int

	SubSearch *ProgramSearch
	Findings  []models.Sid
}

func (s *CandidateSearch) IsEnded() bool {
	return s.Counter == s.Params.Candidate.Generations
}

func (s *CandidateSearch) Terminate() error {
	s.SubSearch.Terminate()
	return nil
}

func (s *CandidateSearch) Iterate() (models.Subjects, error) {

	if s.SubSearch == nil {
		s.SubSearch = NewProgramSearch(s.Evaluator, s.Pool, s.Params)
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
		s.SubSearch = nil // TODO: Deallocate candidates in it?
	}

	return models.Subjects{}, nil
}

func NewCandidateSearch(evaluator *evaluation.Evaluator, pool *pool.Pool, parameters *models.Parameters, sid models.Sid) *CandidateSearch {
	return &CandidateSearch{
		Pool:      pool,
		Subpool:   pool.Sub(),
		Params:    parameters,
		Src:       sid,
		SubSearch: nil,
		Findings:  []models.Sid{},
	}
}
