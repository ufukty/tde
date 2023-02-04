package function_examiner

import (
	"tde/internal/astw/traverse"
	"tde/internal/astw/types"
	"tde/internal/cfg/context_resolution/context"

	"fmt"
	"go/ast"
	"go/token"
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

func examineSingularAssignment(ctx *context.Context, lhs, rhs ast.Expr) {
	if lhs, ok := lhs.(*ast.Ident); ok {
		ctx.AddVariable(lhs)
	}
}

func examineAssignStmt(ctx *context.Context, stmt *ast.AssignStmt) {
	if stmt.Tok == token.DEFINE {
		for i := 0; i < len(stmt.Lhs); i++ {
			examineSingularAssignment(ctx, stmt.Lhs[i], stmt.Rhs[i])
		}
	}
}

func examineImportDeclaration(ctx *context.Context, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		if spec, ok := spec.(*ast.ImportSpec); ok {
			ctx.AddImport(spec)
		}
	}
}

func examineVariableDeclaration(ctx *context.Context, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		if spec, ok := spec.(*ast.ValueSpec); ok {
			for _, name := range spec.Names {
				ctx.AddVariable(name)
			}
		}
	}
}

func examineTypeDeclaration(ctx *context.Context, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		if spec, ok := spec.(*ast.TypeSpec); ok {
			ctx.AddType(spec)
		}
	}
}

func examineGenDecl(ctx *context.Context, decl *ast.GenDecl) {
	switch decl.Tok {
	case token.IMPORT:
		examineImportDeclaration(ctx, decl)
	case token.VAR, token.CONST:
		examineVariableDeclaration(ctx, decl)
	case token.TYPE:
		examineTypeDeclaration(ctx, decl)
	}
}

func examineDeclStmt(ctx *context.Context, declStmt *ast.DeclStmt) {
	switch decl := declStmt.Decl.(type) {
	case *ast.GenDecl:
		examineGenDecl(ctx, decl)
	case *ast.FuncDecl:
		fmt.Println("examineDeclStmt:FuncDecl")
	}
}

func Examine(ctx *context.Context, funcDecl, insertionPoint *traverse.TraversableNode) {
	var isCompleted = false
	traverse.TraverseTwice(funcDecl,
		func(tNodePtr *traverse.TraversableNode) bool {

			if isCompleted || tNodePtr.PointsToNilSpot {
				return false
			}

			if tNodePtr == insertionPoint { // FIXME: is relying on "value comparison" good idea to check if we reach to the same node as choosen spot?
				isCompleted = true
				return false
			}

			switch node := tNodePtr.Value.(type) {

			case *ast.AssignStmt:
				examineAssignStmt(ctx, node)

			// case *ast.ExprStmt:
			// 	ExamineExprStmt(ctx, node)

			case *ast.DeclStmt:
				examineDeclStmt(ctx, node)

			case *ast.BlockStmt:
				ctx.ScopeIn()

			case *ast.FuncDecl:

				// function name itself should be examined earlier with other in-package declarations

				for _, param := range node.Type.Params.List {
					for _, name := range param.Names {
						ctx.AddVariable(name)
					}
				}
				return false
			}

			// if tNodePtr.ExpectedType.IsSliceType() {

			// }

			// examineEnteringNode(&ctx, tNodePtr)
			return true
		},
		func(tNodePtr *traverse.TraversableNode) {
			if tNodePtr.PointsToNilSpot {
				return
			}

			switch tNodePtr.ExpectedType {
			case types.BlockStmt:
				ctx.ScopeOut()
			}
			// examineLeavingNode(&ctx, node)
		},
	)
}
