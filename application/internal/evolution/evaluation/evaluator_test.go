package evaluation

import (
	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone/clean"
	"tde/internal/evolution/evaluation/discovery"
	"tde/internal/evolution/evaluation/inject"
	"tde/internal/evolution/evaluation/list"
	"tde/internal/evolution/evaluation/slotmgr"
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

func prepare() (*Evaluator, []*models.Subject, error) {
	orgPkgAst, err := astwutl.LoadPackageFromDir(ORIGINALPKG)
	if err != nil {
		return nil, nil, fmt.Errorf("prep: %w", err)
	}

	var subjects = []*models.Subject{}
	for i := 0; i < POPULATION; i++ { //
		pkgWorkingCopy := clean.Package(orgPkgAst)
		file, funcdecl, err := astwutl.FindFuncDeclInPkg(pkgWorkingCopy, TARGETNAME)
		if err != nil {
			return nil, nil, fmt.Errorf("locating the func in the working-copy of original target package: %w", err)
		}
		var subject = &models.Subject{
			Sid: models.Sid(uuid.New().String()),
			AST: models.TargetAst{
				Package:  pkgWorkingCopy,
				File:     file,
				FuncDecl: funcdecl,
			},
		}
		subjects = append(subjects, subject)
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
	if err := sm.PlaceSubjectsIntoSlots(subjects); err != nil {
		return nil, nil, fmt.Errorf("passing subjects into slot manager: %w", err)
	}
	syntaxCheckAndProduceCode(subjects)
	evaluator := NewEvaluator(sm)
	return evaluator, subjects, nil
}

// FIXME: check fitness has populated after syntax errors
func Test_SyntaxCheckAndProduceCode(t *testing.T) {
	_, subjects, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	for i, subject := range subjects {
		if len(subject.File) == 0 {
			t.Fatal(fmt.Errorf("assert, subjects[%d].File is empty: %w", i, err))
		}
	}
}

func Test_Compile(t *testing.T) {
	evaluator, subjects, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	if err := evaluator.runSubjects(subjects); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}
}

func Test_Pipeline(t *testing.T) {
	evaluator, subjects, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	evaluator.sm.Print()

	if err := evaluator.Pipeline(subjects); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}
}
