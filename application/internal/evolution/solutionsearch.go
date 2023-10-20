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
	Evaluator  *evaluation.Evaluator
	Params     *models.Parameters
	Pool       *pool.Pool
	Searches   []*search.CandidateSearch
	Context    *models.Context
	offsprings models.Subjects
}

func NewSolutionSearch(e *evaluation.Evaluator, params *models.Parameters, context *models.Context) *SolutionSearch {
	return &SolutionSearch{
		Evaluator: e,
		Params:    params,
		Context:   context,
	}
}

// Pick parents for genetic operations
func (ss *SolutionSearch) pickParents(candidates models.Subjects) (co []*[2]*models.Subject, mu models.Subjects, err error) {
	n := ss.Params.Solution.Evaluations
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

func (ss *SolutionSearch) pruneSearches() error {
	continuing := make([]*search.CandidateSearch, 0, len(ss.Searches))
	ended := make([]*search.CandidateSearch, 0, len(ss.Searches))
	for _, i := range ss.Searches {
		if i.IsEnded() {
			ended = append(ended, i)
		} else {
			continuing = append(continuing, i)
		}
	}
	ss.Searches = continuing
	for _, i := range ended {
		if err := i.Terminate(); err != nil {
			return fmt.Errorf("terminating a candidate search: %w", err)
		}
	}
	return nil
}

func (ss *SolutionSearch) iterateSearches() error {
	allProducts := models.Subjects{}
	for _, search := range ss.Searches {
		products, err := search.Iterate()
		if err != nil {
			return fmt.Errorf("iterating a candidate search (sid: %s): %w", search.Src, err)
		}
		if products != nil && len(products) > 0 {
			allProducts.Join(products)
		}
	}
	ss.Pool.Join(allProducts)
	return nil
}

func (ss *SolutionSearch) diversify() error {
	candidates := ss.Pool.FilterValidIn(models.Candidate)

	co, mu, err := ss.pickParents(candidates)
	if err != nil {
		return fmt.Errorf("picking parents: %w", err)
	}

	ss.offsprings = models.Subjects{}

	// cross overs
	if len(co) > 0 {
		for _, pair := range co {
			oA, oB := pair[0].Clone(), pair[1].Clone()
			if ok := subtreeswitch.SubtreeSwitch(oA.AST, oB.AST); !ok {
				return fmt.Errorf("failed at subtree switch")
			}
			ss.offsprings.Add(oA)
			ss.offsprings.Add(oB)
		}
	}

	// mutations
	if len(mu) > 0 {
		for _, subj := range mu {
			offspring := subj.Clone()
			if err := mutation.Mutate(*ss.Context, offspring, ss.Params.Packages); err != nil {
				return fmt.Errorf("failed at mutation")
			}
			ss.offsprings.Add(offspring)
		}
	}
	return nil
}

func (ss *SolutionSearch) evaluate() error {
	if len(ss.offsprings) > 0 {
		if err := ss.Evaluator.Pipeline(ss.offsprings); err != nil {
			return fmt.Errorf("evaluating offsprings: %w", err)
		}
	}
	return nil
}

// span candidate searches for subjects that doesn't finish the unit test
func (ss *SolutionSearch) startSearches() {
	for _, o := range ss.offsprings {
		if !o.IsValidIn(models.Candidate) {
			cs := search.NewCandidateSearch(ss.Evaluator, ss.Pool, ss.Params, ss.Context, o.Sid)
			ss.Searches = append(ss.Searches, cs)
		}
	}
}

func (ss *SolutionSearch) Loop() error {
	ss.Pool = pool.New(ss.Context.NewSubject())
	for i := 0; i < ss.Params.Generations; i++ {
		if err := ss.iterateSearches(); err != nil {
			return fmt.Errorf("iterating candidate searches: %w", err)
		}
		if err := ss.pruneSearches(); err != nil {
			return fmt.Errorf("pruning failed candidate searches: %w", err)
		}
		if err := ss.diversify(); err != nil {
			return fmt.Errorf("diversifying: %w", err)
		}
		if err := ss.evaluate(); err != nil {
			return fmt.Errorf("evaluating offsprings: %w", err)
		}
		ss.startSearches()
	}
	return nil
}
