package evaluation

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	models "tde/models/program"
)

// returns True for valid syntax
func SyntaxCheckSafe(subject ast.Node) (bool, any) {
	var (
		isValid      = true
		panicMessage any
	)
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicMessage = r
				isValid = false
			}
		}()
		if printer.Fprint(bytes.NewBuffer([]byte{}), token.NewFileSet(), subject) == nil {
			isValid = true
		}
	}()
	return isValid, panicMessage
}

// Compared to SyntaxCheck, this one panics
// returns True for valid syntax.
func SyntaxCheckUnsafe(subject ast.Node) bool {
	isValid := false
	if printer.Fprint(bytes.NewBuffer([]byte{}), token.NewFileSet(), subject) == nil {
		isValid = true
	}
	return isValid
}

// TODO: Write right into the target file instead use memory as intermediate
func ProduceCodeFromASTSafe(context *models.Context, subject *ast.FuncDecl) (*bytes.Buffer, bool, any) {
	var (
		isValid      = true
		panicMessage any
		buffer       = bytes.NewBuffer([]byte{})
	)
	context.Swap(subject)
	defer context.Restore()
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicMessage = r
				isValid = false
			}
		}()

		if printer.Fprint(buffer, token.NewFileSet(), context.File) == nil {
			isValid = true
		}
	}()
	return buffer, isValid, panicMessage
}
