package session_post

import (
	session_manager "tde/cmd/evolver/internal/sessions"
)

//go:generate serdeser bind.go

type (
	Runner struct {
		Address string `json:"address"` // ip:port
		Token   string `json:"token"`
	}

	Target struct {
		Package  string `json:"package"`
		File     string `json:"file"`
		Function string `json:"function"`
	}

	Probabilities struct {
		CrossOver float64 `json:"cross_over"`
		Mutation  float64 `json:"mutation"`
	}

	Request struct {
		ArchiveID string `json:"archive_id"`
		Timeout   string `json:"timeout"` // eg. 1h15m (h)ours (m)inutes (s)econds. Don't put space between.
		Runner    Runner `json:"runner"`
		Iterate   int    `json:"iterate"`

		Probabilities   Probabilities `json:"probabilities"`
		Population      int           `json:"populatiratioon"`
		SizeLimitBytes  int           `json:"size_limit_bytes"`
		AllowedPackages []string      `json:"allowed_packages"`

		TestName string `json:"test_name"`
		Target   Target `json:"target"`
	}
)

type (
	Response struct {
		SessionId string                 `json:"session_id"`
		Status    session_manager.Status `json:"status"`
	}
)
