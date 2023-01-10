package embedding

import (
	"tde/internal/code"
	"tde/internal/utilities"
	"text/template"

	_ "embed"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

//go:embed _assets/main_tde.go
var mainFileContent string

type EmbeddingConfig struct {
	ImplementationPackagePath string // eg. .../userPackage
	ImplementationFilePath    string // eg. .../userPackage/userFile.go
	TestFunctionPath          string // eg. .../userPackage/userFile_tde.go
	TestFunctionName          string // eg. TDE_WordReverse
	TargetPackageImportPath   string // eg. userModule/path/to/userPackageDir/userPackage
	TDEPackagePath            string // eg. .../userPackage/tde
	MainFilePath              string // eg. .../userPackage/tde/main_tde.go
}

func NewEmbeddingConfig(implementationPackageFile, implementationFilePath, testFunctionPath, targetPackageImportPath, testFunctionName string) *EmbeddingConfig {
	return &EmbeddingConfig{
		ImplementationPackagePath: implementationPackageFile,
		ImplementationFilePath:    implementationFilePath,
		TestFunctionPath:          testFunctionPath,
		TestFunctionName:          testFunctionName,
		TargetPackageImportPath:   targetPackageImportPath,
		TDEPackagePath:            fmt.Sprintf("%s/tde", implementationPackageFile),
		MainFilePath:              fmt.Sprintf("%s/tde/main_tde.go", implementationPackageFile),
	}
}

func (ec *EmbeddingConfig) ReadImplementationFileContent() error {
	var err error

	// ec.TargetFileContent, err = os.ReadFile(ec.TargetFunctionPath)
	// if err != nil {
	// 	return errors.Wrap(err, fmt.Sprintf("Could not load the file '%s'", ec.TargetFunctionPath))
	// }

	co := code.Code{}
	err = co.LoadFromFile(ec.ImplementationFilePath)
	if err != nil {
		return errors.Wrap(err, "Could not parse the implementation file")
	}
	co.Print(os.Stdout)

	// ast.Print(fset, astFile)

	// ast.Inspect(astFile, func(n ast.Node) bool {

	// 	if n != nil {
	// 		pos := fset.Position(n.Pos()).String()
	// 		posStripped := pos[strings.Index(pos, ":"):]
	// 		fmt.Printf("Line%s ", posStripped)
	// 	}

	// 	fmt.Print(n, reflect.TypeOf(n))

	// 	var s string
	// 	switch x := n.(type) {
	// 	case *ast.BasicLit:
	// 		s = x.Value
	// 		fmt.Print(" BasicLit ", s)
	// 	case *ast.Ident:
	// 		s = x.Name
	// 		fmt.Print(" Ident ", s)
	// 	case *ast.Package:
	// 		s = x.Name
	// 		fmt.Print(" Package ", s)
	// 	}

	// 	fmt.Printf("\n")
	// 	return true
	// })

	// fmt.Println("==========")
	// // Print the imports from the file's AST.
	// for _, s := range astFile.Scope.Objects {
	// 	fmt.Println(s.Name)
	// 	fmt.Printf("%#v\n", s.Name)
	// }
	// fmt.Println("==========")

	// printer.Fprint(os.Stdout, fset, astFile)

	return nil
}

func (ec *EmbeddingConfig) CreateTDEPackage() error {
	err := os.MkdirAll(ec.TDEPackagePath, 0_755)
	if err != nil {
		return errors.Wrap(err, "could not create 'tde' package folder")
	}
	return nil
}

func PrepareTemplateForMainFile(targetPackageImportPath, testFunctionName, candidateID string) (string, error) {
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

func (ec *EmbeddingConfig) WriteMainFile() error {
	var (
		f   *os.File
		err error
	)
	defer f.Close()

	f, err = os.Create(ec.MainFilePath)
	if err != nil {
		return errors.Wrap(err, "could not place file 'main_tde.go'")
	}
	str, err := PrepareTemplateForMainFile(ec.TargetPackageImportPath, ec.TestFunctionName, "00000000-0000-0000-0000-000000000002")
	if err != nil {
		return errors.Wrap(err, "PrepareTemplateForMainFile() is failed for 'main_tde.go'")
	}
	_, err = fmt.Fprint(f, str)
	if err != nil {
		return errors.Wrap(err, "could not write into file 'main_tde.go'")
	}
	return nil
}

func (ec *EmbeddingConfig) Embed() error {
	var err error

	// create "tde" package inside target package
	err = ec.CreateTDEPackage()
	if err != nil {
		return errors.Wrap(err, "could not create tde package")
	}

	// copy main_tde.go into "tde" package
	err = ec.WriteMainFile()
	if err != nil {
		return errors.Wrap(err, "Could not create tester file")
	}

	// read target file for:
	//    - target function prototype (receiver/method)
	//    - package
	err = ec.ReadImplementationFileContent()
	if err != nil {
		return errors.Wrap(err, "Could not read implementation file")
	}

	return nil
}
