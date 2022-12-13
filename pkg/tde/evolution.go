package tde

import (
	models "tde/models/in_program_models"

	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

type TargetFunctionType any

// Candidate, reference to its embedded function and storage for assertion results
type C struct {
	TestingRef             *E                 `json:"-"`
	Function               TargetFunctionType `json:"-"`
	UUID                   models.CandidateID `json:"uuid"`
	AssertionResults       []bool             `json:"assertion_results"`
	AssertionErrorDistance []float64          `json:"assertion_error_distances"`
	ExecTime               time.Duration      `json:"exec_time"`
	Panicked               bool               `json:"panicked"`
}

func (tc *C) AssertEqual(output, want any) {
	result := output == want
	tc.AssertionResults = append(tc.AssertionResults, result)

	errorDistance := 0.0
	switch wantCasted := want.(type) {
	case string:
		if outputCasted, ok := output.(string); ok {
			errorDistance = StringDistance(wantCasted, outputCasted)
		} else {
			panic("<output> and <want> should be same type and one of string, int, float64")
		}
	case int:
		if outputCasted, ok := output.(int); ok {
			errorDistance = IntegerDistance(wantCasted, outputCasted)
		} else {
			panic("<output> and <want> should be same type and one of string, int, float64")
		}
	case float64:
		if outputCasted, ok := output.(float64); ok {
			errorDistance = FloatDistance(wantCasted, outputCasted)
		} else {
			panic("<output> and <want> should be same type and one of string, int, float64")
		}
	}
	tc.AssertionErrorDistance = append(tc.AssertionErrorDistance, errorDistance)
}

// func (tc *C) AssertNotEqual(output, notWant any) {
// 	result := output != notWant
// 	tc.AssertionResults = append(tc.AssertionResults, result)

// 	errorDistance := 0.0
// 	switch wantCasted := notWant.(type) {
// 	case string:
// 		if outputCasted, ok := output.(string); ok {
// 			errorDistance = StringDistance(wantCasted, outputCasted)
// 		} else {
// 			panic("<output> and <want> should be same type and one of string, int, float64")
// 		}
// 	case int:
// 		if outputCasted, ok := output.(int); ok {
// 			errorDistance = IntegerDistance(wantCasted, outputCasted)
// 		} else {
// 			panic("<output> and <want> should be same type and one of string, int, float64")
// 		}
// 	case float64:
// 		if outputCasted, ok := output.(float64); ok {
// 			errorDistance = FloatDistance(wantCasted, outputCasted)
// 		} else {
// 			panic("<output> and <want> should be same type and one of string, int, float64")
// 		}
// 	}
// 	tc.AssertionErrorDistance = append(tc.AssertionErrorDistance, errorDistance)
// }

type E struct {
	// Candidates     map[models.CandidateID]*models.Candidate
	TestCandidates map[models.CandidateID]*C
	CandidateRefs  map[models.CandidateID]TargetFunctionType
}

func NewE(candidateRefs map[models.CandidateID]TargetFunctionType) *E {
	testCandidates := map[models.CandidateID]*C{}
	e := E{
		TestCandidates: testCandidates,
		CandidateRefs:  candidateRefs,
	}
	for id, ref := range candidateRefs {
		testCandidates[id] = &C{
			UUID:                   id,
			AssertionResults:       []bool{},
			AssertionErrorDistance: []float64{},
			TestingRef:             &e,
			Function:               ref,
		}
	}
	return &e
}

func (e *E) TestCandidate(candidateTesting func(candidate *C)) {
	for _, tc := range e.TestCandidates {
		start := time.Now()

		go func(tc *C) { // BECAUSE: Run in parallel
			defer func() { // If certain candidate panics, pass it by marking for later consideration
				if r := recover(); r != nil {
					fmt.Printf("Panic called by Candidate:%s is recovered: %s\n", tc.UUID, r)
					tc.Panicked = true
				} else {
					tc.Panicked = false
				}
				tc.ExecTime = time.Since(start)
			}()
			candidateTesting(tc) // FIXME: Terminate when candidate exceeds time limit (1sec?)
		}(tc)
	}
}

func (e *E) Export() {
	cleanDynamicKeys := []*C{}
	for _, tc := range e.TestCandidates {
		cleanDynamicKeys = append(cleanDynamicKeys, tc)
	}

	f, err := os.Create("results.json")
	if err != nil {
		panic(errors.Wrap(err, "Could not create 'result.json' for writing"))
	}
	json.NewEncoder(f).Encode(cleanDynamicKeys)
}
