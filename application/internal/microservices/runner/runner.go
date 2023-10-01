package runner

import (
	models "tde/models/program"

	"archive/zip"
)

// TODO: copy the target module into /tmp/created_dir

// TODO: duplicate the package as many as the number of subjects

// TODO: embed a test main package

// TODO: parse target test to get the tde config (limit cpu/memory/etc.)

// TODO: main_package/main_tde.go will be filled with information of target_package and tests

// TODO: run the tests with go run tde_package/main_tde.go --flag="tde"

// TODO: main_tde.go will call the tests in each dir; runs in parallel; accounts for resource usage; prints the results as json

// TODO: read main_tde output from os_stdout and parse.

type Request struct {
	GoModule           zip.File
	EvolutionSessionID string
	Subjects           []models.Subject
}

func Handle(subjects []*models.Subject) {

	// copy whole original module
	// copy the file contains target function

}
