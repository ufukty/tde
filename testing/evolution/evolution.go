package evolution

import (
	"fmt"
	models "models/in_program_models"
)

type TargetFunctionType func(string) string // FIXME: REPLACE BEFORE COMPILE

type TestCandidate struct {
	TestingRef        *E
	CandidateFunction TargetFunctionType
	UUID              models.CandidateID
	AssertionResults  []bool
}

func (tc *TestCandidate) AssertEqual(output, want any) {
	result := output == want
	tc.AssertionResults = append(tc.AssertionResults, result)
}

func (tc *TestCandidate) AssertNotEqual(output, notWant any) {
	result := output != notWant
	tc.AssertionResults = append(tc.AssertionResults, result)
}

type E struct {
	Candidates     map[models.CandidateID]*models.Candidate
	TestCandidates map[models.CandidateID]*TestCandidate
	CandidateRefs  map[models.CandidateID]TargetFunctionType
}

func NewE(candidateRefs map[models.CandidateID]TargetFunctionType) *E {
	e := E{
		CandidateRefs: candidateRefs,
	}
	return &e
}

func (e *E) GetCandidate(candidateTesting func(c *TestCandidate)) {

	for id := range e.Candidates {
		// candidateFunction := t.Candidates[id].
		c := TestCandidate{
			TestingRef:        e,
			CandidateFunction: e.CandidateRefs[id],
			UUID:              id,
			AssertionResults:  []bool{},
		}
		candidateTesting(&c)
	}
}

func (e *E) CalculateFitnesses() { // leave this to "evolve" binary3
	for id, candidate := range e.TestCandidates {
		totalSuccessfulAssertions := 0
		for _, result := range candidate.AssertionResults {
			if result {
				totalSuccessfulAssertions++
			}
		}
		e.Candidates[id].Fitness = float64(totalSuccessfulAssertions) / float64(len(candidate.AssertionResults))
	}
}

func (e *E) PrintStats() {
	fmt.Println("Stats are printed")
	fmt.Println("Number of candidates:", len(e.Candidates))
	for id, cand := range e.Candidates {
		fmt.Printf("%s: %f\n", id, cand.Fitness)

	}
}
