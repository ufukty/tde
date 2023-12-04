package genetics

import (
	"fmt"
	"go/ast"
	"math/rand"
	"slices"
	"tde/internal/evolution/genetics/mutation/v1/stg/nodes"
)

var (
	ErrNoPicks = fmt.Errorf("could not pick a node to grow")
)

// doesn't modify lines (slices)
// TODO: pick nil holding parents
func Grow(fd *ast.FuncDecl) error {
	c, ok := pick(linearizeCursorsForNilValuedFields(fd.Body))
	if !ok {
		return ErrNoPicks
	}

	if err := replaceOnParentWithCursor(c, n); err != nil {

	}

	return nil
}

// TODO:
func AddLine(fd *ast.FuncDecl) error {
	bs, ok := pick(filter[ast.BlockStmt](linearize(fd.Body)))
	if !ok {
		return ErrNoPicks
	}

	l := nodes.InType(nil, 1, typeOf(bs))
	i := rand.Intn(len(bs.List) + 1)

	if i == len(bs.List) {
		bs.List = append(bs.List, l)
	} else {
		bs.List = slices.Insert(bs.List, i, l)
	}

	return nil
}

// TODO:
func DeleteLine(fd *ast.FuncDecl) error {
	bs, ok := pick(filter[ast.BlockStmt](linearize(fd.Body)))
	if !ok {
		return ErrNoPicks
	}
	i := rand.Intn(len(bs.List) + 1)
}

// TODO:
func SwapLines(fd *ast.FuncDecl) {
	s, ok := pick(filter[ast.BlockStmt](linearize(fd.Body)))
}
