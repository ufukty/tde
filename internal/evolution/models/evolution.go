package models

type SearchParameters struct {
	Cap         int // max. number of individuals in search
	Generations int // max. number of generations before calling it failure
	Depth       int // adj. local/global search behaviour. suggestion 1 or 2
	Evaluations int // max evaluations per generation
}

type Parameters struct {
	Population  int
	Generations int
	Size        int // max. code size in bytes
	Packages    []string

	Code      SearchParameters
	Program   SearchParameters
	Candidate SearchParameters
	Solution  SearchParameters
}

type SearchId string
