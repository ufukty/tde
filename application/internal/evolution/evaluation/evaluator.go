package evaluation

import (
	"fmt"
	"tde/internal/evolution/evaluation/slotmgr"
	models "tde/models/program"
)

type Evaluator struct {
	sm  *slotmgr.SlotManager
	ctx *models.Context
}

func NewEvaluator(sm *slotmgr.SlotManager, ctx *models.Context) *Evaluator {
	return &Evaluator{
		sm:  sm,
		ctx: ctx,
	}
}

// TODO: Syntax Check
// TODO: Print code (full package vs changed func body ??)
// TODO: Send whole generation into sandboxed environment
// TODO: Get test results
// TODO: Return test results
// TODO: skip subjects its fitness already set
func (e *Evaluator) Pipeline(subjects models.Subjects) error {
	for _, subj := range subjects {
		if err := e.print(subj); err != nil {
			return fmt.Errorf("printing: %w", err)
		}
		if !subj.IsValidIn(models.Code) {
			continue
		}
		if err := e.sm.PlaceSubjectsIntoSlots(models.SubjectsFrom([]*models.Subject{subj})); err != nil {
			return fmt.Errorf("placing subjects into slots: %w", err)
		}
		if err := e.compile(subj); err != nil {
			return fmt.Errorf("compiling: %w", err)
		}
		if !subj.IsValidIn(models.Program) {
			continue
		}
		if err := e.test(subj); err != nil {
			return fmt.Errorf("testing: %w", err)
		}
	}
	e.ctx.Restore()
	if err := e.sm.FreeAllSlots(); err != nil {
		return fmt.Errorf("restoring slots for next generation: %w", err)
	}
	return nil
}
