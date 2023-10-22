package evaluation

import (
	"fmt"
	models "tde/models/program"
	"testing"
)

func Test_Print(t *testing.T) {
	ctx, err := models.LoadContext("../../..", "testdata", "TDE_WordReverse")
	if err != nil {
		t.Fatal(fmt.Errorf("prep, finding the context for package: %w", err))
	}
	subjects := models.Subjects{}
	subjects.Add(ctx.NewSubject())

	syntaxCheckAndProduceCode(ctx, subjects)

	for i, subject := range subjects {
		if len(subject.Code) == 0 {
			t.Fatal(fmt.Errorf("assert, subjects[%s].File is empty: %w", i, err))
		}
	}
}
