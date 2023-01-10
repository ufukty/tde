package ast_wrapper

import "go/token"

func LineNumberOfPosition(fset *token.FileSet, posToken token.Pos) int {
	return fset.Position(posToken).Line
}
