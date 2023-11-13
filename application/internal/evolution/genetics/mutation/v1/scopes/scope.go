package scopes

import (
	"go/ast"
	"tde/internal/evolution/genetics/mutation/v1/models"
)

func subnodes[T any](n ast.Node) (subnodes []ast.Node) {
	ast.Inspect(n, func(s ast.Node) bool {
		if s == n {
			return true
		}
		if s, ok := s.(T); ok {
			subnodes = append(subnodes, s)
		}
		return false
	})
	return
}

func difference[T comparable](main, subtract []T) (diff []T) {
	mapM := make(map[T]T, len(main))
	mapS := make(map[T]T, len(main))
}

func candidatesForBlockStmt(b *ast.BlockStmt) []ast.Node {
	nodes := subnodes[ast.Node](b)
	blockstmts := subnodes[*ast.BlockStmt](b)
	rest := difference(nodes, blockstmts)
}

func ScopeOut(params models.MutationParameters) error {
	spots := map[*ast.BlockStmt][]ast.Node{}

	ast.Inspect(params.FuncDecl.Body, func(n ast.Node) bool {
		if n, ok := n.(*ast.BlockStmt); ok {

		}
	})
	return nil
}
