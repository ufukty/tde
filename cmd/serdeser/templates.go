package main

import (
	"go/ast"
	"go/token"
)

func templateImports() *ast.GenDecl {
	return &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				Name: nil,
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"bytes\"",
				},
				Comment: nil,
			},
			&ast.ImportSpec{
				Name: nil,
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"encoding/json\"",
				},
				Comment: nil,
			},
			&ast.ImportSpec{
				Name: nil,
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"net/http\"",
				},
				Comment: nil,
			},
			&ast.ImportSpec{
				Name: nil,
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"github.com/pkg/errors\"",
				},
				Comment: nil,
			},
		},
	}
}

func templateNewRequest(structIdent *ast.Ident) *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						{
							Name: "s",
						},
					},
					Type: &ast.StarExpr{
						X: structIdent,
					},
					Tag:     nil,
					Comment: nil,
				},
			},
		},
		Name: &ast.Ident{
			Name: "NewRequest",
		},
		Type: &ast.FuncType{
			Func:       93,
			TypeParams: nil,
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							{
								Name: "method",
							},
							{
								Name: "url",
							},
						},
						Type: &ast.Ident{
							Name: "string",
						},
						Tag:     nil,
						Comment: nil,
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: nil,
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X: &ast.Ident{
									Name: "http",
								},
								Sel: &ast.Ident{
									Name: "Request",
								},
							},
						},
						Tag:     nil,
						Comment: nil,
					},
					{
						Names: nil,
						Type: &ast.Ident{
							Name: "error",
						},
						Tag:     nil,
						Comment: nil,
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.Ident{
							Name: "buffer",
						},
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.Ident{
								Name: "new",
							},
							Args: []ast.Expr{
								&ast.SelectorExpr{
									X: &ast.Ident{
										Name: "bytes",
									},
									Sel: &ast.Ident{
										Name: "Buffer",
									},
								},
							},
						},
					},
				},
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.Ident{
							Name: "err",
						},
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X: &ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X: &ast.Ident{
											Name: "json",
										},
										Sel: &ast.Ident{
											Name: "NewEncoder",
										},
									},
									Args: []ast.Expr{
										&ast.Ident{
											Name: "buffer",
										},
									},
								},
								Sel: &ast.Ident{
									Name: "Encode",
								},
							},
							Args: []ast.Expr{
								&ast.Ident{
									Name: "s",
								},
							},
						},
					},
				},
				&ast.IfStmt{
					If:   261,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 268,
						Op:    token.NEQ,
						Y: &ast.Ident{
							Name: "nil",
						},
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{
									&ast.Ident{
										Name: "nil",
									},
									&ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X: &ast.Ident{
												Name: "errors",
											},
											Sel: &ast.Ident{
												Name: "Wrap",
											},
										},
										Args: []ast.Expr{
											&ast.Ident{
												Name: "err",
											},
											&ast.BasicLit{
												Kind:  token.STRING,
												Value: "\"failed on serialization\"",
											},
										},
									},
								},
							},
						},
					},
					Else: nil,
				},
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.Ident{
							Name: "req",
						},
						&ast.Ident{
							Name: "err",
						},
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X: &ast.Ident{
									Name: "http",
								},
								Sel: &ast.Ident{
									Name: "NewRequest",
								},
							},
							Args: []ast.Expr{
								&ast.Ident{
									Name: "method",
								},
								&ast.Ident{
									Name: "url",
								},
								&ast.Ident{
									Name: "buffer",
								},
							},
						},
					},
				},
				&ast.IfStmt{
					If:   389,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 396,
						Op:    token.NEQ,
						Y: &ast.Ident{
							Name: "nil",
						},
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{
									&ast.Ident{
										Name: "nil",
									},
									&ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X: &ast.Ident{
												Name: "errors",
											},
											Sel: &ast.Ident{
												Name: "Wrap",
											},
										},
										Args: []ast.Expr{
											&ast.Ident{
												Name: "err",
											},
											&ast.BasicLit{
												Kind:  token.STRING,
												Value: "\"failed on creating request object\"",
											},
										},
									},
								},
							},
						},
					},
					Else: nil,
				},
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.Ident{
							Name: "req",
						},
						&ast.Ident{
							Name: "nil",
						},
					},
				},
			},
		},
	}
}

func templateParseRequest(structIdent *ast.Ident) *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						{
							Name: "s",
						},
					},
					Type: &ast.StarExpr{
						X: structIdent,
					},
					Tag:     nil,
					Comment: nil,
				},
			},
		},
		Name: &ast.Ident{
			Name: "ParseRequest",
		},
		Type: &ast.FuncType{
			Func:       290,
			TypeParams: nil,
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							{
								Name: "r",
							},
						},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X: &ast.Ident{
									Name: "http",
								},
								Sel: &ast.Ident{
									Name: "Request",
								},
							},
						},
						Tag:     nil,
						Comment: nil,
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: nil,
						Type: &ast.Ident{
							Name: "error",
						},
						Tag:     nil,
						Comment: nil,
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.Ident{
							Name: "err",
						},
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X: &ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X: &ast.Ident{
											Name: "json",
										},
										Sel: &ast.Ident{
											Name: "NewDecoder",
										},
									},
									Args: []ast.Expr{
										&ast.SelectorExpr{
											X: &ast.Ident{
												Name: "r",
											},
											Sel: &ast.Ident{
												Name: "Body",
											},
										},
									},
								},
								Sel: &ast.Ident{
									Name: "Decode",
								},
							},
							Args: []ast.Expr{
								&ast.Ident{
									Name: "s",
								},
							},
						},
					},
				},
				&ast.IfStmt{
					If:   410,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 417,
						Op:    token.NEQ,
						Y: &ast.Ident{
							Name: "nil",
						},
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{
									&ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X: &ast.Ident{
												Name: "errors",
											},
											Sel: &ast.Ident{
												Name: "Wrap",
											},
										},
										Args: []ast.Expr{
											&ast.Ident{
												Name: "err",
											},
											&ast.BasicLit{
												Kind:  token.STRING,
												Value: "\"failed on serialization\"",
											},
										},
									},
								},
							},
						},
					},
					Else: nil,
				},
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.Ident{
							Name: "nil",
						},
					},
				},
			},
		},
	}
}

func templateSerializeIntoResponseWriter(structIdent *ast.Ident) *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						{
							Name: "s",
						},
					},
					Type: &ast.StarExpr{
						X: structIdent,
					},
					Tag:     nil,
					Comment: nil,
				},
			},
		},
		Name: &ast.Ident{
			Name: "SerializeIntoResponseWriter",
		},
		Type: &ast.FuncType{
			Func:       704,
			TypeParams: nil,
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							{
								Name: "w",
							},
						},
						Type: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: "http",
							},
							Sel: &ast.Ident{
								Name: "ResponseWriter",
							},
						},
						Tag:     nil,
						Comment: nil,
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: nil,
						Type: &ast.Ident{
							Name: "error",
						},
						Tag:     nil,
						Comment: nil,
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.Ident{
							Name: "err",
						},
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X: &ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X: &ast.Ident{
											Name: "json",
										},
										Sel: &ast.Ident{
											Name: "NewEncoder",
										},
									},
									Args: []ast.Expr{
										&ast.Ident{
											Name: "w",
										},
									},
								},
								Sel: &ast.Ident{
									Name: "Encode",
								},
							},
							Args: []ast.Expr{
								&ast.Ident{
									Name: "s",
								},
							},
						},
					},
				},
				&ast.IfStmt{
					If:   842,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 849,
						Op:    token.NEQ,
						Y: &ast.Ident{
							Name: "nil",
						},
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{
									&ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X: &ast.Ident{
												Name: "errors",
											},
											Sel: &ast.Ident{
												Name: "Wrap",
											},
										},
										Args: []ast.Expr{
											&ast.Ident{
												Name: "err",
											},
											&ast.BasicLit{
												Kind:  token.STRING,
												Value: "\"failed on serialization\"",
											},
										},
									},
								},
							},
						},
					},
					Else: nil,
				},
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.Ident{
							Name: "nil",
						},
					},
				},
			},
		},
	}
}

func templateDeserializeResponse(structIdent *ast.Ident) *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						{
							Name: "s",
						},
					},
					Type: &ast.StarExpr{
						X: structIdent,
					},
					Tag:     nil,
					Comment: nil,
				},
			},
		},
		Name: &ast.Ident{
			Name: "DeserializeResponse",
		},
		Type: &ast.FuncType{
			Func:       929,
			TypeParams: nil,
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							{
								Name: "res",
							},
						},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X: &ast.Ident{
									Name: "http",
								},
								Sel: &ast.Ident{
									Name: "Response",
								},
							},
						},
						Tag:     nil,
						Comment: nil,
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: nil,
						Type: &ast.Ident{
							Name: "error",
						},
						Tag:     nil,
						Comment: nil,
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.Ident{
							Name: "err",
						},
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X: &ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X: &ast.Ident{
											Name: "json",
										},
										Sel: &ast.Ident{
											Name: "NewDecoder",
										},
									},
									Args: []ast.Expr{
										&ast.SelectorExpr{
											X: &ast.Ident{
												Name: "res",
											},
											Sel: &ast.Ident{
												Name: "Body",
											},
										},
									},
								},
								Sel: &ast.Ident{
									Name: "Decode",
								},
							},
							Args: []ast.Expr{
								&ast.Ident{
									Name: "s",
								},
							},
						},
					},
				},
				&ast.IfStmt{
					If:   1063,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 1070,
						Op:    token.NEQ,
						Y: &ast.Ident{
							Name: "nil",
						},
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{
									&ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X: &ast.Ident{
												Name: "errors",
											},
											Sel: &ast.Ident{
												Name: "Wrap",
											},
										},
										Args: []ast.Expr{
											&ast.Ident{
												Name: "err",
											},
											&ast.BasicLit{
												Kind:  token.STRING,
												Value: "\"failed on serialization\"",
											},
										},
									},
								},
							},
						},
					},
					Else: nil,
				},
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.Ident{
							Name: "nil",
						},
					},
				},
			},
		},
	}
}
