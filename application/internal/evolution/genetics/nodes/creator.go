package nodes

import (
	"fmt"
	"tde/internal/astw/types"
	"tde/internal/evolution/symbols"
)

type Creator struct {
	sm *symbols.SymbolsMngr
}

func NewCreator(sm *symbols.SymbolsMngr) *Creator {
	return &Creator{
		sm: sm,
	}
}

var (
	ErrLimitReached      = fmt.Errorf("limit reached")
	ErrNoAvailableValues = fmt.Errorf("no available values")
)

var AllowedPackagesToImport = []string{"fmt", "strings", "math"}

func (c *Creator) InType(t types.NodeType, l int) (any, error) {
	switch {
	case t.IsDecl():
		return c.Decl(l)
	case t.IsExpr():
		return c.Expr(l)
	case t.IsSpec():
		return c.Spec(l)
	case t.IsStmt():
		return c.Stmt(l)
	default:
		panic(fmt.Sprintf("unexpected type %q", t))
	}
}
