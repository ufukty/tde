package embedding

import (
	"tde/internal/embedding/models"
	"tde/models/in_program_models"

	"fmt"
	"os"

	"github.com/otiai10/copy"
)

func createTemporaryDirs(config *models.OperationConfig, candidates []*in_program_models.Candidate) {
	for _, candidate := range candidates {
		config.CandidateDirs[candidate.UUID] = os.TempDir()
	}

}

func duplicateOriginalModule(config *models.OperationConfig, candidates []*in_program_models.Candidate) {
	for _, candidate := range candidates {
		fmt.Println("copying module for candidate:", string(candidate.UUID))
		copy.Copy(config.ModulePath, string(candidate.UUID))
	}
}

func InitialEmbed(config *models.OperationConfig, candidates []*in_program_models.Candidate) {
	createTemporaryDirs(config, candidates)
	duplicateOriginalModule(config, candidates)

	// TODO: Copy original module into moduleDir

	// TODO: Unarchive
	// TODO: Rewrite changed file
}

func UpdateEmbedding(moduleDir string, config *models.EmbeddingConfig) {
	// rewrite changed file
}

func Embed(ec *models.EmbeddingConfig) error {
	return nil
}
