package evaluation

import (
	"tde/internal/folders/slots"
	models "tde/models/program"

	"fmt"
	"go/printer"
	"go/token"
	"os"
)

type Evaluator struct {
	SlotManagerSession *slots.Session
}

func NewEvaluator(slotManagerSession *slots.Session) *Evaluator {
	return &Evaluator{
		SlotManagerSession: slotManagerSession,
	}
}

// FIXME: count errors on code creation
func syntaxCheckAndProduceCode(candidates []*models.Candidate) {
	for _, candidate := range candidates {
		fmt.Println(printer.Fprint(os.Stdout, token.NewFileSet(), candidate.AST.File))
		buffer, ok, err := ProduceCodeFromASTSafe(candidate.AST.File) // produce code from ast.File to capture changes in import list too
		if err != nil || !ok {
			candidate.Fitness.AST = 1.0
			continue
		}
		candidate.File = buffer.Bytes()
	}
}

// TODO: Syntax Check
// TODO: Print code (full package vs changed func body ??)
// TODO: Send whole generation into sandboxed environment
// TODO: Get test results
// TODO: Return test results
func (e *Evaluator) Pipeline(candidates []*models.Candidate) {
	syntaxCheckAndProduceCode(candidates)
	e.SlotManagerSession.PlaceCandidatesIntoSlots(candidates)

	// compile each slot
	//
}
