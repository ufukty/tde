package sessions

import (
	"tde/internal/evolution"
)

type Session struct {
	Manager *evolution.SolutionSearch
	Config  *Config
	Status  Status
}

func NewSession(em *evolution.SolutionSearch, config *Config) *Session {
	return &Session{
		Manager: em,
		Config:  config,
		Status:  Initialized,
	}
}
