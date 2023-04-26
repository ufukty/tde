package evolution

import (
	"tde/cmd/evolver-server/internal/case_manager"
	"tde/internal/evolution"
	"tde/models/dto"
)

type ArchiveStatus int

const (
	ArchiveAccessible = ArchiveStatus(iota)
	ArchiveNoAuth
	ArchiveNotFound
)

func validateArchiveID(archiveID string) ArchiveStatus {
	return ArchiveNotFound
}

func Controller(request dto.EvolverService_Evolve_Request) (response dto.EvolverService_Evolve_Response) {
	archiveStatus := validateArchiveID(request.ArchiveID)
	if archiveStatus == ArchiveNoAuth || archiveStatus == ArchiveNotFound {
		response.Started = false
		return
	}

	cs := &case_manager.Case{
		EvolutionManager: evolution.NewEvolutionManager(request.EvolutionTarget),
		EvolutionConfig:  request.EvolutionConfig,
	}
	caseId := caseManager.NewCase(cs)
	response.CaseID = string(caseId)
	response.Started = true
	go caseManager.Iterate(caseId) // async

	return
}
