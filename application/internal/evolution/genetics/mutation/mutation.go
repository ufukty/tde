package mutation

import (
	"tde/internal/evolution/genetics/mutation/common"
	import_path "tde/internal/evolution/genetics/mutation/import-path"
	"tde/internal/evolution/genetics/mutation/literals"
	remove_line "tde/internal/evolution/genetics/mutation/remove-line"
	switch_lines "tde/internal/evolution/genetics/mutation/switch-lines"
	token_shuffle "tde/internal/evolution/genetics/mutation/token-shuffle"
	"tde/internal/utilities"
)

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
