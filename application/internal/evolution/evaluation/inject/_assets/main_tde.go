//go:build tde

package main

import (
	targetPackage "{{ .TargetPackageImportPath }}"
	"tde/internal/evolution/models"
	"tde/pkg/testing"

	"flag"
)

var subjectUUID string

func init() {
	flag.StringVar(
		&subjectUUID,
		"subject-uuid",
		"",
		"Tested subject's identifier. The program will use this only for output and it is not essential for program to run.",
	)
}

func main() {
	flag.Parse()
	var (
		testFunction  = targetPackage.{{ .TestFunctionName }}
		subjectUUID = models.Sid(subjectUUID)
		t			  = testing.NewT(subjectUUID)
	)
	testFunction(t)
	t.Export()
}