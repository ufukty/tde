package evolution

import (
	"tde/internal/evolution"
)

//go:generate serdeser bind.go

type (
	Request struct {
		ArchiveID       string                     `json:"archive_id"`
		File            string                     `json:"file"`
		EvolutionTarget *evolution.EvolutionTarget `json:"evolution_target"`
		EvolutionConfig *evolution.EvolutionConfig `json:"evolution_config"`
	}

	Response struct {
		Started bool
		CaseID  string
	}
)
