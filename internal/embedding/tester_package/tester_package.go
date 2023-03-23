package tester_package

import (
	"tde/internal/embedding/models"
	"tde/internal/utilities"

	_ "embed"
	"fmt"
	"os"
	"text/template"

	"github.com/pkg/errors"
)

//go:embed _assets/main_tde.go
var mainFileContent string

func createTDEPackage(ec *models.EmbeddingConfig) error {
	err := os.MkdirAll(ec.TesterHeadDirPath, 0_755)
	if err != nil {
		return errors.Wrap(err, "could not create 'tde' package folder")
	}
	return nil
}

func prepareTemplateForMainFile(targetPackageImportPath, testFunctionName, candidateID string) (string, error) {
	sw := utilities.NewStringWriter()
	templ := template.Must(template.New("").Parse(mainFileContent))
	err := templ.Execute(sw, struct {
		TargetPackageImportPath string
		TestFunctionName        string
		CandidateID             string
	}{targetPackageImportPath, testFunctionName, candidateID})
	if err != nil {
		return "", errors.Wrap(err, "could not execute the template 'main_tde.go' with given details")
	}
	return sw.String(), nil
}

func writeMainFile(config *models.EmbeddingConfig) error {
	var (
		f   *os.File
		err error
	)
	defer f.Close()

	f, err = os.Create(config.TesterHeadFilePath)
	if err != nil {
		return errors.Wrap(err, "could not place file 'main_tde.go'")
	}
	str, err := prepareTemplateForMainFile(config.PackageImportPath, config.TestFunctionName, "00000000-0000-0000-0000-000000000002")
	if err != nil {
		return errors.Wrap(err, "PrepareTemplateForMainFile() is failed for 'main_tde.go'")
	}
	_, err = fmt.Fprint(f, str)
	if err != nil {
		return errors.Wrap(err, "could not write into file 'main_tde.go'")
	}
	return nil
}

// if the testedPkgDir
func Inject(config *models.EmbeddingConfig) error {

	err := createTDEPackage(config)
	if err != nil {
		return errors.Wrap(err, "could not create tde package")
	}

	err = writeMainFile(config)
	if err != nil {
		return errors.Wrap(err, "Could not create tester file")
	}

	return nil
}
