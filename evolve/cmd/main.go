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
	"GoGP/evolve/file"
	"fmt"
)

func main() {
	fmt.Println(config)

	file := file.NewFile(config.File)

	fmt.Println(file)

	// var e = evolution.Evolution{}
	// e.InitPopulation(n)

	// for i := 0; i < g; i++ {
	// 	e.ExecuteOneGeneration()
	// }
}
