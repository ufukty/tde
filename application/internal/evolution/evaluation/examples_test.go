package evaluation

import (
	"go/ast"
	"go/token"
	models "tde/models/program"
)

var examples = map[models.Layer][]*ast.FuncDecl{
	models.AST: {
		{
			Name: &ast.Ident{Name: "WordReverse"},
			Type: nil, // NOTE: DO NOT FIX
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: "\"\""}},
					},
				},
			},
		},
	},

	models.Code: {
		{
			Name: &ast.Ident{Name: "WordReverse"},
			Type: &ast.FuncType{
				Params:  &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "w"}}, Type: &ast.Ident{Name: "string"}}}},
				Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "string"}}}},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: ""}}, // NOTE: DO NOT FIX
					},
				},
			},
		},
	},

	models.Candidate: {
		{
			Name: &ast.Ident{Name: "WordReverse"},
			Type: &ast.FuncType{
				Params:  &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "w"}}, Type: &ast.Ident{Name: "string"}}}},
				Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "string"}}}},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: "\"\""}},
					},
				},
			},
		},
	},

	models.Solution: {
		{
			Name: &ast.Ident{Name: "WordReverse"},
			Type: &ast.FuncType{
				Params:  &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "w"}}, Type: &ast.Ident{Name: "string"}}}},
				Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "string"}}}},
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
		},
		{
			Name: &ast.Ident{Name: "WordReverse"},
			Type: &ast.FuncType{
				Params:  &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "w"}}, Type: &ast.Ident{Name: "string"}}}},
				Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "string"}}}},
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
							&ast.CallExpr{Fun: &ast.Ident{Name: "string"}, Args: []ast.Expr{&ast.Ident{Name: "inputRunes"}}},
						},
					},
				},
			},
		}},
}
