package evolution

import (
	"fmt"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/genetics/crossover/subtreeswitch"
	"tde/internal/evolution/genetics/mutation/v1"
	models1 "tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/evolution/models"
	"tde/internal/evolution/pool"
	"tde/internal/evolution/selection"
	"tde/internal/utilities/slicew"

	"golang.org/x/exp/maps"
)

type SolutionSearch struct {
	commons    *commons
	Searches   []*candidateSearch
	offsprings models.Subjects
}

func NewSolutionSearch(e *evaluation.Evaluator, params *models.Parameters, context *models.Context) *SolutionSearch {
	return &SolutionSearch{
		commons: &commons{
			Evaluator: e,
			Params:    params,
			Context:   context,
			Pool:      nil,
		},
	}
}

// Pick parents for genetic operations
func (ss *SolutionSearch) pickParents(candidates models.Subjects) (co []*[2]*models.Subject, mu models.Subjects, err error) {
	n := ss.commons.Params.Solution.Evaluations
	parents, err := selection.RouletteWheelToReproduce(candidates, models.Candidate, n)
	if err != nil {
		return nil, nil, fmt.Errorf("running RouletteWheel: %w", err)
	}
	nCo := int((float64(n)) / 20)
	coA, err := selection.RouletteWheelToReproduce(parents, models.Candidate, nCo)
	if err != nil {
		return nil, nil, fmt.Errorf("picking crossover parents from picked parents: %w", err)
	}
	coB, err := selection.RouletteWheelToReproduce(parents, models.Candidate, nCo)
	if err != nil {
		return nil, nil, fmt.Errorf("picking crossover parents from picked parents: %w", err)
	}
	co = slicew.Zip(maps.Values(coA), maps.Values(coB))
	nMu := n - 2*nCo
	mu, err = selection.RouletteWheelToReproduce(parents, models.Candidate, nMu)
	if err != nil {
		return nil, nil, fmt.Errorf("picking mutation parents from picked parents: %w", err)
	}
	return
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
	ss.commons.Pool.Join(allProducts)
	return nil
}

func (ss *SolutionSearch) pruneSearches() error {
	continuing := make([]*candidateSearch, 0, len(ss.Searches))
	ended := make([]*candidateSearch, 0, len(ss.Searches))
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

func (ss *SolutionSearch) reproduce() error {
	candidates := ss.commons.Pool.FilterValidIn(models.Candidate)

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
			opctx := &models1.MutationParameters{
				Package:         ss.commons.Context.Package,
				File:            ss.commons.Context.File,
				FuncDecl:        subj.AST,
				AllowedPackages: ss.commons.Params.Packages,
			}
			if err := mutation.Mutate(opctx); err != nil {
				return fmt.Errorf("applying mutation on offspring: %w", err)
			}
			ss.offsprings.Add(offspring)
		}
	}
	return nil
}

func (ss *SolutionSearch) evaluate() error {
	if len(ss.offsprings) > 0 {
		if err := ss.commons.Evaluator.Pipeline(ss.offsprings); err != nil {
			return fmt.Errorf("evaluating offsprings: %w", err)
		}
	}
	return nil
}

// span candidate searches for subjects that doesn't finish the unit test
func (ss *SolutionSearch) startSearches() {
	for _, o := range ss.offsprings {
		if !o.IsValidIn(models.Candidate) {
			cs := newCandidateSearch(ss.commons, o)
			ss.Searches = append(ss.Searches, cs)
		}
	}
}

func (ss *SolutionSearch) Loop() error {
	ss.commons.Pool = pool.New(ss.commons.Context.NewSubject())
	for i := 0; i < ss.commons.Params.Generations; i++ {
		if err := ss.reproduce(); err != nil {
			return fmt.Errorf("diversifying: %w", err)
		}
		if err := ss.evaluate(); err != nil {
			return fmt.Errorf("evaluating offsprings: %w", err)
		}
		if err := ss.iterateSearches(); err != nil {
			return fmt.Errorf("iterating candidate searches: %w", err)
		}
		if err := ss.pruneSearches(); err != nil {
			return fmt.Errorf("pruning failed candidate searches: %w", err)
		}
		ss.startSearches()
	}
	return nil
}
