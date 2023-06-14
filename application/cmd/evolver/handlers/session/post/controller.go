package session_post

import (
	case_manager "tde/cmd/evolver/internal/case-manager"
	"tde/internal/evolution"
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

func Controller(request Request) (response Response) {
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
