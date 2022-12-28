package tde

import (
	models "tde/models/in_program_models"

	"encoding/json"
	"os"
	"time"

	"github.com/pkg/errors"
)

type Config struct {
	// in kiloflops
	MaxCompute int
	// in kilobytes
	MaxMemory int
	// Maximum program size; in bytes
	MaxSize int
	// Maximum execution time in seconds
	MaxTime int
	// Suggestive value for optimization
	ComputeToMemoryRatio float64
}

type TargetFunctionType any

type E struct {
	UUID                   models.CandidateID `json:"uuid"`
	AssertionResults       []bool             `json:"assertion_results"`
	AssertionErrorDistance []float64          `json:"assertion_error_distances"`
	ExecTime               time.Duration      `json:"exec_time"`
	Panicked               bool               `json:"panicked"`
	Config                 Config             `json:"-"`
}

func NewE(canditateUUID models.CandidateID) *E {
	e := E{
		UUID: canditateUUID,
	}
	return &e
}

func (e *E) SetConfig(config Config) {
	e.Config = config
}

func (e *E) AssertEqual(output, want any) {
	result := output == want
	e.AssertionResults = append(e.AssertionResults, result)

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
	e.AssertionErrorDistance = append(e.AssertionErrorDistance, errorDistance)
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

func (e *E) Export() {
	f, err := os.Create("results.json")
	if err != nil {
		panic(errors.Wrap(err, "Could not create 'result.json' for writing"))
	}
	json.NewEncoder(f).Encode(e)
}
