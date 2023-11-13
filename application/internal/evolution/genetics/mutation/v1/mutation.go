package mutation

import (
	"fmt"
	"log"
	"tde/internal/evolution/genetics/mutation/v1/imports"
	"tde/internal/evolution/genetics/mutation/v1/lines"
	"tde/internal/evolution/genetics/mutation/v1/literals"
	"tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/evolution/genetics/mutation/v1/stg"
	"tde/internal/evolution/genetics/mutation/v1/tokens"
	"tde/internal/utilities"

	"golang.org/x/exp/maps"
)

// TODO: RegenerateSubtree (cfg/node_constructor)
// TODO: Merge declared variables

// var availableOperations = []models.GeneticOperation{
// 	imports.GeneticOperation,
// 	literals.GeneticOperation,
// 	lines.RemoveLine,
// 	lines.SwapLines,
// 	tokens.GeneticOperation,
// }

//go:generate stringer -type MutationOperation

type MutationOperation int

const (
	MOImportPackage = MutationOperation(iota)
	MOLiteralValueAlter
	MORemoveLine
	MOSwapLines
	MOTokenShuffle
	MOSTG
)

var mutators = map[MutationOperation]models.GeneticOperation{
	MOImportPackage:     imports.ImportPackage,
	MOLiteralValueAlter: literals.LiteralValueAlter,
	MORemoveLine:        lines.RemoveLine,
	MOSwapLines:         lines.SwapLines,
	MOTokenShuffle:      tokens.TokenShuffle,
	MOSTG:               stg.Develop,
}

func runMutator(mutator models.GeneticOperation, params *models.MutationParameters) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	return mutator(params)
}

func Mutate(params *models.MutationParameters) error {
	for attempts := 0; attempts < 50; attempts++ {
		MO := *utilities.Pick(maps.Keys(mutators))
		mutator, ok := mutators[MO]
		if !ok {
			return fmt.Errorf("unkown mutation operation %s", MO)
		}
		err := runMutator(mutator, params)
		if err == nil {
			log.Println("applied", MO)
			return nil
		} else {
			log.Println(fmt.Errorf("applying %s: %w", MO, err))
		}
	}
	return fmt.Errorf("Limit to try is reached.")
}
