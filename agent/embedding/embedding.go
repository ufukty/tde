package embedding

import (
	"go/parser"
	"models/in_program_models"

	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

//go:embed assets/main_evolve.go
var mainFileContent string

type EmbeddingConfig struct {
	Path    string
	Content []byte
}

func NewEmbeddingConfig(path string) *EmbeddingConfig {
	f := &EmbeddingConfig{
		Path: path,
	}
	f.LoadContent()
	return f
}

func (f *EmbeddingConfig) LoadContent() {
	var err error

	f.Content, err = os.ReadFile(f.Path)
	if err != nil {
		log.Fatalln(errors.Wrap(err, fmt.Sprintf("Could not load the file '%s'", f.Path)))
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (config *EmbeddingConfig) EmbedIntoFile(candidates []in_program_models.Candidate) {

	// // fmt.Println(mainFileContent)

	// d1 := []byte("hello\ngo\n")
	// err := os.WriteFile("/tmp/dat1", d1, 0_644)
	// check(err)

	// // For more granular writes, open a file for writing.
	// f, err := os.Create("/tmp/dat2")
	// check(err)

	// // It’s idiomatic to defer a Close immediately after opening a file.
	// defer f.Close()

	// // You can Write byte slices as you’d expect.
	// d2 := []byte{115, 111, 109, 101, 10}
	// n2, err := f.Write(d2)
	// check(err)
	// fmt.Printf("wrote %d bytes\n", n2)

	// // A WriteString is also available.
	// n3, err := f.WriteString("writes\n")
	// check(err)
	// fmt.Printf("wrote %d bytes\n", n3)

	// // Issue a Sync to flush writes to stable storage.
	// f.Sync()

	// // bufio provides buffered writers in addition to the buffered readers we saw earlier.
	// w := bufio.NewWriter(f)
	// n4, err := w.WriteString("buffered\n")
	// check(err)
	// fmt.Printf("wrote %d bytes\n", n4)

	// // Use Flush to ensure all buffered operations have been applied to the underlying writer.
	// w.Flush()

	// contentToWrite := ""

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

}

func StampFunctionPrototype(original, stamp string) (string, error) {

	originalClosedCurlyBracet := fmt.Sprintf("%s}", original)
	ast_, err := parser.ParseExpr(originalClosedCurlyBracet)
	if err != nil {
		fmt.Printf("%+#v", ast_)
		return "", errors.Wrap(err, "Could not parse original function")
	}

	fmt.Println(ast_)

	return "", nil
}
