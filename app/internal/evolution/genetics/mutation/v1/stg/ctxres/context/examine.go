package context

import (
	"fmt"
	"go/ast"
	"go/token"
	"tde/internal/astw/traverse"
	"tde/internal/astw/types"
)

// func examineEnteringNode(ctx *context.Context, node ast.Node) {
// 	switch node := node.(type) {
// 	case
// 		*ast.BlockStmt,
// 		*ast.FuncDecl,
// 		*ast.ForStmt,
// 		*ast.IfStmt,
// 		*ast.SwitchStmt,
// 		*ast.TypeSwitchStmt:
// 		ctx.ScopeIn()

// 	case
// 		*ast.Package:
// 		for _, file := range node.Files {
// 			for _, decl := range file.Decls {
// 				if funcDecl, ok := decl.(*ast.FuncDecl); ok {
// 					ctx.AddVariable(funcDecl.Name)
// 				}
// 			}
// 		}

// 	case
// 		*ast.File:
// 		for _, decl := range node.Decls {
// 			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
// 				ctx.AddVariable(funcDecl.Name)
// 			}
// 		}
// 	}
// }

// func examineLeavingNode(ctx *context.Context, node ast.Node) {
// 	switch node := node.(type) {
// 	case
// 		*ast.BlockStmt,
// 		*ast.FuncDecl,
// 		*ast.ForStmt,
// 		*ast.IfStmt,
// 		*ast.SwitchStmt,
// 		*ast.TypeSwitchStmt:
// 		ctx.ScopeOut()

// 	case
// 		*ast.AssignStmt:
// 		if node.Tok == token.DEFINE {
// 			for _, expr := range node.Lhs {
// 				if ident, ok := expr.(*ast.Ident); ok {
// 					ctx.AddVariable(ident)
// 				}
// 			}
// 		}

// 	case
// 		*ast.FuncType:
// 		if fieldList := node.Params.List; fieldList != nil {
// 			for _, field := range fieldList {
// 				for _, ident := range field.Names {
// 					ctx.AddVariable(ident)
// 				}
// 			}
// 		}
// 	}
// }

func (ctx *Context) ExamineFuncDecl(funcdecl, insertionPoint *traverse.TraversableNode) {
	var isCompleted = false
	traverse.TraverseTwice(funcdecl, func(tNodePtr *traverse.TraversableNode) bool {
		if isCompleted || tNodePtr.PointsToNilSpot {
			return false
		}
		if tNodePtr == insertionPoint { // FIXME: is relying on "value comparison" good idea to check if we reach to the same node as choosen spot?
			isCompleted = true
			return false
		}

		switch node := tNodePtr.Value.(type) {
		case *ast.AssignStmt:
			if node.Tok == token.DEFINE {
				for i := 0; i < len(node.Lhs); i++ {
					if lhs, ok := node.Lhs[i].(*ast.Ident); ok {
						ctx.AddVariable(lhs)
					}
				}
			}

		// case *ast.ExprStmt:
		// 	ExamineExprStmt(ctx, node)

		case *ast.DeclStmt:

			switch decl := node.Decl.(type) {
			case *ast.GenDecl:

				switch decl.Tok {
				case token.IMPORT:
					for _, spec := range decl.Specs {
						if spec, ok := spec.(*ast.ImportSpec); ok {
							ctx.AddImport(spec)
						}
					}

				case token.VAR, token.CONST:
					for _, spec := range decl.Specs {
						if spec, ok := spec.(*ast.ValueSpec); ok {
							for _, name := range spec.Names {
								ctx.AddVariable(name)
							}
						}
					}

				case token.TYPE:
					for _, spec := range decl.Specs {
						if spec, ok := spec.(*ast.TypeSpec); ok {
							ctx.AddType(spec)
						}
					}
				}

			case *ast.FuncDecl:
				fmt.Println("examineDeclStmt:FuncDecl") // FIXME:
			}

		case *ast.BlockStmt:
			ctx.ScopeIn()

		case *ast.FuncDecl: // function name itself should be examined earlier with other in-package declarations
			if node.Type != nil && node.Type.Params != nil && node.Type.Params.List != nil { // BECAUSE: invalid ASTs may lack any node
				for _, param := range node.Type.Params.List {
					if param.Names != nil {
						for _, name := range param.Names {
							ctx.AddVariable(name)
						}
					}
				}
			}
			return false
		}

		// if tNodePtr.ExpectedType.IsSliceType() {

		// }

		// examineEnteringNode(&ctx, tNodePtr)
		return true
	}, func(tNodePtr *traverse.TraversableNode) {
		if tNodePtr.PointsToNilSpot {
			return
		}

		switch tNodePtr.ExpectedType {
		case types.BlockStmt:
			ctx.ScopeOut()
		}
		// examineLeavingNode(&ctx, node)
	})
}

// Only adds declarations for functions, imports, types, variables and constants. Won't examine function bodies.
func (ctx *Context) ExamineFile(file *ast.File) {
	for _, decl := range file.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:

			switch decl.Tok {
			case token.IMPORT:
				for _, spec := range decl.Specs {
					if spec, ok := spec.(*ast.ImportSpec); ok {
						ctx.AddImport(spec)
					}
				}

			case token.VAR, token.CONST:
				for _, spec := range decl.Specs {
					if spec, ok := spec.(*ast.ValueSpec); ok {
						for _, name := range spec.Names {
							ctx.AddVariable(name)
						}
					}
				}

			case token.TYPE:
				for _, spec := range decl.Specs {
					if spec, ok := spec.(*ast.TypeSpec); ok {
						ctx.AddType(spec)
					}
				}
			}

		case *ast.FuncDecl:
			if decl.Recv == nil {
				ctx.AddFuncDeclaration(decl)
			} else {
				ctx.AddMethodDeclaration(decl)
			}
		}
	}
}

func (ctx *Context) ExeminePkg(pkg *ast.Package) {
	for _, file := range pkg.Files {
		ctx.ExamineFile(file)
	}
}
