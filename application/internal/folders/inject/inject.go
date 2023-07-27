package inject

import (
	"tde/internal/utilities"

	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/pkg/errors"
)

//go:embed _assets/main_tde.go
var mainFileContent string

type TestInfo struct {
	TargetPackageImportPath string
	TestFunctionName        string
}

func createTesterDir(testerPkgPath string) error {
	err := os.MkdirAll(testerPkgPath, 0_755)
	if err != nil {
		return errors.Wrap(err, "mkdir")
	}
	return nil
}

func prepareTemplateForTesterFile(testInfo *TestInfo) (string, error) {
	sw := utilities.NewStringWriter()
	templ := template.Must(template.New("").Parse(mainFileContent))
	err := templ.Execute(sw, testInfo)
	if err != nil {
		return "", errors.Wrap(err, "execute template")
	}
	return sw.String(), nil
}

func writeTesterFileContent(testerPkgDir string, content string) error {
	f, err := os.Create(testerPkgDir)
	if err != nil {
		return errors.Wrap(err, "create file")
	}
	defer f.Close()

	_, err = fmt.Fprint(f, content)
	if err != nil {
		return errors.Wrap(err, "fprintf")
	}
	return nil
}

// will create the testedPkgDir/tde/main_tde.go
func Inject(testedPkgDir string, testInfo *TestInfo) error {
	testerPkgPath := filepath.Join(testedPkgDir, "tde")
	testerFilePath := filepath.Join(testedPkgDir, "tde/main_tde.go")

	err := createTesterDir(testerPkgPath)
	if err != nil {
		return errors.Wrap(err, "tester package dir")
	}

	content, err := prepareTemplateForTesterFile(testInfo)
	if err != nil {
		return errors.Wrap(err, "templating for tester file")
	}

	err = writeTesterFileContent(testerFilePath, content)
	if err != nil {
		return errors.Wrap(err, "writing tester file content")
	}

	return nil
}
