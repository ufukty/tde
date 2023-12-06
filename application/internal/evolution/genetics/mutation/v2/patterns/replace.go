package patterns

import (
	"fmt"
	"go/ast"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"
	"tde/internal/utilities/pick"

	"golang.org/x/tools/go/ast/astutil"
)

func identReplacer(ctx *context.Context, replacements map[*ast.Ident]*ast.Ident, c *astutil.Cursor, ident *ast.Ident) error {
	if _, ok := replacements[ident]; !ok {
		var err error
		replacements[ident], err = pick.Pick(ctx.GetVariables())
		if err != nil {
			return fmt.Errorf("picking ident from context: %w", err)
		}
	}
	c.Replace(replacements[ident])
	return nil
}

func Replace(ctx *context.Context, template ast.Node) (implementation ast.Node, err error) {
	replacements := map[*ast.Ident]*ast.Ident{}
	implementation = template
	astutil.Apply(template, func(c *astutil.Cursor) bool {
		switch node := c.Node().(type) {
		case *ast.Ident:
			if err := identReplacer(ctx, replacements, c, node); err != nil {
				err = fmt.Errorf("identReplacer: %w", err)
			}
		}
		return true
	}, nil)
	return
}
