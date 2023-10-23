package evaluation

import (
	"fmt"
	"go/ast"
	"go/token"
	"tde/internal/evolution/evaluation/inject"
	"tde/internal/evolution/evaluation/list"
	"tde/internal/evolution/evaluation/slotmgr"
	models "tde/models/program"
	"testing"
)

func prepareEvaluator(ctx *models.Context) (*Evaluator, error) {
	pkgs, err := list.ListPackagesInDir("testdata")
	if err != nil {
		return nil, fmt.Errorf("listing packages in target dir: %w", err)
	}
	sample, err := inject.WithCreatingSample(pkgs.First().Module.Dir, pkgs.First(), "TDE_WordReverse")
	if err != nil {
		return nil, fmt.Errorf("creating sample: %w", err)
	}
	sm := slotmgr.New(sample, "internal/evolution/evaluation/testdata", "words.go")
	evaluator := NewEvaluator(sm, ctx)
	return evaluator, nil
}

func prepareSubjects(ctx *models.Context) (models.Subjects, map[*ast.FuncDecl]models.Layer) {
	var layers = map[*ast.FuncDecl]models.Layer{
		&ast.FuncDecl{
			Name: &ast.Ident{Name: "WordReverse"},
			Type: nil,
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: "\"\""}},
					},
				},
			},
		}: models.AST,
		&ast.FuncDecl{
			Name: &ast.Ident{Name: "WordReverse"},
			Type: &ast.FuncType{
				Params:  &ast.FieldList{List: []*ast.Field{&ast.Field{Names: []*ast.Ident{&ast.Ident{Name: "w"}}, Type: &ast.Ident{Name: "string"}}}},
				Results: &ast.FieldList{List: []*ast.Field{&ast.Field{Type: &ast.Ident{Name: "string"}}}},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "\"\""}},
					},
				},
			},
		}: models.Code,
		&ast.FuncDecl{
			Name: &ast.Ident{Name: "WordReverse"},
			Type: &ast.FuncType{
				Params:  &ast.FieldList{List: []*ast.Field{&ast.Field{Names: []*ast.Ident{&ast.Ident{Name: "w"}}, Type: &ast.Ident{Name: "string"}}}},
				Results: &ast.FieldList{List: []*ast.Field{&ast.Field{Type: &ast.Ident{Name: "string"}}}},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: "\"\""}},
					},
				},
			},
		}: models.Candidate,
		&ast.FuncDecl{
			Name: &ast.Ident{Name: "WordReverse"},
			Type: &ast.FuncType{
				Params:  &ast.FieldList{List: []*ast.Field{&ast.Field{Names: []*ast.Ident{&ast.Ident{Name: "w"}}, Type: &ast.Ident{Name: "string"}}}},
				Results: &ast.FieldList{List: []*ast.Field{&ast.Field{Type: &ast.Ident{Name: "string"}}}},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{&ast.Ident{Name: "runes"}},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun:  &ast.ArrayType{Elt: &ast.Ident{Name: "rune"}},
								Args: []ast.Expr{&ast.Ident{Name: "w"}},
							},
						},
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{&ast.Ident{Name: "reversed"}},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CompositeLit{
								Type:       &ast.ArrayType{Elt: &ast.Ident{Name: "rune"}},
								Incomplete: false,
							},
						},
					},
					&ast.ForStmt{
						For: 153,
						Init: &ast.AssignStmt{
							Lhs: []ast.Expr{&ast.Ident{Name: "i"}},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.BinaryExpr{
									X: &ast.CallExpr{
										Fun:  &ast.Ident{Name: "len"},
										Args: []ast.Expr{&ast.Ident{Name: "runes"}},
									},
									OpPos: 173,
									Op:    token.SUB,
									Y:     &ast.BasicLit{Kind: token.INT, Value: "1"},
								},
							},
						},
						Cond: &ast.BinaryExpr{
							X:     &ast.Ident{Name: "i"},
							OpPos: 180,
							Op:    token.GEQ,
							Y:     &ast.BasicLit{Kind: token.INT, Value: "0"},
						},
						Post: &ast.IncDecStmt{X: &ast.Ident{Name: "i"}, Tok: token.DEC},
						Body: &ast.BlockStmt{
							List: []ast.Stmt{
								&ast.AssignStmt{
									Lhs: []ast.Expr{&ast.Ident{Name: "reversed"}},
									Tok: token.ASSIGN,
									Rhs: []ast.Expr{
										&ast.CallExpr{
											Fun: &ast.Ident{Name: "append"},
											Args: []ast.Expr{
												&ast.Ident{Name: "reversed"},
												&ast.IndexExpr{
													X:      &ast.Ident{Name: "runes"},
													Index:  &ast.Ident{Name: "i"},
													Rbrack: 229,
												},
											},
										},
									},
								},
							},
						},
					},
					&ast.ReturnStmt{
						Results: []ast.Expr{
							&ast.CallExpr{Fun: &ast.Ident{Name: "string"}, Args: []ast.Expr{&ast.Ident{Name: "reversed"}}},
						},
					},
				},
			},
		}: models.Solution,
		&ast.FuncDecl{
			Name: &ast.Ident{Name: "WordReverse"},
			Type: &ast.FuncType{
				Params:  &ast.FieldList{List: []*ast.Field{&ast.Field{Names: []*ast.Ident{&ast.Ident{Name: "w"}}, Type: &ast.Ident{Name: "string"}}}},
				Results: &ast.FieldList{List: []*ast.Field{&ast.Field{Type: &ast.Ident{Name: "string"}}}},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{&ast.Ident{Name: "inputRunes"}},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun:  &ast.ArrayType{Elt: &ast.Ident{Name: "rune"}},
								Args: []ast.Expr{&ast.Ident{Name: "w"}},
							},
						},
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{&ast.Ident{Name: "length"}},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.Ident{Name: "len"},
								Args: []ast.Expr{
									&ast.Ident{Name: "inputRunes"},
								},
							},
						},
					},
					&ast.ForStmt{
						For: 361,
						Init: &ast.AssignStmt{
							Lhs: []ast.Expr{
								&ast.Ident{Name: "i"},
								&ast.Ident{Name: "j"},
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.BasicLit{Kind: token.INT, Value: "0"},
								&ast.BinaryExpr{
									X:     &ast.Ident{Name: "length"},
									OpPos: 382,
									Op:    token.SUB,
									Y:     &ast.BasicLit{Kind: token.INT, Value: "1"},
								},
							},
						},
						Cond: &ast.BinaryExpr{
							X:     &ast.Ident{Name: "i"},
							OpPos: 388,
							Op:    token.LSS,
							Y:     &ast.Ident{Name: "j"},
						},
						Post: &ast.AssignStmt{
							Lhs: []ast.Expr{
								&ast.Ident{Name: "i"},
								&ast.Ident{Name: "j"},
							},
							Tok: token.ASSIGN,
							Rhs: []ast.Expr{
								&ast.BinaryExpr{
									X:     &ast.Ident{Name: "i"},
									OpPos: 401,
									Op:    token.ADD,
									Y:     &ast.BasicLit{Kind: token.INT, Value: "1"},
								},
								&ast.BinaryExpr{
									X:     &ast.Ident{Name: "j"},
									OpPos: 406,
									Op:    token.SUB,
									Y:     &ast.BasicLit{Kind: token.INT, Value: "1"},
								},
							},
						},
						Body: &ast.BlockStmt{
							List: []ast.Stmt{
								&ast.AssignStmt{
									Lhs: []ast.Expr{
										&ast.IndexExpr{
											X:      &ast.Ident{Name: "inputRunes"},
											Index:  &ast.Ident{Name: "i"},
											Rbrack: 425,
										},
										&ast.IndexExpr{
											X:      &ast.Ident{Name: "inputRunes"},
											Index:  &ast.Ident{Name: "j"},
											Rbrack: 440,
										},
									},
									Tok: token.ASSIGN,
									Rhs: []ast.Expr{
										&ast.IndexExpr{
											X:      &ast.Ident{Name: "inputRunes"},
											Index:  &ast.Ident{Name: "j"},
											Rbrack: 456,
										},
										&ast.IndexExpr{
											X:      &ast.Ident{Name: "inputRunes"},
											Index:  &ast.Ident{Name: "i"},
											Rbrack: 471,
										},
									},
								},
							},
						},
					},
					&ast.ReturnStmt{
						Results: []ast.Expr{
							&ast.CallExpr{
								Fun:  &ast.Ident{Name: "string"},
								Args: []ast.Expr{&ast.Ident{Name: "inputRunes"}},
							},
						},
					},
				},
			},
		}: models.Solution,
	}
	subjects := models.Subjects{}
	for ast := range layers {
		subject := ctx.NewSubject()
		subject.AST = ast
		subjects.Add(subject)
	}
	return subjects, layers
}

func prepare() (*Evaluator, models.Subjects, map[*ast.FuncDecl]models.Layer, error) {
	ctx, err := models.LoadContext("../../..", "testdata", "WordReverse")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("finding the context for package: %w", err)
	}
	e, err := prepareEvaluator(ctx)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("creating evaluator: %w", err)
	}
	s, layers := prepareSubjects(ctx)
	return e, s, layers, nil
}

// FIXME: check fitness has populated after syntax errors
func Test_Pipeline(t *testing.T) {
	evaluator, subjects, layers, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	if err := evaluator.Pipeline(subjects); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}

	for _, subj := range subjects {
		expectedLayer := layers[subj.AST]
		if subj.Fitness.Layer() != expectedLayer {
			t.Fatal(fmt.Errorf("subject layer mismatch: expected %q got %q", expectedLayer, subj.Fitness.Layer()))
		}
	}
}
