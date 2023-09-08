package evaluation

import (
	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone/clean"
	"tde/internal/folders/discovery"
	"tde/internal/folders/inject"
	"tde/internal/folders/list"
	"tde/internal/folders/slotmgr"
	models "tde/models/program"

	"fmt"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
)

const ( // applies to all test cases in this package
	MODULEROOT  = "../../../"
	ORIGINALPKG = "../../../examples/words"
	PKGINMOD    = "examples/words"
	TESTNAME    = "TDE_WordReverse"
	TARGETNAME  = "WordReverse"
	POPULATION  = 3
)

func prepare() (*Evaluator, []*models.Candidate, error) {
	orgPkgAst, err := astwutl.LoadPackageFromDir(ORIGINALPKG)
	if err != nil {
		return nil, nil, fmt.Errorf("prep: %w", err)
	}

	var candidates = []*models.Candidate{}
	for i := 0; i < POPULATION; i++ { //
		pkgWorkingCopy := clean.Package(orgPkgAst)
		file, funcdecl, err := astwutl.FindFuncDeclInPkg(pkgWorkingCopy, TARGETNAME)
		if err != nil {
			return nil, nil, fmt.Errorf("locating the func in the working-copy of original target package: %w", err)
		}
		var candidate = &models.Candidate{
			UUID: models.CandidateID(uuid.New().String()),
			AST: models.TargetAst{
				Package:  pkgWorkingCopy,
				File:     file,
				FuncDecl: funcdecl,
			},
		}
		candidates = append(candidates, candidate)
	}

	pkgs, err := list.ListPackagesInDir(ORIGINALPKG)
	if err != nil {
		return nil, nil, fmt.Errorf("listing packages in target dir: %w", err)
	}
	sample, err := inject.WithCreatingSample(MODULEROOT, pkgs.First(), TESTNAME)
	if err != nil {
		return nil, nil, fmt.Errorf("creating sample: %w", err)
	}
	combined, err := discovery.CombinedDetailsForTest(filepath.Join(sample, PKGINMOD), TESTNAME)
	if err != nil {
		return nil, nil, fmt.Errorf("retriving combined details: %w", err)
	}
	sm := slotmgr.New(sample, combined)
	if err := sm.PlaceCandidatesIntoSlots(candidates); err != nil {
		return nil, nil, fmt.Errorf("passing candidates into slot manager: %w", err)
	}
	syntaxCheckAndProduceCode(candidates)
	evaluator := NewEvaluator(sm)
	return evaluator, candidates, nil
}

// FIXME: check fitness has populated after syntax errors
func Test_SyntaxCheckAndProduceCode(t *testing.T) {
	_, candidates, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	for i, candidate := range candidates {
		if len(candidate.File) == 0 {
			t.Fatal(fmt.Errorf("assert, candidates[%d].File is empty: %w", i, err))
		}
	}
}

func Test_Compile(t *testing.T) {
	evaluator, candidates, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	if err := evaluator.runCandidates(candidates); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}
}

func Test_Pipeline(t *testing.T) {
	evaluator, candidates, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	evaluator.sm.Print()

	if err := evaluator.Pipeline(candidates); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}
}
