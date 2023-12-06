package genetics

import (
	"fmt"
	"go/ast"
	"tde/internal/astw/types"
	"tde/internal/utilities/setops"
)

func marktypes(fd *ast.FuncDecl) []types.NodeType {
	flags := map[types.NodeType]bool{
		types.CommentSlice: false,
		types.FieldSlice:   false,
		types.IdentSlice:   false,
		types.DeclSlice:    false,
		types.ExprSlice:    false,
		types.SpecSlice:    false,
		types.StmtSlice:    false,
		types.BasicLit:     false,
		types.BlockStmt:    false,
		types.CallExpr:     false,
		types.CommentGroup: false,
		types.FieldList:    false,
		types.FuncType:     false,
		types.Ident:        false,
		types.Decl:         false,
		types.Expr:         false,
		types.Stmt:         false,
	}
	inspect(fd, func(c *cursor) bool {
		if !flags[c.field.expected] {
			flags[c.field.expected] = true
		}
		return true
	})
	l := []types.NodeType{}
	for flag, seen := range flags {
		if seen {
			l = append(l, flag)
		}
	}
	return l
}

func mutualFieldTypes(fd1, fd2 *ast.FuncDecl) ([]types.NodeType, error) {
	t1, t2 := marktypes(fd1), marktypes(fd2)
	mutuals := setops.Intersect(t1, t2)
	if len(mutuals) == 0 {
		return nil, fmt.Errorf("There is no single field type mutually existing in the entire subtree of two function declarations")
	}
	return mutuals, nil
}
