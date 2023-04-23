package evolution

import (
	"context"
	"tde/internal/evolution"
)

// TODO: Listen endpoints for:
//   TODO: Creating a generation (request from client)
//   TODO: Download test results from runner

func GetParametersForUser() {}

func StartEvolution(input string, test string, populationSize int) { // FIXME: actual input is a package, contains tests and target function headers along with irrelevant code
	e := evolution.Evolution{}
	e.InitPopulation(populationSize)
	e.IterateLoop(context.Background())
}
