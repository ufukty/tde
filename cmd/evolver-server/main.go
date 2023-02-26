package main

import (
	"context"
	"tde/internal/evolution"
	"tde/internal/server"
)

func GetParametersForUser() {}

func StartEvolution(input string, test string, populationSize int) { // FIXME: actual input is a package, contains tests and target function headers along with irrelevant code
	e := evolution.Evolution{}
	e.InitPopulation(populationSize)
	e.IterateLoop(context.Background())
}

func main() {

	server.NewServer(6000)

	// t := Testing{}

	// provisionEnvironmentImage [x]
	// run tests

	// fitness := t.Calculate()
	// fmt.Println(fitness)

}
