package context

import (
	"fmt"
	"go/ast"
	"tde/internal/astw"
)

// Parent node is where the new node will be appended into. So, new node can access variables only the parentNode can access.
func findTrace(pkg *ast.Package, parentNode ast.Node) []int {
	var trace []int = nil
	astw.InspectWithTrace(pkg, func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		if currentNode == parentNode {
			trace = childIndexTrace
		}
		return trace == nil
	})
	return trace
}

func ResolveContextAtPackageLevel(ctx *Context, pkg *ast.Package) {
	for _, file := range pkg.Files {
		ExamineFile(ctx, file)
	}
}

func ResolveContextAtLowerScopes(ctx *Context, root ast.Node, indicesInsertionPoint []int) {
	const (
		ANCESTOR = iota
		ON_PATH
		WONT_SEE
	)
	var getRelation = func(current, insertionPoint []int) int {
		var cLastIndex = current[len(current)-1]
		var ipSameDepthIndex = insertionPoint[len(current)-1]
		if cLastIndex == ipSameDepthIndex {
			return ANCESTOR
		} else if cLastIndex < ipSameDepthIndex {
			return ON_PATH
		} else {
			return WONT_SEE
		}
	}
	astw.InspectWithTrace(root, func(node ast.Node, parents []ast.Node, indicesCurrent []int) bool {
		relation := getRelation(indicesCurrent, indicesInsertionPoint)
		if relation == ON_PATH {
			switch stmt := node.(type) {

			case *ast.AssignStmt:
				ExamineAssignStmt(ctx, stmt)

			case *ast.ExprStmt:
				ExamineExprStmt(ctx, stmt)

			case *ast.DeclStmt:
				ExamineDeclStmt(ctx, stmt)
			}

		} else if relation == ANCESTOR {
			switch node.(type) {

			case *ast.BlockStmt:
				ctx.ScopeIn()
			}
		}

		return relation == ANCESTOR
	})
}

func Resolve(pkg *ast.Package, insertionPoint ast.Node) Context {
	indices := findTrace(pkg, insertionPoint)
	fmt.Println("indices trace: ", indices)
	ctx := NewContext()
	// 	fmt.Println(reflect.TypeOf(node), reflect.ValueOf(node))

	var funcDecl *ast.FuncDecl
	astw.Retrace(pkg, indices, func(node ast.Node, depth int) {
		if funcDecl == nil {
			if node, ok := node.(*ast.FuncDecl); ok {
				funcDecl = node
			}
		}
	})

	ResolveContextAtPackageLevel(&ctx, pkg)
	ResolveContextAtLowerScopes(&ctx, funcDecl, indices)

	return ctx

}
