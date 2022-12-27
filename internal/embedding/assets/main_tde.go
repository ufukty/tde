//go:build tde
// +build tde

package main

import (
	"{{ .TargetPackageImportPath }}"
	"tde/models/in_program_models"
	"tde/pkg/tde"
)

func main() {
	var (
		testFunction  = "{{ .TestFunctionName }}"
		candidateUUID = in_program_models.CandidateID("{{ .CandidateID }}")
		e 			  = tde.NewE(candidateUUID)
	)
	testFunction(e)
	e.Export()
}