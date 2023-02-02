package pattern_constructor

import (
	"tde/internal/cfg/context"
	"tde/internal/utilities"

	"go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func Implement(ctx context.Context, template ast.Node) (implementation ast.Node) {
	replacements := map[*ast.Ident]*ast.Ident{}
	implementation = template

	astutil.Apply(template,
		func(c *astutil.Cursor) bool {

			if ident, ok := c.Node().(*ast.Ident); ok {

				if ident, ok := replacements[ident]; ok {
					c.Replace(ident)
				} else {
					replacementIdent := *utilities.Pick(ctx.GetVariables())
					replacements[ident] = replacementIdent
					c.Replace(replacementIdent)
				}

			}

			return true
		},
		nil,
	)

	return
}
