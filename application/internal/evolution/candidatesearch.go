package evolution

import (
	"fmt"
	"tde/internal/evolution/pool"
	models "tde/models/program"
)

// Searches for a AST that compile, run & solve some user problems
type candidateSearch struct {
	*commons
	Subpool  *pool.Pool
	Src      models.Sid
	Counter  int
	Search   *programSearch
	Findings []models.Sid
}

func (cs *candidateSearch) IsEnded() bool {
	return cs.Counter == cs.Params.Candidate.Generations
}

func (s *candidateSearch) Terminate() error {
	s.Search.Terminate()
	return nil
}

func (s *candidateSearch) Iterate() (models.Subjects, error) {

	if s.Search == nil {
		// s.Search = newProgramSearch(s.commons, sid)
	}

	products, err := s.Search.Iterate()
	if err != nil {
		return nil, fmt.Errorf("iterating program search: %w", err)
	}

	if len(products) > 0 {
		return products, nil
	}

	if s.Search.IsEnded() {
		s.Counter++
		s.Search = nil // TODO: Deallocate candidates in it?
	}

	return models.Subjects{}, nil
}

func newCandidateSearch(commons *commons, root *models.Subject) *candidateSearch {
	return &candidateSearch{
		commons:  commons,
		Subpool:  commons.Pool.Sub(root),
		Src:      root.Sid,
		Search:   nil,
		Findings: []models.Sid{},
	}
}
