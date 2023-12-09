package nodes

import (
	"fmt"
	"tde/internal/astw/types"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"
)

var (
	ErrLimitReached      = fmt.Errorf("limit reached")
	ErrNoAvailableValues = fmt.Errorf("no available values")
)

var AllowedPackagesToImport = []string{"fmt", "strings", "math"}

func InType(t types.NodeType, ctx *context.Context, limit int) any {
	switch {
	case t.IsDecl():
		return Decl(ctx, limit)
	case t.IsExpr():
		return Expr(ctx, limit)
	case t.IsSpec():
		return Spec(ctx, limit)
	case t.IsStmt():
		return Stmt(ctx, limit)
	default:
		panic(fmt.Sprintf("unexpected type %q", t))
	}
}
