package mutation

import (
	"tde/internal/genetics/common"
	"tde/internal/genetics/mutation/import_path"
	"tde/internal/genetics/mutation/literals"
	"tde/internal/genetics/mutation/remove_line"
	"tde/internal/genetics/mutation/switch_lines"
	"tde/internal/genetics/mutation/token_shuffle"
	"tde/internal/utilities"
)

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
