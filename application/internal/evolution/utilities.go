package evolution

import (
	"fmt"
	"go/ast"
	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone"
	models "tde/models/program"

	"github.com/google/uuid"
)

func newCandidate(pkg *ast.Package, file *ast.File, funcDecl *ast.FuncDecl) (*models.Subject, error) {
	cloneFile := clone.File(file)
	clondeFuncDecl, err := astwutl.FindFuncDecl(cloneFile, funcDecl.Name.Name)
	if err != nil {
		return nil, fmt.Errorf("finding the target function in file ast: %w", err)
	}
	return &models.Subject{
		Sid: models.Sid(uuid.New().String()), // UUID v4
		AST: models.TargetAst{
			Package:  pkg,
			File:     cloneFile,
			FuncDecl: clondeFuncDecl,
		},
	}, nil
}
