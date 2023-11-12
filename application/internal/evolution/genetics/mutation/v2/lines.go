package mutation

import (
	"go/ast"
	"tde/internal/evolution/models"
	"tde/internal/utilities"

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
	choosen := *utilities.Pick(nodes)

	switch choosen := choosen.(type) {
	case *ast.BlockStmt:
		cut := utilities.URandIntN(len(choosen.List))
		choosen.List = slices.Delete(slices.Clone(choosen.List), cut, cut)
	case *ast.CommClause:
		cut := utilities.URandIntN(len(choosen.Body))
		choosen.Body = slices.Delete(slices.Clone(choosen.Body), cut, cut)
	case *ast.CaseClause:
		cut := utilities.URandIntN(len(choosen.Body))
		choosen.Body = slices.Delete(slices.Clone(choosen.Body), cut, cut)
	}

	return nil
}

func SwapTwoLines(params models.Parameters, context models.Context, subj models.Subject) error {
	nodes := listBlockStatements(subj.AST, 2)
	if len(nodes) == 0 {
		return ErrUnavailable
	}
	choosen := *utilities.Pick(nodes)

	switch choosen := choosen.(type) {
	case *ast.BlockStmt:
		cut := utilities.URandIntN(len(choosen.List) - 1)
		choosen.List[cut], choosen.List[cut+1] = choosen.List[cut+1], choosen.List[cut]
	case *ast.CommClause:
		cut := utilities.URandIntN(len(choosen.Body) - 1)
		choosen.Body[cut], choosen.Body[cut+1] = choosen.Body[cut+1], choosen.Body[cut]
	case *ast.CaseClause:
		cut := utilities.URandIntN(len(choosen.Body) - 1)
		choosen.Body[cut], choosen.Body[cut+1] = choosen.Body[cut+1], choosen.Body[cut]
	}

	return nil
}
