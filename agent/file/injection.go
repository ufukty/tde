package file

import (
	"models/in_program_models"

	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

//go:embed assets/main_evolve.go
var mainFileContent string

type File struct {
	Path    string
	Content []byte
}

func NewFile(path string) *File {
	f := &File{
		Path: path,
	}
	f.LoadContent()
	return f
}

func (f *File) LoadContent() {
	var err error

	f.Content, err = os.ReadFile(f.Path)
	if err != nil {
		log.Fatalln(errors.Wrap(err, fmt.Sprintf("Could not load the file '%s'", f.Path)))
	}
}

func (f *File) InjectCandidates(candidates []in_program_models.Candidate) {

	fmt.Println(mainFileContent)
	// os.OpenFile(f.Path, write, 0_666)
	// for _, ind := range candidates {

	// 	_, err := file.Write([]byte(*ind.Program))
	// 	if err != nil {
	// 		errors.Wrap(err, "Could not inject candidate program to compilation file.")
	// 	}
	// }

}
