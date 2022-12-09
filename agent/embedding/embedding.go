package embedding

import (
	"bufio"
	"go/parser"
	"go/token"
	"models/in_program_models"
	"os"

	_ "embed"
	"fmt"

	"github.com/pkg/errors"
)

//go:embed assets/main_tde.go
var mainFileContent string

type EmbeddingConfig struct {
	ImplementationPackagePath string
	ImplementationFilePath    string
	TestFunctionPath          string

	TDEPackagePath     string
	MainFilePath       string
	CandidatesFilePath string

	TargetFileContent []byte
}

func NewEmbeddingConfig(implementationPackageFile, implementationFilePath, testFunctionPath string) *EmbeddingConfig {
	f := &EmbeddingConfig{
		ImplementationPackagePath: implementationPackageFile,
		ImplementationFilePath:    implementationFilePath,
		TestFunctionPath:          testFunctionPath,
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

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, ec.ImplementationFilePath, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrap(err, "Could not parse the implementation file")
	}

	fmt.Println("==========")
	// Print the imports from the file's AST.
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}
	fmt.Println("==========")

	return nil
}

func (ec *EmbeddingConfig) WriteMainFile() error {
	var (
		f   *os.File
		err error
	)
	defer f.Close()

	f, err = os.Create(ec.MainFilePath)
	if err != nil {
		return errors.Wrap(err, "Could not place file 'main_tde.go'")
	}
	_, err = fmt.Fprint(f, mainFileContent)
	if err != nil {
		return errors.Wrap(err, "Could not write into file 'main_tde.go'")
	}
	return nil
}

func (ec *EmbeddingConfig) WriteCandidatesIntoFile(candidates []in_program_models.Candidate) error {
	var (
		f   *os.File
		err error
	)
	defer f.Close()

	f, err = os.Create(ec.CandidatesFilePath)
	if err != nil {
		return errors.Wrap(err, "Could not create a file for writing candidates")
	}
	// It’s idiomatic to defer a Close immediately after opening a file.

	// You can Write byte slices as you’d expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	if err != nil {
		return errors.Wrap(err, "")
	}
	fmt.Printf("wrote %d bytes\n", n2)

	// A WriteString is also available.
	n3, err := f.WriteString("writes\n")
	if err != nil {
		return errors.Wrap(err, "")
	}
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a Sync to flush writes to stable storage.
	f.Sync()

	// bufio provides buffered writers in addition to the buffered readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	if err != nil {
		return errors.Wrap(err, "")
	}
	fmt.Printf("wrote %d bytes\n", n4)

	// Use Flush to ensure all buffered operations have been applied to the underlying writer.
	w.Flush()

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

	// copy main_tde.go into "tde" package
	err = ec.WriteMainFile()
	if err != nil {
		return errors.Wrap(err, "")
	}

	// read target file for:
	//    - target function prototype (receiver/method)
	//    - package
	err = ec.ReadImplementationFileContent()
	if err != nil {
		errors.Wrap(err, "Could not read target file content")
	}

	// embed candidates into a new file target package
	ec.WriteCandidatesIntoFile([]in_program_models.Candidate{{
		UUID: "414ecb09-3025-51c9-918a-e5f278e1a0f4",
		Body: []byte(`s := ""
return s`),
	}})
	//

	return nil
}
