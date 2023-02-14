package genetics

import (
	"bytes"
	"tde/internal/astw/clone"
	ast_utl "tde/internal/astw/utilities"
	"tde/internal/cfg"
	"tde/internal/evaluation"

	"fmt"
	"go/ast"
	"go/token"
	"testing"

	"github.com/pkg/errors"
)

func loadTestPackage() (*ast.Package, *ast.File, *ast.FuncDecl, error) {
	_, astPkgs, err := ast_utl.LoadDir("../test_package")
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "could not load test package")
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["../test_package/walk.go"]
	funcDecl, err := ast_utl.FindFuncDecl(astPkg, "WalkWithNils")
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "could not find test function")
	}
	return astPkg, astFile, funcDecl, nil
}

func Test_DevelopWithoutSyntaxError(t *testing.T) {
	astPkg, astFile, funcDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on prep"))
	}

	bufferOrg := bytes.NewBuffer([]byte{})
	ast.Fprint(bufferOrg, token.NewFileSet(), funcDecl, nil)

	err = cfg.Develop(astPkg, astFile, funcDecl, 2)
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on call to Develop"))
	}

	if !evaluation.SyntaxCheckUnsafe(funcDecl) {
		t.Error("fails on SyntaxCheck")
	}
}

func Benchmark_DevelopFindUnbreakingChange(b *testing.B) {
	astPkg, astFile, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		b.Error(errors.Wrapf(err, "failed on prep"))
	}

	// nonBreakingCandidates := []*ast.FuncDecl{}
	// candidates := []*ast.FuncDecl{}
	var best *ast.FuncDecl = originalFuncDecl

	pool := []*ast.FuncDecl{}

	for i := 0; i < 10000; i++ {
		pool = append(pool, clone.FuncDecl(originalFuncDecl))
	}

	for gen := 0; gen < 20; gen++ {

		for _, candidate := range pool {
			cfg.Develop(astPkg, astFile, candidate, 1)
		}

		// sort
		for idx, candidate := range pool {
			if ok, _ := evaluation.SyntaxCheck(candidate); ok {
				best = candidate
				fmt.Println("non-breaking change on gen-idx", gen, idx)

				bufferNew := bytes.NewBuffer([]byte{})
				ast.Fprint(bufferNew, token.NewFileSet(), candidate, nil)
				pool = append(append([]*ast.FuncDecl{pool[idx]}, pool[:idx]...), pool[idx+1:]...)
			}
		}

		fmt.Println("gen:", gen)

		for i := 0; i < 1000000; i++ {
			candidate := clone.FuncDecl(best)
			cfg.Develop(astPkg, astFile, candidate, 1)

			if ok, _ := evaluation.SyntaxCheck(candidate); ok {
				best = candidate
				fmt.Println("non-breaking change on #", i)

				bufferNew := bytes.NewBuffer([]byte{})
				ast.Fprint(bufferNew, token.NewFileSet(), candidate, nil)
			} else {
				// fmt.Println("Panic:", msg)
				// break
			}
		}
	}

	// for i := 0; i < 100; i++ {
	// 	candidates = append(candidates, clone.FuncDecl(originalFuncDecl))
	// }

	// for gen := 0; gen < 200000; gen++ {
	// 	if gen%100 == 0 {
	// 		fmt.Println("gen:", gen)
	// 	}
	// 	for candidateId, candidate := range candidates {
	// 		err = cfg.Develop(astPkg, astFile, candidate, 1)
	// 		if err != nil {
	// 			b.Error(errors.Wrapf(err, "failed on Develop"))
	// 		}
	// 		if evaluation.SyntaxCheck(candidate) {
	// 			fmt.Println("non-breaking change on #", candidateId)
	// 			nonBreakingCandidates = append(nonBreakingCandidates, candidate)
	// 		}
	// 	}
	// }

	// if len(nonBreakingCandidates) == 0 {
	// 	b.Error("No non-breaking candidates found.")
	// }
}
