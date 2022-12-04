//go:build tde

package main

import (
	wordreverse "GoGP/examples/word-reverse"
	"GoGP/testing/evolution"
	"models/in_program_models"

	"fmt"
)

func main() {
	fmt.Println("Hello world")
	defer fmt.Println("Good bye world")

	e := evolution.NewE(map[in_program_models.CandidateID]evolution.TargetFunctionType{
		"00000000-0000-0000-000000000000": wordreverse.WordReverse,
		"00000000-0000-0000-000000000001": wordreverse.WordReverse,
		"00000000-0000-0000-000000000002": wordreverse.WordReverse,
		"00000000-0000-0000-000000000003": wordreverse.WordReverse,
		"00000000-0000-0000-000000000004": wordreverse.WordReverse,
		"00000000-0000-0000-000000000005": wordreverse.WordReverse,
	})
	fmt.Println(e)

	wordreverse.TDE_WordReverse(e)
	e.PrintStats()
}
