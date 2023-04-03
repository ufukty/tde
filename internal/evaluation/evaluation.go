package evaluation

import (
	"tde/internal/folders/slot_manager"
	"tde/models/in_program_models"
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
func syntaxCheckAndProduceCode(candidates []*in_program_models.Candidate) {
	for _, candidate := range candidates {
		code, ok, err := ProduceCodeFromASTSafe(candidate.AST.File) // produce code from ast.File to capture changes in import list too
		if err != nil || !ok {
			candidate.Fitness.AST = 1.0
		} else {
			candidate.File = code
		}
	}
}

// func sendCandidatesToRunnerAndGetResults(evolution *in_program_models.EvolverParameters, candidates []*in_program_models.Candidate) {

// }

// type dsds struct {
// 	originalPkg zip.File
// 	candidates  map[in_program_models.CandidateID][]byte
// }

// TODO: Syntax Check
// TODO: Print code (full package vs changed func body ??)
// TODO: Send whole generation into sandboxed environment
// TODO: Get test results
// TODO: Return test results
func (e *Evaluator) Pipeline(candidates []*in_program_models.Candidate) {
	syntaxCheckAndProduceCode(candidates)
	e.SlotManagerSession.PlaceCandidatesIntoSlots(candidates)

	// compile each slot
	// 
}
