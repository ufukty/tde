package main

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Wrong number of arguments. First argument should be the name of input file.")
	}
	file, err := parseFile(os.Args[1])
	if err != nil {
		log.Fatalln(errors.Wrapf(err, "Could not read the file: '%s'", os.Args[1]))
	}

	var (
		packageName            = file.Name
		reqStructs, resStructs = discoverFileForStructDefinitions(file)
		genDecl                = templateImports()
	)

	if len(reqStructs) == 0 && len(resStructs) == 0 {
		fmt.Println("No type definition found in the file which ends with either Request or Response.")
		return
	}

	var genFile = &ast.File{
		Name:  packageName,
		Decls: []ast.Decl{genDecl},
	}

	for _, entity := range reqStructs {
		genFile.Decls = append(genFile.Decls,
			templateNewRequest(entity),
			templateParseRequest(entity),
		)
	}

	for _, entity := range resStructs {
		genFile.Decls = append(genFile.Decls,
			templateSerializeIntoResponseWriter(entity),
			templateDeserializeResponse(entity),
		)
	}

	newFileName := createNewFileName(os.Args[1])
	target, err := os.Create(newFileName)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not create file in current directory"))
	}
	err = printer.Fprint(target, token.NewFileSet(), genFile)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not write into newly created file"))
	}
}
