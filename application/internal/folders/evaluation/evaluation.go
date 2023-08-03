package evaluation

import (
	"tde/internal/folders/slotmgr"
	models "tde/models/program"

	"fmt"
	"go/printer"
	"go/token"
	"os"
)

type Evaluator struct {
	Sm *slotmgr.SlotManager
}

func NewEvaluator(sm *slotmgr.SlotManager) *Evaluator {
	return &Evaluator{
		Sm: sm,
	}
}

// FIXME: count errors on code creation
func syntaxCheckAndProduceCode(candidates []*models.Candidate) {
	for _, candidate := range candidates {
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
	fmt.Println("printing candidate function bodies")
	syntaxCheckAndProduceCode(candidates)

	fmt.Println("placing candidates into slots")
	e.Sm.PlaceCandidatesIntoSlots(candidates)
	e.Sm.Print()

	fmt.Println("releasing all slots for next generation")
	e.Sm.FreeAllSlots()
}
