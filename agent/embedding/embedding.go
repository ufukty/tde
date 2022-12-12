package embedding

import (
	"bufio"
	"go/parser"
	"go/printer"
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
	ImplementationPackagePath string // .../userPackage
	ImplementationFilePath    string // .../userPackage/userFile.go
	TestFunctionPath          string // .../userPackage/userFile_tde.go
	TDEPackagePath            string // .../userPackage/tde
	MainFilePath              string // .../userPackage/tde/main_tde.go
	CandidatesFilePath        string // .../userPackage/tde/candidates_tde.go
	TargetFileContent         []byte
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
	astFile, err := parser.ParseFile(fset, ec.ImplementationFilePath, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrap(err, "Could not parse the implementation file")
	}

	// fmt.Println("=", SearchFunctionDeclaration(astFile, "WordReverse"))

	fd, err := SearchFunctionDeclaration(astFile, "WordReverse")
	if err != nil {
		errors.Wrap(err, "Could not get the function definition")
	}
	ReplaceFunctionName(fd, "Kilimcinin")
	printer.Fprint(os.Stdout, fset, fd)

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

func (ec *EmbeddingConfig) CreateMainPackage() error {
	os.Mkdir(ec.TDEPackagePath)
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
	fmt.Fprint(f, s) // 

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
	err = ec.CreateMainPackage()

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
	err = ec.WriteCandidatesIntoFile([]in_program_models.Candidate{{
		UUID: "414ecb09-3025-51c9-918a-e5f278e1a0f4",
		Body: []byte(`s := ""
return s`),
	}})
	if err != nil {
		return errors.Wrap(err, "Could not write candidates")
	}

	return nil
}
