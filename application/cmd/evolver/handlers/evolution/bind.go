package evolution

import common_models "tde/models/common-models"

//go:generate serdeser bind.go

type (
	Request struct {
		ArchiveID       string                         `json:"archive_id"`
		File            string                         `json:"file"`
		EvolutionTarget *common_models.EvolutionTarget `json:"evolution_target"`
		EvolutionConfig *common_models.EvolutionConfig `json:"evolution_config"`
	}

	Response struct {
		Started bool
		CaseID  string
	}
)
