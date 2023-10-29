package evaluation

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	models "tde/models/program"
	"time"
)

// // returns True for valid syntax
// func SyntaxCheckSafe(context *models.Context, subject *ast.FuncDecl) (bool, any) {
// 	var (
// 		isValid      = true
// 		panicMessage any
// 	)
// 	context.Swap(subject)
// 	defer context.Restore()
// 	func() {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				panicMessage = r
// 				isValid = false
// 			}
// 		}()
// 		if printer.Fprint(bytes.NewBuffer([]byte{}), token.NewFileSet(), subject) == nil {
// 			isValid = true
// 		}
// 	}()
// 	return isValid, panicMessage
// }

// // FIXME: count errors on code creation
// // TODO: ...and populate fitness component for it
// func syntaxCheckAndProduceCode(context *models.Context, subjects models.Subjects) {
// 	for _, subject := range subjects {
// 		buffer, ok, err := ProduceCodeFromASTSafe(context, subject.AST) // produce code from ast.File to capture changes in import list too
// 		// FIXME:
// 		if err != nil || !ok {
// 			subject.Fitness.AST = 1.0
// 			continue
// 		}
// 		subject.Code = buffer.Bytes()
// 	}
// }

func fprintSafe(file *ast.File) (code []byte, err error) {
	code = []byte{}
	buffer := bytes.NewBuffer([]byte{})
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %q", r)
		}
	}()
	if printer.Fprint(buffer, token.NewFileSet(), file) != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (e *Evaluator) print(subj *models.Subject) error {
	if subj.AST == nil {
		return fmt.Errorf("Empty AST")
	}
	start := time.Now()
	e.ctx.Swap(subj.AST)
	code, err := fprintSafe(e.ctx.File)
	if err != nil {
		dur := time.Now().Sub(start).Nanoseconds()
		inv := 1.0 / float64(max(dur, 1))
		subj.Fitness.AST = inv
	} else {
		subj.Fitness.AST = 0.0
		subj.Code = code
	}
	return nil
}
