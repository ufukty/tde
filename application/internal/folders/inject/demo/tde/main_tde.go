//go:build tde
// +build tde

package main

import (
	targetPackage "tde/internal/folders/inject/demo"
	models "tde/models/program"
	"tde/pkg/tde"

	"flag"
)

var candidateUUID string

func init() {
	flag.StringVar(
		&candidateUUID,
		"candidate-uuid",
		"",
		"Tested candidate's identifier. The program will use this only for output and it is not essential for program to run.",
	)
}

func main() {
	flag.Parse()
	var (
		testFunction  = targetPackage.TDE_WordReverse
		candidateUUID = models.CandidateID(candidateUUID)
		e             = tde.NewE(candidateUUID)
	)
	testFunction(e)
	e.Export()
}
