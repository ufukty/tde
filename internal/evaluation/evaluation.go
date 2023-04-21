package evaluation

import (
	"fmt"
	"go/printer"
	"go/token"
	"os"
	"tde/internal/folders/slot_manager"
	models "tde/models/program"
)

type Evaluator struct {
	SlotManagerSession *slot_manager.Session
}

func NewEvaluator(slotManagerSession *slot_manager.Session) *Evaluator {
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
