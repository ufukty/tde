package sessions

import (
	"tde/internal/evolution"

	"github.com/google/uuid"
)

type SessionId string

func NewSessionId() SessionId {
	return SessionId(uuid.NewString())
}

type Session struct {
	EvolutionManager *evolution.Manager
	Config           *Config
	Status           Status
}

func NewSession(em *evolution.Manager, config *Config) *Session {
	return &Session{
		Status:           Initialized,
		EvolutionManager: em,
		Config:           config,
	}
}
