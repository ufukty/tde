package inject

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed _assets/main_tde.go
var mainFileContent string

type TestInfo struct {
	TargetPackageImportPath string
	TestFunctionName        string
}

func createTesterDir(testerPkgPath string) error {
	err := os.MkdirAll(testerPkgPath, 0o_755)
	if err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}
	return nil
}

func prepareTemplateForTesterFile(testInfo *TestInfo) (string, error) {
	buf := bytes.NewBuffer([]byte{})
	templ := template.Must(template.New("").Parse(mainFileContent))
	err := templ.Execute(buf, testInfo)
	if err != nil {
		return "", fmt.Errorf("execute template: %w", err)
	}
	return buf.String(), nil
}

func writeTesterFileContent(testerPkgDir string, content string) error {
	f, err := os.Create(testerPkgDir)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer f.Close()

	_, err = fmt.Fprint(f, content)
	if err != nil {
		return fmt.Errorf("fprintf: %w", err)
	}
	return nil
}

// will create the testedPkgDir/tde/main_tde.go
func Inject(testedPkgDir string, testInfo *TestInfo) error {
	testerPkgPath := filepath.Join(testedPkgDir, "tde")
	testerFilePath := filepath.Join(testedPkgDir, "tde/main_tde.go")

	err := createTesterDir(testerPkgPath)
	if err != nil {
		return fmt.Errorf("tester package dir: %w", err)
	}

	content, err := prepareTemplateForTesterFile(testInfo)
	if err != nil {
		return fmt.Errorf("templating for tester file: %w", err)
	}

	err = writeTesterFileContent(testerFilePath, content)
	if err != nil {
		return fmt.Errorf("writing tester file content: %w", err)
	}

	return nil
}
