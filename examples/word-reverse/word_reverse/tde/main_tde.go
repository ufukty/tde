//go:build tde
// +build tde

package main

import (
	wordreverse "GoGP/examples/word-reverse/word_reverse"
	"GoGP/testing/evolution"
	"models/in_program_models"
)

func main() {
	e := evolution.NewE(map[in_program_models.CandidateID]evolution.TargetFunctionType{
		"00000000-0000-0000-000000000000": func(s string) string {
			r := ""
			for i := len(s) - 1; i >= 0; i-- {
				r += string(s[i])
			}
			return r
		},

		"00000000-0000-0000-000000000001": func(s string) string {
			if s == "Hello world" {
				panic("panic message")
			}
			return ""
		},
		"00000000-0000-0000-000000000002": wordreverse.WordReverse,
		"00000000-0000-0000-000000000003": wordreverse.WordReverse,
		"00000000-0000-0000-000000000004": wordreverse.WordReverse,
		"00000000-0000-0000-000000000005": wordreverse.WordReverse,
		"00000000-0000-0000-000000000010": wordreverse.WordReverse,
		"00000000-0000-0000-000000000011": wordreverse.WordReverse,
		"00000000-0000-0000-000000000012": wordreverse.WordReverse,
		"00000000-0000-0000-000000000013": wordreverse.WordReverse,
		"00000000-0000-0000-000000000014": wordreverse.WordReverse,
		"00000000-0000-0000-000000000015": wordreverse.WordReverse,
		"00000000-0000-0000-000000000020": wordreverse.WordReverse,
		"00000000-0000-0000-000000000021": wordreverse.WordReverse,
		"00000000-0000-0000-000000000022": wordreverse.WordReverse,
		"00000000-0000-0000-000000000023": wordreverse.WordReverse,
		"00000000-0000-0000-000000000024": wordreverse.WordReverse,
		"00000000-0000-0000-000000000025": wordreverse.WordReverse,
		"00000000-0000-0000-000000000030": wordreverse.WordReverse,
		"00000000-0000-0000-000000000031": wordreverse.WordReverse,
		"00000000-0000-0000-000000000032": wordreverse.WordReverse,
		"00000000-0000-0000-000000000033": wordreverse.WordReverse,
		"00000000-0000-0000-000000000034": wordreverse.WordReverse,
		"00000000-0000-0000-000000000035": wordreverse.WordReverse,
	})
	// fmt.Println(len(e.TestCandidates))

	wordreverse.TDE_WordReverse(e)
	// fmt.Println(len(e.TestCandidates))
	e.Export()
}
