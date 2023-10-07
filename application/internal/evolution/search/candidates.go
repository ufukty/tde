package search

import "fmt"

type CandidateSearches []*CandidateSearch

func (s *CandidateSearches) Append(cs *CandidateSearch) {
	(*s) = append((*s), cs)
}

func (s CandidateSearches) Prune() error {
	continuing := make(CandidateSearches, 0, len(s))
	ended := make(CandidateSearches, 0, len(s))
	for _, i := range s {
		if i.IsEnded() {
			ended = append(ended, i)
		} else {
			continuing = append(continuing, i)
		}
	}
	s = continuing
	for _, i := range ended {
		if err := i.Terminate(); err != nil {
			return fmt.Errorf("terminating a candidate search: %w", err)
		}
	}
	return nil
}
