package mutation

import (
	"fmt"
	import_path "tde/internal/evolution/genetics/mutation/v1/import-path"
	"tde/internal/evolution/genetics/mutation/v1/literals"
	"tde/internal/evolution/genetics/mutation/v1/models"
	remove_line "tde/internal/evolution/genetics/mutation/v1/remove-line"
	switch_lines "tde/internal/evolution/genetics/mutation/v1/switch-lines"
	token_shuffle "tde/internal/evolution/genetics/mutation/v1/token-shuffle"
	"tde/internal/utilities"
)

// TODO: RegenerateSubtree (cfg/node_constructor)
// TODO: Merge declared variables

var availableOperations = []models.GeneticOperation{
	import_path.GeneticOperation,
	literals.GeneticOperation,
	remove_line.GeneticOperation,
	switch_lines.GeneticOperation,
	token_shuffle.GeneticOperation,
}

func Pick() models.GeneticOperation {
	return *utilities.Pick(availableOperations)
}

func mutate(ctx *models.GeneticOperationContext) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	op := Pick()
	if ok := op(ctx); !ok {
		err = fmt.Errorf("operation (%v) returned false", op)
	}
	return
}

func Mutate(ctx *models.GeneticOperationContext) error {
	for attempts := 0; true; attempts++ {
		if attempts == 10 {
			return fmt.Errorf("Limit to try is reached.")
		}
		if err := mutate(ctx); err != nil {
			return fmt.Errorf("mutating: %w", err)
		}
	}
	return nil
}
