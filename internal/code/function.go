package code

import (
	"go/ast"
	"go/printer"
	"tde/internal/utilities"
)

type Function struct {
	Root *ast.FuncDecl
}

func (f *Function) OverwriteBody(content string) {

	
	stringWriter := utilities.NewStringWriter()

	printer.Fprint(stringWriter)

	// [:f.Body.Lbrace] + " " + string(code_fragment.ProduceRandomFragment()) + " " + input[f.Body.Rbrace-1:]
}
