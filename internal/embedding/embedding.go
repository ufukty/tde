package embedding

import (
	"tde/internal/code"
	"tde/internal/utilities"
	"tde/models/in_program_models"
	"text/template"

	"bytes"
	_ "embed"
	"fmt"
	"go/printer"
	"go/token"
	"os"

	"github.com/pkg/errors"
)

//go:embed assets/main_tde.go
var mainFileContent string

type EmbeddingConfig struct {
	ImplementationPackagePath string // .../userPackage
	ImplementationFilePath    string // .../userPackage/userFile.go
	TestFunctionPath          string // .../userPackage/userFile_tde.go
	TargetPackageImportPath   string // userModule/path/to/userPackageDir/userPackage
	TDEPackagePath            string // .../userPackage/tde
	MainFilePath              string // .../userPackage/tde/main_tde.go
	CandidatesFilePath        string // .../userPackage/tde/candidates_tde.go
	TargetFileContent         []byte
}

func NewEmbeddingConfig(implementationPackageFile, implementationFilePath, testFunctionPath, targetPackageImportPath string) *EmbeddingConfig {
	f := &EmbeddingConfig{
		ImplementationPackagePath: implementationPackageFile,
		ImplementationFilePath:    implementationFilePath,
		TestFunctionPath:          testFunctionPath,
		TargetPackageImportPath:   targetPackageImportPath,
		TDEPackagePath:            fmt.Sprintf("%s/tde", implementationPackageFile),
		MainFilePath:              fmt.Sprintf("%s/tde/main_tde.go", implementationPackageFile),
		CandidatesFilePath:        fmt.Sprintf("%s/tde/candidates.go", implementationPackageFile),
	}
	return f
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
	err := os.Mkdir(ec.TDEPackagePath, 0_755)
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
	_, err = fmt.Fprint(f, mainFileContent)
	if err != nil {
		return errors.Wrap(err, "could not write into file 'main_tde.go'")
	}
	return nil
}

func (ec *EmbeddingConfig) WriteCandidatesIntoFile(candidates []*in_program_models.Candidate) error {
	var (
		f   *os.File
		err error
	)
	defer f.Close()

	f, err = os.Create(ec.CandidatesFilePath)
	if err != nil {
		return errors.Wrap(err, "Could not create a file for writing candidates")
	}
	fset := token.NewFileSet()
	buf := bytes.NewBuffer([]byte{})

	for _, cand := range candidates {

		// d := ast.File{}

		// funcName := fmt.Sprintf("candidate%s", strings.ReplaceAll(string(cand.UUID), "-", ""))

		printer.Fprint(buf, fset, cand) // printer to buffer
	}

	s := buf.String() // buffer to string
	fmt.Fprint(f, s)  // string to file

	// 	fileHeader := `//go:build tde
	// package main`

	// // os.OpenFile(f.Path, write, 0_666)
	// for _, candidate := range candidates {

	// 	var (
	// 		functionOriginalPrototype = "func (config *EmbeddingConfig) EmbedIntoFile(candidates []in_program_models.Candidate)"
	// 		functionEditedPrototype   = "func ()"
	// 		functionNameSafe          = strings.ReplaceAll(string(candidate.UUID), "-", "_")
	// 		functionParameterList     = ""
	// 		function                  = fmt.Sprintf("func %sCandidate_%s(%s) {%s}}\n\n", functionNameSafe, functionParameterList, candidate.Body)
	// 	)
	// 	_, err := file.Write([]byte(*candidate.Program))
	// 	if err != nil {
	// 		errors.Wrap(err, "Could not inject candidate program to compilation file.")
	// 	}
	// }

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

	// embed candidates into a new file target package
	err = ec.WriteCandidatesIntoFile([]*in_program_models.Candidate{{
		UUID: "414ecb09-3025-51c9-918a-e5f278e1a0f4",
		Body: []byte(`s := ""
return s`),
	}})
	if err != nil {
		return errors.Wrap(err, "Could not write candidates")
	}

	return nil
}
