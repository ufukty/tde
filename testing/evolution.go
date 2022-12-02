package testing

import (
	models "models/in_program_models"
	"reflect"
)

type TargetFunctionType func(string) string // FIXME: REPLACE BEFORE COMPILE

type Evolution struct {
	Candidates     map[models.CandidateID]*models.Candidate
	TestCandidates map[models.CandidateID]*TestCandidate
}

func CreateTesting(reflect.Type) *Evolution {
	return &Evolution{}
}

type TestCandidate struct {
	TestingRef        *Evolution
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

func (t *Evolution) GetCandidate(candidateTesting func(c *TestCandidate)) {

	
	for id := range t.Candidates {
		candidateFunction := t.Candidates[id].
		c := TestCandidate{
			TestingRef:        t,
			CandidateFunction: ,
			UUID:              id,
			AssertionResults:  []bool{},
		}
		candidateTesting(&c)
	}
}
