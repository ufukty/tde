package evaluation

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

// returns True for valid syntax
func SyntaxCheckSafe(candidate ast.Node) (bool, any) {
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
		if printer.Fprint(bytes.NewBuffer([]byte{}), token.NewFileSet(), candidate) == nil {
			isValid = true
		}
	}()
	return isValid, panicMessage
}

// Compared to SyntaxCheck, this one panics
// returns True for valid syntax.
func SyntaxCheckUnsafe(candidate ast.Node) bool {
	isValid := false
	if printer.Fprint(bytes.NewBuffer([]byte{}), token.NewFileSet(), candidate) == nil {
		isValid = true
	}
	return isValid
}

// TODO: Write right into the target file instead use memory as intermediate
func ProduceCodeFromASTSafe(candidate ast.Node) (*bytes.Buffer, bool, any) {
	var (
		isValid      = true
		panicMessage any
		buffer       = bytes.NewBuffer([]byte{})
	)
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicMessage = r
				isValid = false
			}
		}()

		if printer.Fprint(buffer, token.NewFileSet(), candidate) == nil {
			isValid = true
		}
	}()
	return buffer, isValid, panicMessage
}
