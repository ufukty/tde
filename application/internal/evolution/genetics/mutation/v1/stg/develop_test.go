package stg

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone"
	models1 "tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/evolution/models"
	"tde/internal/utilities"
	"testing"
)

// NOTE: empty branch statement always leads fail in ast->code convertion

type tcase struct {
	ctx    *models.Context
	params *models1.MutationParameters
}

func loadTestCases() (map[string]tcase, error) {
	tcases := map[string]tcase{}
	if ctx, err := models.LoadContext("", "testdata/evolution", "WalkWithNils"); err != nil {
		return nil, fmt.Errorf("loading context for testdata/evolution: %w", err)
	} else {
		subj := ctx.NewSubject()
		tcases["evolution"] = tcase{
			ctx: ctx,
			params: &models1.MutationParameters{
				Package:  ctx.Package,
				File:     ctx.File,
				FuncDecl: subj.AST,
			},
		}
	}
	if ctx, err := models.LoadContext("", "testdata/words", "WordReverse"); err != nil {
		return nil, fmt.Errorf("loading context for testdata/words: %w", err)
	} else {
		subj := ctx.NewSubject()
		tcases["words"] = tcase{
			ctx: ctx,
			params: &models1.MutationParameters{
				Package:  ctx.Package,
				File:     ctx.File,
				FuncDecl: subj.AST,
			},
		}
	}
	return tcases, nil
}

func fprintSafe(fdecl *ast.FuncDecl) (code []byte, err error) {
	code = []byte{}
	buffer := bytes.NewBuffer([]byte{})
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %q", r)
		}
	}()
	if printer.Fprint(buffer, token.NewFileSet(), fdecl) != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Test_Develop(t *testing.T) {
	tcases, err := loadTestCases()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	for tname, tcase := range tcases {
		t.Run(tname, func(t *testing.T) {
			err := Develop(tcase.params)
			if err != nil {
				t.Fatal(fmt.Errorf("act: %w", err))
			}
			if astwutl.CompareRecursivelyWithAddresses(tcase.params.FuncDecl, tcase.ctx.FuncDecl) {
				t.Error("Failed to see change on subject")
			}
		})
	}
}

func Benchmark_Develop(b *testing.B) {
	tcases, err := loadTestCases()
	if err != nil {
		b.Fatal(fmt.Errorf("prep: %w", err))
	}
	for tname, tcase := range tcases {
		b.Run(tname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := Develop(tcase.params)
				if err != nil {
					b.Fatal(fmt.Errorf("act: %w", err))
				}
				if astwutl.CompareRecursivelyWithAddresses(tcase.params.FuncDecl, tcase.ctx.FuncDecl) {
					b.Error("Failed to see change on subject")
				}
			}
		})
	}

}

func Test_DevelopProgressively(t *testing.T) {
	tcases, err := loadTestCases()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	for tname, tcase := range tcases {
		t.Run(tname, func(t *testing.T) {
			var best *ast.FuncDecl
			for i := 0; i < 2000; i++ {
				subject := clone.FuncDecl(best)
				err := Develop(tcase.params)
				if err != nil {
					t.Error(fmt.Errorf("act, i=%d: %w", i, err))
				}
				if code, err := fprintSafe(best); err == nil {
					fmt.Printf("Code:\n%s\n", utilities.IndentLines(string(code), 4))
					best = subject
				}
			}
		})
	}
}
