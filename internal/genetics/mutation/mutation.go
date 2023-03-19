package mutation

import (
	"go/ast"
	"tde/internal/genetics/mutation/common"
	"tde/internal/genetics/mutation/import_path"
	"tde/internal/genetics/mutation/literals"
	"tde/internal/genetics/mutation/remove_line"
	"tde/internal/genetics/mutation/switch_lines"
	"tde/internal/genetics/mutation/token_shuffle"
	"tde/internal/utilities"
)

type GeneticOperationContext struct {
	Package         *ast.Package
	File            *ast.File
	FuncDecl        *ast.FuncDecl
	AllowedPackages []string
}

type GeneticOperation func(*GeneticOperationContext) bool

// TODO: RegenerateSubtree (cfg/node_constructor)
// TODO: Merge declared variables

var availableOperations = []common.GeneticOperation{
	import_path.GeneticOperation,
	literals.GeneticOperation,
	remove_line.GeneticOperation,
	switch_lines.GeneticOperation,
	token_shuffle.GeneticOperation,
}

func Pick() common.GeneticOperation {
	return *utilities.Pick(availableOperations)
}
