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
		"00000000-0000-0000-000000000002": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000003": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000004": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000005": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000010": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000011": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000012": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000013": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000014": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000015": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000020": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000021": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000022": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000023": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000024": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000025": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000030": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000031": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000032": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000033": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000034": func(s string) string { return "wordreverse.WordReverse" },
		"00000000-0000-0000-000000000035": func(s string) string { return "wordreverse.WordReverse" },
	})
	// fmt.Println(len(e.TestCandidates))

	wordreverse.TDE_WordReverse(e)
	// fmt.Println(len(e.TestCandidates))
	e.Export()
}
