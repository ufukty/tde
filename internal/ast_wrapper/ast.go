package ast_wrapper

import (
	"tde/internal/utilities"

	"go/ast"
	"go/printer"
	"go/token"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/ast/astutil"
)

func ListSubnodes(root ast.Node) (subnodes []*ast.Node) {
	nodes := []*ast.Node{}
	ast.Inspect(root, func(n ast.Node) bool {
		if n != nil {
			nodes = append(nodes, &n)
		}
		return true
	})
	return nodes
}

func FilterSubNodes(root ast.Node, filter func(n ast.Node) bool) (filteredNodes []ast.Node) {
	list := []ast.Node{}
	ast.Inspect(root, func(n ast.Node) bool {
		if filter(n) {
			list = append(list, n)
		}
		return true
	})
	return list
}

func ChildNodes(root ast.Node) (children []ast.Node) {
	list := []ast.Node{}
	astutil.Apply(root, func(c *astutil.Cursor) bool {
		if c.Parent() == root {
			list = append(list, c.Node())
		}
		if c.Node() == root {
			return true
		} else {
			return false
		}
	}, nil)
	return list
}

// panics in case the tree is invalid
func StringAST(node ast.Node) string {
	fset := token.NewFileSet()
	sw := utilities.NewStringWriter()
	err := printer.Fprint(sw, fset, node)
	if err != nil {
		panic(errors.Wrapf(err, "failed print"))
	}
	return sw.String()
}

func Inspect(node ast.Node, callback func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool) {
	var (
		parentTrace     []ast.Node
		childIndexTrace []int
	)
	var (
		updateParentTrace = func(n ast.Node) {
			if n != nil {
				parentTrace = append(parentTrace, n)
			} else {
				parentTrace = parentTrace[:len(parentTrace)-1]
			}
		}
		updateChildIndexTrace = func(n ast.Node) {
			if n != nil {
				childIndexTrace = append(childIndexTrace, 0)
			} else {
				l := len(childIndexTrace)
				childIndexTrace = childIndexTrace[:l-1]
				childIndexTrace[l-2]++
			}
		}
	)
	ast.Inspect(node, func(n ast.Node) bool {
		updateParentTrace(n)
		updateChildIndexTrace(n)
		return callback(n, parentTrace[:len(parentTrace)-1], childIndexTrace[:len(childIndexTrace)-1])
	})
}
