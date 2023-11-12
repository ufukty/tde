package mutation

import (
	"fmt"
	"tde/internal/evolution/genetics/mutation/v1/imports"
	"tde/internal/evolution/genetics/mutation/v1/lines"
	"tde/internal/evolution/genetics/mutation/v1/literals"
	"tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/evolution/genetics/mutation/v1/tokens"
	"tde/internal/utilities"
)

// TODO: RegenerateSubtree (cfg/node_constructor)
// TODO: Merge declared variables

var availableOperations = []models.GeneticOperation{
	imports.GeneticOperation,
	literals.GeneticOperation,
	lines.RemoveLine,
	lines.SwapLines,
	tokens.GeneticOperation,
}

func Pick() models.GeneticOperation {
	return *utilities.Pick(availableOperations)
}

func mutate(ctx *models.MutationParameters) (err error) {
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

func Mutate(ctx *models.MutationParameters) error {
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
