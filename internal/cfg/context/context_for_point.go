package context

import (
	"go/ast"
	"tde/internal/astw"
)


func ExamineFile(ctx Context, file *ast.File) {
	for _, decl := range file.Decls {
		switch node := decl.(type) {
		case *ast.FuncDecl:
			if node.Recv == nil {
				ctx.AddVariable(*node.Name)
			}

		}
	}
}


func findTrace(astPkg *ast.Package, insertionPoint ast.Node) []int {
	var trace []int = nil
	astw.InspectWithTrace(astPkg, func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		if currentNode == insertionPoint {
			trace = childIndexTrace
		}
		return trace == nil
	})
	return trace
}

func Resolve(astPkg *ast.Package, insertionPoint ast.Node) Context {
	trace := findTrace(astPkg, insertionPoint)

	ctx := NewContext()

	astw.Retrace(astPkg, trace, func(node ast.Node, depth int) {

		astw.InspectChildren(node, func(node ast.Node, indices int) {

			switch node := node.(type) {
			case *ast.Package:
				// add declarations of every file
				for _, file := range node.Files {
					ExamineFile(file)
				}
				
			case *ast.File:
				ExamineFile(file)
			}

			ExamineEnteringNode(&ctx, node)
		})
	})

	return ctx

}
