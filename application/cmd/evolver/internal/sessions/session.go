package sessions

import (
	"tde/internal/evolution"
)

type Session struct {
	Manager *evolution.Manager
	Config  *Config
	Status  Status
}

func NewSession(em *evolution.Manager, config *Config) *Session {
	return &Session{
		Manager: em,
		Config:  config,
		Status:  Initialized,
	}
}
