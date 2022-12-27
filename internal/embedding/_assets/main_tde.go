//go:build tde
// +build tde

package main

import (
	targetPackage "{{ .TargetPackageImportPath }}"
	"tde/models/in_program_models"
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
		testFunction  = targetPackage.{{ .TestFunctionName }}
		candidateUUID = in_program_models.CandidateID(candidateUUID)
		e 			  = tde.NewE(candidateUUID)
	)
	testFunction(e)
	e.Export()
}