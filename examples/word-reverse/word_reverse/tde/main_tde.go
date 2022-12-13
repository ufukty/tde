//go:build tde
// +build tde

package main

import (
	"tde/examples/word-reverse/word_reverse"
	models "tde/models/in_program_models"
	"tde/pkg/evolution"
)

var candidates = map[models.CandidateID]evolution.TargetFunctionType{}

func main() {
	e := evolution.NewE(candidates)
	word_reverse.TDE_WordReverse(e)
	e.Export()
}
