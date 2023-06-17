package session_post

import (
	sessions "tde/cmd/evolver/internal/sessions"
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
		response.Status = sessions.Initialized
		return
	}

	var (
		sc        = &sessions.Config{}
		em        = evolution.NewManager((*evolution.Target)(&request.EvolutionTarget))
		session   = sessions.NewSession(em, sc)
		sessionId = sessionStore.Add(session)
	)

	response.SessionId = string(sessionId)

	// go sessionManager.Iterate(caseId) // async

	return
}
