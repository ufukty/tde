package main

/*
Usage:
	cd ./example
	gp Knappsack pop=1000 gen=10 bloat=1024

Main progress:
	Discovery -> starting and ending lines of
				 the blocks in the implementation
				 file and test file
	Initialization : Random code segments will be
	      		     created for empty bodies when
					 they are given as input.
	Evolve : Iterates the evolution by 1 generation.
*/

import (
	"log"
)

const (
	n = 1000 // population
	g = 100  // generation limit
	c = 1000 // character limit
)

func main() {
	log.Printf("Parameters: n=%d g=%d c=%d", n, g, c)

	// var e = evolution.Evolution{}
	// e.InitPopulation(n)

	// for i := 0; i < g; i++ {
	// 	e.ExecuteOneGeneration()
	// }
}
