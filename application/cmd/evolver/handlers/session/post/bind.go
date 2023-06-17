package session_post

import (
	"go/ast"
	session_manager "tde/cmd/evolver/internal/sessions"
	sessions "tde/cmd/evolver/internal/sessions"
)

//go:generate serdeser bind.go

type (
	Runner struct {
		Address string `json:"address"`
		Port    string `json:"port"`
		Token   string `json:"token"`
	}

	EvolutionTarget struct {
		// TODO: Module          map[string]*ast.Package // import path -> package
		Package  *ast.Package  `json:"package"`
		File     *ast.File     `json:"file"`
		FuncDecl *ast.FuncDecl `json:"func_decl"`
	}

	Probabilities struct {
		CrossOver float64 `json:"cross_over"`
		Mutation  float64 `json:"mutation"`
	}

	EvolutionConfig struct {
		Timeout int    `json:"timeout"` // in seconds
		Runner  string `json:"runner"`  // ip address

		Probabilities   Probabilities `json:"probabilities"`
		Population      int           `json:"population"`
		SizeLimit       int           `json:"size_limit"`
		AllowedPackages []string      `json:"allowed_packages"` // packages allowed to import

		Iterate  int    `json:"iterate"`
		TestName string `json:"test_name"`
	}

	Request struct {
		ArchiveID       string          `json:"archive_id"`
		Runner          Runner          `json:"runner"`
		EvolutionTarget EvolutionTarget `json:"evolution_target"`
		EvolutionConfig EvolutionConfig `json:"evolution_config"`
	}

	Response struct {
		SessionId string                 `json:"session_id"`
		Status    session_manager.Status `json:"status"`
	}
)
