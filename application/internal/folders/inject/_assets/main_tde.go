//go:build tde

package main

import (
	targetPackage "{{ .TargetPackageImportPath }}"
	models "tde/models/program"
	"tde/pkg/testing"

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
		candidateUUID = models.CandidateID(candidateUUID)
		t			  = testing.NewT(candidateUUID)
	)
	testFunction(t)
	t.Export()
}