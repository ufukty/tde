//go:build tde
// +build tde

package main

import (
	"tde/examples/word-reverse/word_reverse"
	models "tde/models/in_program_models"
	"tde/pkg/tde"
)

var candidates = map[models.CandidateID]tde.TargetFunctionType{}

func main() {
	e := tde.NewE(candidates)
	word_reverse.TDE_WordReverse(e)
	e.Export()
}
