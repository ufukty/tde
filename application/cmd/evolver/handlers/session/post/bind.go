package session_post

import (
	"go/ast"
	"tde/internal/evolution"
)

//go:generate serdeser bind.go

type (
	OnPremisesRunnerConfig struct {
		Address string
		Port    string
		Token   string
	}

	EvolutionTarget struct {
		// TODO: Module          map[string]*ast.Package // import path -> package
		Package  *ast.Package
		File     *ast.File
		FuncDecl *ast.FuncDecl
	}

	EvolutionConfig struct {
		Timeout    int      // in seconds
		Runner     string   // ip address
		Continue   string   // session
		Model      string   //
		Ratios     string   //
		Population int      //
		Iterate    int      //
		Size       int      //
		Package    []string // packages allowed to import
		TestName   string
	}

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
