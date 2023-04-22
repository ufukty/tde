package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func createNewFileName(inputPath string) string {
	filename := filepath.Base(inputPath)
	fileExt := filepath.Ext(filename)
	filenameWithoutExt := filename[0 : len(filename)-len(fileExt)]
	return filepath.Join(filepath.Dir(inputPath), filenameWithoutExt+".sd.go")
}

func templateSerializer(structIdent *ast.Ident) *ast.FuncDecl {
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
			Name: "Serialize",
		},
		Type: &ast.FuncType{
			Func:       84,
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
					If:   203,
					Init: nil,
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						OpPos: 210,
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

func templateDeserializer(structIdent *ast.Ident) *ast.FuncDecl {
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
			Name: "Deserialize",
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

func checkSuffix(str, suffix string) bool {
	return strings.LastIndex(str, suffix) == len(str)-len(suffix)
}

func discoverFileForStructDefinitions(file *ast.File) (structs []*ast.Ident) {
	ast.Inspect(file, func(n ast.Node) bool {
		if n == nil || n == file {
			return true
		}
		switch n := n.(type) {
		case *ast.File, *ast.GenDecl:
			return true
		case *ast.TypeSpec:
			if _, ok := n.Type.(*ast.StructType); ok {
				if checkSuffix(n.Name.Name, "Request") ||
					checkSuffix(n.Name.Name, "Response") {
					structs = append(structs, n.Name)
				}
			}
			return false
		}
		return false
	})
	return
}

func parseFile(path string) (*ast.File, error) {
	file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.AllErrors)
	if err != nil {
		return nil, errors.Wrap(err, "Outside error")
	}
	return file, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Wrong number of arguments. First argument should be the name of input file.")
	}
	file, err := parseFile(os.Args[1])
	if err != nil {
		log.Fatalln(errors.Wrapf(err, "Could not read the file: '%s'", os.Args[1]))
	}

	var (
		packageName = file.Name
		types       = discoverFileForStructDefinitions(file)
		genDecl     = &ast.GenDecl{
			Tok: token.IMPORT,
			Specs: []ast.Spec{
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
	)

	if len(types) == 0 {
		fmt.Println("No type definition found in the file which ends with either Request or Response.")
		return
	}

	var genFile = &ast.File{
		Name:  packageName,
		Decls: []ast.Decl{genDecl},
	}

	for _, entity := range types {
		genFile.Decls = append(genFile.Decls,
			templateSerializer(entity),
			templateDeserializer(entity),
		)
	}

	newFileName := createNewFileName(os.Args[1])
	target, err := os.Create(newFileName)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not create file in current directory"))
	}
	err = printer.Fprint(target, token.NewFileSet(), genFile)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not write into newly created file"))
	}
}
