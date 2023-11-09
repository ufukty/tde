package evolution

import (
	"fmt"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/evaluation/inject"
	"tde/internal/evolution/evaluation/list"
	"tde/internal/evolution/evaluation/slotmgr"
	"tde/internal/evolution/pool"
	models "tde/models/program"
	"testing"
)

func prepareCodeSearch() (*codeSearch, error) {
	pkgs, err := list.ListPackagesInDir("testdata/words")
	if err != nil {
		return nil, fmt.Errorf("listing packages in the testdata package: %w", err)
	}
	sample, err := inject.WithCreatingSample(pkgs.First().Module.Dir, pkgs.First(), "TDE_WordReverse")
	if err != nil {
		return nil, fmt.Errorf("creating the sample (injected one): %w", err)
	}
	sm := slotmgr.New(sample, pkgs.First().PathInModule(), "words.go")
	ctx, err := models.LoadContext(pkgs.First().Module.Dir, "testdata/words", "WordReverse")
	if err != nil {
		return nil, fmt.Errorf("loading the context: %w", err)
	}
	ev := evaluation.NewEvaluator(sm, ctx)
	subj := ctx.NewSubject()
	subj.AST = examples[models.AST][0]
	cmns := &commons{
		Evaluator: ev,
		Params:    defaults,
		Context:   ctx,
		Pool:      pool.New(subj),
	}
	em := newCodeSearch(cmns, subj)
	return em, nil
}

func Test_CodeSearch(t *testing.T) {
	cs, err := prepareCodeSearch()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	for !cs.IsEnded() {
		products, err := cs.Iterate()
		if err != nil {
			t.Fatal(fmt.Errorf("act: %w", err))
		}

		for _, subj := range products {
			fmt.Println(subj.Code)
		}
	}
}

func Test_CodeSearch_Probabilistic(t *testing.T) {

}
