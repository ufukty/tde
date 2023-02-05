package pattern_constructor

import (
	"tde/internal/cfg/context_resolution/context"
	"tde/internal/utilities"

	"go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func identReplacer(ctx *context.Context, replacements map[*ast.Ident]*ast.Ident, c *astutil.Cursor, ident *ast.Ident) {
	if _, ok := replacements[ident]; !ok {
		replacements[ident] = *utilities.Pick(ctx.GetVariables())
	}
	c.Replace(replacements[ident])
}

func createReplacer(ctx *context.Context, replacements map[*ast.Ident]*ast.Ident) func(c *astutil.Cursor) bool {
	return func(c *astutil.Cursor) bool {
		switch node := c.Node().(type) {
		case *ast.Ident:
			identReplacer(ctx, replacements, c, node)
		}
		return true
	}
}

func Replace(ctx *context.Context, template ast.Node) (implementation ast.Node) {
	replacements := map[*ast.Ident]*ast.Ident{}
	implementation = template
	astutil.Apply(template, createReplacer(ctx, replacements), nil)
	return
}
