package mutation

import (
	"fmt"
	"go/ast"
	"tde/internal/evolution/models"
	"tde/internal/utilities/pick"
	"tde/internal/utilities/randoms"

	"golang.org/x/exp/slices"
)

func listBlockStatements(n ast.Node, subnodes int) (slice []ast.Node) {
	ast.Inspect(n, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.BlockStmt:
			if len(n.List) >= subnodes {
				slice = append(slice, n)
			}
		case *ast.CommClause:
			if len(n.Body) >= subnodes {
				slice = append(slice, n)
			}
		case *ast.CaseClause:
			if len(n.Body) >= subnodes {
				slice = append(slice, n)
			}
		}

		return true
	})
	return
}

func RemoveALine(params models.Parameters, context models.Context, subj models.Subject) error {
	nodes := listBlockStatements(subj.AST, 1)
	if len(nodes) == 0 {
		return ErrUnavailable
	}
	choosen, err := pick.Pick(nodes)
	if err != nil {
		return fmt.Errorf("picking the block statement: %w", err)
	}

	switch choosen := choosen.(type) {
	case *ast.BlockStmt:
		cut := randoms.UniformIntN(len(choosen.List))
		choosen.List = slices.Delete(slices.Clone(choosen.List), cut, cut)
	case *ast.CommClause:
		cut := randoms.UniformIntN(len(choosen.Body))
		choosen.Body = slices.Delete(slices.Clone(choosen.Body), cut, cut)
	case *ast.CaseClause:
		cut := randoms.UniformIntN(len(choosen.Body))
		choosen.Body = slices.Delete(slices.Clone(choosen.Body), cut, cut)
	}

	return nil
}

func SwapTwoLines(params models.Parameters, context models.Context, subj models.Subject) error {
	nodes := listBlockStatements(subj.AST, 2)
	if len(nodes) == 0 {
		return ErrUnavailable
	}
	choosen, err := pick.Pick(nodes)
	if err != nil {
		return fmt.Errorf("picking the block statement: %w", err)
	}

	switch choosen := choosen.(type) {
	case *ast.BlockStmt:
		cut := randoms.UniformIntN(len(choosen.List) - 1)
		choosen.List[cut], choosen.List[cut+1] = choosen.List[cut+1], choosen.List[cut]
	case *ast.CommClause:
		cut := randoms.UniformIntN(len(choosen.Body) - 1)
		choosen.Body[cut], choosen.Body[cut+1] = choosen.Body[cut+1], choosen.Body[cut]
	case *ast.CaseClause:
		cut := randoms.UniformIntN(len(choosen.Body) - 1)
		choosen.Body[cut], choosen.Body[cut+1] = choosen.Body[cut+1], choosen.Body[cut]
	}

	return nil
}
