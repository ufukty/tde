package evaluation

import (
	"fmt"
	"tde/internal/evolution/evaluation/slotmgr"
	"tde/internal/evolution/models"
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

func (e *Evaluator) mount(subject *models.Subject) error {
	if err := e.sm.PlaceSubjectsIntoSlots(models.Subjects{subject.Sid: subject}); err != nil {
		return fmt.Errorf("calling PlaceSubjectsIntoSlots: %w", err)
	}
	return nil
}

func (e *Evaluator) unmount(subject *models.Subject) error {
	if err := e.sm.FreeAllSlots(); err != nil {
		return fmt.Errorf("calling FreeAllSlots: %w", err)
	}
	return nil
}

// TODO: Syntax Check
// TODO: Print code (full package vs changed func body ??)
// TODO: Send whole generation into sandboxed environment
// TODO: Get test results
// TODO: Return test results
// TODO: skip subjects its fitness already set

func (e *Evaluator) pipeline(subj *models.Subject) (err error) {
	if err = e.print(subj); err != nil {
		err = fmt.Errorf("printing: %w", err)
		return
	}
	if !subj.IsValidIn(models.Code) {
		return nil
	}
	if err = e.mount(subj); err != nil {
		err = fmt.Errorf("mounting the subject: %w", err)
		return
	}
	defer func() {
		if err != nil {
			if derr := e.unmount(subj); derr != nil {
				err = fmt.Errorf("%s\n\n(also) unmounting the subject: %w", err.Error(), derr)
			}
		}
	}()
	if err = e.compile(subj); err != nil {
		err = fmt.Errorf("compiling: %w", err)
		return
	}
	if !subj.IsValidIn(models.Program) {
		return nil
	}
	if err = e.test(subj); err != nil {
		err = fmt.Errorf("testing: %w", err)
		return
	}
	return
}

func (e *Evaluator) Pipeline(subjects models.Subjects) error {
	defer e.ctx.Restore()
	for _, subj := range subjects {
		e.pipeline(subj)
	}
	return nil
}
