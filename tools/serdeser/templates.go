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

func templateNewRequest(modelIdent *ast.Ident) *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						{
							Name: "req",
						},
					},
					Type: &ast.StarExpr{
						X: modelIdent,
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
			Func:       92,
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
									Name: "req",
								},
							},
						},
					},
				},
				&ast.IfStmt{
					If:   264,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 271,
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
							Name: "httpRequest",
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
					If:   400,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 407,
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
							Name: "httpRequest",
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

func templateSend(requestModelIdent, responseModelIdent *ast.Ident) *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						{
							Name: "req",
						},
					},
					Type: &ast.StarExpr{
						X: requestModelIdent,
					},
					Tag:     nil,
					Comment: nil,
				},
			},
		},
		Name: &ast.Ident{
			Name: "Send",
		},
		Type: &ast.FuncType{
			Func:       516,
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
							X: responseModelIdent,
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
							Name: "httpRequest",
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
									Name: "req",
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
							},
						},
					},
				},
				&ast.IfStmt{
					If:   676,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 683,
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
												Value: "\"failed on creating an object for request\"",
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
							Name: "httpResponse",
						},
						&ast.Ident{
							Name: "err",
						},
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X: &ast.SelectorExpr{
									X: &ast.Ident{
										Name: "http",
									},
									Sel: &ast.Ident{
										Name: "DefaultClient",
									},
								},
								Sel: &ast.Ident{
									Name: "Do",
								},
							},
							Args: []ast.Expr{
								&ast.Ident{
									Name: "httpRequest",
								},
							},
						},
					},
				},
				&ast.IfStmt{
					If:   828,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 835,
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
												Value: "\"failed on sending the request\"",
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
							Name: "res",
						},
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CompositeLit{
							Type:       responseModelIdent,
							Elts:       nil,
							Incomplete: false,
						},
					},
				},
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.Ident{
							Name: "err",
						},
					},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X: &ast.Ident{
									Name: "res",
								},
								Sel: &ast.Ident{
									Name: "DeserializeResponse",
								},
							},
							Args: []ast.Expr{
								&ast.Ident{
									Name: "httpResponse",
								},
							},
						},
					},
				},
				&ast.IfStmt{
					If:   998,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 1005,
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
												Value: "\"failed on parsing the response body\"",
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
						&ast.UnaryExpr{
							OpPos: 1095,
							Op:    token.AND,
							X: &ast.Ident{
								Name: "res",
							},
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

func templateParseRequest(modelIdent *ast.Ident) *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						{
							Name: "req",
						},
					},
					Type: &ast.StarExpr{
						X: modelIdent,
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
			Func:       1105,
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
									Name: "req",
								},
							},
						},
					},
				},
				&ast.IfStmt{
					If:   1230,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 1237,
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
												Value: "\"failed on parsing the response body\"",
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

func templateSerializeIntoResponseWriter(modelIdent *ast.Ident) *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						{
							Name: "res",
						},
					},
					Type: &ast.StarExpr{
						X: modelIdent,
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
			Func:       1316,
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
									Name: "res",
								},
							},
						},
					},
				},
				&ast.IfStmt{
					If:   1458,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 1465,
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

func templateDeserializeResponse(modelIdent *ast.Ident) *ast.FuncDecl {
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
						X: modelIdent,
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
			Func:       1544,
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
					If:   1678,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 1685,
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
												Value: "\"failed on parsing the response body\"",
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
